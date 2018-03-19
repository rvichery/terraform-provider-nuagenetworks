package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVSDComponent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVSDComponentRead,
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
			"management_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
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
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vsd": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVSDComponentRead(d *schema.ResourceData, m interface{}) error {
	filteredVSDComponents := vspk.VSDComponentsList{}
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
	parent := &vspk.VSD{ID: d.Get("parent_vsd").(string)}
	filteredVSDComponents, err = parent.VSDComponents(fetchFilter)
	if err != nil {
		return err
	}

	VSDComponent := &vspk.VSDComponent{}

	if len(filteredVSDComponents) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVSDComponents) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VSDComponent = filteredVSDComponents[0]

	d.Set("name", VSDComponent.Name)
	d.Set("management_ip", VSDComponent.ManagementIP)
	d.Set("address", VSDComponent.Address)
	d.Set("description", VSDComponent.Description)
	d.Set("entity_scope", VSDComponent.EntityScope)
	d.Set("location", VSDComponent.Location)
	d.Set("product_version", VSDComponent.ProductVersion)
	d.Set("status", VSDComponent.Status)
	d.Set("external_id", VSDComponent.ExternalID)
	d.Set("type", VSDComponent.Type)

	d.Set("id", VSDComponent.Identifier())
	d.Set("parent_id", VSDComponent.ParentID)
	d.Set("parent_type", VSDComponent.ParentType)
	d.Set("owner", VSDComponent.Owner)

	d.SetId(VSDComponent.Identifier())

	return nil
}
