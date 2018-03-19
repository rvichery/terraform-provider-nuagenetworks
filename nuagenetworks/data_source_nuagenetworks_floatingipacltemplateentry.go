package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceFloatingIPACLTemplateEntry() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFloatingIPACLTemplateEntryRead,
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
			"associated_live_entity_id": {
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
			"parent_floating_ipacl_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceFloatingIPACLTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
	filteredFloatingIPACLTemplateEntries := vspk.FloatingIPACLTemplateEntriesList{}
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
	parent := &vspk.FloatingIPACLTemplate{ID: d.Get("parent_floating_ipacl_template").(string)}
	filteredFloatingIPACLTemplateEntries, err = parent.FloatingIPACLTemplateEntries(fetchFilter)
	if err != nil {
		return err
	}

	FloatingIPACLTemplateEntry := &vspk.FloatingIPACLTemplateEntry{}

	if len(filteredFloatingIPACLTemplateEntries) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredFloatingIPACLTemplateEntries) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	FloatingIPACLTemplateEntry = filteredFloatingIPACLTemplateEntries[0]

	d.Set("acl_template_name", FloatingIPACLTemplateEntry.ACLTemplateName)
	d.Set("icmp_code", FloatingIPACLTemplateEntry.ICMPCode)
	d.Set("icmp_type", FloatingIPACLTemplateEntry.ICMPType)
	d.Set("ipv6_address_override", FloatingIPACLTemplateEntry.IPv6AddressOverride)
	d.Set("dscp", FloatingIPACLTemplateEntry.DSCP)
	d.Set("last_updated_by", FloatingIPACLTemplateEntry.LastUpdatedBy)
	d.Set("action", FloatingIPACLTemplateEntry.Action)
	d.Set("address_override", FloatingIPACLTemplateEntry.AddressOverride)
	d.Set("description", FloatingIPACLTemplateEntry.Description)
	d.Set("destination_port", FloatingIPACLTemplateEntry.DestinationPort)
	d.Set("network_id", FloatingIPACLTemplateEntry.NetworkID)
	d.Set("network_type", FloatingIPACLTemplateEntry.NetworkType)
	d.Set("mirror_destination_id", FloatingIPACLTemplateEntry.MirrorDestinationID)
	d.Set("flow_logging_enabled", FloatingIPACLTemplateEntry.FlowLoggingEnabled)
	d.Set("enterprise_name", FloatingIPACLTemplateEntry.EnterpriseName)
	d.Set("entity_scope", FloatingIPACLTemplateEntry.EntityScope)
	d.Set("location_id", FloatingIPACLTemplateEntry.LocationID)
	d.Set("location_type", FloatingIPACLTemplateEntry.LocationType)
	d.Set("policy_state", FloatingIPACLTemplateEntry.PolicyState)
	d.Set("domain_name", FloatingIPACLTemplateEntry.DomainName)
	d.Set("source_port", FloatingIPACLTemplateEntry.SourcePort)
	d.Set("priority", FloatingIPACLTemplateEntry.Priority)
	d.Set("protocol", FloatingIPACLTemplateEntry.Protocol)
	d.Set("associated_live_entity_id", FloatingIPACLTemplateEntry.AssociatedLiveEntityID)
	d.Set("stateful", FloatingIPACLTemplateEntry.Stateful)
	d.Set("stats_id", FloatingIPACLTemplateEntry.StatsID)
	d.Set("stats_logging_enabled", FloatingIPACLTemplateEntry.StatsLoggingEnabled)
	d.Set("ether_type", FloatingIPACLTemplateEntry.EtherType)
	d.Set("external_id", FloatingIPACLTemplateEntry.ExternalID)

	d.Set("id", FloatingIPACLTemplateEntry.Identifier())
	d.Set("parent_id", FloatingIPACLTemplateEntry.ParentID)
	d.Set("parent_type", FloatingIPACLTemplateEntry.ParentType)
	d.Set("owner", FloatingIPACLTemplateEntry.Owner)

	d.SetId(FloatingIPACLTemplateEntry.Identifier())

	return nil
}
