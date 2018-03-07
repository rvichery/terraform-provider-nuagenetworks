package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceInfrastructureGatewayProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInfrastructureGatewayProfileCreate,
		Read:   resourceInfrastructureGatewayProfileRead,
		Update: resourceInfrastructureGatewayProfileUpdate,
		Delete: resourceInfrastructureGatewayProfileDelete,
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
			"ntp_server_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server_key_id": &schema.Schema{
				Type:     schema.TypeInt,
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
			"datapath_sync_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"dead_timer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dead_timer_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"remote_log_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
			},
			"remote_log_server_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_log_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata_upgrade_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_eviction_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2500,
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
			"controller_less_duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "P7DT0H0M",
			},
			"controller_less_forwarding_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
			},
			"controller_less_remote_duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "P3DT0H0M",
			},
			"force_immediate_system_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"open_flow_audit_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"upgrade_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DOWNLOAD_AND_UPGRADE_AT_WINDOW",
			},
			"proxy_dns_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"use_two_factor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"stats_collector_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  39090,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_sync_scheduler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0 0 * * *",
			},
		},
	}
}

func resourceInfrastructureGatewayProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize InfrastructureGatewayProfile object
	o := &vspk.InfrastructureGatewayProfile{
		Name:         d.Get("name").(string),
		ProxyDNSName: d.Get("proxy_dns_name").(string),
	}
	if attr, ok := d.GetOk("ntp_server_key"); ok {
		o.NTPServerKey = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server_key_id"); ok {
		o.NTPServerKeyID = attr.(int)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		o.DatapathSyncTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("dead_timer"); ok {
		o.DeadTimer = attr.(string)
	}
	if attr, ok := d.GetOk("dead_timer_enabled"); ok {
		o.DeadTimerEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("remote_log_mode"); ok {
		o.RemoteLogMode = attr.(string)
	}
	if attr, ok := d.GetOk("remote_log_server_address"); ok {
		o.RemoteLogServerAddress = attr.(string)
	}
	if attr, ok := d.GetOk("remote_log_server_port"); ok {
		o.RemoteLogServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_upgrade_path"); ok {
		o.MetadataUpgradePath = attr.(string)
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		o.FlowEvictionThreshold = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_duration"); ok {
		o.ControllerLessDuration = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_forwarding_mode"); ok {
		o.ControllerLessForwardingMode = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_remote_duration"); ok {
		o.ControllerLessRemoteDuration = attr.(string)
	}
	if attr, ok := d.GetOk("force_immediate_system_sync"); ok {
		o.ForceImmediateSystemSync = attr.(bool)
	}
	if attr, ok := d.GetOk("open_flow_audit_timer"); ok {
		o.OpenFlowAuditTimer = attr.(int)
	}
	if attr, ok := d.GetOk("upgrade_action"); ok {
		o.UpgradeAction = attr.(string)
	}
	if attr, ok := d.GetOk("use_two_factor"); ok {
		o.UseTwoFactor = attr.(bool)
	}
	if attr, ok := d.GetOk("stats_collector_port"); ok {
		o.StatsCollectorPort = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("system_sync_scheduler"); ok {
		o.SystemSyncScheduler = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateInfrastructureGatewayProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceInfrastructureGatewayProfileRead(d, m)
}

func resourceInfrastructureGatewayProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureGatewayProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ntp_server_key", o.NTPServerKey)
	d.Set("ntp_server_key_id", o.NTPServerKeyID)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("datapath_sync_timeout", o.DatapathSyncTimeout)
	d.Set("dead_timer", o.DeadTimer)
	d.Set("dead_timer_enabled", o.DeadTimerEnabled)
	d.Set("remote_log_mode", o.RemoteLogMode)
	d.Set("remote_log_server_address", o.RemoteLogServerAddress)
	d.Set("remote_log_server_port", o.RemoteLogServerPort)
	d.Set("description", o.Description)
	d.Set("metadata_upgrade_path", o.MetadataUpgradePath)
	d.Set("flow_eviction_threshold", o.FlowEvictionThreshold)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("controller_less_duration", o.ControllerLessDuration)
	d.Set("controller_less_forwarding_mode", o.ControllerLessForwardingMode)
	d.Set("controller_less_remote_duration", o.ControllerLessRemoteDuration)
	d.Set("force_immediate_system_sync", o.ForceImmediateSystemSync)
	d.Set("open_flow_audit_timer", o.OpenFlowAuditTimer)
	d.Set("upgrade_action", o.UpgradeAction)
	d.Set("proxy_dns_name", o.ProxyDNSName)
	d.Set("use_two_factor", o.UseTwoFactor)
	d.Set("stats_collector_port", o.StatsCollectorPort)
	d.Set("external_id", o.ExternalID)
	d.Set("system_sync_scheduler", o.SystemSyncScheduler)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceInfrastructureGatewayProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureGatewayProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.ProxyDNSName = d.Get("proxy_dns_name").(string)

	if attr, ok := d.GetOk("ntp_server_key"); ok {
		o.NTPServerKey = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server_key_id"); ok {
		o.NTPServerKeyID = attr.(int)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		o.DatapathSyncTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("dead_timer"); ok {
		o.DeadTimer = attr.(string)
	}
	if attr, ok := d.GetOk("dead_timer_enabled"); ok {
		o.DeadTimerEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("remote_log_mode"); ok {
		o.RemoteLogMode = attr.(string)
	}
	if attr, ok := d.GetOk("remote_log_server_address"); ok {
		o.RemoteLogServerAddress = attr.(string)
	}
	if attr, ok := d.GetOk("remote_log_server_port"); ok {
		o.RemoteLogServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_upgrade_path"); ok {
		o.MetadataUpgradePath = attr.(string)
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		o.FlowEvictionThreshold = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_duration"); ok {
		o.ControllerLessDuration = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_forwarding_mode"); ok {
		o.ControllerLessForwardingMode = attr.(string)
	}
	if attr, ok := d.GetOk("controller_less_remote_duration"); ok {
		o.ControllerLessRemoteDuration = attr.(string)
	}
	if attr, ok := d.GetOk("force_immediate_system_sync"); ok {
		o.ForceImmediateSystemSync = attr.(bool)
	}
	if attr, ok := d.GetOk("open_flow_audit_timer"); ok {
		o.OpenFlowAuditTimer = attr.(int)
	}
	if attr, ok := d.GetOk("upgrade_action"); ok {
		o.UpgradeAction = attr.(string)
	}
	if attr, ok := d.GetOk("use_two_factor"); ok {
		o.UseTwoFactor = attr.(bool)
	}
	if attr, ok := d.GetOk("stats_collector_port"); ok {
		o.StatsCollectorPort = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("system_sync_scheduler"); ok {
		o.SystemSyncScheduler = attr.(string)
	}

	o.Save()

	return nil
}

func resourceInfrastructureGatewayProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureGatewayProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
