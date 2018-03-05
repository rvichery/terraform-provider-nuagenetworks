package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceOSPFInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceOSPFInterfaceCreate,
		Read:   resourceOSPFInterfaceRead,
		Update: resourceOSPFInterfaceUpdate,
		Delete: resourceOSPFInterfaceDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"passive_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UP",
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  40,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message_digest_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "BROADCAST",
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"associated_subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authentication_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_ospf_area": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOSPFInterfaceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize OSPFInterface object
	o := &vspk.OSPFInterface{
		Name:               d.Get("name").(string),
		AssociatedSubnetID: d.Get("associated_subnet_id").(string),
	}
	if attr, ok := d.GetOk("passive_enabled"); ok {
		o.PassiveEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("admin_state"); ok {
		o.AdminState = attr.(string)
	}
	if attr, ok := d.GetOk("dead_interval"); ok {
		o.DeadInterval = attr.(int)
	}
	if attr, ok := d.GetOk("hello_interval"); ok {
		o.HelloInterval = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("message_digest_keys"); ok {
		o.MessageDigestKeys = attr.([]interface{})
	}
	if attr, ok := d.GetOk("metric"); ok {
		o.Metric = attr.(int)
	}
	if attr, ok := d.GetOk("interface_type"); ok {
		o.InterfaceType = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("authentication_key"); ok {
		o.AuthenticationKey = attr.(string)
	}
	if attr, ok := d.GetOk("authentication_type"); ok {
		o.AuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.OSPFArea{ID: d.Get("parent_ospf_area").(string)}
	err := parent.CreateOSPFInterface(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceOSPFInterfaceRead(d, m)
}

func resourceOSPFInterfaceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("passive_enabled", o.PassiveEnabled)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("admin_state", o.AdminState)
	d.Set("dead_interval", o.DeadInterval)
	d.Set("hello_interval", o.HelloInterval)
	d.Set("description", o.Description)
	d.Set("message_digest_keys", o.MessageDigestKeys)
	d.Set("metric", o.Metric)
	d.Set("interface_type", o.InterfaceType)
	d.Set("entity_scope", o.EntityScope)
	d.Set("priority", o.Priority)
	d.Set("associated_subnet_id", o.AssociatedSubnetID)
	d.Set("mtu", o.Mtu)
	d.Set("authentication_key", o.AuthenticationKey)
	d.Set("authentication_type", o.AuthenticationType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceOSPFInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.AssociatedSubnetID = d.Get("associated_subnet_id").(string)

	if attr, ok := d.GetOk("passive_enabled"); ok {
		o.PassiveEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("admin_state"); ok {
		o.AdminState = attr.(string)
	}
	if attr, ok := d.GetOk("dead_interval"); ok {
		o.DeadInterval = attr.(int)
	}
	if attr, ok := d.GetOk("hello_interval"); ok {
		o.HelloInterval = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("message_digest_keys"); ok {
		o.MessageDigestKeys = attr.([]interface{})
	}
	if attr, ok := d.GetOk("metric"); ok {
		o.Metric = attr.(int)
	}
	if attr, ok := d.GetOk("interface_type"); ok {
		o.InterfaceType = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("authentication_key"); ok {
		o.AuthenticationKey = attr.(string)
	}
	if attr, ok := d.GetOk("authentication_type"); ok {
		o.AuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceOSPFInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInterface{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
