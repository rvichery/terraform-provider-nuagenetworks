package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceDomainFIPAclTemplateEntry() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceDomainFIPAclTemplateEntryRead,
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
            "action_details": &schema.Schema{
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
            "dest_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dest_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_value": &schema.Schema{
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
            "source_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_value": &schema.Schema{
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
            "associated_live_entity_id": &schema.Schema{
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain_fip_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceDomainFIPAclTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
    filteredDomainFIPAclTemplateEntries := vspk.DomainFIPAclTemplateEntriesList{}
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
    parent := &vspk.DomainFIPAclTemplate{ID: d.Get("parent_domain_fip_acl_template").(string)}
    filteredDomainFIPAclTemplateEntries, err = parent.DomainFIPAclTemplateEntries(fetchFilter)
    if err != nil {
        return err
    }

    DomainFIPAclTemplateEntry := &vspk.DomainFIPAclTemplateEntry{}

    if len(filteredDomainFIPAclTemplateEntries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredDomainFIPAclTemplateEntries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    DomainFIPAclTemplateEntry = filteredDomainFIPAclTemplateEntries[0]

    d.Set("acl_template_name", DomainFIPAclTemplateEntry.ACLTemplateName)
    d.Set("icmp_code", DomainFIPAclTemplateEntry.ICMPCode)
    d.Set("icmp_type", DomainFIPAclTemplateEntry.ICMPType)
    d.Set("ipv6_address_override", DomainFIPAclTemplateEntry.IPv6AddressOverride)
    d.Set("dscp", DomainFIPAclTemplateEntry.DSCP)
    d.Set("last_updated_by", DomainFIPAclTemplateEntry.LastUpdatedBy)
    d.Set("action", DomainFIPAclTemplateEntry.Action)
    d.Set("action_details", DomainFIPAclTemplateEntry.ActionDetails)
    d.Set("address_override", DomainFIPAclTemplateEntry.AddressOverride)
    d.Set("description", DomainFIPAclTemplateEntry.Description)
    d.Set("dest_pg_id", DomainFIPAclTemplateEntry.DestPgId)
    d.Set("dest_pg_type", DomainFIPAclTemplateEntry.DestPgType)
    d.Set("destination_port", DomainFIPAclTemplateEntry.DestinationPort)
    d.Set("destination_type", DomainFIPAclTemplateEntry.DestinationType)
    d.Set("destination_value", DomainFIPAclTemplateEntry.DestinationValue)
    d.Set("network_id", DomainFIPAclTemplateEntry.NetworkID)
    d.Set("network_type", DomainFIPAclTemplateEntry.NetworkType)
    d.Set("mirror_destination_id", DomainFIPAclTemplateEntry.MirrorDestinationID)
    d.Set("flow_logging_enabled", DomainFIPAclTemplateEntry.FlowLoggingEnabled)
    d.Set("enterprise_name", DomainFIPAclTemplateEntry.EnterpriseName)
    d.Set("entity_scope", DomainFIPAclTemplateEntry.EntityScope)
    d.Set("location_id", DomainFIPAclTemplateEntry.LocationID)
    d.Set("location_type", DomainFIPAclTemplateEntry.LocationType)
    d.Set("policy_state", DomainFIPAclTemplateEntry.PolicyState)
    d.Set("domain_name", DomainFIPAclTemplateEntry.DomainName)
    d.Set("source_pg_id", DomainFIPAclTemplateEntry.SourcePgId)
    d.Set("source_pg_type", DomainFIPAclTemplateEntry.SourcePgType)
    d.Set("source_port", DomainFIPAclTemplateEntry.SourcePort)
    d.Set("source_type", DomainFIPAclTemplateEntry.SourceType)
    d.Set("source_value", DomainFIPAclTemplateEntry.SourceValue)
    d.Set("priority", DomainFIPAclTemplateEntry.Priority)
    d.Set("protocol", DomainFIPAclTemplateEntry.Protocol)
    d.Set("associated_live_entity_id", DomainFIPAclTemplateEntry.AssociatedLiveEntityID)
    d.Set("stateful", DomainFIPAclTemplateEntry.Stateful)
    d.Set("stats_id", DomainFIPAclTemplateEntry.StatsID)
    d.Set("stats_logging_enabled", DomainFIPAclTemplateEntry.StatsLoggingEnabled)
    d.Set("ether_type", DomainFIPAclTemplateEntry.EtherType)
    d.Set("external_id", DomainFIPAclTemplateEntry.ExternalID)
    
    d.Set("id", DomainFIPAclTemplateEntry.Identifier())
    d.Set("parent_id", DomainFIPAclTemplateEntry.ParentID)
    d.Set("parent_type", DomainFIPAclTemplateEntry.ParentType)
    d.Set("owner", DomainFIPAclTemplateEntry.Owner)

    d.SetId(DomainFIPAclTemplateEntry.Identifier())
    
    return nil
}