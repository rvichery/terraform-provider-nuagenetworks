package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceNSPortTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSPortTemplateCreate,
		Read:   resourceNSPortTemplateRead,
		Update: resourceNSPortTemplateUpdate,
		Delete: resourceNSPortTemplateDelete,
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"physical_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"infrastructure_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"associated_egress_qos_policy_id": {
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
			"parent_ns_gateway_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNSPortTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NSPortTemplate object
	o := &vspk.NSPortTemplate{
		Name:         d.Get("name").(string),
		PhysicalName: d.Get("physical_name").(string),
		PortType:     d.Get("port_type").(string),
	}
	if attr, ok := d.GetOk("vlan_range"); ok {
		o.VLANRange = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("infrastructure_profile_id"); ok {
		o.InfrastructureProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("speed"); ok {
		o.Speed = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.NSGatewayTemplate{ID: d.Get("parent_ns_gateway_template").(string)}
	err := parent.CreateNSPortTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceNSPortTemplateRead(d, m)
}

func resourceNSPortTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPortTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("vlan_range", o.VLANRange)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("physical_name", o.PhysicalName)
	d.Set("infrastructure_profile_id", o.InfrastructureProfileID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("port_type", o.PortType)
	d.Set("speed", o.Speed)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("mtu", o.Mtu)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNSPortTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPortTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.PhysicalName = d.Get("physical_name").(string)
	o.PortType = d.Get("port_type").(string)

	if attr, ok := d.GetOk("vlan_range"); ok {
		o.VLANRange = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("infrastructure_profile_id"); ok {
		o.InfrastructureProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("speed"); ok {
		o.Speed = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
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

func resourceNSPortTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSPortTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
