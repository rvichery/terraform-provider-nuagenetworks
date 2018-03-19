package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEgressAdvFwdEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEgressAdvFwdEntryTemplateRead,
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
			"redirect_vport_tag_id": {
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
			"uplink_preference": {
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
			"associated_live_entity_id": {
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
			"parent_egress_adv_fwd_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceEgressAdvFwdEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredEgressAdvFwdEntryTemplates := vspk.EgressAdvFwdEntryTemplatesList{}
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
	parent := &vspk.EgressAdvFwdTemplate{ID: d.Get("parent_egress_adv_fwd_template").(string)}
	filteredEgressAdvFwdEntryTemplates, err = parent.EgressAdvFwdEntryTemplates(fetchFilter)
	if err != nil {
		return err
	}

	EgressAdvFwdEntryTemplate := &vspk.EgressAdvFwdEntryTemplate{}

	if len(filteredEgressAdvFwdEntryTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEgressAdvFwdEntryTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EgressAdvFwdEntryTemplate = filteredEgressAdvFwdEntryTemplates[0]

	d.Set("acl_template_name", EgressAdvFwdEntryTemplate.ACLTemplateName)
	d.Set("icmp_code", EgressAdvFwdEntryTemplate.ICMPCode)
	d.Set("icmp_type", EgressAdvFwdEntryTemplate.ICMPType)
	d.Set("fc_override", EgressAdvFwdEntryTemplate.FCOverride)
	d.Set("ipv6_address_override", EgressAdvFwdEntryTemplate.IPv6AddressOverride)
	d.Set("dscp", EgressAdvFwdEntryTemplate.DSCP)
	d.Set("failsafe_datapath", EgressAdvFwdEntryTemplate.FailsafeDatapath)
	d.Set("name", EgressAdvFwdEntryTemplate.Name)
	d.Set("last_updated_by", EgressAdvFwdEntryTemplate.LastUpdatedBy)
	d.Set("action", EgressAdvFwdEntryTemplate.Action)
	d.Set("address_override", EgressAdvFwdEntryTemplate.AddressOverride)
	d.Set("redirect_vport_tag_id", EgressAdvFwdEntryTemplate.RedirectVPortTagID)
	d.Set("description", EgressAdvFwdEntryTemplate.Description)
	d.Set("destination_port", EgressAdvFwdEntryTemplate.DestinationPort)
	d.Set("network_id", EgressAdvFwdEntryTemplate.NetworkID)
	d.Set("network_type", EgressAdvFwdEntryTemplate.NetworkType)
	d.Set("mirror_destination_id", EgressAdvFwdEntryTemplate.MirrorDestinationID)
	d.Set("flow_logging_enabled", EgressAdvFwdEntryTemplate.FlowLoggingEnabled)
	d.Set("enterprise_name", EgressAdvFwdEntryTemplate.EnterpriseName)
	d.Set("entity_scope", EgressAdvFwdEntryTemplate.EntityScope)
	d.Set("location_id", EgressAdvFwdEntryTemplate.LocationID)
	d.Set("location_type", EgressAdvFwdEntryTemplate.LocationType)
	d.Set("policy_state", EgressAdvFwdEntryTemplate.PolicyState)
	d.Set("domain_name", EgressAdvFwdEntryTemplate.DomainName)
	d.Set("source_port", EgressAdvFwdEntryTemplate.SourcePort)
	d.Set("uplink_preference", EgressAdvFwdEntryTemplate.UplinkPreference)
	d.Set("priority", EgressAdvFwdEntryTemplate.Priority)
	d.Set("protocol", EgressAdvFwdEntryTemplate.Protocol)
	d.Set("associated_live_entity_id", EgressAdvFwdEntryTemplate.AssociatedLiveEntityID)
	d.Set("stats_id", EgressAdvFwdEntryTemplate.StatsID)
	d.Set("stats_logging_enabled", EgressAdvFwdEntryTemplate.StatsLoggingEnabled)
	d.Set("ether_type", EgressAdvFwdEntryTemplate.EtherType)
	d.Set("external_id", EgressAdvFwdEntryTemplate.ExternalID)

	d.Set("id", EgressAdvFwdEntryTemplate.Identifier())
	d.Set("parent_id", EgressAdvFwdEntryTemplate.ParentID)
	d.Set("parent_type", EgressAdvFwdEntryTemplate.ParentType)
	d.Set("owner", EgressAdvFwdEntryTemplate.Owner)

	d.SetId(EgressAdvFwdEntryTemplate.Identifier())

	return nil
}
