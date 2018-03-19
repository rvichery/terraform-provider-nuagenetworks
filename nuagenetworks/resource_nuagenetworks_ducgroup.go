package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceDUCGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceDUCGroupCreate,
        Read:   resourceDUCGroupRead,
        Update: resourceDUCGroupUpdate,
        Delete: resourceDUCGroupDelete,
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
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_performance_monitor_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceDUCGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DUCGroup object
    o := &vspk.DUCGroup{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_performance_monitor_id"); ok {
        o.AssociatedPerformanceMonitorID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateDUCGroup(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("nsgateways"); ok {
        o.AssignNSGateways(attr.(vspk.NSGatewaysList))
    }
    return resourceDUCGroupRead(d, m)
}

func resourceDUCGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroup{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("description", o.Description)
    d.Set("associated_performance_monitor_id", o.AssociatedPerformanceMonitorID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDUCGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroup{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_performance_monitor_id"); ok {
        o.AssociatedPerformanceMonitorID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDUCGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}