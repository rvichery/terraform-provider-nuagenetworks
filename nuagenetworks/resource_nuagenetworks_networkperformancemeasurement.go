package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceNetworkPerformanceMeasurement() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetworkPerformanceMeasurementCreate,
        Read:   resourceNetworkPerformanceMeasurementRead,
        Update: resourceNetworkPerformanceMeasurementUpdate,
        Delete: resourceNetworkPerformanceMeasurementDelete,
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
            "npm_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NONE",
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_performance_monitor_id": &schema.Schema{
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

func resourceNetworkPerformanceMeasurementCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetworkPerformanceMeasurement object
    o := &vspk.NetworkPerformanceMeasurement{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("npm_type"); ok {
        o.NPMType = attr.(string)
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_performance_monitor_id"); ok {
        o.AssociatedPerformanceMonitorID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateNetworkPerformanceMeasurement(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("networkperformancebindings"); ok {
        o.AssignNetworkPerformanceBindings(attr.(vspk.NetworkPerformanceBindingsList))
    }
    return resourceNetworkPerformanceMeasurementRead(d, m)
}

func resourceNetworkPerformanceMeasurementRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceMeasurement{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("npm_type", o.NPMType)
    d.Set("name", o.Name)
    d.Set("read_only", o.ReadOnly)
    d.Set("description", o.Description)
    d.Set("associated_performance_monitor_id", o.AssociatedPerformanceMonitorID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetworkPerformanceMeasurementUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceMeasurement{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("npm_type"); ok {
        o.NPMType = attr.(string)
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
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

func resourceNetworkPerformanceMeasurementDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceMeasurement{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}