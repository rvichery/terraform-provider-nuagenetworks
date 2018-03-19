package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceDestinationurl() *schema.Resource {
	return &schema.Resource{
		Create: resourceDestinationurlCreate,
		Read:   resourceDestinationurlRead,
		Update: resourceDestinationurlUpdate,
		Delete: resourceDestinationurlDelete,
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
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HEAD",
			},
			"packet_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"percentage_weight": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3000,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"down_threshold_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"probe_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_tier": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDestinationurlCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Destinationurl object
	o := &vspk.Destinationurl{}
	if attr, ok := d.GetOk("url"); ok {
		o.URL = attr.(string)
	}
	if attr, ok := d.GetOk("http_method"); ok {
		o.HTTPMethod = attr.(string)
	}
	if attr, ok := d.GetOk("packet_count"); ok {
		o.PacketCount = attr.(int)
	}
	if attr, ok := d.GetOk("percentage_weight"); ok {
		o.PercentageWeight = attr.(int)
	}
	if attr, ok := d.GetOk("timeout"); ok {
		o.Timeout = attr.(int)
	}
	if attr, ok := d.GetOk("down_threshold_count"); ok {
		o.DownThresholdCount = attr.(int)
	}
	if attr, ok := d.GetOk("probe_interval"); ok {
		o.ProbeInterval = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Tier{ID: d.Get("parent_tier").(string)}
	err := parent.CreateDestinationurl(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceDestinationurlRead(d, m)
}

func resourceDestinationurlRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Destinationurl{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("url", o.URL)
	d.Set("http_method", o.HTTPMethod)
	d.Set("packet_count", o.PacketCount)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("percentage_weight", o.PercentageWeight)
	d.Set("timeout", o.Timeout)
	d.Set("entity_scope", o.EntityScope)
	d.Set("down_threshold_count", o.DownThresholdCount)
	d.Set("probe_interval", o.ProbeInterval)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDestinationurlUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Destinationurl{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("url"); ok {
		o.URL = attr.(string)
	}
	if attr, ok := d.GetOk("http_method"); ok {
		o.HTTPMethod = attr.(string)
	}
	if attr, ok := d.GetOk("packet_count"); ok {
		o.PacketCount = attr.(int)
	}
	if attr, ok := d.GetOk("percentage_weight"); ok {
		o.PercentageWeight = attr.(int)
	}
	if attr, ok := d.GetOk("timeout"); ok {
		o.Timeout = attr.(int)
	}
	if attr, ok := d.GetOk("down_threshold_count"); ok {
		o.DownThresholdCount = attr.(int)
	}
	if attr, ok := d.GetOk("probe_interval"); ok {
		o.ProbeInterval = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDestinationurlDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Destinationurl{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
