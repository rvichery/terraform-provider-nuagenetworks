package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceVNFMetadata() *schema.Resource {
    return &schema.Resource{
        Create: resourceVNFMetadataCreate,
        Read:   resourceVNFMetadataRead,
        Update: resourceVNFMetadataUpdate,
        Delete: resourceVNFMetadataDelete,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "blob": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceVNFMetadataCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VNFMetadata object
    o := &vspk.VNFMetadata{
        Name: d.Get("name").(string),
        Blob: d.Get("blob").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateVNFMetadata(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateVNFMetadata(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceVNFMetadataRead(d, m)
}

func resourceVNFMetadataRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFMetadata{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("description", o.Description)
    d.Set("blob", o.Blob)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVNFMetadataUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFMetadata{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Blob = d.Get("blob").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVNFMetadataDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFMetadata{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}