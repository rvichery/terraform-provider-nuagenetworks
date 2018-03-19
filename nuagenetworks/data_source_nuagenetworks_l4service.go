package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceL4Service() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceL4ServiceRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_service": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_l4_service_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l4_service_group"},
			},
		},
	}
}

func dataSourceL4ServiceRead(d *schema.ResourceData, m interface{}) error {
	filteredL4Services := vspk.L4ServicesList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	if attr, ok := d.GetOk("parent_l4_service_group"); ok {
		parent := &vspk.L4ServiceGroup{ID: attr.(string)}
		filteredL4Services, err = parent.L4Services(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredL4Services, err = parent.L4Services(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredL4Services, err = parent.L4Services(fetchFilter)
		if err != nil {
			return err
		}
	}

	L4Service := &vspk.L4Service{}

	if len(filteredL4Services) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredL4Services) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	L4Service = filteredL4Services[0]

	d.Set("name", L4Service.Name)
	d.Set("last_updated_by", L4Service.LastUpdatedBy)
	d.Set("default_service", L4Service.DefaultService)
	d.Set("description", L4Service.Description)
	d.Set("entity_scope", L4Service.EntityScope)
	d.Set("ports", L4Service.Ports)
	d.Set("protocol", L4Service.Protocol)
	d.Set("external_id", L4Service.ExternalID)

	d.Set("id", L4Service.Identifier())
	d.Set("parent_id", L4Service.ParentID)
	d.Set("parent_type", L4Service.ParentType)
	d.Set("owner", L4Service.Owner)

	d.SetId(L4Service.Identifier())

	return nil
}
