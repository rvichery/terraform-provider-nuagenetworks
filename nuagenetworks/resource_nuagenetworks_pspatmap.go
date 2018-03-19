package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourcePSPATMap() *schema.Resource {
    return &schema.Resource{
        Create: resourcePSPATMapCreate,
        Read:   resourcePSPATMapRead,
        Update: resourcePSPATMapUpdate,
        Delete: resourcePSPATMapDelete,
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
            "reserved_spatips": &schema.Schema{
                Type:     schema.TypeList,
                Required: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "associated_spat_sources_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_psnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourcePSPATMapCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PSPATMap object
    o := &vspk.PSPATMap{
        Name: d.Get("name").(string),
        ReservedSPATIPs: d.Get("reserved_spatips").([]interface{}),
        AssociatedSPATSourcesPoolID: d.Get("associated_spat_sources_pool_id").(string),
    }
    parent := &vspk.PSNATPool{ID: d.Get("parent_psnat_pool").(string)}
    err := parent.CreatePSPATMap(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePSPATMapRead(d, m)
}

func resourcePSPATMapRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("reserved_spatips", o.ReservedSPATIPs)
    d.Set("associated_spat_sources_pool_id", o.AssociatedSPATSourcesPoolID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePSPATMapUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.ReservedSPATIPs = d.Get("reserved_spatips").([]interface{})
    o.AssociatedSPATSourcesPoolID = d.Get("associated_spat_sources_pool_id").(string)
    

    o.Save()

    return nil
}

func resourcePSPATMapDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}