package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNetworkMacroGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkMacroGroupRead,
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
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_macros": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise_network": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise_network"},
			},
		},
	}
}

func dataSourceNetworkMacroGroupRead(d *schema.ResourceData, m interface{}) error {
	filteredNetworkMacroGroups := vspk.NetworkMacroGroupsList{}
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
	if attr, ok := d.GetOk("parent_enterprise_network"); ok {
		parent := &vspk.EnterpriseNetwork{ID: attr.(string)}
		filteredNetworkMacroGroups, err = parent.NetworkMacroGroups(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredNetworkMacroGroups, err = parent.NetworkMacroGroups(fetchFilter)
		if err != nil {
			return err
		}
	}

	NetworkMacroGroup := &vspk.NetworkMacroGroup{}

	if len(filteredNetworkMacroGroups) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNetworkMacroGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NetworkMacroGroup = filteredNetworkMacroGroups[0]

	d.Set("name", NetworkMacroGroup.Name)
	d.Set("last_updated_by", NetworkMacroGroup.LastUpdatedBy)
	d.Set("description", NetworkMacroGroup.Description)
	d.Set("network_macros", NetworkMacroGroup.NetworkMacros)
	d.Set("entity_scope", NetworkMacroGroup.EntityScope)
	d.Set("external_id", NetworkMacroGroup.ExternalID)

	d.Set("id", NetworkMacroGroup.Identifier())
	d.Set("parent_id", NetworkMacroGroup.ParentID)
	d.Set("parent_type", NetworkMacroGroup.ParentType)
	d.Set("owner", NetworkMacroGroup.Owner)

	d.SetId(NetworkMacroGroup.Identifier())

	return nil
}
