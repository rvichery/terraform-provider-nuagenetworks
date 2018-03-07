package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIKEGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKEGatewayRead,
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
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ik_ev1_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceIKEGatewayRead(d *schema.ResourceData, m interface{}) error {
	filteredIKEGateways := vspk.IKEGatewaysList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredIKEGateways, err = parent.IKEGateways(fetchFilter)
	if err != nil {
		return err
	}

	IKEGateway := &vspk.IKEGateway{}

	if len(filteredIKEGateways) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKEGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		IKEGateway = filteredIKEGateways[0]
	}

	d.Set("ike_version", IKEGateway.IKEVersion)
	d.Set("ik_ev1_mode", IKEGateway.IKEv1Mode)
	d.Set("ip_address", IKEGateway.IPAddress)
	d.Set("name", IKEGateway.Name)
	d.Set("last_updated_by", IKEGateway.LastUpdatedBy)
	d.Set("description", IKEGateway.Description)
	d.Set("entity_scope", IKEGateway.EntityScope)
	d.Set("associated_enterprise_id", IKEGateway.AssociatedEnterpriseID)
	d.Set("external_id", IKEGateway.ExternalID)

	d.Set("id", IKEGateway.Identifier())
	d.Set("parent_id", IKEGateway.ParentID)
	d.Set("parent_type", IKEGateway.ParentType)
	d.Set("owner", IKEGateway.Owner)

	d.SetId(IKEGateway.Identifier())

	return nil
}
