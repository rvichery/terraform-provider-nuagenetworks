package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceWANService() *schema.Resource {
	return &schema.Resource{
		Create: resourceWANServiceCreate,
		Read:   resourceWANServiceRead,
		Update: resourceWANServiceUpdate,
		Delete: resourceWANServiceDelete,

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
			"wan_service_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"irb_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permitted_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vn_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enterprise_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"orphan": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_user_mnemonic": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_mnemonic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_vpn_connect_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_route_target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_redundancy_group": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_gateway"},
			},
			"parent_gateway": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group"},
			},
		},
	}
}

func resourceWANServiceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize WANService object
	o := &vspk.WANService{
		Name:        d.Get("name").(string),
		ServiceType: d.Get("service_type").(string),
	}
	if attr, ok := d.GetOk("wan_service_identifier"); ok {
		o.WANServiceIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("irb_enabled"); ok {
		o.IRBEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("service_policy"); ok {
		o.ServicePolicy = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("vn_id"); ok {
		o.VnId = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_name"); ok {
		o.EnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("config_type"); ok {
		o.ConfigType = attr.(string)
	}
	if attr, ok := d.GetOk("orphan"); ok {
		o.Orphan = attr.(bool)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		o.UseUserMnemonic = attr.(bool)
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_domain_id"); ok {
		o.AssociatedDomainID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vpn_connect_id"); ok {
		o.AssociatedVPNConnectID = attr.(string)
	}
	if attr, ok := d.GetOk("tunnel_type"); ok {
		o.TunnelType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("external_route_target"); ok {
		o.ExternalRouteTarget = attr.(string)
	}
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		err := parent.CreateWANService(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		err := parent.CreateWANService(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceWANServiceRead(d, m)
}

func resourceWANServiceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.WANService{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("wan_service_identifier", o.WANServiceIdentifier)
	d.Set("irb_enabled", o.IRBEnabled)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("service_policy", o.ServicePolicy)
	d.Set("service_type", o.ServiceType)
	d.Set("description", o.Description)
	d.Set("vn_id", o.VnId)
	d.Set("enterprise_name", o.EnterpriseName)
	d.Set("entity_scope", o.EntityScope)
	d.Set("domain_name", o.DomainName)
	d.Set("config_type", o.ConfigType)
	d.Set("orphan", o.Orphan)
	d.Set("use_user_mnemonic", o.UseUserMnemonic)
	d.Set("user_mnemonic", o.UserMnemonic)
	d.Set("associated_domain_id", o.AssociatedDomainID)
	d.Set("associated_vpn_connect_id", o.AssociatedVPNConnectID)
	d.Set("tunnel_type", o.TunnelType)
	d.Set("external_id", o.ExternalID)
	d.Set("external_route_target", o.ExternalRouteTarget)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceWANServiceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.WANService{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.ServiceType = d.Get("service_type").(string)

	if attr, ok := d.GetOk("wan_service_identifier"); ok {
		o.WANServiceIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("irb_enabled"); ok {
		o.IRBEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("service_policy"); ok {
		o.ServicePolicy = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("vn_id"); ok {
		o.VnId = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_name"); ok {
		o.EnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("config_type"); ok {
		o.ConfigType = attr.(string)
	}
	if attr, ok := d.GetOk("orphan"); ok {
		o.Orphan = attr.(bool)
	}
	if attr, ok := d.GetOk("use_user_mnemonic"); ok {
		o.UseUserMnemonic = attr.(bool)
	}
	if attr, ok := d.GetOk("user_mnemonic"); ok {
		o.UserMnemonic = attr.(string)
	}
	if attr, ok := d.GetOk("associated_domain_id"); ok {
		o.AssociatedDomainID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vpn_connect_id"); ok {
		o.AssociatedVPNConnectID = attr.(string)
	}
	if attr, ok := d.GetOk("tunnel_type"); ok {
		o.TunnelType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("external_route_target"); ok {
		o.ExternalRouteTarget = attr.(string)
	}

	o.Save()

	return nil
}

func resourceWANServiceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.WANService{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
