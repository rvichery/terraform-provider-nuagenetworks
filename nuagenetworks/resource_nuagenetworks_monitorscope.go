package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceMonitorscope() *schema.Resource {
    return &schema.Resource{
        Create: resourceMonitorscopeCreate,
        Read:   resourceMonitorscopeRead,
        Update: resourceMonitorscopeUpdate,
        Delete: resourceMonitorscopeDelete,
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
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "destination_nsgs": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "allow_all_destination_nsgs": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "allow_all_source_nsgs": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "source_nsgs": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_application": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_network_performance_measurement"},
            },
            "parent_network_performance_measurement": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_application"},
            },
        },
    }
}

func resourceMonitorscopeCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Monitorscope object
    o := &vspk.Monitorscope{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("destination_nsgs"); ok {
        o.DestinationNSGs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("allow_all_destination_nsgs"); ok {
        o.AllowAllDestinationNSGs = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_all_source_nsgs"); ok {
        o.AllowAllSourceNSGs = attr.(bool)
    }
    if attr, ok := d.GetOk("source_nsgs"); ok {
        o.SourceNSGs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("parent_application"); ok {
        parent := &vspk.Application{ID: attr.(string)}
        err := parent.CreateMonitorscope(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_network_performance_measurement"); ok {
        parent := &vspk.NetworkPerformanceMeasurement{ID: attr.(string)}
        err := parent.CreateMonitorscope(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceMonitorscopeRead(d, m)
}

func resourceMonitorscopeRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Monitorscope{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("read_only", o.ReadOnly)
    d.Set("destination_nsgs", o.DestinationNSGs)
    d.Set("allow_all_destination_nsgs", o.AllowAllDestinationNSGs)
    d.Set("allow_all_source_nsgs", o.AllowAllSourceNSGs)
    d.Set("source_nsgs", o.SourceNSGs)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceMonitorscopeUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Monitorscope{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("destination_nsgs"); ok {
        o.DestinationNSGs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("allow_all_destination_nsgs"); ok {
        o.AllowAllDestinationNSGs = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_all_source_nsgs"); ok {
        o.AllowAllSourceNSGs = attr.(bool)
    }
    if attr, ok := d.GetOk("source_nsgs"); ok {
        o.SourceNSGs = attr.([]interface{})
    }

    o.Save()

    return nil
}

func resourceMonitorscopeDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Monitorscope{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}