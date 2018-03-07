package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEnterprise() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnterpriseRead,
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
			"ldap_authorization_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ldap_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"bgp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dhcp_lease_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vnf_management_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"receive_multi_cast_list_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_multi_cast_list_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared_enterprise": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dictionary_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allow_advanced_qos_configuration": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_gateway_management": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_trusted_forwarding_class": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allowed_forwarding_classes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floating_ips_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"floating_ips_used": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"flow_collection_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_application_performance_management": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"encryption_management_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_enterprise_security_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_group_key_encryption_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_key_server_monitor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"avatar_data": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceEnterpriseRead(d *schema.ResourceData, m interface{}) error {
	filteredEnterprises := vspk.EnterprisesList{}
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
	if attr, ok := d.GetOk("parent_enterprise_profile"); ok {
		parent := &vspk.EnterpriseProfile{ID: attr.(string)}
		filteredEnterprises, err = parent.Enterprises(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredEnterprises, err = parent.Enterprises(fetchFilter)
		if err != nil {
			return err
		}
	}

	Enterprise := &vspk.Enterprise{}

	if len(filteredEnterprises) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEnterprises) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Enterprise = filteredEnterprises[0]
	}

	d.Set("ldap_authorization_enabled", Enterprise.LDAPAuthorizationEnabled)
	d.Set("ldap_enabled", Enterprise.LDAPEnabled)
	d.Set("bgp_enabled", Enterprise.BGPEnabled)
	d.Set("dhcp_lease_interval", Enterprise.DHCPLeaseInterval)
	d.Set("vnf_management_enabled", Enterprise.VNFManagementEnabled)
	d.Set("name", Enterprise.Name)
	d.Set("last_updated_by", Enterprise.LastUpdatedBy)
	d.Set("receive_multi_cast_list_id", Enterprise.ReceiveMultiCastListID)
	d.Set("send_multi_cast_list_id", Enterprise.SendMultiCastListID)
	d.Set("description", Enterprise.Description)
	d.Set("shared_enterprise", Enterprise.SharedEnterprise)
	d.Set("dictionary_version", Enterprise.DictionaryVersion)
	d.Set("allow_advanced_qos_configuration", Enterprise.AllowAdvancedQOSConfiguration)
	d.Set("allow_gateway_management", Enterprise.AllowGatewayManagement)
	d.Set("allow_trusted_forwarding_class", Enterprise.AllowTrustedForwardingClass)
	d.Set("allowed_forwarding_classes", Enterprise.AllowedForwardingClasses)
	d.Set("floating_ips_quota", Enterprise.FloatingIPsQuota)
	d.Set("floating_ips_used", Enterprise.FloatingIPsUsed)
	d.Set("flow_collection_enabled", Enterprise.FlowCollectionEnabled)
	d.Set("enable_application_performance_management", Enterprise.EnableApplicationPerformanceManagement)
	d.Set("encryption_management_mode", Enterprise.EncryptionManagementMode)
	d.Set("enterprise_profile_id", Enterprise.EnterpriseProfileID)
	d.Set("entity_scope", Enterprise.EntityScope)
	d.Set("local_as", Enterprise.LocalAS)
	d.Set("associated_enterprise_security_id", Enterprise.AssociatedEnterpriseSecurityID)
	d.Set("associated_group_key_encryption_profile_id", Enterprise.AssociatedGroupKeyEncryptionProfileID)
	d.Set("associated_key_server_monitor_id", Enterprise.AssociatedKeyServerMonitorID)
	d.Set("customer_id", Enterprise.CustomerID)
	d.Set("avatar_data", Enterprise.AvatarData)
	d.Set("avatar_type", Enterprise.AvatarType)
	d.Set("external_id", Enterprise.ExternalID)

	d.Set("id", Enterprise.Identifier())
	d.Set("parent_id", Enterprise.ParentID)
	d.Set("parent_type", Enterprise.ParentType)
	d.Set("owner", Enterprise.Owner)

	d.SetId(Enterprise.Identifier())

	return nil
}
