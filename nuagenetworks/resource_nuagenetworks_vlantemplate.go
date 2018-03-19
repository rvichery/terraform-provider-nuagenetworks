package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceVLANTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVLANTemplateCreate,
		Read:   resourceVLANTemplateRead,
		Update: resourceVLANTemplateUpdate,
		Delete: resourceVLANTemplateDelete,
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
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_uplink": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"associated_connection_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ingress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_uplink_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_vsc_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"duc_vlan": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_ns_port_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_port_template"},
			},
			"parent_port_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ns_port_template"},
			},
		},
	}
}

func resourceVLANTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VLANTemplate object
	o := &vspk.VLANTemplate{
		Value: d.Get("value").(int),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_qos_policy_id"); ok {
		o.AssociatedIngressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
		o.AssociatedUplinkConnectionID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
		o.AssociatedVSCProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("duc_vlan"); ok {
		o.DucVlan = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	if attr, ok := d.GetOk("parent_ns_port_template"); ok {
		parent := &vspk.NSPortTemplate{ID: attr.(string)}
		err := parent.CreateVLANTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_port_template"); ok {
		parent := &vspk.PortTemplate{ID: attr.(string)}
		err := parent.CreateVLANTemplate(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceVLANTemplateRead(d, m)
}

func resourceVLANTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLANTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("value", o.Value)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("is_uplink", o.IsUplink)
	d.Set("associated_connection_type", o.AssociatedConnectionType)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("associated_ingress_qos_policy_id", o.AssociatedIngressQOSPolicyID)
	d.Set("associated_uplink_connection_id", o.AssociatedUplinkConnectionID)
	d.Set("associated_vsc_profile_id", o.AssociatedVSCProfileID)
	d.Set("duc_vlan", o.DucVlan)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVLANTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLANTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Value = d.Get("value").(int)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ingress_qos_policy_id"); ok {
		o.AssociatedIngressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
		o.AssociatedUplinkConnectionID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
		o.AssociatedVSCProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("duc_vlan"); ok {
		o.DucVlan = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVLANTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VLANTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
