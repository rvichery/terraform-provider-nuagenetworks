package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceSharedNetworkResource() *schema.Resource {
    return &schema.Resource{
        Create: resourceSharedNetworkResourceCreate,
        Read:   resourceSharedNetworkResourceRead,
        Update: resourceSharedNetworkResourceUpdate,
        Delete: resourceSharedNetworkResourceDelete,
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
            "ecmp_count": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "dhcp_managed": &schema.Schema{
                Type:     schema.TypeBool,
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
            "back_haul_vnid": &schema.Schema{
                Type:     schema.TypeInt,
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
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "gateway_mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "access_restriction_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "permitted_action_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "shared_resource_parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "vn_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "underlay": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "domain_route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uplink_gw_vlan_attachment_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uplink_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uplink_interface_mac": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uplink_vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "use_global_mac": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_pat_mapper_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dynamic_pat_allocation_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceSharedNetworkResourceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SharedNetworkResource object
    o := &vspk.SharedNetworkResource{
        Name: d.Get("name").(string),
        Address: d.Get("address").(string),
        Netmask: d.Get("netmask").(string),
        Type: d.Get("type").(string),
    }
    if attr, ok := d.GetOk("ecmp_count"); ok {
        o.ECMPCount = attr.(int)
    }
    if attr, ok := d.GetOk("dhcp_managed"); ok {
        o.DHCPManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("back_haul_route_distinguisher"); ok {
        o.BackHaulRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_route_target"); ok {
        o.BackHaulRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_vnid"); ok {
        o.BackHaulVNID = attr.(int)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("access_restriction_enabled"); ok {
        o.AccessRestrictionEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("permitted_action_type"); ok {
        o.PermittedActionType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("shared_resource_parent_id"); ok {
        o.SharedResourceParentID = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnID = attr.(int)
    }
    if attr, ok := d.GetOk("underlay"); ok {
        o.Underlay = attr.(bool)
    }
    if attr, ok := d.GetOk("domain_route_distinguisher"); ok {
        o.DomainRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("domain_route_target"); ok {
        o.DomainRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_gw_vlan_attachment_id"); ok {
        o.UplinkGWVlanAttachmentID = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_interface_ip"); ok {
        o.UplinkInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_interface_mac"); ok {
        o.UplinkInterfaceMAC = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_vport_name"); ok {
        o.UplinkVPortName = attr.(string)
    }
    if attr, ok := d.GetOk("use_global_mac"); ok {
        o.UseGlobalMAC = attr.(string)
    }
    if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
        o.AssociatedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_pat_allocation_enabled"); ok {
        o.DynamicPATAllocationEnabled = attr.(bool)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateSharedNetworkResource(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceSharedNetworkResourceRead(d, m)
}

func resourceSharedNetworkResourceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SharedNetworkResource{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ecmp_count", o.ECMPCount)
    d.Set("dhcp_managed", o.DHCPManaged)
    d.Set("back_haul_route_distinguisher", o.BackHaulRouteDistinguisher)
    d.Set("back_haul_route_target", o.BackHaulRouteTarget)
    d.Set("back_haul_vnid", o.BackHaulVNID)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("gateway_mac_address", o.GatewayMACAddress)
    d.Set("access_restriction_enabled", o.AccessRestrictionEnabled)
    d.Set("address", o.Address)
    d.Set("permitted_action_type", o.PermittedActionType)
    d.Set("description", o.Description)
    d.Set("netmask", o.Netmask)
    d.Set("shared_resource_parent_id", o.SharedResourceParentID)
    d.Set("vn_id", o.VnID)
    d.Set("underlay", o.Underlay)
    d.Set("entity_scope", o.EntityScope)
    d.Set("domain_route_distinguisher", o.DomainRouteDistinguisher)
    d.Set("domain_route_target", o.DomainRouteTarget)
    d.Set("uplink_gw_vlan_attachment_id", o.UplinkGWVlanAttachmentID)
    d.Set("uplink_interface_ip", o.UplinkInterfaceIP)
    d.Set("uplink_interface_mac", o.UplinkInterfaceMAC)
    d.Set("uplink_vport_name", o.UplinkVPortName)
    d.Set("use_global_mac", o.UseGlobalMAC)
    d.Set("associated_pat_mapper_id", o.AssociatedPATMapperID)
    d.Set("external_id", o.ExternalID)
    d.Set("dynamic_pat_allocation_enabled", o.DynamicPATAllocationEnabled)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSharedNetworkResourceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SharedNetworkResource{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Address = d.Get("address").(string)
    o.Netmask = d.Get("netmask").(string)
    o.Type = d.Get("type").(string)
    
    if attr, ok := d.GetOk("ecmp_count"); ok {
        o.ECMPCount = attr.(int)
    }
    if attr, ok := d.GetOk("dhcp_managed"); ok {
        o.DHCPManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("back_haul_route_distinguisher"); ok {
        o.BackHaulRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_route_target"); ok {
        o.BackHaulRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("back_haul_vnid"); ok {
        o.BackHaulVNID = attr.(int)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("access_restriction_enabled"); ok {
        o.AccessRestrictionEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("permitted_action_type"); ok {
        o.PermittedActionType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("shared_resource_parent_id"); ok {
        o.SharedResourceParentID = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnID = attr.(int)
    }
    if attr, ok := d.GetOk("underlay"); ok {
        o.Underlay = attr.(bool)
    }
    if attr, ok := d.GetOk("domain_route_distinguisher"); ok {
        o.DomainRouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("domain_route_target"); ok {
        o.DomainRouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_gw_vlan_attachment_id"); ok {
        o.UplinkGWVlanAttachmentID = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_interface_ip"); ok {
        o.UplinkInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_interface_mac"); ok {
        o.UplinkInterfaceMAC = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_vport_name"); ok {
        o.UplinkVPortName = attr.(string)
    }
    if attr, ok := d.GetOk("use_global_mac"); ok {
        o.UseGlobalMAC = attr.(string)
    }
    if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
        o.AssociatedPATMapperID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_pat_allocation_enabled"); ok {
        o.DynamicPATAllocationEnabled = attr.(bool)
    }

    o.Save()

    return nil
}

func resourceSharedNetworkResourceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SharedNetworkResource{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}