package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceSSIDConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSIDConnectionCreate,
		Read:   resourceSSIDConnectionRead,
		Update: resourceSSIDConnectionUpdate,
		Delete: resourceSSIDConnectionDelete,
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
			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_option": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ORIGINAL_REQUEST",
			},
			"redirect_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"generic_config": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"white_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"black_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interface_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vport_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_ssid": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"associated_captive_portal_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPEN",
			},
			"parent_wireless_port": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSSIDConnectionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize SSIDConnection object
	o := &vspk.SSIDConnection{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("passphrase"); ok {
		o.Passphrase = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_option"); ok {
		o.RedirectOption = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_url"); ok {
		o.RedirectURL = attr.(string)
	}
	if attr, ok := d.GetOk("generic_config"); ok {
		o.GenericConfig = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("white_list"); ok {
		o.WhiteList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("black_list"); ok {
		o.BlackList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("broadcast_ssid"); ok {
		o.BroadcastSSID = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_captive_portal_profile_id"); ok {
		o.AssociatedCaptivePortalProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("authentication_mode"); ok {
		o.AuthenticationMode = attr.(string)
	}
	parent := &vspk.WirelessPort{ID: d.Get("parent_wireless_port").(string)}
	err := parent.CreateSSIDConnection(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("captiveportalprofiles"); ok {
		o.AssignCaptivePortalProfiles(attr.(vspk.CaptivePortalProfilesList))
	}
	return resourceSSIDConnectionRead(d, m)
}

func resourceSSIDConnectionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SSIDConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("passphrase", o.Passphrase)
	d.Set("redirect_option", o.RedirectOption)
	d.Set("redirect_url", o.RedirectURL)
	d.Set("generic_config", o.GenericConfig)
	d.Set("description", o.Description)
	d.Set("white_list", o.WhiteList)
	d.Set("black_list", o.BlackList)
	d.Set("interface_name", o.InterfaceName)
	d.Set("vport_id", o.VportID)
	d.Set("broadcast_ssid", o.BroadcastSSID)
	d.Set("associated_captive_portal_profile_id", o.AssociatedCaptivePortalProfileID)
	d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
	d.Set("authentication_mode", o.AuthenticationMode)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceSSIDConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SSIDConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("passphrase"); ok {
		o.Passphrase = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_option"); ok {
		o.RedirectOption = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_url"); ok {
		o.RedirectURL = attr.(string)
	}
	if attr, ok := d.GetOk("generic_config"); ok {
		o.GenericConfig = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("white_list"); ok {
		o.WhiteList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("black_list"); ok {
		o.BlackList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("broadcast_ssid"); ok {
		o.BroadcastSSID = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_captive_portal_profile_id"); ok {
		o.AssociatedCaptivePortalProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
		o.AssociatedEgressQOSPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("authentication_mode"); ok {
		o.AuthenticationMode = attr.(string)
	}

	o.Save()

	return nil
}

func resourceSSIDConnectionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SSIDConnection{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
