package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCreate,
		Read:   resourceCertificateRead,
		Update: resourceCertificateUpdate,
		Delete: resourceCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pem_encoded": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCertificateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Certificate object
	o := &vspk.Certificate{}
	if attr, ok := d.GetOk("pem_encoded"); ok {
		o.PemEncoded = attr.(string)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(int)
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
	err := parent.CreateCertificate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceCertificateRead(d, m)
}

func resourceCertificateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Certificate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("pem_encoded", o.PemEncoded)
	d.Set("serial_number", o.SerialNumber)
	d.Set("entity_scope", o.EntityScope)
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

func resourceCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Certificate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("pem_encoded"); ok {
		o.PemEncoded = attr.(string)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(int)
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

func resourceCertificateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Certificate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
