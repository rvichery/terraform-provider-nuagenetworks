package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceBootstrapActivation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBootstrapActivationCreate,
		Read:   resourceBootstrapActivationRead,
		Update: resourceBootstrapActivationUpdate,
		Delete: resourceBootstrapActivationDelete,
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
			"cacert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tpm_owner_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tpm_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"srk_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsd_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_bootstrap": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_ns_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBootstrapActivationCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize BootstrapActivation object
	o := &vspk.BootstrapActivation{}
	if attr, ok := d.GetOk("cacert"); ok {
		o.Cacert = attr.(string)
	}
	if attr, ok := d.GetOk("hash"); ok {
		o.Hash = attr.(string)
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
	}
	if attr, ok := d.GetOk("seed"); ok {
		o.Seed = attr.(string)
	}
	if attr, ok := d.GetOk("cert"); ok {
		o.Cert = attr.(string)
	}
	if attr, ok := d.GetOk("config_url"); ok {
		o.ConfigURL = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_owner_password"); ok {
		o.TpmOwnerPassword = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_state"); ok {
		o.TpmState = attr.(int)
	}
	if attr, ok := d.GetOk("srk_password"); ok {
		o.SrkPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vsd_time"); ok {
		o.VsdTime = attr.(int)
	}
	if attr, ok := d.GetOk("csr"); ok {
		o.Csr = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("auto_bootstrap"); ok {
		o.AutoBootstrap = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	err := parent.CreateBootstrapActivation(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceBootstrapActivationRead(d, m)
}

func resourceBootstrapActivationRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BootstrapActivation{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("cacert", o.Cacert)
	d.Set("hash", o.Hash)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("action", o.Action)
	d.Set("seed", o.Seed)
	d.Set("cert", o.Cert)
	d.Set("entity_scope", o.EntityScope)
	d.Set("config_url", o.ConfigURL)
	d.Set("tpm_owner_password", o.TpmOwnerPassword)
	d.Set("tpm_state", o.TpmState)
	d.Set("srk_password", o.SrkPassword)
	d.Set("vsd_time", o.VsdTime)
	d.Set("csr", o.Csr)
	d.Set("status", o.Status)
	d.Set("auto_bootstrap", o.AutoBootstrap)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceBootstrapActivationUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BootstrapActivation{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("cacert"); ok {
		o.Cacert = attr.(string)
	}
	if attr, ok := d.GetOk("hash"); ok {
		o.Hash = attr.(string)
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
	}
	if attr, ok := d.GetOk("seed"); ok {
		o.Seed = attr.(string)
	}
	if attr, ok := d.GetOk("cert"); ok {
		o.Cert = attr.(string)
	}
	if attr, ok := d.GetOk("config_url"); ok {
		o.ConfigURL = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_owner_password"); ok {
		o.TpmOwnerPassword = attr.(string)
	}
	if attr, ok := d.GetOk("tpm_state"); ok {
		o.TpmState = attr.(int)
	}
	if attr, ok := d.GetOk("srk_password"); ok {
		o.SrkPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vsd_time"); ok {
		o.VsdTime = attr.(int)
	}
	if attr, ok := d.GetOk("csr"); ok {
		o.Csr = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("auto_bootstrap"); ok {
		o.AutoBootstrap = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceBootstrapActivationDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BootstrapActivation{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
