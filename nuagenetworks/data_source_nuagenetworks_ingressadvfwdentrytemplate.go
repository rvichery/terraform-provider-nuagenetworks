package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceIngressAdvFwdEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIngressAdvFwdEntryTemplateRead,
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
			"acl_template_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fc_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dscp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failsafe_datapath": {
				Type:     schema.TypeString,
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
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redirect_rewrite_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redirect_rewrite_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redirect_vport_tag_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_uplink_preference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_destination_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_logging_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_preference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_sla_aware": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"associated_application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_live_entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_traffic_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_traffic_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stats_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stats_logging_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ether_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ingress_adv_fwd_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_mirror_destination"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ingress_adv_fwd_template", "parent_mirror_destination"},
			},
			"parent_mirror_destination": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ingress_adv_fwd_template", "parent_vport"},
			},
		},
	}
}

func dataSourceIngressAdvFwdEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredIngressAdvFwdEntryTemplates := vspk.IngressAdvFwdEntryTemplatesList{}
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
	if attr, ok := d.GetOk("parent_ingress_adv_fwd_template"); ok {
		parent := &vspk.IngressAdvFwdTemplate{ID: attr.(string)}
		filteredIngressAdvFwdEntryTemplates, err = parent.IngressAdvFwdEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredIngressAdvFwdEntryTemplates, err = parent.IngressAdvFwdEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_mirror_destination"); ok {
		parent := &vspk.MirrorDestination{ID: attr.(string)}
		filteredIngressAdvFwdEntryTemplates, err = parent.IngressAdvFwdEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredIngressAdvFwdEntryTemplates, err = parent.IngressAdvFwdEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	IngressAdvFwdEntryTemplate := &vspk.IngressAdvFwdEntryTemplate{}

	if len(filteredIngressAdvFwdEntryTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIngressAdvFwdEntryTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IngressAdvFwdEntryTemplate = filteredIngressAdvFwdEntryTemplates[0]

	d.Set("acl_template_name", IngressAdvFwdEntryTemplate.ACLTemplateName)
	d.Set("icmp_code", IngressAdvFwdEntryTemplate.ICMPCode)
	d.Set("icmp_type", IngressAdvFwdEntryTemplate.ICMPType)
	d.Set("fc_override", IngressAdvFwdEntryTemplate.FCOverride)
	d.Set("ipv6_address_override", IngressAdvFwdEntryTemplate.IPv6AddressOverride)
	d.Set("dscp", IngressAdvFwdEntryTemplate.DSCP)
	d.Set("failsafe_datapath", IngressAdvFwdEntryTemplate.FailsafeDatapath)
	d.Set("name", IngressAdvFwdEntryTemplate.Name)
	d.Set("last_updated_by", IngressAdvFwdEntryTemplate.LastUpdatedBy)
	d.Set("action", IngressAdvFwdEntryTemplate.Action)
	d.Set("address_override", IngressAdvFwdEntryTemplate.AddressOverride)
	d.Set("redirect_rewrite_type", IngressAdvFwdEntryTemplate.RedirectRewriteType)
	d.Set("redirect_rewrite_value", IngressAdvFwdEntryTemplate.RedirectRewriteValue)
	d.Set("redirect_vport_tag_id", IngressAdvFwdEntryTemplate.RedirectVPortTagID)
	d.Set("remote_uplink_preference", IngressAdvFwdEntryTemplate.RemoteUplinkPreference)
	d.Set("description", IngressAdvFwdEntryTemplate.Description)
	d.Set("destination_port", IngressAdvFwdEntryTemplate.DestinationPort)
	d.Set("network_id", IngressAdvFwdEntryTemplate.NetworkID)
	d.Set("network_type", IngressAdvFwdEntryTemplate.NetworkType)
	d.Set("mirror_destination_id", IngressAdvFwdEntryTemplate.MirrorDestinationID)
	d.Set("vlan_range", IngressAdvFwdEntryTemplate.VlanRange)
	d.Set("flow_logging_enabled", IngressAdvFwdEntryTemplate.FlowLoggingEnabled)
	d.Set("enterprise_name", IngressAdvFwdEntryTemplate.EnterpriseName)
	d.Set("entity_scope", IngressAdvFwdEntryTemplate.EntityScope)
	d.Set("location_id", IngressAdvFwdEntryTemplate.LocationID)
	d.Set("location_type", IngressAdvFwdEntryTemplate.LocationType)
	d.Set("policy_state", IngressAdvFwdEntryTemplate.PolicyState)
	d.Set("domain_name", IngressAdvFwdEntryTemplate.DomainName)
	d.Set("source_port", IngressAdvFwdEntryTemplate.SourcePort)
	d.Set("uplink_preference", IngressAdvFwdEntryTemplate.UplinkPreference)
	d.Set("app_type", IngressAdvFwdEntryTemplate.AppType)
	d.Set("priority", IngressAdvFwdEntryTemplate.Priority)
	d.Set("protocol", IngressAdvFwdEntryTemplate.Protocol)
	d.Set("is_sla_aware", IngressAdvFwdEntryTemplate.IsSLAAware)
	d.Set("associated_application_id", IngressAdvFwdEntryTemplate.AssociatedApplicationID)
	d.Set("associated_live_entity_id", IngressAdvFwdEntryTemplate.AssociatedLiveEntityID)
	d.Set("associated_traffic_type", IngressAdvFwdEntryTemplate.AssociatedTrafficType)
	d.Set("associated_traffic_type_id", IngressAdvFwdEntryTemplate.AssociatedTrafficTypeID)
	d.Set("stats_id", IngressAdvFwdEntryTemplate.StatsID)
	d.Set("stats_logging_enabled", IngressAdvFwdEntryTemplate.StatsLoggingEnabled)
	d.Set("ether_type", IngressAdvFwdEntryTemplate.EtherType)
	d.Set("external_id", IngressAdvFwdEntryTemplate.ExternalID)

	d.Set("id", IngressAdvFwdEntryTemplate.Identifier())
	d.Set("parent_id", IngressAdvFwdEntryTemplate.ParentID)
	d.Set("parent_type", IngressAdvFwdEntryTemplate.ParentType)
	d.Set("owner", IngressAdvFwdEntryTemplate.Owner)

	d.SetId(IngressAdvFwdEntryTemplate.Identifier())

	return nil
}
