package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNSRedundantGatewayGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSRedundantGatewayGroupRead,
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_autodiscovered_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer1_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_autodiscovered_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_peer2_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heartbeat_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"heartbeat_vlanid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"redundancy_port_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"redundant_gateway_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"consecutive_failures_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNSRedundantGatewayGroupRead(d *schema.ResourceData, m interface{}) error {
	filteredNSRedundantGatewayGroups := vspk.NSRedundantGatewayGroupsList{}
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
		filteredNSRedundantGatewayGroups, err = parent.NSRedundantGatewayGroups(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredNSRedundantGatewayGroups, err = parent.NSRedundantGatewayGroups(fetchFilter)
		if err != nil {
			return err
		}
	}

	NSRedundantGatewayGroup := &vspk.NSRedundantGatewayGroup{}

	if len(filteredNSRedundantGatewayGroups) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSRedundantGatewayGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NSRedundantGatewayGroup = filteredNSRedundantGatewayGroups[0]

	d.Set("name", NSRedundantGatewayGroup.Name)
	d.Set("last_updated_by", NSRedundantGatewayGroup.LastUpdatedBy)
	d.Set("gateway_peer1_autodiscovered_gateway_id", NSRedundantGatewayGroup.GatewayPeer1AutodiscoveredGatewayID)
	d.Set("gateway_peer1_id", NSRedundantGatewayGroup.GatewayPeer1ID)
	d.Set("gateway_peer1_name", NSRedundantGatewayGroup.GatewayPeer1Name)
	d.Set("gateway_peer2_autodiscovered_gateway_id", NSRedundantGatewayGroup.GatewayPeer2AutodiscoveredGatewayID)
	d.Set("gateway_peer2_id", NSRedundantGatewayGroup.GatewayPeer2ID)
	d.Set("gateway_peer2_name", NSRedundantGatewayGroup.GatewayPeer2Name)
	d.Set("heartbeat_interval", NSRedundantGatewayGroup.HeartbeatInterval)
	d.Set("heartbeat_vlanid", NSRedundantGatewayGroup.HeartbeatVLANID)
	d.Set("redundancy_port_ids", NSRedundantGatewayGroup.RedundancyPortIDs)
	d.Set("redundant_gateway_status", NSRedundantGatewayGroup.RedundantGatewayStatus)
	d.Set("permitted_action", NSRedundantGatewayGroup.PermittedAction)
	d.Set("personality", NSRedundantGatewayGroup.Personality)
	d.Set("description", NSRedundantGatewayGroup.Description)
	d.Set("enterprise_id", NSRedundantGatewayGroup.EnterpriseID)
	d.Set("entity_scope", NSRedundantGatewayGroup.EntityScope)
	d.Set("consecutive_failures_count", NSRedundantGatewayGroup.ConsecutiveFailuresCount)
	d.Set("external_id", NSRedundantGatewayGroup.ExternalID)

	d.Set("id", NSRedundantGatewayGroup.Identifier())
	d.Set("parent_id", NSRedundantGatewayGroup.ParentID)
	d.Set("parent_type", NSRedundantGatewayGroup.ParentType)
	d.Set("owner", NSRedundantGatewayGroup.Owner)

	d.SetId(NSRedundantGatewayGroup.Identifier())

	return nil
}
