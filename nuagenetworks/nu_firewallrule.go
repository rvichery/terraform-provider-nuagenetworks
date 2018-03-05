package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceFirewallRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceFirewallRuleCreate,
        Read:   resourceFirewallRuleRead,
        Update: resourceFirewallRuleUpdate,
        Delete: resourceFirewallRuleDelete,

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
            "acl_template_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "icmp_code": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "icmp_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ipv6_address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dest_network": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dest_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dest_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_ipv6_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "network_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "location_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_ipv6_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_network": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_application_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associatedfirewall_aclid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "stateful": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize FirewallRule object
    o := &vspk.FirewallRule{
    }
    if attr, ok := d.GetOk("acl_template_name"); ok {
        o.ACLTemplateName = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dest_network"); ok {
        o.DestNetwork = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_id"); ok {
        o.DestPgId = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_type"); ok {
        o.DestPgType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_ipv6_value"); ok {
        o.DestinationIpv6Value = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("destination_type"); ok {
        o.DestinationType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_value"); ok {
        o.DestinationValue = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("enterprise_name"); ok {
        o.EnterpriseName = attr.(string)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("location_type"); ok {
        o.LocationType = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("source_ipv6_value"); ok {
        o.SourceIpv6Value = attr.(string)
    }
    if attr, ok := d.GetOk("source_network"); ok {
        o.SourceNetwork = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_id"); ok {
        o.SourcePgId = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_type"); ok {
        o.SourcePgType = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("source_type"); ok {
        o.SourceType = attr.(string)
    }
    if attr, ok := d.GetOk("source_value"); ok {
        o.SourceValue = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associatedfirewall_aclid"); ok {
        o.AssociatedfirewallACLID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_id"); ok {
        o.StatsID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateFirewallRule(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceFirewallRuleRead(d, m)
}

func resourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallRule{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("acl_template_name", o.ACLTemplateName)
    d.Set("icmp_code", o.ICMPCode)
    d.Set("icmp_type", o.ICMPType)
    d.Set("ipv6_address_override", o.IPv6AddressOverride)
    d.Set("dscp", o.DSCP)
    d.Set("action", o.Action)
    d.Set("address_override", o.AddressOverride)
    d.Set("description", o.Description)
    d.Set("dest_network", o.DestNetwork)
    d.Set("dest_pg_id", o.DestPgId)
    d.Set("dest_pg_type", o.DestPgType)
    d.Set("destination_ipv6_value", o.DestinationIpv6Value)
    d.Set("destination_port", o.DestinationPort)
    d.Set("destination_type", o.DestinationType)
    d.Set("destination_value", o.DestinationValue)
    d.Set("network_id", o.NetworkID)
    d.Set("network_type", o.NetworkType)
    d.Set("mirror_destination_id", o.MirrorDestinationID)
    d.Set("flow_logging_enabled", o.FlowLoggingEnabled)
    d.Set("enterprise_name", o.EnterpriseName)
    d.Set("location_id", o.LocationID)
    d.Set("location_type", o.LocationType)
    d.Set("domain_name", o.DomainName)
    d.Set("source_ipv6_value", o.SourceIpv6Value)
    d.Set("source_network", o.SourceNetwork)
    d.Set("source_pg_id", o.SourcePgId)
    d.Set("source_pg_type", o.SourcePgType)
    d.Set("source_port", o.SourcePort)
    d.Set("source_type", o.SourceType)
    d.Set("source_value", o.SourceValue)
    d.Set("priority", o.Priority)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_application_object_id", o.AssociatedApplicationObjectID)
    d.Set("associatedfirewall_aclid", o.AssociatedfirewallACLID)
    d.Set("stateful", o.Stateful)
    d.Set("stats_id", o.StatsID)
    d.Set("stats_logging_enabled", o.StatsLoggingEnabled)
    d.Set("ether_type", o.EtherType)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallRule{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("acl_template_name"); ok {
        o.ACLTemplateName = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dest_network"); ok {
        o.DestNetwork = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_id"); ok {
        o.DestPgId = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_type"); ok {
        o.DestPgType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_ipv6_value"); ok {
        o.DestinationIpv6Value = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("destination_type"); ok {
        o.DestinationType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_value"); ok {
        o.DestinationValue = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("enterprise_name"); ok {
        o.EnterpriseName = attr.(string)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("location_type"); ok {
        o.LocationType = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("source_ipv6_value"); ok {
        o.SourceIpv6Value = attr.(string)
    }
    if attr, ok := d.GetOk("source_network"); ok {
        o.SourceNetwork = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_id"); ok {
        o.SourcePgId = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_type"); ok {
        o.SourcePgType = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("source_type"); ok {
        o.SourceType = attr.(string)
    }
    if attr, ok := d.GetOk("source_value"); ok {
        o.SourceValue = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associatedfirewall_aclid"); ok {
        o.AssociatedfirewallACLID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_id"); ok {
        o.StatsID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }

    o.Save()

    return nil
}

func resourceFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallRule{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}