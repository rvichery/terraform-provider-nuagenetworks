package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVLAN() *schema.Resource {
	return &schema.Resource{
		Create: resourceVLANCreate,
		Read:   resourceVLANRead,
		Update: resourceVLANUpdate,
		Delete: resourceVLANDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeInt,
				Required: true,
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
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
				Computed: true,
			},
			"user_mnemonic": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_bgp_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_connection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_ingress_overlay_qo_s_policer_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_ingress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_ingress_underlay_qo_s_policer_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_uplink_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_vsc_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duc_vlan": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceVLANCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VLAN object
	Value := d.Get("value").(int)
	o := &vspk.VLAN{
		Value: &Value,
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		UseUserMnemonic := attr.(bool)
		o.UseUserMnemonic = &UseUserMnemonic
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
		o.AssociatedBGPProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_overlay_qo_s_policer_id"); ok {
		o.AssociatedIngressOverlayQoSPolicerID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_qos_policy_id"); ok {
		o.AssociatedIngressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_underlay_qo_s_policer_id"); ok {
		o.AssociatedIngressUnderlayQoSPolicerID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
		o.AssociatedUplinkConnectionID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
		o.AssociatedVSCProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("duc_vlan"); ok {
		DucVlan := attr.(bool)
		o.DucVlan = &DucVlan
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		err := parent.CreateVLAN(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_redundant_port"); ok {
		parent := &vspk.RedundantPort{ID: attr.(string)}
		err := parent.CreateVLAN(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		err := parent.CreateVLAN(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		err := parent.CreateVLAN(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("patnatpools"); ok {
		o.AssignPATNATPools(attr.(vspk.PATNATPoolsList))
	}
	return resourceVLANRead(d, m)
}

func resourceVLANRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLAN{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("value", o.Value)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("gateway_id", o.GatewayID)
	d.Set("readonly", o.Readonly)
	d.Set("template_id", o.TemplateID)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("description", o.Description)
	d.Set("restricted", o.Restricted)
	d.Set("entity_scope", o.EntityScope)
	d.Set("vport_id", o.VportID)
	d.Set("is_uplink", o.IsUplink)
	d.Set("use_user_mnemonic", o.UseUserMnemonic)
	d.Set("user_mnemonic", o.UserMnemonic)
	d.Set("associated_bgp_profile_id", o.AssociatedBGPProfileID)
	d.Set("associated_connection_type", o.AssociatedConnectionType)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("associated_ingress_overlay_qo_s_policer_id", o.AssociatedIngressOverlayQoSPolicerID)
	d.Set("associated_ingress_qos_policy_id", o.AssociatedIngressQOSPolicyID)
	d.Set("associated_ingress_underlay_qo_s_policer_id", o.AssociatedIngressUnderlayQoSPolicerID)
	d.Set("associated_uplink_connection_id", o.AssociatedUplinkConnectionID)
	d.Set("associated_vsc_profile_id", o.AssociatedVSCProfileID)
	d.Set("status", o.Status)
	d.Set("duc_vlan", o.DucVlan)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVLANUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLAN{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	Value := d.Get("value").(int)
	o.Value = &Value

	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		UseUserMnemonic := attr.(bool)
		o.UseUserMnemonic = &UseUserMnemonic
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
		o.AssociatedBGPProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_overlay_qo_s_policer_id"); ok {
		o.AssociatedIngressOverlayQoSPolicerID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_qos_policy_id"); ok {
		o.AssociatedIngressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_underlay_qo_s_policer_id"); ok {
		o.AssociatedIngressUnderlayQoSPolicerID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
		o.AssociatedUplinkConnectionID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
		o.AssociatedVSCProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("duc_vlan"); ok {
		DucVlan := attr.(bool)
		o.DucVlan = &DucVlan
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVLANDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLAN{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
