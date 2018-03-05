package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceKeyServerMember() *schema.Resource {
    return &schema.Resource{
        Create: resourceKeyServerMemberCreate,
        Read:   resourceKeyServerMemberRead,
        Update: resourceKeyServerMemberUpdate,
        Delete: resourceKeyServerMemberDelete,

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
            "pem_encoded": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "certificate_serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "fqdn": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "issuer_dn": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "subject_dn": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "public_key": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceKeyServerMemberCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize KeyServerMember object
    o := &vspk.KeyServerMember{
    }
    if attr, ok := d.GetOk("pem_encoded"); ok {
        o.PemEncoded = attr.(string)
    }
    if attr, ok := d.GetOk("certificate_serial_number"); ok {
        o.CertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("fqdn"); ok {
        o.Fqdn = attr.(string)
    }
    if attr, ok := d.GetOk("issuer_dn"); ok {
        o.IssuerDN = attr.(string)
    }
    if attr, ok := d.GetOk("subject_dn"); ok {
        o.SubjectDN = attr.(string)
    }
    if attr, ok := d.GetOk("public_key"); ok {
        o.PublicKey = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateKeyServerMember(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceKeyServerMemberRead(d, m)
}

func resourceKeyServerMemberRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMember{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("pem_encoded", o.PemEncoded)
    d.Set("certificate_serial_number", o.CertificateSerialNumber)
    d.Set("entity_scope", o.EntityScope)
    d.Set("fqdn", o.Fqdn)
    d.Set("issuer_dn", o.IssuerDN)
    d.Set("subject_dn", o.SubjectDN)
    d.Set("public_key", o.PublicKey)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceKeyServerMemberUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMember{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("pem_encoded"); ok {
        o.PemEncoded = attr.(string)
    }
    if attr, ok := d.GetOk("certificate_serial_number"); ok {
        o.CertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("fqdn"); ok {
        o.Fqdn = attr.(string)
    }
    if attr, ok := d.GetOk("issuer_dn"); ok {
        o.IssuerDN = attr.(string)
    }
    if attr, ok := d.GetOk("subject_dn"); ok {
        o.SubjectDN = attr.(string)
    }
    if attr, ok := d.GetOk("public_key"); ok {
        o.PublicKey = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceKeyServerMemberDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMember{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}