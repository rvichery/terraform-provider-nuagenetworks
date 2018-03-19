package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceUserContext() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserContextRead,
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
			"aar_flow_stats_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"aar_probe_stats_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vss_stats_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"flow_collection_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"statistics_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"stats_tsdb_server_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserContextRead(d *schema.ResourceData, m interface{}) error {
	filteredUserContexts := vspk.UserContextsList{}
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
	filteredUserContexts, err = parent.UserContexts(fetchFilter)
	if err != nil {
		return err
	}

	UserContext := &vspk.UserContext{}

	if len(filteredUserContexts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredUserContexts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	UserContext = filteredUserContexts[0]

	d.Set("aar_flow_stats_interval", UserContext.AARFlowStatsInterval)
	d.Set("aar_probe_stats_interval", UserContext.AARProbeStatsInterval)
	d.Set("vss_stats_interval", UserContext.VSSStatsInterval)
	d.Set("flow_collection_enabled", UserContext.FlowCollectionEnabled)
	d.Set("statistics_enabled", UserContext.StatisticsEnabled)
	d.Set("stats_tsdb_server_address", UserContext.StatsTSDBServerAddress)

	d.Set("id", UserContext.Identifier())
	d.Set("parent_id", UserContext.ParentID)
	d.Set("parent_type", UserContext.ParentType)
	d.Set("owner", UserContext.Owner)

	d.SetId(UserContext.Identifier())

	return nil
}
