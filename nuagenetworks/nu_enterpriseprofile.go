package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceEnterpriseProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnterpriseProfileCreate,
		Read:   resourceEnterpriseProfileRead,
		Update: resourceEnterpriseProfileUpdate,
		Delete: resourceEnterpriseProfileDelete,

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
			"bgp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dhcp_lease_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vnf_management_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"receive_multi_cast_list_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_multi_cast_list_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_advanced_qos_configuration": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_gateway_management": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_trusted_forwarding_class": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allowed_forwarding_classes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floating_ips_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enable_application_performance_management": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"encryption_management_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEnterpriseProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EnterpriseProfile object
	o := &vspk.EnterpriseProfile{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("bgp_enabled"); ok {
		o.BGPEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
		o.DHCPLeaseInterval = attr.(int)
	}
	if attr, ok := d.GetOk("vnf_management_enabled"); ok {
		o.VNFManagementEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("receive_multi_cast_list_id"); ok {
		o.ReceiveMultiCastListID = attr.(string)
	}
	if attr, ok := d.GetOk("send_multi_cast_list_id"); ok {
		o.SendMultiCastListID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("allow_advanced_qos_configuration"); ok {
		o.AllowAdvancedQOSConfiguration = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_gateway_management"); ok {
		o.AllowGatewayManagement = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
		o.AllowTrustedForwardingClass = attr.(bool)
	}
	if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
		o.AllowedForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("floating_ips_quota"); ok {
		o.FloatingIPsQuota = attr.(int)
	}
	if attr, ok := d.GetOk("enable_application_performance_management"); ok {
		o.EnableApplicationPerformanceManagement = attr.(bool)
	}
	if attr, ok := d.GetOk("encryption_management_mode"); ok {
		o.EncryptionManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateEnterpriseProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceEnterpriseProfileRead(d, m)
}

func resourceEnterpriseProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("bgp_enabled", o.BGPEnabled)
	d.Set("dhcp_lease_interval", o.DHCPLeaseInterval)
	d.Set("vnf_management_enabled", o.VNFManagementEnabled)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("receive_multi_cast_list_id", o.ReceiveMultiCastListID)
	d.Set("send_multi_cast_list_id", o.SendMultiCastListID)
	d.Set("description", o.Description)
	d.Set("allow_advanced_qos_configuration", o.AllowAdvancedQOSConfiguration)
	d.Set("allow_gateway_management", o.AllowGatewayManagement)
	d.Set("allow_trusted_forwarding_class", o.AllowTrustedForwardingClass)
	d.Set("allowed_forwarding_classes", o.AllowedForwardingClasses)
	d.Set("floating_ips_quota", o.FloatingIPsQuota)
	d.Set("enable_application_performance_management", o.EnableApplicationPerformanceManagement)
	d.Set("encryption_management_mode", o.EncryptionManagementMode)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceEnterpriseProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("bgp_enabled"); ok {
		o.BGPEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
		o.DHCPLeaseInterval = attr.(int)
	}
	if attr, ok := d.GetOk("vnf_management_enabled"); ok {
		o.VNFManagementEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("receive_multi_cast_list_id"); ok {
		o.ReceiveMultiCastListID = attr.(string)
	}
	if attr, ok := d.GetOk("send_multi_cast_list_id"); ok {
		o.SendMultiCastListID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("allow_advanced_qos_configuration"); ok {
		o.AllowAdvancedQOSConfiguration = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_gateway_management"); ok {
		o.AllowGatewayManagement = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
		o.AllowTrustedForwardingClass = attr.(bool)
	}
	if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
		o.AllowedForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("floating_ips_quota"); ok {
		o.FloatingIPsQuota = attr.(int)
	}
	if attr, ok := d.GetOk("enable_application_performance_management"); ok {
		o.EnableApplicationPerformanceManagement = attr.(bool)
	}
	if attr, ok := d.GetOk("encryption_management_mode"); ok {
		o.EncryptionManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEnterpriseProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
