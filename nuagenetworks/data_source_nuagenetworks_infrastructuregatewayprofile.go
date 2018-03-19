package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceInfrastructureGatewayProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfrastructureGatewayProfileRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server_key_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dead_timer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_timer_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"remote_log_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_log_server_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_log_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_upgrade_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_eviction_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_less_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_less_forwarding_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_less_remote_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_immediate_system_sync": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"open_flow_audit_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_two_factor": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"stats_collector_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_sync_scheduler": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceInfrastructureGatewayProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredInfrastructureGatewayProfiles := vspk.InfrastructureGatewayProfilesList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	parent := m.(*vspk.Me)
	filteredInfrastructureGatewayProfiles, err = parent.InfrastructureGatewayProfiles(fetchFilter)
	if err != nil {
		return err
	}

	InfrastructureGatewayProfile := &vspk.InfrastructureGatewayProfile{}

	if len(filteredInfrastructureGatewayProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredInfrastructureGatewayProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	InfrastructureGatewayProfile = filteredInfrastructureGatewayProfiles[0]

	d.Set("ntp_server_key", InfrastructureGatewayProfile.NTPServerKey)
	d.Set("ntp_server_key_id", InfrastructureGatewayProfile.NTPServerKeyID)
	d.Set("name", InfrastructureGatewayProfile.Name)
	d.Set("last_updated_by", InfrastructureGatewayProfile.LastUpdatedBy)
	d.Set("datapath_sync_timeout", InfrastructureGatewayProfile.DatapathSyncTimeout)
	d.Set("dead_timer", InfrastructureGatewayProfile.DeadTimer)
	d.Set("dead_timer_enabled", InfrastructureGatewayProfile.DeadTimerEnabled)
	d.Set("remote_log_mode", InfrastructureGatewayProfile.RemoteLogMode)
	d.Set("remote_log_server_address", InfrastructureGatewayProfile.RemoteLogServerAddress)
	d.Set("remote_log_server_port", InfrastructureGatewayProfile.RemoteLogServerPort)
	d.Set("description", InfrastructureGatewayProfile.Description)
	d.Set("metadata_upgrade_path", InfrastructureGatewayProfile.MetadataUpgradePath)
	d.Set("flow_eviction_threshold", InfrastructureGatewayProfile.FlowEvictionThreshold)
	d.Set("enterprise_id", InfrastructureGatewayProfile.EnterpriseID)
	d.Set("entity_scope", InfrastructureGatewayProfile.EntityScope)
	d.Set("controller_less_duration", InfrastructureGatewayProfile.ControllerLessDuration)
	d.Set("controller_less_forwarding_mode", InfrastructureGatewayProfile.ControllerLessForwardingMode)
	d.Set("controller_less_remote_duration", InfrastructureGatewayProfile.ControllerLessRemoteDuration)
	d.Set("force_immediate_system_sync", InfrastructureGatewayProfile.ForceImmediateSystemSync)
	d.Set("open_flow_audit_timer", InfrastructureGatewayProfile.OpenFlowAuditTimer)
	d.Set("upgrade_action", InfrastructureGatewayProfile.UpgradeAction)
	d.Set("proxy_dns_name", InfrastructureGatewayProfile.ProxyDNSName)
	d.Set("use_two_factor", InfrastructureGatewayProfile.UseTwoFactor)
	d.Set("stats_collector_port", InfrastructureGatewayProfile.StatsCollectorPort)
	d.Set("external_id", InfrastructureGatewayProfile.ExternalID)
	d.Set("system_sync_scheduler", InfrastructureGatewayProfile.SystemSyncScheduler)

	d.Set("id", InfrastructureGatewayProfile.Identifier())
	d.Set("parent_id", InfrastructureGatewayProfile.ParentID)
	d.Set("parent_type", InfrastructureGatewayProfile.ParentType)
	d.Set("owner", InfrastructureGatewayProfile.Owner)

	d.SetId(InfrastructureGatewayProfile.Identifier())

	return nil
}
