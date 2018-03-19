package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceNSPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSPortCreate,
		Read:   resourceNSPortRead,
		Update: resourceNSPortUpdate,
		Delete: resourceNSPortDelete,
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
			"nat_traversal": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"vlan_range": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0-4094",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"physical_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_nat_probes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"speed": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTONEGOTIATE",
			},
			"traffic_through_ubr_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_user_mnemonic": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_mnemonic": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_redundant_port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_ns_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNSPortCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NSPort object
	o := &vspk.NSPort{
		Name:         d.Get("name").(string),
		PhysicalName: d.Get("physical_name").(string),
		PortType:     d.Get("port_type").(string),
	}
	if attr, ok := d.GetOk("nat_traversal"); ok {
		o.NATTraversal = attr.(string)
	}
	if attr, ok := d.GetOk("vlan_range"); ok {
		o.VLANRange = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enable_nat_probes"); ok {
		o.EnableNATProbes = attr.(bool)
	}
	if attr, ok := d.GetOk("speed"); ok {
		o.Speed = attr.(string)
	}
	if attr, ok := d.GetOk("traffic_through_ubr_only"); ok {
		o.TrafficThroughUBROnly = attr.(bool)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		o.UseUserMnemonic = attr.(bool)
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_redundant_port_id"); ok {
		o.AssociatedRedundantPortID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	err := parent.CreateNSPort(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceNSPortRead(d, m)
}

func resourceNSPortRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPort{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("nat_traversal", o.NATTraversal)
	d.Set("vlan_range", o.VLANRange)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("template_id", o.TemplateID)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("description", o.Description)
	d.Set("physical_name", o.PhysicalName)
	d.Set("enable_nat_probes", o.EnableNATProbes)
	d.Set("entity_scope", o.EntityScope)
	d.Set("port_type", o.PortType)
	d.Set("speed", o.Speed)
	d.Set("traffic_through_ubr_only", o.TrafficThroughUBROnly)
	d.Set("use_user_mnemonic", o.UseUserMnemonic)
	d.Set("user_mnemonic", o.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("associated_redundant_port_id", o.AssociatedRedundantPortID)
	d.Set("status", o.Status)
	d.Set("mtu", o.Mtu)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNSPortUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPort{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.PhysicalName = d.Get("physical_name").(string)
	o.PortType = d.Get("port_type").(string)

	if attr, ok := d.GetOk("nat_traversal"); ok {
		o.NATTraversal = attr.(string)
	}
	if attr, ok := d.GetOk("vlan_range"); ok {
		o.VLANRange = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enable_nat_probes"); ok {
		o.EnableNATProbes = attr.(bool)
	}
	if attr, ok := d.GetOk("speed"); ok {
		o.Speed = attr.(string)
	}
	if attr, ok := d.GetOk("traffic_through_ubr_only"); ok {
		o.TrafficThroughUBROnly = attr.(bool)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		o.UseUserMnemonic = attr.(bool)
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_redundant_port_id"); ok {
		o.AssociatedRedundantPortID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNSPortDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPort{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
