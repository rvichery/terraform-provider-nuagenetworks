package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceGatewaySecurity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGatewaySecurityRead,
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
			"gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision": {
				Type:     schema.TypeInt,
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
			"parent_ns_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGatewaySecurityRead(d *schema.ResourceData, m interface{}) error {
	filteredGatewaySecurities := vspk.GatewaySecuritiesList{}
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
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	filteredGatewaySecurities, err = parent.GatewaySecurities(fetchFilter)
	if err != nil {
		return err
	}

	GatewaySecurity := &vspk.GatewaySecurity{}

	if len(filteredGatewaySecurities) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredGatewaySecurities) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	GatewaySecurity = filteredGatewaySecurities[0]

	d.Set("last_updated_by", GatewaySecurity.LastUpdatedBy)
	d.Set("gateway_id", GatewaySecurity.GatewayID)
	d.Set("revision", GatewaySecurity.Revision)
	d.Set("entity_scope", GatewaySecurity.EntityScope)
	d.Set("external_id", GatewaySecurity.ExternalID)

	d.Set("id", GatewaySecurity.Identifier())
	d.Set("parent_id", GatewaySecurity.ParentID)
	d.Set("parent_type", GatewaySecurity.ParentType)
	d.Set("owner", GatewaySecurity.Owner)

	d.SetId(GatewaySecurity.Identifier())

	return nil
}
