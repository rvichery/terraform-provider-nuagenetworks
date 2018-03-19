package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceEgressACLEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEgressACLEntryTemplateRead,
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
			"ipv6_address_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dscp": {
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
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_l7_application_signature_id": {
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
			"stateful": {
				Type:     schema.TypeBool,
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
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_mirror_destination", "parent_l2_domain", "parent_egress_acl_template"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_mirror_destination", "parent_l2_domain", "parent_egress_acl_template"},
			},
			"parent_mirror_destination": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_l2_domain", "parent_egress_acl_template"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_mirror_destination", "parent_egress_acl_template"},
			},
			"parent_egress_acl_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_mirror_destination", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceEgressACLEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredEgressACLEntryTemplates := vspk.EgressACLEntryTemplatesList{}
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
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_mirror_destination"); ok {
		parent := &vspk.MirrorDestination{ID: attr.(string)}
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_egress_acl_template"); ok {
		parent := &vspk.EgressACLTemplate{ID: attr.(string)}
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredEgressACLEntryTemplates, err = parent.EgressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	EgressACLEntryTemplate := &vspk.EgressACLEntryTemplate{}

	if len(filteredEgressACLEntryTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEgressACLEntryTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EgressACLEntryTemplate = filteredEgressACLEntryTemplates[0]

	d.Set("acl_template_name", EgressACLEntryTemplate.ACLTemplateName)
	d.Set("icmp_code", EgressACLEntryTemplate.ICMPCode)
	d.Set("icmp_type", EgressACLEntryTemplate.ICMPType)
	d.Set("ipv6_address_override", EgressACLEntryTemplate.IPv6AddressOverride)
	d.Set("dscp", EgressACLEntryTemplate.DSCP)
	d.Set("last_updated_by", EgressACLEntryTemplate.LastUpdatedBy)
	d.Set("action", EgressACLEntryTemplate.Action)
	d.Set("address_override", EgressACLEntryTemplate.AddressOverride)
	d.Set("description", EgressACLEntryTemplate.Description)
	d.Set("destination_port", EgressACLEntryTemplate.DestinationPort)
	d.Set("network_id", EgressACLEntryTemplate.NetworkID)
	d.Set("network_type", EgressACLEntryTemplate.NetworkType)
	d.Set("mirror_destination_id", EgressACLEntryTemplate.MirrorDestinationID)
	d.Set("flow_logging_enabled", EgressACLEntryTemplate.FlowLoggingEnabled)
	d.Set("enterprise_name", EgressACLEntryTemplate.EnterpriseName)
	d.Set("entity_scope", EgressACLEntryTemplate.EntityScope)
	d.Set("location_id", EgressACLEntryTemplate.LocationID)
	d.Set("location_type", EgressACLEntryTemplate.LocationType)
	d.Set("policy_state", EgressACLEntryTemplate.PolicyState)
	d.Set("domain_name", EgressACLEntryTemplate.DomainName)
	d.Set("source_port", EgressACLEntryTemplate.SourcePort)
	d.Set("priority", EgressACLEntryTemplate.Priority)
	d.Set("protocol", EgressACLEntryTemplate.Protocol)
	d.Set("associated_l7_application_signature_id", EgressACLEntryTemplate.AssociatedL7ApplicationSignatureID)
	d.Set("associated_live_entity_id", EgressACLEntryTemplate.AssociatedLiveEntityID)
	d.Set("associated_traffic_type", EgressACLEntryTemplate.AssociatedTrafficType)
	d.Set("associated_traffic_type_id", EgressACLEntryTemplate.AssociatedTrafficTypeID)
	d.Set("stateful", EgressACLEntryTemplate.Stateful)
	d.Set("stats_id", EgressACLEntryTemplate.StatsID)
	d.Set("stats_logging_enabled", EgressACLEntryTemplate.StatsLoggingEnabled)
	d.Set("ether_type", EgressACLEntryTemplate.EtherType)
	d.Set("external_id", EgressACLEntryTemplate.ExternalID)

	d.Set("id", EgressACLEntryTemplate.Identifier())
	d.Set("parent_id", EgressACLEntryTemplate.ParentID)
	d.Set("parent_type", EgressACLEntryTemplate.ParentType)
	d.Set("owner", EgressACLEntryTemplate.Owner)

	d.SetId(EgressACLEntryTemplate.Identifier())

	return nil
}
