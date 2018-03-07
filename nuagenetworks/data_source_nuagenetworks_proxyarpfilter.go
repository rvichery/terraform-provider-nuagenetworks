package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceProxyARPFilter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyARPFilterRead,
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
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceProxyARPFilterRead(d *schema.ResourceData, m interface{}) error {
	filteredProxyARPFilters := vspk.ProxyARPFiltersList{}
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
	parent := &vspk.Subnet{ID: d.Get("parent_subnet").(string)}
	filteredProxyARPFilters, err = parent.ProxyARPFilters(fetchFilter)
	if err != nil {
		return err
	}

	ProxyARPFilter := &vspk.ProxyARPFilter{}

	if len(filteredProxyARPFilters) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredProxyARPFilters) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		ProxyARPFilter = filteredProxyARPFilters[0]
	}

	d.Set("ip_type", ProxyARPFilter.IPType)
	d.Set("last_updated_by", ProxyARPFilter.LastUpdatedBy)
	d.Set("max_address", ProxyARPFilter.MaxAddress)
	d.Set("min_address", ProxyARPFilter.MinAddress)
	d.Set("entity_scope", ProxyARPFilter.EntityScope)
	d.Set("external_id", ProxyARPFilter.ExternalID)

	d.Set("id", ProxyARPFilter.Identifier())
	d.Set("parent_id", ProxyARPFilter.ParentID)
	d.Set("parent_type", ProxyARPFilter.ParentType)
	d.Set("owner", ProxyARPFilter.Owner)

	d.SetId(ProxyARPFilter.Identifier())

	return nil
}
