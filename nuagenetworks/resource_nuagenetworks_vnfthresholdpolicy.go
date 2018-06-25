package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVNFThresholdPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVNFThresholdPolicyCreate,
		Read:   resourceVNFThresholdPolicyRead,
		Update: resourceVNFThresholdPolicyUpdate,
		Delete: resourceVNFThresholdPolicyDelete,
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
			"cpu_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"memory_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_occurrence": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"monit_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"storage_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVNFThresholdPolicyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VNFThresholdPolicy object
	o := &vspk.VNFThresholdPolicy{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("cpu_threshold"); ok {
		CPUThreshold := attr.(int)
		o.CPUThreshold = &CPUThreshold
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
	}
	if attr, ok := d.GetOk("memory_threshold"); ok {
		MemoryThreshold := attr.(int)
		o.MemoryThreshold = &MemoryThreshold
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("min_occurrence"); ok {
		MinOccurrence := attr.(int)
		o.MinOccurrence = &MinOccurrence
	}
	if attr, ok := d.GetOk("monit_interval"); ok {
		MonitInterval := attr.(int)
		o.MonitInterval = &MonitInterval
	}
	if attr, ok := d.GetOk("storage_threshold"); ok {
		StorageThreshold := attr.(int)
		o.StorageThreshold = &StorageThreshold
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateVNFThresholdPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateVNFThresholdPolicy(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceVNFThresholdPolicyRead(d, m)
}

func resourceVNFThresholdPolicyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFThresholdPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("cpu_threshold", o.CPUThreshold)
	d.Set("name", o.Name)
	d.Set("action", o.Action)
	d.Set("memory_threshold", o.MemoryThreshold)
	d.Set("description", o.Description)
	d.Set("min_occurrence", o.MinOccurrence)
	d.Set("monit_interval", o.MonitInterval)
	d.Set("storage_threshold", o.StorageThreshold)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVNFThresholdPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFThresholdPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("cpu_threshold"); ok {
		CPUThreshold := attr.(int)
		o.CPUThreshold = &CPUThreshold
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
	}
	if attr, ok := d.GetOk("memory_threshold"); ok {
		MemoryThreshold := attr.(int)
		o.MemoryThreshold = &MemoryThreshold
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("min_occurrence"); ok {
		MinOccurrence := attr.(int)
		o.MinOccurrence = &MinOccurrence
	}
	if attr, ok := d.GetOk("monit_interval"); ok {
		MonitInterval := attr.(int)
		o.MonitInterval = &MonitInterval
	}
	if attr, ok := d.GetOk("storage_threshold"); ok {
		StorageThreshold := attr.(int)
		o.StorageThreshold = &StorageThreshold
	}

	o.Save()

	return nil
}

func resourceVNFThresholdPolicyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFThresholdPolicy{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
