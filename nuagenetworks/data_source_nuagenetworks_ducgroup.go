package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDUCGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDUCGroupRead,
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
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_performance_monitor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDUCGroupRead(d *schema.ResourceData, m interface{}) error {
	filteredDUCGroups := vspk.DUCGroupsList{}
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
	filteredDUCGroups, err = parent.DUCGroups(fetchFilter)
	if err != nil {
		return err
	}

	DUCGroup := &vspk.DUCGroup{}

	if len(filteredDUCGroups) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDUCGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	DUCGroup = filteredDUCGroups[0]

	d.Set("name", DUCGroup.Name)
	d.Set("description", DUCGroup.Description)
	d.Set("associated_performance_monitor_id", DUCGroup.AssociatedPerformanceMonitorID)

	d.Set("id", DUCGroup.Identifier())
	d.Set("parent_id", DUCGroup.ParentID)
	d.Set("parent_type", DUCGroup.ParentType)
	d.Set("owner", DUCGroup.Owner)

	d.SetId(DUCGroup.Identifier())

	return nil
}
