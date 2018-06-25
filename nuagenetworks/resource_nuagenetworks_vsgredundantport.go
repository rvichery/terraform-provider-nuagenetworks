package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVsgRedundantPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceVsgRedundantPortCreate,
		Read:   resourceVsgRedundantPortRead,
		Update: resourceVsgRedundantPortUpdate,
		Delete: resourceVsgRedundantPortDelete,
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
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
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
			"physical_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_peer1_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_peer2_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_type": {
				Type:     schema.TypeString,
				Required: true,
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
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_redundancy_group": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVsgRedundantPortCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VsgRedundantPort object
	o := &vspk.VsgRedundantPort{
		Name:         d.Get("name").(string),
		PhysicalName: d.Get("physical_name").(string),
		PortType:     d.Get("port_type").(string),
	}
	if attr, ok := d.GetOk("vlan_range"); ok {
		o.VLANRange = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("port_peer1_id"); ok {
		o.PortPeer1ID = attr.(string)
	}
	if attr, ok := d.GetOk("port_peer2_id"); ok {
		o.PortPeer2ID = attr.(string)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		UseUserMnemonic := attr.(bool)
		o.UseUserMnemonic = &UseUserMnemonic
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.RedundancyGroup{ID: d.Get("parent_redundancy_group").(string)}
	err := parent.CreateVsgRedundantPort(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceVsgRedundantPortRead(d, m)
}

func resourceVsgRedundantPortRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VsgRedundantPort{
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
	d.Set("permitted_action", o.PermittedAction)
	d.Set("description", o.Description)
	d.Set("physical_name", o.PhysicalName)
	d.Set("entity_scope", o.EntityScope)
	d.Set("port_peer1_id", o.PortPeer1ID)
	d.Set("port_peer2_id", o.PortPeer2ID)
	d.Set("port_type", o.PortType)
	d.Set("use_user_mnemonic", o.UseUserMnemonic)
	d.Set("user_mnemonic", o.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("status", o.Status)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVsgRedundantPortUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VsgRedundantPort{
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
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("port_peer1_id"); ok {
		o.PortPeer1ID = attr.(string)
	}
	if attr, ok := d.GetOk("port_peer2_id"); ok {
		o.PortPeer2ID = attr.(string)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		UseUserMnemonic := attr.(bool)
		o.UseUserMnemonic = &UseUserMnemonic
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVsgRedundantPortDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VsgRedundantPort{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
