package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_authentication_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"seed_payload_encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_key_server_monitor": {
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
		Lifetime := attr.(int)
		o.Lifetime = &Lifetime
	}
	if attr, ok := d.GetOk("creation_time"); ok {
		CreationTime := attr.(int)
		o.CreationTime = &CreationTime
	}
	if attr, ok := d.GetOk("start_time"); ok {
		StartTime := attr.(int)
		o.StartTime = &StartTime
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
		Lifetime := attr.(int)
		o.Lifetime = &Lifetime
	}
	if attr, ok := d.GetOk("creation_time"); ok {
		CreationTime := attr.(int)
		o.CreationTime = &CreationTime
	}
	if attr, ok := d.GetOk("start_time"); ok {
		StartTime := attr.(int)
		o.StartTime = &StartTime
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
