package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDefaultGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDefaultGatewayRead,
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
			"gateway_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mac_address": &schema.Schema{
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

func dataSourceDefaultGatewayRead(d *schema.ResourceData, m interface{}) error {
	filteredDefaultGateways := vspk.DefaultGatewaysList{}
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
	filteredDefaultGateways, err = parent.DefaultGateways(fetchFilter)
	if err != nil {
		return err
	}

	DefaultGateway := &vspk.DefaultGateway{}

	if len(filteredDefaultGateways) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDefaultGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		DefaultGateway = filteredDefaultGateways[0]
	}

	d.Set("name", DefaultGateway.Name)
	d.Set("gateway_ip_address", DefaultGateway.GatewayIPAddress)
	d.Set("gateway_mac_address", DefaultGateway.GatewayMACAddress)

	d.Set("id", DefaultGateway.Identifier())
	d.Set("parent_id", DefaultGateway.ParentID)
	d.Set("parent_type", DefaultGateway.ParentType)
	d.Set("owner", DefaultGateway.Owner)

	d.SetId(DefaultGateway.Identifier())

	return nil
}
