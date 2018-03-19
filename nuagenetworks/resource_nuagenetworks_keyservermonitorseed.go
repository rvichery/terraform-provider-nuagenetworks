package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceKeyServerMonitorSeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyServerMonitorSeedCreate,
		Read:   resourceKeyServerMonitorSeedRead,
		Update: resourceKeyServerMonitorSeedUpdate,
		Delete: resourceKeyServerMonitorSeedDelete,
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
				Optional: true,
				Computed: true,
			},
			"seed_traffic_authentication_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"seed_traffic_encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"seed_traffic_encryption_key_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_key_server_monitor": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceKeyServerMonitorSeedCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize KeyServerMonitorSeed object
	o := &vspk.KeyServerMonitorSeed{}
	if attr, ok := d.GetOk("seed_traffic_authentication_algorithm"); ok {
		o.SeedTrafficAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_traffic_encryption_algorithm"); ok {
		o.SeedTrafficEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_traffic_encryption_key_lifetime"); ok {
		o.SeedTrafficEncryptionKeyLifetime = attr.(int)
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
	err := parent.CreateKeyServerMonitorSeed(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceKeyServerMonitorSeedRead(d, m)
}

func resourceKeyServerMonitorSeedRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSeed{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("seed_traffic_authentication_algorithm", o.SeedTrafficAuthenticationAlgorithm)
	d.Set("seed_traffic_encryption_algorithm", o.SeedTrafficEncryptionAlgorithm)
	d.Set("seed_traffic_encryption_key_lifetime", o.SeedTrafficEncryptionKeyLifetime)
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

func resourceKeyServerMonitorSeedUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSeed{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("seed_traffic_authentication_algorithm"); ok {
		o.SeedTrafficAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_traffic_encryption_algorithm"); ok {
		o.SeedTrafficEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("seed_traffic_encryption_key_lifetime"); ok {
		o.SeedTrafficEncryptionKeyLifetime = attr.(int)
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

func resourceKeyServerMonitorSeedDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.KeyServerMonitorSeed{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
