package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceBFDSession() *schema.Resource {
    return &schema.Resource{
        Create: resourceBFDSessionCreate,
        Read:   resourceBFDSessionRead,
        Update: resourceBFDSessionUpdate,
        Delete: resourceBFDSessionDelete,

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
            "bfd_destination_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "bfd_multiplier": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 3,
            },
            "bfd_timer": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 500,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multi_hop_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_uplink_connection": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_br_connection"},
            },
            "parent_br_connection": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_uplink_connection"},
            },
        },
    }
}

func resourceBFDSessionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize BFDSession object
    o := &vspk.BFDSession{
    }
    if attr, ok := d.GetOk("bfd_destination_ip"); ok {
        o.BFDDestinationIP = attr.(string)
    }
    if attr, ok := d.GetOk("bfd_multiplier"); ok {
        o.BFDMultiplier = attr.(int)
    }
    if attr, ok := d.GetOk("bfd_timer"); ok {
        o.BFDTimer = attr.(int)
    }
    if attr, ok := d.GetOk("multi_hop_enabled"); ok {
        o.MultiHopEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_uplink_connection"); ok {
        parent := &vspk.UplinkConnection{ID: attr.(string)}
        err := parent.CreateBFDSession(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_br_connection"); ok {
        parent := &vspk.BRConnection{ID: attr.(string)}
        err := parent.CreateBFDSession(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceBFDSessionRead(d, m)
}

func resourceBFDSessionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BFDSession{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("bfd_destination_ip", o.BFDDestinationIP)
    d.Set("bfd_multiplier", o.BFDMultiplier)
    d.Set("bfd_timer", o.BFDTimer)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("multi_hop_enabled", o.MultiHopEnabled)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceBFDSessionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BFDSession{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("bfd_destination_ip"); ok {
        o.BFDDestinationIP = attr.(string)
    }
    if attr, ok := d.GetOk("bfd_multiplier"); ok {
        o.BFDMultiplier = attr.(int)
    }
    if attr, ok := d.GetOk("bfd_timer"); ok {
        o.BFDTimer = attr.(int)
    }
    if attr, ok := d.GetOk("multi_hop_enabled"); ok {
        o.MultiHopEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceBFDSessionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BFDSession{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}