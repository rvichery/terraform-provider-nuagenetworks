package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceMultiCastRange() *schema.Resource {
    return &schema.Resource{
        Create: resourceMultiCastRangeCreate,
        Read:   resourceMultiCastRangeRead,
        Update: resourceMultiCastRangeUpdate,
        Delete: resourceMultiCastRangeDelete,

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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "max_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "min_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_multi_cast_channel_map": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceMultiCastRangeCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize MultiCastRange object
    o := &vspk.MultiCastRange{
        MaxAddress: d.Get("max_address").(string),
        MinAddress: d.Get("min_address").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.MultiCastChannelMap{ID: d.Get("parent_multi_cast_channel_map").(string)}
    err := parent.CreateMultiCastRange(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceMultiCastRangeRead(d, m)
}

func resourceMultiCastRangeRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.MultiCastRange{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("max_address", o.MaxAddress)
    d.Set("min_address", o.MinAddress)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceMultiCastRangeUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.MultiCastRange{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.MaxAddress = d.Get("max_address").(string)
    o.MinAddress = d.Get("min_address").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceMultiCastRangeDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.MultiCastRange{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}