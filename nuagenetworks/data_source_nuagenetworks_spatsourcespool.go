package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceSPATSourcesPool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSPATSourcesPoolRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSPATSourcesPoolRead(d *schema.ResourceData, m interface{}) error {
	filteredSPATSourcesPools := vspk.SPATSourcesPoolsList{}
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
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	filteredSPATSourcesPools, err = parent.SPATSourcesPools(fetchFilter)
	if err != nil {
		return err
	}

	SPATSourcesPool := &vspk.SPATSourcesPool{}

	if len(filteredSPATSourcesPools) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredSPATSourcesPools) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		SPATSourcesPool = filteredSPATSourcesPools[0]
	}

	d.Set("name", SPATSourcesPool.Name)
	d.Set("family", SPATSourcesPool.Family)
	d.Set("address_list", SPATSourcesPool.AddressList)

	d.Set("id", SPATSourcesPool.Identifier())
	d.Set("parent_id", SPATSourcesPool.ParentID)
	d.Set("parent_type", SPATSourcesPool.ParentType)
	d.Set("owner", SPATSourcesPool.Owner)

	d.SetId(SPATSourcesPool.Identifier())

	return nil
}
