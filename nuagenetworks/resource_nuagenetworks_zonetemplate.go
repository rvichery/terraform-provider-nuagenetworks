package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceZoneTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceZoneTemplateCreate,
        Read:   resourceZoneTemplateRead,
        Update: resourceZoneTemplateUpdate,
        Delete: resourceZoneTemplateDelete,
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
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "INHERITED",
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ipv6_address": &schema.Schema{
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
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "encryption": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "public_zone": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "number_of_hosts_in_subnets": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "dynamic_ipv6_address": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceZoneTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ZoneTemplate object
    o := &vspk.ZoneTemplate{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("public_zone"); ok {
        o.PublicZone = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("number_of_hosts_in_subnets"); ok {
        o.NumberOfHostsInSubnets = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
        o.DynamicIpv6Address = attr.(bool)
    }
    parent := &vspk.DomainTemplate{ID: d.Get("parent_domain_template").(string)}
    err := parent.CreateZoneTemplate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceZoneTemplateRead(d, m)
}

func resourceZoneTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ZoneTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dpi", o.DPI)
    d.Set("ip_type", o.IPType)
    d.Set("ipv6_address", o.IPv6Address)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("address", o.Address)
    d.Set("description", o.Description)
    d.Set("netmask", o.Netmask)
    d.Set("encryption", o.Encryption)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
    d.Set("public_zone", o.PublicZone)
    d.Set("multicast", o.Multicast)
    d.Set("number_of_hosts_in_subnets", o.NumberOfHostsInSubnets)
    d.Set("external_id", o.ExternalID)
    d.Set("dynamic_ipv6_address", o.DynamicIpv6Address)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceZoneTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ZoneTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("public_zone"); ok {
        o.PublicZone = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("number_of_hosts_in_subnets"); ok {
        o.NumberOfHostsInSubnets = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
        o.DynamicIpv6Address = attr.(bool)
    }

    o.Save()

    return nil
}

func resourceZoneTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ZoneTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}