package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceRedundancyGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRedundancyGroupRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_autodiscovered_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_autodiscovered_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundant_gateway_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"personality": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vtep": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceRedundancyGroupRead(d *schema.ResourceData, m interface{}) error {
	filteredRedundancyGroups := vspk.RedundancyGroupsList{}
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
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredRedundancyGroups, err = parent.RedundancyGroups(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredRedundancyGroups, err = parent.RedundancyGroups(fetchFilter)
		if err != nil {
			return err
		}
	}

	RedundancyGroup := &vspk.RedundancyGroup{}

	if len(filteredRedundancyGroups) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredRedundancyGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		RedundancyGroup = filteredRedundancyGroups[0]
	}

	d.Set("name", RedundancyGroup.Name)
	d.Set("last_updated_by", RedundancyGroup.LastUpdatedBy)
	d.Set("gateway_peer1_autodiscovered_gateway_id", RedundancyGroup.GatewayPeer1AutodiscoveredGatewayID)
	d.Set("gateway_peer1_id", RedundancyGroup.GatewayPeer1ID)
	d.Set("gateway_peer1_name", RedundancyGroup.GatewayPeer1Name)
	d.Set("gateway_peer2_autodiscovered_gateway_id", RedundancyGroup.GatewayPeer2AutodiscoveredGatewayID)
	d.Set("gateway_peer2_id", RedundancyGroup.GatewayPeer2ID)
	d.Set("gateway_peer2_name", RedundancyGroup.GatewayPeer2Name)
	d.Set("redundant_gateway_status", RedundancyGroup.RedundantGatewayStatus)
	d.Set("permitted_action", RedundancyGroup.PermittedAction)
	d.Set("personality", RedundancyGroup.Personality)
	d.Set("description", RedundancyGroup.Description)
	d.Set("enterprise_id", RedundancyGroup.EnterpriseID)
	d.Set("entity_scope", RedundancyGroup.EntityScope)
	d.Set("vtep", RedundancyGroup.Vtep)
	d.Set("external_id", RedundancyGroup.ExternalID)

	d.Set("id", RedundancyGroup.Identifier())
	d.Set("parent_id", RedundancyGroup.ParentID)
	d.Set("parent_type", RedundancyGroup.ParentType)
	d.Set("owner", RedundancyGroup.Owner)

	d.SetId(RedundancyGroup.Identifier())

	return nil
}
