package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceCloudMgmtSystem() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudMgmtSystemRead,
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
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudMgmtSystemRead(d *schema.ResourceData, m interface{}) error {
	filteredCloudMgmtSystems := vspk.CloudMgmtSystemsList{}
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
	parent := m.(*vspk.Me)
	filteredCloudMgmtSystems, err = parent.CloudMgmtSystems(fetchFilter)
	if err != nil {
		return err
	}

	CloudMgmtSystem := &vspk.CloudMgmtSystem{}

	if len(filteredCloudMgmtSystems) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCloudMgmtSystems) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	CloudMgmtSystem = filteredCloudMgmtSystems[0]

	d.Set("name", CloudMgmtSystem.Name)
	d.Set("last_updated_by", CloudMgmtSystem.LastUpdatedBy)
	d.Set("entity_scope", CloudMgmtSystem.EntityScope)
	d.Set("external_id", CloudMgmtSystem.ExternalID)

	d.Set("id", CloudMgmtSystem.Identifier())
	d.Set("parent_id", CloudMgmtSystem.ParentID)
	d.Set("parent_type", CloudMgmtSystem.ParentType)
	d.Set("owner", CloudMgmtSystem.Owner)

	d.SetId(CloudMgmtSystem.Identifier())

	return nil
}
