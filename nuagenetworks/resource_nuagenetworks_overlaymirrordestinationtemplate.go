package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceOverlayMirrorDestinationTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceOverlayMirrorDestinationTemplateCreate,
        Read:   resourceOverlayMirrorDestinationTemplateRead,
        Update: resourceOverlayMirrorDestinationTemplateUpdate,
        Delete: resourceOverlayMirrorDestinationTemplateDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redundancy_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "end_point_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "trigger_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceOverlayMirrorDestinationTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize OverlayMirrorDestinationTemplate object
    o := &vspk.OverlayMirrorDestinationTemplate{
        Name: d.Get("name").(string),
        EndPointType: d.Get("end_point_type").(string),
    }
    if attr, ok := d.GetOk("redundancy_enabled"); ok {
        o.RedundancyEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("trigger_type"); ok {
        o.TriggerType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.L2DomainTemplate{ID: d.Get("parent_l2_domain_template").(string)}
    err := parent.CreateOverlayMirrorDestinationTemplate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceOverlayMirrorDestinationTemplateRead(d, m)
}

func resourceOverlayMirrorDestinationTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestinationTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("redundancy_enabled", o.RedundancyEnabled)
    d.Set("description", o.Description)
    d.Set("end_point_type", o.EndPointType)
    d.Set("entity_scope", o.EntityScope)
    d.Set("trigger_type", o.TriggerType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceOverlayMirrorDestinationTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestinationTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.EndPointType = d.Get("end_point_type").(string)
    
    if attr, ok := d.GetOk("redundancy_enabled"); ok {
        o.RedundancyEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("trigger_type"); ok {
        o.TriggerType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceOverlayMirrorDestinationTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestinationTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}