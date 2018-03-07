package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceGatewaySecuredData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGatewaySecuredDataRead,
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
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_cert_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keyserver_cert_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"signed_data": &schema.Schema{
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
			"parent_gateway_security": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGatewaySecuredDataRead(d *schema.ResourceData, m interface{}) error {
	filteredGatewaySecuredDatas := vspk.GatewaySecuredDatasList{}
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
	parent := &vspk.GatewaySecurity{ID: d.Get("parent_gateway_security").(string)}
	filteredGatewaySecuredDatas, err = parent.GatewaySecuredDatas(fetchFilter)
	if err != nil {
		return err
	}

	GatewaySecuredData := &vspk.GatewaySecuredData{}

	if len(filteredGatewaySecuredDatas) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredGatewaySecuredDatas) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		GatewaySecuredData = filteredGatewaySecuredDatas[0]
	}

	d.Set("last_updated_by", GatewaySecuredData.LastUpdatedBy)
	d.Set("data", GatewaySecuredData.Data)
	d.Set("gateway_cert_serial_number", GatewaySecuredData.GatewayCertSerialNumber)
	d.Set("keyserver_cert_serial_number", GatewaySecuredData.KeyserverCertSerialNumber)
	d.Set("signed_data", GatewaySecuredData.SignedData)
	d.Set("entity_scope", GatewaySecuredData.EntityScope)
	d.Set("external_id", GatewaySecuredData.ExternalID)

	d.Set("id", GatewaySecuredData.Identifier())
	d.Set("parent_id", GatewaySecuredData.ParentID)
	d.Set("parent_type", GatewaySecuredData.ParentType)
	d.Set("owner", GatewaySecuredData.Owner)

	d.SetId(GatewaySecuredData.Identifier())

	return nil
}
