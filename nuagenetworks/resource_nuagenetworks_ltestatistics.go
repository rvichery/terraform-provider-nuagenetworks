package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceLtestatistics() *schema.Resource {
    return &schema.Resource{
        Create: resourceLtestatisticsCreate,
        Read:   resourceLtestatisticsRead,
        Update: resourceLtestatisticsUpdate,
        Delete: resourceLtestatisticsDelete,
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
            "version": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "end_time": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "start_time": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "stats_data": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceLtestatisticsCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Ltestatistics object
    o := &vspk.Ltestatistics{
    }
    if attr, ok := d.GetOk("version"); ok {
        o.Version = attr.(int)
    }
    if attr, ok := d.GetOk("end_time"); ok {
        o.EndTime = attr.(int)
    }
    if attr, ok := d.GetOk("start_time"); ok {
        o.StartTime = attr.(int)
    }
    if attr, ok := d.GetOk("stats_data"); ok {
        o.StatsData = attr.([]interface{})
    }
    parent := &vspk.VLAN{ID: d.Get("parent_vlan").(string)}
    err := parent.CreateLtestatistics(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceLtestatisticsRead(d, m)
}

func resourceLtestatisticsRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Ltestatistics{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("version", o.Version)
    d.Set("end_time", o.EndTime)
    d.Set("start_time", o.StartTime)
    d.Set("stats_data", o.StatsData)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceLtestatisticsUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Ltestatistics{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("version"); ok {
        o.Version = attr.(int)
    }
    if attr, ok := d.GetOk("end_time"); ok {
        o.EndTime = attr.(int)
    }
    if attr, ok := d.GetOk("start_time"); ok {
        o.StartTime = attr.(int)
    }
    if attr, ok := d.GetOk("stats_data"); ok {
        o.StatsData = attr.([]interface{})
    }

    o.Save()

    return nil
}

func resourceLtestatisticsDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Ltestatistics{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}