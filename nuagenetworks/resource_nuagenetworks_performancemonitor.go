package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourcePerformanceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourcePerformanceMonitorCreate,
		Read:   resourcePerformanceMonitorRead,
		Update: resourcePerformanceMonitorUpdate,
		Delete: resourcePerformanceMonitorDelete,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"payload_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  137,
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "H",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_target_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hold_down_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
			},
			"probe_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ONEWAY",
			},
			"number_of_packets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePerformanceMonitorCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PerformanceMonitor object
	o := &vspk.PerformanceMonitor{
		Name:            d.Get("name").(string),
		Interval:        d.Get("interval").(int),
		NumberOfPackets: d.Get("number_of_packets").(int),
	}
	if attr, ok := d.GetOk("payload_size"); ok {
		o.PayloadSize = attr.(int)
	}
	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_target_list"); ok {
		o.DestinationTargetList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("timeout"); ok {
		o.Timeout = attr.(int)
	}
	if attr, ok := d.GetOk("hold_down_timer"); ok {
		o.HoldDownTimer = attr.(int)
	}
	if attr, ok := d.GetOk("probe_type"); ok {
		o.ProbeType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreatePerformanceMonitor(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreatePerformanceMonitor(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("nsgateways"); ok {
		o.AssignNSGateways(attr.(vspk.NSGatewaysList))
	}
	return resourcePerformanceMonitorRead(d, m)
}

func resourcePerformanceMonitorRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PerformanceMonitor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("payload_size", o.PayloadSize)
	d.Set("read_only", o.ReadOnly)
	d.Set("service_class", o.ServiceClass)
	d.Set("description", o.Description)
	d.Set("destination_target_list", o.DestinationTargetList)
	d.Set("timeout", o.Timeout)
	d.Set("interval", o.Interval)
	d.Set("entity_scope", o.EntityScope)
	d.Set("hold_down_timer", o.HoldDownTimer)
	d.Set("probe_type", o.ProbeType)
	d.Set("number_of_packets", o.NumberOfPackets)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePerformanceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PerformanceMonitor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Interval = d.Get("interval").(int)
	o.NumberOfPackets = d.Get("number_of_packets").(int)

	if attr, ok := d.GetOk("payload_size"); ok {
		o.PayloadSize = attr.(int)
	}
	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_target_list"); ok {
		o.DestinationTargetList = attr.([]interface{})
	}
	if attr, ok := d.GetOk("timeout"); ok {
		o.Timeout = attr.(int)
	}
	if attr, ok := d.GetOk("hold_down_timer"); ok {
		o.HoldDownTimer = attr.(int)
	}
	if attr, ok := d.GetOk("probe_type"); ok {
		o.ProbeType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourcePerformanceMonitorDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PerformanceMonitor{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
