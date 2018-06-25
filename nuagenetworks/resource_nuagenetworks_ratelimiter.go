package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceRateLimiter() *schema.Resource {
	return &schema.Resource{
		Create: resourceRateLimiterCreate,
		Read:   resourceRateLimiterRead,
		Update: resourceRateLimiterUpdate,
		Delete: resourceRateLimiterDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peak_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peak_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"committed_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRateLimiterCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize RateLimiter object
	o := &vspk.RateLimiter{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("peak_burst_size"); ok {
		o.PeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("peak_information_rate"); ok {
		o.PeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("committed_information_rate"); ok {
		o.CommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateRateLimiter(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateRateLimiter(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceRateLimiterRead(d, m)
}

func resourceRateLimiterRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RateLimiter{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("peak_burst_size", o.PeakBurstSize)
	d.Set("peak_information_rate", o.PeakInformationRate)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("committed_information_rate", o.CommittedInformationRate)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceRateLimiterUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RateLimiter{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("peak_burst_size"); ok {
		o.PeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("peak_information_rate"); ok {
		o.PeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("committed_information_rate"); ok {
		o.CommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceRateLimiterDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RateLimiter{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
