package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceNextHop() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNextHopRead,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_distinguisher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceNextHopRead(d *schema.ResourceData, m interface{}) error {
	filteredNextHops := vspk.NextHopsList{}
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
	parent := &vspk.Link{ID: d.Get("parent_link").(string)}
	filteredNextHops, err = parent.NextHops(fetchFilter)
	if err != nil {
		return err
	}

	NextHop := &vspk.NextHop{}

	if len(filteredNextHops) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNextHops) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NextHop = filteredNextHops[0]

	d.Set("last_updated_by", NextHop.LastUpdatedBy)
	d.Set("entity_scope", NextHop.EntityScope)
	d.Set("route_distinguisher", NextHop.RouteDistinguisher)
	d.Set("ip", NextHop.Ip)
	d.Set("external_id", NextHop.ExternalID)

	d.Set("id", NextHop.Identifier())
	d.Set("parent_id", NextHop.ParentID)
	d.Set("parent_type", NextHop.ParentType)
	d.Set("owner", NextHop.Owner)

	d.SetId(NextHop.Identifier())

	return nil
}
