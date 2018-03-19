package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceKeyServerMonitorEncryptedSeed() *schema.Resource {
    return &schema.Resource{
        Create: resourceKeyServerMonitorEncryptedSeedCreate,
        Read:   resourceKeyServerMonitorEncryptedSeedRead,
        Update: resourceKeyServerMonitorEncryptedSeedUpdate,
        Delete: resourceKeyServerMonitorEncryptedSeedDelete,
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
            "sek_creation_time": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "key_server_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "enterprise_secured_data_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_key_server_monitor_sek_creation_time": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "associated_key_server_monitor_sekid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_key_server_monitor_seed_creation_time": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "associated_key_server_monitor_seed_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_key_server_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceKeyServerMonitorEncryptedSeedCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize KeyServerMonitorEncryptedSeed object
    o := &vspk.KeyServerMonitorEncryptedSeed{
    }
    if attr, ok := d.GetOk("sek_creation_time"); ok {
        o.SEKCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("key_server_certificate_serial_number"); ok {
        o.KeyServerCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("enterprise_secured_data_id"); ok {
        o.EnterpriseSecuredDataID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sek_creation_time"); ok {
        o.AssociatedKeyServerMonitorSEKCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sekid"); ok {
        o.AssociatedKeyServerMonitorSEKID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_seed_creation_time"); ok {
        o.AssociatedKeyServerMonitorSeedCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_seed_id"); ok {
        o.AssociatedKeyServerMonitorSeedID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
    err := parent.CreateKeyServerMonitorEncryptedSeed(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceKeyServerMonitorEncryptedSeedRead(d, m)
}

func resourceKeyServerMonitorEncryptedSeedRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSeed{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("sek_creation_time", o.SEKCreationTime)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("key_server_certificate_serial_number", o.KeyServerCertificateSerialNumber)
    d.Set("enterprise_secured_data_id", o.EnterpriseSecuredDataID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_key_server_monitor_sek_creation_time", o.AssociatedKeyServerMonitorSEKCreationTime)
    d.Set("associated_key_server_monitor_sekid", o.AssociatedKeyServerMonitorSEKID)
    d.Set("associated_key_server_monitor_seed_creation_time", o.AssociatedKeyServerMonitorSeedCreationTime)
    d.Set("associated_key_server_monitor_seed_id", o.AssociatedKeyServerMonitorSeedID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceKeyServerMonitorEncryptedSeedUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSeed{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("sek_creation_time"); ok {
        o.SEKCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("key_server_certificate_serial_number"); ok {
        o.KeyServerCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("enterprise_secured_data_id"); ok {
        o.EnterpriseSecuredDataID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sek_creation_time"); ok {
        o.AssociatedKeyServerMonitorSEKCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sekid"); ok {
        o.AssociatedKeyServerMonitorSEKID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_seed_creation_time"); ok {
        o.AssociatedKeyServerMonitorSeedCreationTime = attr.(int)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_seed_id"); ok {
        o.AssociatedKeyServerMonitorSeedID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceKeyServerMonitorEncryptedSeedDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSeed{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}