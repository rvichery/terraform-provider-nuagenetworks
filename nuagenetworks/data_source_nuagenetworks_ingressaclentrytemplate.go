package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIngressACLEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIngressACLEntryTemplateRead,
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
			"acl_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dscp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_logging_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enterprise_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_l7_application_signature_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_live_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_traffic_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stateful": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"stats_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stats_logging_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ether_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overlay_mirror_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_mirror_destination", "parent_l2_domain", "parent_ingress_acl_template"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_mirror_destination", "parent_l2_domain", "parent_ingress_acl_template"},
			},
			"parent_mirror_destination": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_l2_domain", "parent_ingress_acl_template"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_mirror_destination", "parent_ingress_acl_template"},
			},
			"parent_ingress_acl_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_mirror_destination", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceIngressACLEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredIngressACLEntryTemplates := vspk.IngressACLEntryTemplatesList{}
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
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_mirror_destination"); ok {
		parent := &vspk.MirrorDestination{ID: attr.(string)}
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
		parent := &vspk.IngressACLTemplate{ID: attr.(string)}
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredIngressACLEntryTemplates, err = parent.IngressACLEntryTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	IngressACLEntryTemplate := &vspk.IngressACLEntryTemplate{}

	if len(filteredIngressACLEntryTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIngressACLEntryTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		IngressACLEntryTemplate = filteredIngressACLEntryTemplates[0]
	}

	d.Set("acl_template_name", IngressACLEntryTemplate.ACLTemplateName)
	d.Set("icmp_code", IngressACLEntryTemplate.ICMPCode)
	d.Set("icmp_type", IngressACLEntryTemplate.ICMPType)
	d.Set("ipv6_address_override", IngressACLEntryTemplate.IPv6AddressOverride)
	d.Set("dscp", IngressACLEntryTemplate.DSCP)
	d.Set("last_updated_by", IngressACLEntryTemplate.LastUpdatedBy)
	d.Set("action", IngressACLEntryTemplate.Action)
	d.Set("address_override", IngressACLEntryTemplate.AddressOverride)
	d.Set("description", IngressACLEntryTemplate.Description)
	d.Set("destination_port", IngressACLEntryTemplate.DestinationPort)
	d.Set("network_id", IngressACLEntryTemplate.NetworkID)
	d.Set("network_type", IngressACLEntryTemplate.NetworkType)
	d.Set("mirror_destination_id", IngressACLEntryTemplate.MirrorDestinationID)
	d.Set("flow_logging_enabled", IngressACLEntryTemplate.FlowLoggingEnabled)
	d.Set("enterprise_name", IngressACLEntryTemplate.EnterpriseName)
	d.Set("entity_scope", IngressACLEntryTemplate.EntityScope)
	d.Set("location_id", IngressACLEntryTemplate.LocationID)
	d.Set("location_type", IngressACLEntryTemplate.LocationType)
	d.Set("policy_state", IngressACLEntryTemplate.PolicyState)
	d.Set("domain_name", IngressACLEntryTemplate.DomainName)
	d.Set("source_port", IngressACLEntryTemplate.SourcePort)
	d.Set("priority", IngressACLEntryTemplate.Priority)
	d.Set("protocol", IngressACLEntryTemplate.Protocol)
	d.Set("associated_l7_application_signature_id", IngressACLEntryTemplate.AssociatedL7ApplicationSignatureID)
	d.Set("associated_live_entity_id", IngressACLEntryTemplate.AssociatedLiveEntityID)
	d.Set("associated_traffic_type", IngressACLEntryTemplate.AssociatedTrafficType)
	d.Set("associated_traffic_type_id", IngressACLEntryTemplate.AssociatedTrafficTypeID)
	d.Set("stateful", IngressACLEntryTemplate.Stateful)
	d.Set("stats_id", IngressACLEntryTemplate.StatsID)
	d.Set("stats_logging_enabled", IngressACLEntryTemplate.StatsLoggingEnabled)
	d.Set("ether_type", IngressACLEntryTemplate.EtherType)
	d.Set("overlay_mirror_destination_id", IngressACLEntryTemplate.OverlayMirrorDestinationID)
	d.Set("external_id", IngressACLEntryTemplate.ExternalID)

	d.Set("id", IngressACLEntryTemplate.Identifier())
	d.Set("parent_id", IngressACLEntryTemplate.ParentID)
	d.Set("parent_type", IngressACLEntryTemplate.ParentType)
	d.Set("owner", IngressACLEntryTemplate.Owner)

	d.SetId(IngressACLEntryTemplate.Identifier())

	return nil
}
