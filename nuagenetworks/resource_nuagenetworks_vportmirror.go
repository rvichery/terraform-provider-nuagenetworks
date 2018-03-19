package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceVPortMirror() *schema.Resource {
    return &schema.Resource{
        Create: resourceVPortMirrorCreate,
        Read:   resourceVPortMirrorRead,
        Update: resourceVPortMirrorUpdate,
        Delete: resourceVPortMirrorDelete,
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
            "vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "mirror_destination_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "mirror_direction": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "enterpise_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "attached_network_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceVPortMirrorCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VPortMirror object
    o := &vspk.VPortMirror{
    }
    if attr, ok := d.GetOk("vport_name"); ok {
        o.VPortName = attr.(string)
    }
    if attr, ok := d.GetOk("network_name"); ok {
        o.NetworkName = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_name"); ok {
        o.MirrorDestinationName = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_direction"); ok {
        o.MirrorDirection = attr.(string)
    }
    if attr, ok := d.GetOk("enterpise_name"); ok {
        o.EnterpiseName = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VportId = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_type"); ok {
        o.AttachedNetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.VPort{ID: d.Get("parent_vport").(string)}
    err := parent.CreateVPortMirror(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVPortMirrorRead(d, m)
}

func resourceVPortMirrorRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPortMirror{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vport_name", o.VPortName)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("network_name", o.NetworkName)
    d.Set("mirror_destination_id", o.MirrorDestinationID)
    d.Set("mirror_destination_name", o.MirrorDestinationName)
    d.Set("mirror_direction", o.MirrorDirection)
    d.Set("enterpise_name", o.EnterpiseName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("domain_name", o.DomainName)
    d.Set("vport_id", o.VportId)
    d.Set("attached_network_type", o.AttachedNetworkType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVPortMirrorUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPortMirror{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("vport_name"); ok {
        o.VPortName = attr.(string)
    }
    if attr, ok := d.GetOk("network_name"); ok {
        o.NetworkName = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_name"); ok {
        o.MirrorDestinationName = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_direction"); ok {
        o.MirrorDirection = attr.(string)
    }
    if attr, ok := d.GetOk("enterpise_name"); ok {
        o.EnterpiseName = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VportId = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_type"); ok {
        o.AttachedNetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVPortMirrorDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPortMirror{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}