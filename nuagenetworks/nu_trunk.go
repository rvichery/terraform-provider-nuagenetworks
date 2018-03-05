package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceTrunk() *schema.Resource {
    return &schema.Resource{
        Create: resourceTrunkCreate,
        Read:   resourceTrunkRead,
        Update: resourceTrunkUpdate,
        Delete: resourceTrunkDelete,

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
            "associated_vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceTrunkCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Trunk object
    o := &vspk.Trunk{
        Name: d.Get("name").(string),
        AssociatedVPortID: d.Get("associated_vport_id").(string),
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateTrunk(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceTrunkRead(d, m)
}

func resourceTrunkRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Trunk{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("name", o.Name)
    d.Set("associated_vport_id", o.AssociatedVPortID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceTrunkUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Trunk{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.AssociatedVPortID = d.Get("associated_vport_id").(string)
    

    o.Save()

    return nil
}

func resourceTrunkDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Trunk{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}