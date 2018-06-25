package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceEnterpriseProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnterpriseProfileCreate,
		Read:   resourceEnterpriseProfileRead,
		Update: resourceEnterpriseProfileUpdate,
		Delete: resourceEnterpriseProfileDelete,
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
			"bgp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dhcp_lease_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vnf_management_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"receive_multi_cast_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_multi_cast_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_advanced_qos_configuration": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_gateway_management": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_trusted_forwarding_class": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allowed_forwarding_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floating_ips_quota": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enable_application_performance_management": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"encryption_management_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		BGPEnabled := attr.(bool)
		o.BGPEnabled = &BGPEnabled
	}
	if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
		DHCPLeaseInterval := attr.(int)
		o.DHCPLeaseInterval = &DHCPLeaseInterval
	}
	if attr, ok := d.GetOk("vnf_management_enabled"); ok {
		VNFManagementEnabled := attr.(bool)
		o.VNFManagementEnabled = &VNFManagementEnabled
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
		AllowAdvancedQOSConfiguration := attr.(bool)
		o.AllowAdvancedQOSConfiguration = &AllowAdvancedQOSConfiguration
	}
	if attr, ok := d.GetOk("allow_gateway_management"); ok {
		AllowGatewayManagement := attr.(bool)
		o.AllowGatewayManagement = &AllowGatewayManagement
	}
	if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
		AllowTrustedForwardingClass := attr.(bool)
		o.AllowTrustedForwardingClass = &AllowTrustedForwardingClass
	}
	if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
		o.AllowedForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("floating_ips_quota"); ok {
		FloatingIPsQuota := attr.(int)
		o.FloatingIPsQuota = &FloatingIPsQuota
	}
	if attr, ok := d.GetOk("enable_application_performance_management"); ok {
		EnableApplicationPerformanceManagement := attr.(bool)
		o.EnableApplicationPerformanceManagement = &EnableApplicationPerformanceManagement
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
		BGPEnabled := attr.(bool)
		o.BGPEnabled = &BGPEnabled
	}
	if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
		DHCPLeaseInterval := attr.(int)
		o.DHCPLeaseInterval = &DHCPLeaseInterval
	}
	if attr, ok := d.GetOk("vnf_management_enabled"); ok {
		VNFManagementEnabled := attr.(bool)
		o.VNFManagementEnabled = &VNFManagementEnabled
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
		AllowAdvancedQOSConfiguration := attr.(bool)
		o.AllowAdvancedQOSConfiguration = &AllowAdvancedQOSConfiguration
	}
	if attr, ok := d.GetOk("allow_gateway_management"); ok {
		AllowGatewayManagement := attr.(bool)
		o.AllowGatewayManagement = &AllowGatewayManagement
	}
	if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
		AllowTrustedForwardingClass := attr.(bool)
		o.AllowTrustedForwardingClass = &AllowTrustedForwardingClass
	}
	if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
		o.AllowedForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("floating_ips_quota"); ok {
		FloatingIPsQuota := attr.(int)
		o.FloatingIPsQuota = &FloatingIPsQuota
	}
	if attr, ok := d.GetOk("enable_application_performance_management"); ok {
		EnableApplicationPerformanceManagement := attr.(bool)
		o.EnableApplicationPerformanceManagement = &EnableApplicationPerformanceManagement
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
