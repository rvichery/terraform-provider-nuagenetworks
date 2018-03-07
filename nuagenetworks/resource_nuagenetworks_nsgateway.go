package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceNSGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGatewayCreate,
		Read:   resourceNSGatewayRead,
		Update: resourceNSGatewayUpdate,
		Delete: resourceNSGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_traversal_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcpmss_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tcp_maximum_segment_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1330,
			},
			"bios_release_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bios_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sku": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tpm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNKNOWN",
			},
			"tpm_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsg_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_configuration_reload_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datapath_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patches": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redundancy_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"derived_ssh_service_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"personality": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"libraries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inherited_ssh_service_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration_reload_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNKNOWN",
			},
			"configuration_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNKNOWN",
			},
			"control_traffic_cos_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  7,
			},
			"control_traffic_dscp_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  56,
			},
			"bootstrap_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bootstrap_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_gateway_security_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_gateway_security_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_nsg_info_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_nsg_upgrade_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_disc_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNSGatewayCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NSGateway object
	o := &vspk.NSGateway{
		Name:       d.Get("name").(string),
		TemplateID: d.Get("template_id").(string),
	}
	if attr, ok := d.GetOk("mac_address"); ok {
		o.MACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("nat_traversal_enabled"); ok {
		o.NATTraversalEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("tcpmss_enabled"); ok {
		o.TCPMSSEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("tcp_maximum_segment_size"); ok {
		o.TCPMaximumSegmentSize = attr.(int)
	}
	if attr, ok := d.GetOk("bios_release_date"); ok {
		o.BIOSReleaseDate = attr.(string)
	}
	if attr, ok := d.GetOk("bios_version"); ok {
		o.BIOSVersion = attr.(string)
	}
	if attr, ok := d.GetOk("sku"); ok {
		o.SKU = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_status"); ok {
		o.TPMStatus = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_type"); ok {
		o.CPUType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_version"); ok {
		o.NSGVersion = attr.(string)
	}
	if attr, ok := d.GetOk("ssh_service"); ok {
		o.SSHService = attr.(string)
	}
	if attr, ok := d.GetOk("uuid"); ok {
		o.UUID = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("redundancy_group_id"); ok {
		o.RedundancyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("pending"); ok {
		o.Pending = attr.(bool)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("derived_ssh_service_state"); ok {
		o.DerivedSSHServiceState = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("libraries"); ok {
		o.Libraries = attr.(string)
	}
	if attr, ok := d.GetOk("inherited_ssh_service_state"); ok {
		o.InheritedSSHServiceState = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("configuration_reload_state"); ok {
		o.ConfigurationReloadState = attr.(string)
	}
	if attr, ok := d.GetOk("configuration_status"); ok {
		o.ConfigurationStatus = attr.(string)
	}
	if attr, ok := d.GetOk("control_traffic_cos_value"); ok {
		o.ControlTrafficCOSValue = attr.(int)
	}
	if attr, ok := d.GetOk("control_traffic_dscp_value"); ok {
		o.ControlTrafficDSCPValue = attr.(int)
	}
	if attr, ok := d.GetOk("bootstrap_id"); ok {
		o.BootstrapID = attr.(string)
	}
	if attr, ok := d.GetOk("bootstrap_status"); ok {
		o.BootstrapStatus = attr.(string)
	}
	if attr, ok := d.GetOk("operation_mode"); ok {
		o.OperationMode = attr.(string)
	}
	if attr, ok := d.GetOk("operation_status"); ok {
		o.OperationStatus = attr.(string)
	}
	if attr, ok := d.GetOk("product_name"); ok {
		o.ProductName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_nsg_info_id"); ok {
		o.AssociatedNSGInfoID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_nsg_upgrade_profile_id"); ok {
		o.AssociatedNSGUpgradeProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("system_id"); ok {
		o.SystemID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateNSGateway(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("patnatpools"); ok {
		o.AssignPATNATPools(attr.(vspk.PATNATPoolsList))
	}
	return resourceNSGatewayRead(d, m)
}

func resourceNSGatewayRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac_address", o.MACAddress)
	d.Set("nat_traversal_enabled", o.NATTraversalEnabled)
	d.Set("tcpmss_enabled", o.TCPMSSEnabled)
	d.Set("tcp_maximum_segment_size", o.TCPMaximumSegmentSize)
	d.Set("bios_release_date", o.BIOSReleaseDate)
	d.Set("bios_version", o.BIOSVersion)
	d.Set("sku", o.SKU)
	d.Set("tpm_status", o.TPMStatus)
	d.Set("tpm_version", o.TPMVersion)
	d.Set("cpu_type", o.CPUType)
	d.Set("nsg_version", o.NSGVersion)
	d.Set("ssh_service", o.SSHService)
	d.Set("uuid", o.UUID)
	d.Set("name", o.Name)
	d.Set("family", o.Family)
	d.Set("last_configuration_reload_timestamp", o.LastConfigurationReloadTimestamp)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("datapath_id", o.DatapathID)
	d.Set("patches", o.Patches)
	d.Set("redundancy_group_id", o.RedundancyGroupID)
	d.Set("template_id", o.TemplateID)
	d.Set("pending", o.Pending)
	d.Set("serial_number", o.SerialNumber)
	d.Set("derived_ssh_service_state", o.DerivedSSHServiceState)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("personality", o.Personality)
	d.Set("description", o.Description)
	d.Set("libraries", o.Libraries)
	d.Set("inherited_ssh_service_state", o.InheritedSSHServiceState)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("location_id", o.LocationID)
	d.Set("configuration_reload_state", o.ConfigurationReloadState)
	d.Set("configuration_status", o.ConfigurationStatus)
	d.Set("control_traffic_cos_value", o.ControlTrafficCOSValue)
	d.Set("control_traffic_dscp_value", o.ControlTrafficDSCPValue)
	d.Set("bootstrap_id", o.BootstrapID)
	d.Set("bootstrap_status", o.BootstrapStatus)
	d.Set("operation_mode", o.OperationMode)
	d.Set("operation_status", o.OperationStatus)
	d.Set("product_name", o.ProductName)
	d.Set("associated_gateway_security_id", o.AssociatedGatewaySecurityID)
	d.Set("associated_gateway_security_profile_id", o.AssociatedGatewaySecurityProfileID)
	d.Set("associated_nsg_info_id", o.AssociatedNSGInfoID)
	d.Set("associated_nsg_upgrade_profile_id", o.AssociatedNSGUpgradeProfileID)
	d.Set("auto_disc_gateway_id", o.AutoDiscGatewayID)
	d.Set("external_id", o.ExternalID)
	d.Set("system_id", o.SystemID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNSGatewayUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.TemplateID = d.Get("template_id").(string)

	if attr, ok := d.GetOk("mac_address"); ok {
		o.MACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("nat_traversal_enabled"); ok {
		o.NATTraversalEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("tcpmss_enabled"); ok {
		o.TCPMSSEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("tcp_maximum_segment_size"); ok {
		o.TCPMaximumSegmentSize = attr.(int)
	}
	if attr, ok := d.GetOk("bios_release_date"); ok {
		o.BIOSReleaseDate = attr.(string)
	}
	if attr, ok := d.GetOk("bios_version"); ok {
		o.BIOSVersion = attr.(string)
	}
	if attr, ok := d.GetOk("sku"); ok {
		o.SKU = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_status"); ok {
		o.TPMStatus = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_type"); ok {
		o.CPUType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_version"); ok {
		o.NSGVersion = attr.(string)
	}
	if attr, ok := d.GetOk("ssh_service"); ok {
		o.SSHService = attr.(string)
	}
	if attr, ok := d.GetOk("uuid"); ok {
		o.UUID = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("redundancy_group_id"); ok {
		o.RedundancyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("pending"); ok {
		o.Pending = attr.(bool)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("derived_ssh_service_state"); ok {
		o.DerivedSSHServiceState = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("libraries"); ok {
		o.Libraries = attr.(string)
	}
	if attr, ok := d.GetOk("inherited_ssh_service_state"); ok {
		o.InheritedSSHServiceState = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("configuration_reload_state"); ok {
		o.ConfigurationReloadState = attr.(string)
	}
	if attr, ok := d.GetOk("configuration_status"); ok {
		o.ConfigurationStatus = attr.(string)
	}
	if attr, ok := d.GetOk("control_traffic_cos_value"); ok {
		o.ControlTrafficCOSValue = attr.(int)
	}
	if attr, ok := d.GetOk("control_traffic_dscp_value"); ok {
		o.ControlTrafficDSCPValue = attr.(int)
	}
	if attr, ok := d.GetOk("bootstrap_id"); ok {
		o.BootstrapID = attr.(string)
	}
	if attr, ok := d.GetOk("bootstrap_status"); ok {
		o.BootstrapStatus = attr.(string)
	}
	if attr, ok := d.GetOk("operation_mode"); ok {
		o.OperationMode = attr.(string)
	}
	if attr, ok := d.GetOk("operation_status"); ok {
		o.OperationStatus = attr.(string)
	}
	if attr, ok := d.GetOk("product_name"); ok {
		o.ProductName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_nsg_info_id"); ok {
		o.AssociatedNSGInfoID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_nsg_upgrade_profile_id"); ok {
		o.AssociatedNSGUpgradeProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("system_id"); ok {
		o.SystemID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNSGatewayDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGateway{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
