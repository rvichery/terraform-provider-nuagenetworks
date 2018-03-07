package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceInfrastructureConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfrastructureConfigRead,
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
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ns_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceInfrastructureConfigRead(d *schema.ResourceData, m interface{}) error {
	filteredInfrastructureConfigs := vspk.InfrastructureConfigsList{}
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
	filteredInfrastructureConfigs, err = parent.InfrastructureConfigs(fetchFilter)
	if err != nil {
		return err
	}

	InfrastructureConfig := &vspk.InfrastructureConfig{}

	if len(filteredInfrastructureConfigs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredInfrastructureConfigs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		InfrastructureConfig = filteredInfrastructureConfigs[0]
	}

	d.Set("last_updated_by", InfrastructureConfig.LastUpdatedBy)
	d.Set("entity_scope", InfrastructureConfig.EntityScope)
	d.Set("config", InfrastructureConfig.Config)
	d.Set("config_status", InfrastructureConfig.ConfigStatus)
	d.Set("external_id", InfrastructureConfig.ExternalID)

	d.Set("id", InfrastructureConfig.Identifier())
	d.Set("parent_id", InfrastructureConfig.ParentID)
	d.Set("parent_type", InfrastructureConfig.ParentType)
	d.Set("owner", InfrastructureConfig.Owner)

	d.SetId(InfrastructureConfig.Identifier())

	return nil
}
