package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceRedundantPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRedundantPortRead,
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
			"vlan_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
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
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"physical_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"infrastructure_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_peer1_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_peer2_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"speed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_untagged_heartbeat_vlan": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_user_mnemonic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_mnemonic": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ns_redundant_gateway_group": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRedundantPortRead(d *schema.ResourceData, m interface{}) error {
	filteredRedundantPorts := vspk.RedundantPortsList{}
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
	parent := &vspk.NSRedundantGatewayGroup{ID: d.Get("parent_ns_redundant_gateway_group").(string)}
	filteredRedundantPorts, err = parent.RedundantPorts(fetchFilter)
	if err != nil {
		return err
	}

	RedundantPort := &vspk.RedundantPort{}

	if len(filteredRedundantPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredRedundantPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	RedundantPort = filteredRedundantPorts[0]

	d.Set("vlan_range", RedundantPort.VLANRange)
	d.Set("mtu", RedundantPort.MTU)
	d.Set("name", RedundantPort.Name)
	d.Set("last_updated_by", RedundantPort.LastUpdatedBy)
	d.Set("permitted_action", RedundantPort.PermittedAction)
	d.Set("description", RedundantPort.Description)
	d.Set("physical_name", RedundantPort.PhysicalName)
	d.Set("infrastructure_profile_id", RedundantPort.InfrastructureProfileID)
	d.Set("entity_scope", RedundantPort.EntityScope)
	d.Set("port_peer1_id", RedundantPort.PortPeer1ID)
	d.Set("port_peer2_id", RedundantPort.PortPeer2ID)
	d.Set("port_type", RedundantPort.PortType)
	d.Set("speed", RedundantPort.Speed)
	d.Set("use_untagged_heartbeat_vlan", RedundantPort.UseUntaggedHeartbeatVlan)
	d.Set("use_user_mnemonic", RedundantPort.UseUserMnemonic)
	d.Set("user_mnemonic", RedundantPort.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", RedundantPort.AssociatedEgressQOSPolicyID)
	d.Set("status", RedundantPort.Status)
	d.Set("external_id", RedundantPort.ExternalID)

	d.Set("id", RedundantPort.Identifier())
	d.Set("parent_id", RedundantPort.ParentID)
	d.Set("parent_type", RedundantPort.ParentType)
	d.Set("owner", RedundantPort.Owner)

	d.SetId(RedundantPort.Identifier())

	return nil
}
