package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIKEGatewayConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKEGatewayConfigRead,
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
			"config": {
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

func dataSourceIKEGatewayConfigRead(d *schema.ResourceData, m interface{}) error {
	filteredIKEGatewayConfigs := vspk.IKEGatewayConfigsList{}
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
	filteredIKEGatewayConfigs, err = parent.IKEGatewayConfigs(fetchFilter)
	if err != nil {
		return err
	}

	IKEGatewayConfig := &vspk.IKEGatewayConfig{}

	if len(filteredIKEGatewayConfigs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKEGatewayConfigs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IKEGatewayConfig = filteredIKEGatewayConfigs[0]

	d.Set("last_updated_by", IKEGatewayConfig.LastUpdatedBy)
	d.Set("entity_scope", IKEGatewayConfig.EntityScope)
	d.Set("config", IKEGatewayConfig.Config)
	d.Set("external_id", IKEGatewayConfig.ExternalID)

	d.Set("id", IKEGatewayConfig.Identifier())
	d.Set("parent_id", IKEGatewayConfig.ParentID)
	d.Set("parent_type", IKEGatewayConfig.ParentType)
	d.Set("owner", IKEGatewayConfig.Owner)

	d.SetId(IKEGatewayConfig.Identifier())

	return nil
}
