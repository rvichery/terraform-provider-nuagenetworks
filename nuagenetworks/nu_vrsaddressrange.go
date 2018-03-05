package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceVRSAddressRange() *schema.Resource {
    return &schema.Resource{
        Create: resourceVRSAddressRangeCreate,
        Read:   resourceVRSAddressRangeRead,
        Update: resourceVRSAddressRangeUpdate,
        Delete: resourceVRSAddressRangeDelete,

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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "max_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "min_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_vrs_config": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_hypervisor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config"},
            },
        },
    }
}

func resourceVRSAddressRangeCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VRSAddressRange object
    o := &vspk.VRSAddressRange{
        MaxAddress: d.Get("max_address").(string),
        MinAddress: d.Get("min_address").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        err := parent.CreateVRSAddressRange(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        err := parent.CreateVRSAddressRange(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter"); ok {
        parent := &vspk.VCenter{ID: attr.(string)}
        err := parent.CreateVRSAddressRange(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_vrs_config"); ok {
        parent := &vspk.VCenterVRSConfig{ID: attr.(string)}
        err := parent.CreateVRSAddressRange(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
        parent := &vspk.VCenterHypervisor{ID: attr.(string)}
        err := parent.CreateVRSAddressRange(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceVRSAddressRangeRead(d, m)
}

func resourceVRSAddressRangeRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSAddressRange{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("max_address", o.MaxAddress)
    d.Set("min_address", o.MinAddress)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVRSAddressRangeUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSAddressRange{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.MaxAddress = d.Get("max_address").(string)
    o.MinAddress = d.Get("min_address").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVRSAddressRangeDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSAddressRange{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}