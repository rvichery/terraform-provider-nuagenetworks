package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceKeyServerMonitorSEK() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyServerMonitorSEKCreate,
		Read:   resourceKeyServerMonitorSEKRead,
		Update: resourceKeyServerMonitorSEKUpdate,
		Delete: resourceKeyServerMonitorSEKDelete,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"seed_payload_authentication_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seed_payload_encryption_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creation_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceKeyServerMonitorSEKCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize KeyServerMonitorSEK object
	o := &vspk.KeyServerMonitorSEK{}
	if attr, ok := d.GetOk("seed_payload_authentication_algorithm"); ok {
		o.SeedPayloadAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_payload_encryption_algorithm"); ok {
		o.SeedPayloadEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("lifetime"); ok {
		o.Lifetime = attr.(int)
	}
	if attr, ok := d.GetOk("creation_time"); ok {
		o.CreationTime = attr.(int)
	}
	if attr, ok := d.GetOk("start_time"); ok {
		o.StartTime = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
	err := parent.CreateKeyServerMonitorSEK(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceKeyServerMonitorSEKRead(d, m)
}

func resourceKeyServerMonitorSEKRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSEK{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("seed_payload_authentication_algorithm", o.SeedPayloadAuthenticationAlgorithm)
	d.Set("seed_payload_encryption_algorithm", o.SeedPayloadEncryptionAlgorithm)
	d.Set("lifetime", o.Lifetime)
	d.Set("entity_scope", o.EntityScope)
	d.Set("creation_time", o.CreationTime)
	d.Set("start_time", o.StartTime)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceKeyServerMonitorSEKUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSEK{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("seed_payload_authentication_algorithm"); ok {
		o.SeedPayloadAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_payload_encryption_algorithm"); ok {
		o.SeedPayloadEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("lifetime"); ok {
		o.Lifetime = attr.(int)
	}
	if attr, ok := d.GetOk("creation_time"); ok {
		o.CreationTime = attr.(int)
	}
	if attr, ok := d.GetOk("start_time"); ok {
		o.StartTime = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceKeyServerMonitorSEKDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSEK{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
