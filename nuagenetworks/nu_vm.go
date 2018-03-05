package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceVM() *schema.Resource {
    return &schema.Resource{
        Create: resourceVMCreate,
        Read:   resourceVMRead,
        Update: resourceVMUpdate,
        Delete: resourceVMDelete,

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
            "l2_domain_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "vrsid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "reason_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "delete_expiry": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "delete_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "resync_info": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "site_identifier": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "interfaces": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "zone_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "orchestration_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "user_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "user_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "subnet_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "hypervisor_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VM object
    o := &vspk.VM{
        UUID: d.Get("uuid").(string),
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("l2_domain_ids"); ok {
        o.L2DomainIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("vrsid"); ok {
        o.VRSID = attr.(string)
    }
    if attr, ok := d.GetOk("reason_type"); ok {
        o.ReasonType = attr.(string)
    }
    if attr, ok := d.GetOk("delete_expiry"); ok {
        o.DeleteExpiry = attr.(int)
    }
    if attr, ok := d.GetOk("delete_mode"); ok {
        o.DeleteMode = attr.(string)
    }
    if attr, ok := d.GetOk("resync_info"); ok {
        o.ResyncInfo = attr.(interface{})
    }
    if attr, ok := d.GetOk("site_identifier"); ok {
        o.SiteIdentifier = attr.(string)
    }
    if attr, ok := d.GetOk("interfaces"); ok {
        o.Interfaces = attr.([]interface{})
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_name"); ok {
        o.EnterpriseName = attr.(string)
    }
    if attr, ok := d.GetOk("domain_ids"); ok {
        o.DomainIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("zone_ids"); ok {
        o.ZoneIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("orchestration_id"); ok {
        o.OrchestrationID = attr.(string)
    }
    if attr, ok := d.GetOk("user_id"); ok {
        o.UserID = attr.(string)
    }
    if attr, ok := d.GetOk("user_name"); ok {
        o.UserName = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("subnet_ids"); ok {
        o.SubnetIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("hypervisor_ip"); ok {
        o.HypervisorIP = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateVM(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVMRead(d, m)
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VM{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("l2_domain_ids", o.L2DomainIDs)
    d.Set("vrsid", o.VRSID)
    d.Set("uuid", o.UUID)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("reason_type", o.ReasonType)
    d.Set("delete_expiry", o.DeleteExpiry)
    d.Set("delete_mode", o.DeleteMode)
    d.Set("resync_info", o.ResyncInfo)
    d.Set("site_identifier", o.SiteIdentifier)
    d.Set("interfaces", o.Interfaces)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("enterprise_name", o.EnterpriseName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("domain_ids", o.DomainIDs)
    d.Set("zone_ids", o.ZoneIDs)
    d.Set("orchestration_id", o.OrchestrationID)
    d.Set("user_id", o.UserID)
    d.Set("user_name", o.UserName)
    d.Set("status", o.Status)
    d.Set("subnet_ids", o.SubnetIDs)
    d.Set("external_id", o.ExternalID)
    d.Set("hypervisor_ip", o.HypervisorIP)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VM{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.UUID = d.Get("uuid").(string)
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("l2_domain_ids"); ok {
        o.L2DomainIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("vrsid"); ok {
        o.VRSID = attr.(string)
    }
    if attr, ok := d.GetOk("reason_type"); ok {
        o.ReasonType = attr.(string)
    }
    if attr, ok := d.GetOk("delete_expiry"); ok {
        o.DeleteExpiry = attr.(int)
    }
    if attr, ok := d.GetOk("delete_mode"); ok {
        o.DeleteMode = attr.(string)
    }
    if attr, ok := d.GetOk("resync_info"); ok {
        o.ResyncInfo = attr.(interface{})
    }
    if attr, ok := d.GetOk("site_identifier"); ok {
        o.SiteIdentifier = attr.(string)
    }
    if attr, ok := d.GetOk("interfaces"); ok {
        o.Interfaces = attr.([]interface{})
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_name"); ok {
        o.EnterpriseName = attr.(string)
    }
    if attr, ok := d.GetOk("domain_ids"); ok {
        o.DomainIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("zone_ids"); ok {
        o.ZoneIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("orchestration_id"); ok {
        o.OrchestrationID = attr.(string)
    }
    if attr, ok := d.GetOk("user_id"); ok {
        o.UserID = attr.(string)
    }
    if attr, ok := d.GetOk("user_name"); ok {
        o.UserName = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("subnet_ids"); ok {
        o.SubnetIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("hypervisor_ip"); ok {
        o.HypervisorIP = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VM{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}