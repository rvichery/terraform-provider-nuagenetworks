package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDemarcationService() *schema.Resource {
    return &schema.Resource{
        Create: resourceDemarcationServiceCreate,
        Read:   resourceDemarcationServiceRead,
        Update: resourceDemarcationServiceUpdate,
        Delete: resourceDemarcationServiceDelete,

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
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceDemarcationServiceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DemarcationService object
    o := &vspk.DemarcationService{
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vlanid"); ok {
        o.AssociatedVLANID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    err := parent.CreateDemarcationService(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceDemarcationServiceRead(d, m)
}

func resourceDemarcationServiceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DemarcationService{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("priority", o.Priority)
    d.Set("associated_gateway_id", o.AssociatedGatewayID)
    d.Set("associated_vlanid", o.AssociatedVLANID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDemarcationServiceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DemarcationService{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vlanid"); ok {
        o.AssociatedVLANID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDemarcationServiceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DemarcationService{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}