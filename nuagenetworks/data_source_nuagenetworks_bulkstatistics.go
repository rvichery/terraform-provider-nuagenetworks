package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceBulkStatistics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBulkStatisticsRead,
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
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_data_points": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_patnat_pool": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBulkStatisticsRead(d *schema.ResourceData, m interface{}) error {
	filteredBulkStatistics := vspk.BulkStatisticsList{}
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
	parent := &vspk.PATNATPool{ID: d.Get("parent_patnat_pool").(string)}
	filteredBulkStatistics, err = parent.BulkStatistics(fetchFilter)
	if err != nil {
		return err
	}

	BulkStatistics := &vspk.BulkStatistics{}

	if len(filteredBulkStatistics) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredBulkStatistics) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	BulkStatistics = filteredBulkStatistics[0]

	d.Set("data", BulkStatistics.Data)
	d.Set("version", BulkStatistics.Version)
	d.Set("end_time", BulkStatistics.EndTime)
	d.Set("start_time", BulkStatistics.StartTime)
	d.Set("number_of_data_points", BulkStatistics.NumberOfDataPoints)

	d.Set("id", BulkStatistics.Identifier())
	d.Set("parent_id", BulkStatistics.ParentID)
	d.Set("parent_type", BulkStatistics.ParentType)
	d.Set("owner", BulkStatistics.Owner)

	d.SetId(BulkStatistics.Identifier())

	return nil
}
