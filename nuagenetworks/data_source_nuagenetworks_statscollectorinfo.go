package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceStatsCollectorInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStatsCollectorInfoRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proto_buf_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceStatsCollectorInfoRead(d *schema.ResourceData, m interface{}) error {
	filteredStatsCollectorInfos := vspk.StatsCollectorInfosList{}
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
	filteredStatsCollectorInfos, err = parent.StatsCollectorInfos(fetchFilter)
	if err != nil {
		return err
	}

	StatsCollectorInfo := &vspk.StatsCollectorInfo{}

	if len(filteredStatsCollectorInfos) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredStatsCollectorInfos) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		StatsCollectorInfo = filteredStatsCollectorInfos[0]
	}

	d.Set("last_updated_by", StatsCollectorInfo.LastUpdatedBy)
	d.Set("address_type", StatsCollectorInfo.AddressType)
	d.Set("entity_scope", StatsCollectorInfo.EntityScope)
	d.Set("port", StatsCollectorInfo.Port)
	d.Set("ip_address", StatsCollectorInfo.IpAddress)
	d.Set("proto_buf_port", StatsCollectorInfo.ProtoBufPort)
	d.Set("external_id", StatsCollectorInfo.ExternalID)

	d.Set("id", StatsCollectorInfo.Identifier())
	d.Set("parent_id", StatsCollectorInfo.ParentID)
	d.Set("parent_type", StatsCollectorInfo.ParentType)
	d.Set("owner", StatsCollectorInfo.Owner)

	d.SetId(StatsCollectorInfo.Identifier())

	return nil
}
