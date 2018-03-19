package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceApplication() *schema.Resource {
    return &schema.Resource{
        Create: resourceApplicationCreate,
        Read:   resourceApplicationRead,
        Update: resourceApplicationUpdate,
        Delete: resourceApplicationDelete,
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
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "bandwidth": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "performance_monitor_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "FIRST_PACKET",
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "enable_pps": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "one_way_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "one_way_jitter": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "one_way_loss": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "post_classification_path": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "ANY",
            },
            "source_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "app_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "optimize_path_selection": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "pre_classification_path": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DEFAULT",
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NONE",
            },
            "associated_l7_application_signature_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "symmetry": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Application object
    o := &vspk.Application{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("bandwidth"); ok {
        o.Bandwidth = attr.(int)
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("performance_monitor_type"); ok {
        o.PerformanceMonitorType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_ip"); ok {
        o.DestinationIP = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("enable_pps"); ok {
        o.EnablePPS = attr.(bool)
    }
    if attr, ok := d.GetOk("one_way_delay"); ok {
        o.OneWayDelay = attr.(int)
    }
    if attr, ok := d.GetOk("one_way_jitter"); ok {
        o.OneWayJitter = attr.(int)
    }
    if attr, ok := d.GetOk("one_way_loss"); ok {
        o.OneWayLoss = attr.(float64)
    }
    if attr, ok := d.GetOk("post_classification_path"); ok {
        o.PostClassificationPath = attr.(string)
    }
    if attr, ok := d.GetOk("source_ip"); ok {
        o.SourceIP = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("optimize_path_selection"); ok {
        o.OptimizePathSelection = attr.(string)
    }
    if attr, ok := d.GetOk("pre_classification_path"); ok {
        o.PreClassificationPath = attr.(string)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
        o.AssociatedL7ApplicationSignatureID = attr.(string)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("symmetry"); ok {
        o.Symmetry = attr.(bool)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateApplication(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceApplicationRead(d, m)
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Application{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dscp", o.DSCP)
    d.Set("name", o.Name)
    d.Set("bandwidth", o.Bandwidth)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("read_only", o.ReadOnly)
    d.Set("performance_monitor_type", o.PerformanceMonitorType)
    d.Set("description", o.Description)
    d.Set("destination_ip", o.DestinationIP)
    d.Set("destination_port", o.DestinationPort)
    d.Set("enable_pps", o.EnablePPS)
    d.Set("one_way_delay", o.OneWayDelay)
    d.Set("one_way_jitter", o.OneWayJitter)
    d.Set("one_way_loss", o.OneWayLoss)
    d.Set("entity_scope", o.EntityScope)
    d.Set("post_classification_path", o.PostClassificationPath)
    d.Set("source_ip", o.SourceIP)
    d.Set("source_port", o.SourcePort)
    d.Set("app_id", o.AppId)
    d.Set("optimize_path_selection", o.OptimizePathSelection)
    d.Set("pre_classification_path", o.PreClassificationPath)
    d.Set("protocol", o.Protocol)
    d.Set("associated_l7_application_signature_id", o.AssociatedL7ApplicationSignatureID)
    d.Set("ether_type", o.EtherType)
    d.Set("external_id", o.ExternalID)
    d.Set("symmetry", o.Symmetry)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Application{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("bandwidth"); ok {
        o.Bandwidth = attr.(int)
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("performance_monitor_type"); ok {
        o.PerformanceMonitorType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_ip"); ok {
        o.DestinationIP = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("enable_pps"); ok {
        o.EnablePPS = attr.(bool)
    }
    if attr, ok := d.GetOk("one_way_delay"); ok {
        o.OneWayDelay = attr.(int)
    }
    if attr, ok := d.GetOk("one_way_jitter"); ok {
        o.OneWayJitter = attr.(int)
    }
    if attr, ok := d.GetOk("one_way_loss"); ok {
        o.OneWayLoss = attr.(float64)
    }
    if attr, ok := d.GetOk("post_classification_path"); ok {
        o.PostClassificationPath = attr.(string)
    }
    if attr, ok := d.GetOk("source_ip"); ok {
        o.SourceIP = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("optimize_path_selection"); ok {
        o.OptimizePathSelection = attr.(string)
    }
    if attr, ok := d.GetOk("pre_classification_path"); ok {
        o.PreClassificationPath = attr.(string)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
        o.AssociatedL7ApplicationSignatureID = attr.(string)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("symmetry"); ok {
        o.Symmetry = attr.(bool)
    }

    o.Save()

    return nil
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Application{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}