package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceIKESubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKESubnetRead,
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
			"prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ike_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ike_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceIKESubnetRead(d *schema.ResourceData, m interface{}) error {
	filteredIKESubnets := vspk.IKESubnetsList{}
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
	parent := &vspk.IKEGateway{ID: d.Get("parent_ike_gateway").(string)}
	filteredIKESubnets, err = parent.IKESubnets(fetchFilter)
	if err != nil {
		return err
	}

	IKESubnet := &vspk.IKESubnet{}

	if len(filteredIKESubnets) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKESubnets) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IKESubnet = filteredIKESubnets[0]

	d.Set("last_updated_by", IKESubnet.LastUpdatedBy)
	d.Set("entity_scope", IKESubnet.EntityScope)
	d.Set("prefix", IKESubnet.Prefix)
	d.Set("associated_ike_gateway_id", IKESubnet.AssociatedIKEGatewayID)
	d.Set("external_id", IKESubnet.ExternalID)

	d.Set("id", IKESubnet.Identifier())
	d.Set("parent_id", IKESubnet.ParentID)
	d.Set("parent_type", IKESubnet.ParentType)
	d.Set("owner", IKESubnet.Owner)

	d.SetId(IKESubnet.Identifier())

	return nil
}
