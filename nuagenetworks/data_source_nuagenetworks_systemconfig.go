package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSystemConfig() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSystemConfigRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "aar_flow_stats_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "aar_probe_stats_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "acl_allow_origin": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ecmp_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "ldap_sync_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "ldap_trust_store_certifcate": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ldap_trust_store_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ad_gateway_purge_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rd_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rd_public_network_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rd_public_network_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rd_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "zfb_bootstrap_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "zfb_request_retry_timer": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "zfb_scheduler_stale_request_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "pgid_lower_limit": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "pgid_upper_limit": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dhcp_option_size": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vlanid_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vlanid_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_cache_size": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_purge_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_resync_deletion_wait_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_resync_outstanding_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_unreachable_cleanup_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_unreachable_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vnf_task_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vnid_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vnid_public_network_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vnid_public_network_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vnid_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "api_key_renewal_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "api_key_validity": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vport_init_stateful_timer": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "lru_cache_size_per_subnet": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vsc_on_same_version_as_vsd": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vsd_read_only_mode": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vsd_upgrade_is_complete": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "as_number": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vss_stats_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rt_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rt_public_network_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rt_public_network_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "rt_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "evpnbgp_community_tag_as_number": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "evpnbgp_community_tag_lower_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "evpnbgp_community_tag_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "page_max_size": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "page_size": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "max_failed_logins": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "max_response": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "accumulate_licenses_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vcin_load_balancer_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "per_domain_vlan_id_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "performance_path_selection_vnid": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "service_id_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "key_server_monitor_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "key_server_vsd_data_synchronization_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "offset_customer_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "offset_service_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "virtual_firewall_rules_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "ejbca_nsg_certificate_profile": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ejbca_nsg_end_entity_profile": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ejbca_ocsp_responder_cn": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ejbca_ocsp_responder_uri": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ejbca_vsp_root_ca": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "alarms_max_per_object": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "elastic_cluster_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "allow_enterprise_avatar_on_nsg": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "global_mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "flow_collection_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "inactive_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "infrastructure_bgpas_number": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_tunnel_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "post_processor_threads_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_default_sek_generation_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_default_sek_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_default_sek_payload_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_sek_payload_signing_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_seed_generation_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_default_seed_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_default_seed_payload_authentication_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_seed_payload_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_seed_payload_signing_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_traffic_authentication_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_traffic_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "group_key_default_traffic_encryption_key_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_generation_interval_on_forced_re_key": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_generation_interval_on_revoke": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_minimum_sek_generation_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_minimum_sek_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_minimum_seed_generation_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_minimum_seed_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "group_key_minimum_traffic_encryption_key_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nsg_bootstrap_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_config_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_local_ui_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "esi_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "csproot_authentication_method": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stack_trace_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "stateful_acl_non_tcp_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stateful_acltcp_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "static_wan_service_purge_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "statistics_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "stats_collector_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_collector_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_collector_proto_buf_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_max_data_points": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stats_min_duration": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stats_number_of_data_points": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stats_tsdb_server_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sticky_ecmp_idle_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "attach_probe_to_ipsec_npm": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "attach_probe_to_vxlannpm": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "subnet_resync_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "subnet_resync_outstanding_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "customer_id_upper_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "customer_key": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "avatar_base_path": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "avatar_base_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "event_log_cleanup_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_log_entry_max_age": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_processor_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_processor_max_events_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_processor_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "two_factor_code_expiry": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "two_factor_code_length": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "two_factor_code_seed_length": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dynamic_wan_service_diff_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "syslog_destination_host": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "syslog_destination_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "sysmon_cleanup_task_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "sysmon_node_presence_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "sysmon_probe_response_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "system_avatar_data": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_avatar_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceSystemConfigRead(d *schema.ResourceData, m interface{}) error {
    filteredSystemConfigs := vspk.SystemConfigsList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    parent := m.(*vspk.Me)
    filteredSystemConfigs, err = parent.SystemConfigs(fetchFilter)
    if err != nil {
        return err
    }

    SystemConfig := &vspk.SystemConfig{}

    if len(filteredSystemConfigs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSystemConfigs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SystemConfig = filteredSystemConfigs[0]

    d.Set("aar_flow_stats_interval", SystemConfig.AARFlowStatsInterval)
    d.Set("aar_probe_stats_interval", SystemConfig.AARProbeStatsInterval)
    d.Set("acl_allow_origin", SystemConfig.ACLAllowOrigin)
    d.Set("ecmp_count", SystemConfig.ECMPCount)
    d.Set("ldap_sync_interval", SystemConfig.LDAPSyncInterval)
    d.Set("ldap_trust_store_certifcate", SystemConfig.LDAPTrustStoreCertifcate)
    d.Set("ldap_trust_store_password", SystemConfig.LDAPTrustStorePassword)
    d.Set("ad_gateway_purge_time", SystemConfig.ADGatewayPurgeTime)
    d.Set("rd_lower_limit", SystemConfig.RDLowerLimit)
    d.Set("rd_public_network_lower_limit", SystemConfig.RDPublicNetworkLowerLimit)
    d.Set("rd_public_network_upper_limit", SystemConfig.RDPublicNetworkUpperLimit)
    d.Set("rd_upper_limit", SystemConfig.RDUpperLimit)
    d.Set("zfb_bootstrap_enabled", SystemConfig.ZFBBootstrapEnabled)
    d.Set("zfb_request_retry_timer", SystemConfig.ZFBRequestRetryTimer)
    d.Set("zfb_scheduler_stale_request_timeout", SystemConfig.ZFBSchedulerStaleRequestTimeout)
    d.Set("pgid_lower_limit", SystemConfig.PGIDLowerLimit)
    d.Set("pgid_upper_limit", SystemConfig.PGIDUpperLimit)
    d.Set("dhcp_option_size", SystemConfig.DHCPOptionSize)
    d.Set("vlanid_lower_limit", SystemConfig.VLANIDLowerLimit)
    d.Set("vlanid_upper_limit", SystemConfig.VLANIDUpperLimit)
    d.Set("vm_cache_size", SystemConfig.VMCacheSize)
    d.Set("vm_purge_time", SystemConfig.VMPurgeTime)
    d.Set("vm_resync_deletion_wait_time", SystemConfig.VMResyncDeletionWaitTime)
    d.Set("vm_resync_outstanding_interval", SystemConfig.VMResyncOutstandingInterval)
    d.Set("vm_unreachable_cleanup_time", SystemConfig.VMUnreachableCleanupTime)
    d.Set("vm_unreachable_time", SystemConfig.VMUnreachableTime)
    d.Set("vnf_task_timeout", SystemConfig.VNFTaskTimeout)
    d.Set("vnid_lower_limit", SystemConfig.VNIDLowerLimit)
    d.Set("vnid_public_network_lower_limit", SystemConfig.VNIDPublicNetworkLowerLimit)
    d.Set("vnid_public_network_upper_limit", SystemConfig.VNIDPublicNetworkUpperLimit)
    d.Set("vnid_upper_limit", SystemConfig.VNIDUpperLimit)
    d.Set("api_key_renewal_interval", SystemConfig.APIKeyRenewalInterval)
    d.Set("api_key_validity", SystemConfig.APIKeyValidity)
    d.Set("vport_init_stateful_timer", SystemConfig.VPortInitStatefulTimer)
    d.Set("lru_cache_size_per_subnet", SystemConfig.LRUCacheSizePerSubnet)
    d.Set("vsc_on_same_version_as_vsd", SystemConfig.VSCOnSameVersionAsVSD)
    d.Set("vsd_read_only_mode", SystemConfig.VSDReadOnlyMode)
    d.Set("vsd_upgrade_is_complete", SystemConfig.VSDUpgradeIsComplete)
    d.Set("as_number", SystemConfig.ASNumber)
    d.Set("vss_stats_interval", SystemConfig.VSSStatsInterval)
    d.Set("rt_lower_limit", SystemConfig.RTLowerLimit)
    d.Set("rt_public_network_lower_limit", SystemConfig.RTPublicNetworkLowerLimit)
    d.Set("rt_public_network_upper_limit", SystemConfig.RTPublicNetworkUpperLimit)
    d.Set("rt_upper_limit", SystemConfig.RTUpperLimit)
    d.Set("evpnbgp_community_tag_as_number", SystemConfig.EVPNBGPCommunityTagASNumber)
    d.Set("evpnbgp_community_tag_lower_limit", SystemConfig.EVPNBGPCommunityTagLowerLimit)
    d.Set("evpnbgp_community_tag_upper_limit", SystemConfig.EVPNBGPCommunityTagUpperLimit)
    d.Set("page_max_size", SystemConfig.PageMaxSize)
    d.Set("page_size", SystemConfig.PageSize)
    d.Set("last_updated_by", SystemConfig.LastUpdatedBy)
    d.Set("max_failed_logins", SystemConfig.MaxFailedLogins)
    d.Set("max_response", SystemConfig.MaxResponse)
    d.Set("accumulate_licenses_enabled", SystemConfig.AccumulateLicensesEnabled)
    d.Set("vcin_load_balancer_ip", SystemConfig.VcinLoadBalancerIP)
    d.Set("per_domain_vlan_id_enabled", SystemConfig.PerDomainVlanIdEnabled)
    d.Set("performance_path_selection_vnid", SystemConfig.PerformancePathSelectionVNID)
    d.Set("service_id_upper_limit", SystemConfig.ServiceIDUpperLimit)
    d.Set("key_server_monitor_enabled", SystemConfig.KeyServerMonitorEnabled)
    d.Set("key_server_vsd_data_synchronization_interval", SystemConfig.KeyServerVSDDataSynchronizationInterval)
    d.Set("offset_customer_id", SystemConfig.OffsetCustomerID)
    d.Set("offset_service_id", SystemConfig.OffsetServiceID)
    d.Set("virtual_firewall_rules_enabled", SystemConfig.VirtualFirewallRulesEnabled)
    d.Set("ejbca_nsg_certificate_profile", SystemConfig.EjbcaNSGCertificateProfile)
    d.Set("ejbca_nsg_end_entity_profile", SystemConfig.EjbcaNSGEndEntityProfile)
    d.Set("ejbca_ocsp_responder_cn", SystemConfig.EjbcaOCSPResponderCN)
    d.Set("ejbca_ocsp_responder_uri", SystemConfig.EjbcaOCSPResponderURI)
    d.Set("ejbca_vsp_root_ca", SystemConfig.EjbcaVspRootCa)
    d.Set("alarms_max_per_object", SystemConfig.AlarmsMaxPerObject)
    d.Set("elastic_cluster_name", SystemConfig.ElasticClusterName)
    d.Set("allow_enterprise_avatar_on_nsg", SystemConfig.AllowEnterpriseAvatarOnNSG)
    d.Set("global_mac_address", SystemConfig.GlobalMACAddress)
    d.Set("flow_collection_enabled", SystemConfig.FlowCollectionEnabled)
    d.Set("inactive_timeout", SystemConfig.InactiveTimeout)
    d.Set("infrastructure_bgpas_number", SystemConfig.InfrastructureBGPASNumber)
    d.Set("entity_scope", SystemConfig.EntityScope)
    d.Set("domain_tunnel_type", SystemConfig.DomainTunnelType)
    d.Set("post_processor_threads_count", SystemConfig.PostProcessorThreadsCount)
    d.Set("group_key_default_sek_generation_interval", SystemConfig.GroupKeyDefaultSEKGenerationInterval)
    d.Set("group_key_default_sek_lifetime", SystemConfig.GroupKeyDefaultSEKLifetime)
    d.Set("group_key_default_sek_payload_encryption_algorithm", SystemConfig.GroupKeyDefaultSEKPayloadEncryptionAlgorithm)
    d.Set("group_key_default_sek_payload_signing_algorithm", SystemConfig.GroupKeyDefaultSEKPayloadSigningAlgorithm)
    d.Set("group_key_default_seed_generation_interval", SystemConfig.GroupKeyDefaultSeedGenerationInterval)
    d.Set("group_key_default_seed_lifetime", SystemConfig.GroupKeyDefaultSeedLifetime)
    d.Set("group_key_default_seed_payload_authentication_algorithm", SystemConfig.GroupKeyDefaultSeedPayloadAuthenticationAlgorithm)
    d.Set("group_key_default_seed_payload_encryption_algorithm", SystemConfig.GroupKeyDefaultSeedPayloadEncryptionAlgorithm)
    d.Set("group_key_default_seed_payload_signing_algorithm", SystemConfig.GroupKeyDefaultSeedPayloadSigningAlgorithm)
    d.Set("group_key_default_traffic_authentication_algorithm", SystemConfig.GroupKeyDefaultTrafficAuthenticationAlgorithm)
    d.Set("group_key_default_traffic_encryption_algorithm", SystemConfig.GroupKeyDefaultTrafficEncryptionAlgorithm)
    d.Set("group_key_default_traffic_encryption_key_lifetime", SystemConfig.GroupKeyDefaultTrafficEncryptionKeyLifetime)
    d.Set("group_key_generation_interval_on_forced_re_key", SystemConfig.GroupKeyGenerationIntervalOnForcedReKey)
    d.Set("group_key_generation_interval_on_revoke", SystemConfig.GroupKeyGenerationIntervalOnRevoke)
    d.Set("group_key_minimum_sek_generation_interval", SystemConfig.GroupKeyMinimumSEKGenerationInterval)
    d.Set("group_key_minimum_sek_lifetime", SystemConfig.GroupKeyMinimumSEKLifetime)
    d.Set("group_key_minimum_seed_generation_interval", SystemConfig.GroupKeyMinimumSeedGenerationInterval)
    d.Set("group_key_minimum_seed_lifetime", SystemConfig.GroupKeyMinimumSeedLifetime)
    d.Set("group_key_minimum_traffic_encryption_key_lifetime", SystemConfig.GroupKeyMinimumTrafficEncryptionKeyLifetime)
    d.Set("nsg_bootstrap_endpoint", SystemConfig.NsgBootstrapEndpoint)
    d.Set("nsg_config_endpoint", SystemConfig.NsgConfigEndpoint)
    d.Set("nsg_local_ui_url", SystemConfig.NsgLocalUiUrl)
    d.Set("esi_id", SystemConfig.EsiID)
    d.Set("csproot_authentication_method", SystemConfig.CsprootAuthenticationMethod)
    d.Set("stack_trace_enabled", SystemConfig.StackTraceEnabled)
    d.Set("stateful_acl_non_tcp_timeout", SystemConfig.StatefulACLNonTCPTimeout)
    d.Set("stateful_acltcp_timeout", SystemConfig.StatefulACLTCPTimeout)
    d.Set("static_wan_service_purge_time", SystemConfig.StaticWANServicePurgeTime)
    d.Set("statistics_enabled", SystemConfig.StatisticsEnabled)
    d.Set("stats_collector_address", SystemConfig.StatsCollectorAddress)
    d.Set("stats_collector_port", SystemConfig.StatsCollectorPort)
    d.Set("stats_collector_proto_buf_port", SystemConfig.StatsCollectorProtoBufPort)
    d.Set("stats_max_data_points", SystemConfig.StatsMaxDataPoints)
    d.Set("stats_min_duration", SystemConfig.StatsMinDuration)
    d.Set("stats_number_of_data_points", SystemConfig.StatsNumberOfDataPoints)
    d.Set("stats_tsdb_server_address", SystemConfig.StatsTSDBServerAddress)
    d.Set("sticky_ecmp_idle_timeout", SystemConfig.StickyECMPIdleTimeout)
    d.Set("attach_probe_to_ipsec_npm", SystemConfig.AttachProbeToIPsecNPM)
    d.Set("attach_probe_to_vxlannpm", SystemConfig.AttachProbeToVXLANNPM)
    d.Set("subnet_resync_interval", SystemConfig.SubnetResyncInterval)
    d.Set("subnet_resync_outstanding_interval", SystemConfig.SubnetResyncOutstandingInterval)
    d.Set("customer_id_upper_limit", SystemConfig.CustomerIDUpperLimit)
    d.Set("customer_key", SystemConfig.CustomerKey)
    d.Set("avatar_base_path", SystemConfig.AvatarBasePath)
    d.Set("avatar_base_url", SystemConfig.AvatarBaseURL)
    d.Set("event_log_cleanup_interval", SystemConfig.EventLogCleanupInterval)
    d.Set("event_log_entry_max_age", SystemConfig.EventLogEntryMaxAge)
    d.Set("event_processor_interval", SystemConfig.EventProcessorInterval)
    d.Set("event_processor_max_events_count", SystemConfig.EventProcessorMaxEventsCount)
    d.Set("event_processor_timeout", SystemConfig.EventProcessorTimeout)
    d.Set("two_factor_code_expiry", SystemConfig.TwoFactorCodeExpiry)
    d.Set("two_factor_code_length", SystemConfig.TwoFactorCodeLength)
    d.Set("two_factor_code_seed_length", SystemConfig.TwoFactorCodeSeedLength)
    d.Set("external_id", SystemConfig.ExternalID)
    d.Set("dynamic_wan_service_diff_time", SystemConfig.DynamicWANServiceDiffTime)
    d.Set("syslog_destination_host", SystemConfig.SyslogDestinationHost)
    d.Set("syslog_destination_port", SystemConfig.SyslogDestinationPort)
    d.Set("sysmon_cleanup_task_interval", SystemConfig.SysmonCleanupTaskInterval)
    d.Set("sysmon_node_presence_timeout", SystemConfig.SysmonNodePresenceTimeout)
    d.Set("sysmon_probe_response_timeout", SystemConfig.SysmonProbeResponseTimeout)
    d.Set("system_avatar_data", SystemConfig.SystemAvatarData)
    d.Set("system_avatar_type", SystemConfig.SystemAvatarType)
    
    d.Set("id", SystemConfig.Identifier())
    d.Set("parent_id", SystemConfig.ParentID)
    d.Set("parent_type", SystemConfig.ParentType)
    d.Set("owner", SystemConfig.Owner)

    d.SetId(SystemConfig.Identifier())
    
    return nil
}