package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourcePATNATPool() *schema.Resource {
	return &schema.Resource{
		Create: resourcePATNATPoolCreate,
		Read:   resourcePATNATPoolRead,
		Update: resourcePATNATPoolUpdate,
		Delete: resourcePATNATPoolDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_patip": {
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
			"end_address_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_source_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_gateway_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_address_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_source_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_source_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourcePATNATPoolCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PATNATPool object
	o := &vspk.PATNATPool{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("address_range"); ok {
		o.AddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("default_patip"); ok {
		o.DefaultPATIP = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("end_address_range"); ok {
		o.EndAddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("end_source_address"); ok {
		o.EndSourceAddress = attr.(string)
	}
	if attr, ok := d.GetOk("associated_gateway_id"); ok {
		o.AssociatedGatewayId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_gateway_type"); ok {
		o.AssociatedGatewayType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_subnet_id"); ok {
		o.AssociatedSubnetId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vlan_id"); ok {
		o.AssociatedVlanId = attr.(string)
	}
	if attr, ok := d.GetOk("start_address_range"); ok {
		o.StartAddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("start_source_address"); ok {
		o.StartSourceAddress = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_source_enabled"); ok {
		o.DynamicSourceEnabled = attr.(bool)
	}
	parent := m.(*vspk.Me)
	err := parent.CreatePATNATPool(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourcePATNATPoolRead(d, m)
}

func resourcePATNATPoolRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATNATPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("address_range", o.AddressRange)
	d.Set("default_patip", o.DefaultPATIP)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("description", o.Description)
	d.Set("end_address_range", o.EndAddressRange)
	d.Set("end_source_address", o.EndSourceAddress)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_gateway_id", o.AssociatedGatewayId)
	d.Set("associated_gateway_type", o.AssociatedGatewayType)
	d.Set("associated_subnet_id", o.AssociatedSubnetId)
	d.Set("associated_vlan_id", o.AssociatedVlanId)
	d.Set("start_address_range", o.StartAddressRange)
	d.Set("start_source_address", o.StartSourceAddress)
	d.Set("external_id", o.ExternalID)
	d.Set("dynamic_source_enabled", o.DynamicSourceEnabled)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePATNATPoolUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATNATPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("address_range"); ok {
		o.AddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("default_patip"); ok {
		o.DefaultPATIP = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("end_address_range"); ok {
		o.EndAddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("end_source_address"); ok {
		o.EndSourceAddress = attr.(string)
	}
	if attr, ok := d.GetOk("associated_gateway_id"); ok {
		o.AssociatedGatewayId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_gateway_type"); ok {
		o.AssociatedGatewayType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_subnet_id"); ok {
		o.AssociatedSubnetId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vlan_id"); ok {
		o.AssociatedVlanId = attr.(string)
	}
	if attr, ok := d.GetOk("start_address_range"); ok {
		o.StartAddressRange = attr.(string)
	}
	if attr, ok := d.GetOk("start_source_address"); ok {
		o.StartSourceAddress = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_source_enabled"); ok {
		o.DynamicSourceEnabled = attr.(bool)
	}

	o.Save()

	return nil
}

func resourcePATNATPoolDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATNATPool{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
