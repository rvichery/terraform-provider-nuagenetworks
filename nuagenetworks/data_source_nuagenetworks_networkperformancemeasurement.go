package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNetworkPerformanceMeasurement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkPerformanceMeasurementRead,
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
			"npm_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"read_only": {
				Type:     schema.TypeBool,
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
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceNetworkPerformanceMeasurementRead(d *schema.ResourceData, m interface{}) error {
	filteredNetworkPerformanceMeasurements := vspk.NetworkPerformanceMeasurementsList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredNetworkPerformanceMeasurements, err = parent.NetworkPerformanceMeasurements(fetchFilter)
	if err != nil {
		return err
	}

	NetworkPerformanceMeasurement := &vspk.NetworkPerformanceMeasurement{}

	if len(filteredNetworkPerformanceMeasurements) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNetworkPerformanceMeasurements) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NetworkPerformanceMeasurement = filteredNetworkPerformanceMeasurements[0]

	d.Set("npm_type", NetworkPerformanceMeasurement.NPMType)
	d.Set("name", NetworkPerformanceMeasurement.Name)
	d.Set("read_only", NetworkPerformanceMeasurement.ReadOnly)
	d.Set("description", NetworkPerformanceMeasurement.Description)
	d.Set("associated_performance_monitor_id", NetworkPerformanceMeasurement.AssociatedPerformanceMonitorID)

	d.Set("id", NetworkPerformanceMeasurement.Identifier())
	d.Set("parent_id", NetworkPerformanceMeasurement.ParentID)
	d.Set("parent_type", NetworkPerformanceMeasurement.ParentType)
	d.Set("owner", NetworkPerformanceMeasurement.Owner)

	d.SetId(NetworkPerformanceMeasurement.Identifier())

	return nil
}
