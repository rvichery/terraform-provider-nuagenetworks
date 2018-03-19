package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVLAN() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVLANRead,
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
			"value": {
				Type:     schema.TypeInt,
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
			"readonly": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"template_id": {
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
			"restricted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vport_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_uplink": {
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
			"associated_bgp_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_connection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ingress_overlay_qo_s_policer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ingress_qos_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ingress_underlay_qo_s_policer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_uplink_connection_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_vsc_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"duc_vlan": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vsg_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundant_port", "parent_port", "parent_ns_port"},
			},
			"parent_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vsg_redundant_port", "parent_port", "parent_ns_port"},
			},
			"parent_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vsg_redundant_port", "parent_redundant_port", "parent_ns_port"},
			},
			"parent_ns_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vsg_redundant_port", "parent_redundant_port", "parent_port"},
			},
		},
	}
}

func dataSourceVLANRead(d *schema.ResourceData, m interface{}) error {
	filteredVLANs := vspk.VLANsList{}
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
	if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		filteredVLANs, err = parent.VLANs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_redundant_port"); ok {
		parent := &vspk.RedundantPort{ID: attr.(string)}
		filteredVLANs, err = parent.VLANs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		filteredVLANs, err = parent.VLANs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		filteredVLANs, err = parent.VLANs(fetchFilter)
		if err != nil {
			return err
		}
	}

	VLAN := &vspk.VLAN{}

	if len(filteredVLANs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVLANs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VLAN = filteredVLANs[0]

	d.Set("value", VLAN.Value)
	d.Set("last_updated_by", VLAN.LastUpdatedBy)
	d.Set("gateway_id", VLAN.GatewayID)
	d.Set("readonly", VLAN.Readonly)
	d.Set("template_id", VLAN.TemplateID)
	d.Set("permitted_action", VLAN.PermittedAction)
	d.Set("description", VLAN.Description)
	d.Set("restricted", VLAN.Restricted)
	d.Set("entity_scope", VLAN.EntityScope)
	d.Set("vport_id", VLAN.VportID)
	d.Set("is_uplink", VLAN.IsUplink)
	d.Set("use_user_mnemonic", VLAN.UseUserMnemonic)
	d.Set("user_mnemonic", VLAN.UserMnemonic)
	d.Set("associated_bgp_profile_id", VLAN.AssociatedBGPProfileID)
	d.Set("associated_connection_type", VLAN.AssociatedConnectionType)
	d.Set("associated_egress_qos_policy_id", VLAN.AssociatedEgressQOSPolicyID)
	d.Set("associated_ingress_overlay_qo_s_policer_id", VLAN.AssociatedIngressOverlayQoSPolicerID)
	d.Set("associated_ingress_qos_policy_id", VLAN.AssociatedIngressQOSPolicyID)
	d.Set("associated_ingress_underlay_qo_s_policer_id", VLAN.AssociatedIngressUnderlayQoSPolicerID)
	d.Set("associated_uplink_connection_id", VLAN.AssociatedUplinkConnectionID)
	d.Set("associated_vsc_profile_id", VLAN.AssociatedVSCProfileID)
	d.Set("status", VLAN.Status)
	d.Set("duc_vlan", VLAN.DucVlan)
	d.Set("external_id", VLAN.ExternalID)
	d.Set("type", VLAN.Type)

	d.Set("id", VLAN.Identifier())
	d.Set("parent_id", VLAN.ParentID)
	d.Set("parent_type", VLAN.ParentType)
	d.Set("owner", VLAN.Owner)

	d.SetId(VLAN.Identifier())

	return nil
}
