package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceIKEGatewayConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceIKEGatewayConnectionCreate,
		Read:   resourceIKEGatewayConnectionRead,
		Update: resourceIKEGatewayConnectionUpdate,
		Delete: resourceIKEGatewayConnectionDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsg_identifier_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ID_KEY_ID",
			},
			"nsg_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mark": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"allow_any_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"unencrypted_psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"associated_ike_authentication_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_authentication_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_encryption_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_gateway_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_vlanid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIKEGatewayConnectionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IKEGatewayConnection object
	o := &vspk.IKEGatewayConnection{}
	if attr, ok := d.GetOk("nsg_identifier"); ok {
		o.NSGIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_identifier_type"); ok {
		o.NSGIdentifierType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_role"); ok {
		o.NSGRole = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("mark"); ok {
		o.Mark = attr.(int)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		o.Sequence = attr.(int)
	}
	if attr, ok := d.GetOk("allow_any_subnet"); ok {
		o.AllowAnySubnet = attr.(bool)
	}
	if attr, ok := d.GetOk("unencrypted_psk"); ok {
		o.UnencryptedPSK = attr.(string)
	}
	if attr, ok := d.GetOk("port_vlan_name"); ok {
		o.PortVLANName = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_id"); ok {
		o.AssociatedIKEAuthenticationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_type"); ok {
		o.AssociatedIKEAuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_encryption_profile_id"); ok {
		o.AssociatedIKEEncryptionProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_profile_id"); ok {
		o.AssociatedIKEGatewayProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vlanid"); ok {
		o.AssociatedVLANID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.VLAN{ID: d.Get("parent_vlan").(string)}
	err := parent.CreateIKEGatewayConnection(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("performancemonitors"); ok {
		o.AssignPerformanceMonitors(attr.(vspk.PerformanceMonitorsList))
	}
	if attr, ok := d.GetOk("subnets"); ok {
		o.AssignSubnets(attr.(vspk.SubnetsList))
	}
	return resourceIKEGatewayConnectionRead(d, m)
}

func resourceIKEGatewayConnectionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("nsg_identifier", o.NSGIdentifier)
	d.Set("nsg_identifier_type", o.NSGIdentifierType)
	d.Set("nsg_role", o.NSGRole)
	d.Set("name", o.Name)
	d.Set("mark", o.Mark)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("sequence", o.Sequence)
	d.Set("allow_any_subnet", o.AllowAnySubnet)
	d.Set("unencrypted_psk", o.UnencryptedPSK)
	d.Set("entity_scope", o.EntityScope)
	d.Set("port_vlan_name", o.PortVLANName)
	d.Set("priority", o.Priority)
	d.Set("associated_ike_authentication_id", o.AssociatedIKEAuthenticationID)
	d.Set("associated_ike_authentication_type", o.AssociatedIKEAuthenticationType)
	d.Set("associated_ike_encryption_profile_id", o.AssociatedIKEEncryptionProfileID)
	d.Set("associated_ike_gateway_profile_id", o.AssociatedIKEGatewayProfileID)
	d.Set("associated_vlanid", o.AssociatedVLANID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIKEGatewayConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("nsg_identifier"); ok {
		o.NSGIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_identifier_type"); ok {
		o.NSGIdentifierType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_role"); ok {
		o.NSGRole = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("mark"); ok {
		o.Mark = attr.(int)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		o.Sequence = attr.(int)
	}
	if attr, ok := d.GetOk("allow_any_subnet"); ok {
		o.AllowAnySubnet = attr.(bool)
	}
	if attr, ok := d.GetOk("unencrypted_psk"); ok {
		o.UnencryptedPSK = attr.(string)
	}
	if attr, ok := d.GetOk("port_vlan_name"); ok {
		o.PortVLANName = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_id"); ok {
		o.AssociatedIKEAuthenticationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_type"); ok {
		o.AssociatedIKEAuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_encryption_profile_id"); ok {
		o.AssociatedIKEEncryptionProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_profile_id"); ok {
		o.AssociatedIKEGatewayProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vlanid"); ok {
		o.AssociatedVLANID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIKEGatewayConnectionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayConnection{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
