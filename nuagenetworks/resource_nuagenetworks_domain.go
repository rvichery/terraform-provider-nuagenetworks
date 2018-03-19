package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceDomain() *schema.Resource {
    return &schema.Resource{
        Create: resourceDomainCreate,
        Read:   resourceDomainRead,
        Update: resourceDomainUpdate,
        Delete: resourceDomainDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
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
            "pat_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ecmp_count": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "bgp_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "dhcp_behavior": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dhcp_server_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "fip_underlay": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DISABLED",
            },
            "label_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "back_haul_route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "back_haul_route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "back_haul_service_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "back_haul_subnet_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "back_haul_subnet_mask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "back_haul_vnid": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "maintenance_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "advertise_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "leaking_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "secondary_dhcp_server_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "service_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dhcp_server_addresses": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "global_routing_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "flow_collection_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "INHERITED",
            },
            "import_route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "encryption": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "underlay_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "local_as": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "policy_change_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "domain_vlanid": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_bgp_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_pat_mapper_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_shared_pat_mapper_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_underlay_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "stretched": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "tunnel_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "customer_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "export_route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
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

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Domain object
    o := &vspk.Domain{
        Name: d.Get("name").(string),
        TemplateID: d.Get("template_id").(string),
    }
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("ecmp_count"); ok {
        o.ECMPCount = attr.(int)
    }
    if attr, ok := d.GetOk("bgp_enabled"); ok {
        o.BGPEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("dhcp_behavior"); ok {
        o.DHCPBehavior = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_server_address"); ok {
        o.DHCPServerAddress = attr.(string)
    }
    if attr, ok := d.GetOk("fip_underlay"); ok {
        o.FIPUnderlay = attr.(bool)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("label_id"); ok {
        o.LabelID = attr.(int)
    }
    if attr, ok := d.GetOk("back_haul_route_distinguisher"); ok {
        o.BackHaulRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_route_target"); ok {
        o.BackHaulRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_subnet_ip_address"); ok {
        o.BackHaulSubnetIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_subnet_mask"); ok {
        o.BackHaulSubnetMask = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_vnid"); ok {
        o.BackHaulVNID = attr.(int)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("advertise_criteria"); ok {
        o.AdvertiseCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("leaking_enabled"); ok {
        o.LeakingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("secondary_dhcp_server_address"); ok {
        o.SecondaryDHCPServerAddress = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_server_addresses"); ok {
        o.DhcpServerAddresses = attr.([]interface{})
    }
    if attr, ok := d.GetOk("global_routing_enabled"); ok {
        o.GlobalRoutingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_collection_enabled"); ok {
        o.FlowCollectionEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("import_route_target"); ok {
        o.ImportRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("local_as"); ok {
        o.LocalAS = attr.(int)
    }
    if attr, ok := d.GetOk("policy_change_status"); ok {
        o.PolicyChangeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(int)
    }
    if attr, ok := d.GetOk("domain_vlanid"); ok {
        o.DomainVLANID = attr.(int)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
        o.AssociatedBGPProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
        o.AssociatedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_pat_mapper_id"); ok {
        o.AssociatedSharedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_id"); ok {
        o.AssociatedUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("stretched"); ok {
        o.Stretched = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("tunnel_type"); ok {
        o.TunnelType = attr.(string)
    }
    if attr, ok := d.GetOk("customer_id"); ok {
        o.CustomerID = attr.(int)
    }
    if attr, ok := d.GetOk("export_route_target"); ok {
        o.ExportRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateDomain(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("domains"); ok {
        o.AssignDomains(attr.(vspk.DomainsList))
    }
    return resourceDomainRead(d, m)
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Domain{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("pat_enabled", o.PATEnabled)
    d.Set("ecmp_count", o.ECMPCount)
    d.Set("bgp_enabled", o.BGPEnabled)
    d.Set("dhcp_behavior", o.DHCPBehavior)
    d.Set("dhcp_server_address", o.DHCPServerAddress)
    d.Set("fip_underlay", o.FIPUnderlay)
    d.Set("dpi", o.DPI)
    d.Set("label_id", o.LabelID)
    d.Set("back_haul_route_distinguisher", o.BackHaulRouteDistinguisher)
    d.Set("back_haul_route_target", o.BackHaulRouteTarget)
    d.Set("back_haul_service_id", o.BackHaulServiceID)
    d.Set("back_haul_subnet_ip_address", o.BackHaulSubnetIPAddress)
    d.Set("back_haul_subnet_mask", o.BackHaulSubnetMask)
    d.Set("back_haul_vnid", o.BackHaulVNID)
    d.Set("maintenance_mode", o.MaintenanceMode)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("advertise_criteria", o.AdvertiseCriteria)
    d.Set("leaking_enabled", o.LeakingEnabled)
    d.Set("secondary_dhcp_server_address", o.SecondaryDHCPServerAddress)
    d.Set("template_id", o.TemplateID)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("service_id", o.ServiceID)
    d.Set("description", o.Description)
    d.Set("dhcp_server_addresses", o.DhcpServerAddresses)
    d.Set("global_routing_enabled", o.GlobalRoutingEnabled)
    d.Set("flow_collection_enabled", o.FlowCollectionEnabled)
    d.Set("import_route_target", o.ImportRouteTarget)
    d.Set("encryption", o.Encryption)
    d.Set("underlay_enabled", o.UnderlayEnabled)
    d.Set("entity_scope", o.EntityScope)
    d.Set("local_as", o.LocalAS)
    d.Set("policy_change_status", o.PolicyChangeStatus)
    d.Set("domain_id", o.DomainID)
    d.Set("domain_vlanid", o.DomainVLANID)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("route_target", o.RouteTarget)
    d.Set("uplink_preference", o.UplinkPreference)
    d.Set("associated_bgp_profile_id", o.AssociatedBGPProfileID)
    d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
    d.Set("associated_pat_mapper_id", o.AssociatedPATMapperID)
    d.Set("associated_shared_pat_mapper_id", o.AssociatedSharedPATMapperID)
    d.Set("associated_underlay_id", o.AssociatedUnderlayID)
    d.Set("stretched", o.Stretched)
    d.Set("multicast", o.Multicast)
    d.Set("tunnel_type", o.TunnelType)
    d.Set("customer_id", o.CustomerID)
    d.Set("export_route_target", o.ExportRouteTarget)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDomainUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Domain{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.TemplateID = d.Get("template_id").(string)
    
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("ecmp_count"); ok {
        o.ECMPCount = attr.(int)
    }
    if attr, ok := d.GetOk("bgp_enabled"); ok {
        o.BGPEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("dhcp_behavior"); ok {
        o.DHCPBehavior = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_server_address"); ok {
        o.DHCPServerAddress = attr.(string)
    }
    if attr, ok := d.GetOk("fip_underlay"); ok {
        o.FIPUnderlay = attr.(bool)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("label_id"); ok {
        o.LabelID = attr.(int)
    }
    if attr, ok := d.GetOk("back_haul_route_distinguisher"); ok {
        o.BackHaulRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_route_target"); ok {
        o.BackHaulRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_subnet_ip_address"); ok {
        o.BackHaulSubnetIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_subnet_mask"); ok {
        o.BackHaulSubnetMask = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_vnid"); ok {
        o.BackHaulVNID = attr.(int)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("advertise_criteria"); ok {
        o.AdvertiseCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("leaking_enabled"); ok {
        o.LeakingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("secondary_dhcp_server_address"); ok {
        o.SecondaryDHCPServerAddress = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_server_addresses"); ok {
        o.DhcpServerAddresses = attr.([]interface{})
    }
    if attr, ok := d.GetOk("global_routing_enabled"); ok {
        o.GlobalRoutingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_collection_enabled"); ok {
        o.FlowCollectionEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("import_route_target"); ok {
        o.ImportRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("local_as"); ok {
        o.LocalAS = attr.(int)
    }
    if attr, ok := d.GetOk("policy_change_status"); ok {
        o.PolicyChangeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(int)
    }
    if attr, ok := d.GetOk("domain_vlanid"); ok {
        o.DomainVLANID = attr.(int)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
        o.AssociatedBGPProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
        o.AssociatedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_pat_mapper_id"); ok {
        o.AssociatedSharedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_id"); ok {
        o.AssociatedUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("stretched"); ok {
        o.Stretched = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("tunnel_type"); ok {
        o.TunnelType = attr.(string)
    }
    if attr, ok := d.GetOk("customer_id"); ok {
        o.CustomerID = attr.(int)
    }
    if attr, ok := d.GetOk("export_route_target"); ok {
        o.ExportRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Domain{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}