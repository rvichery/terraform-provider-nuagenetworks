package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceOSPFArea() *schema.Resource {
	return &schema.Resource{
		Create: resourceOSPFAreaCreate,
		Read:   resourceOSPFAreaRead,
		Update: resourceOSPFAreaUpdate,
		Delete: resourceOSPFAreaDelete,
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
			"redistribute_external_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_metric": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_originate_option": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate_area_range": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"aggregate_area_range_nssa": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"area_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"area_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NORMAL",
			},
			"summaries_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"suppress_area_range": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"suppress_area_range_nssa": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_ospf_instance": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOSPFAreaCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize OSPFArea object
	AreaID := d.Get("area_id").(int)
	o := &vspk.OSPFArea{
		AreaID: &AreaID,
	}
	if attr, ok := d.GetOk("redistribute_external_enabled"); ok {
		RedistributeExternalEnabled := attr.(bool)
		o.RedistributeExternalEnabled = &RedistributeExternalEnabled
	}
	if attr, ok := d.GetOk("default_metric"); ok {
		DefaultMetric := attr.(int)
		o.DefaultMetric = &DefaultMetric
	}
	if attr, ok := d.GetOk("default_originate_option"); ok {
		o.DefaultOriginateOption = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("aggregate_area_range"); ok {
		o.AggregateAreaRange = attr.([]interface{})
	}
	if attr, ok := d.GetOk("aggregate_area_range_nssa"); ok {
		o.AggregateAreaRangeNSSA = attr.([]interface{})
	}
	if attr, ok := d.GetOk("area_type"); ok {
		o.AreaType = attr.(string)
	}
	if attr, ok := d.GetOk("summaries_enabled"); ok {
		SummariesEnabled := attr.(bool)
		o.SummariesEnabled = &SummariesEnabled
	}
	if attr, ok := d.GetOk("suppress_area_range"); ok {
		o.SuppressAreaRange = attr.([]interface{})
	}
	if attr, ok := d.GetOk("suppress_area_range_nssa"); ok {
		o.SuppressAreaRangeNSSA = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.OSPFInstance{ID: d.Get("parent_ospf_instance").(string)}
	err := parent.CreateOSPFArea(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceOSPFAreaRead(d, m)
}

func resourceOSPFAreaRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFArea{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("redistribute_external_enabled", o.RedistributeExternalEnabled)
	d.Set("default_metric", o.DefaultMetric)
	d.Set("default_originate_option", o.DefaultOriginateOption)
	d.Set("description", o.Description)
	d.Set("aggregate_area_range", o.AggregateAreaRange)
	d.Set("aggregate_area_range_nssa", o.AggregateAreaRangeNSSA)
	d.Set("entity_scope", o.EntityScope)
	d.Set("area_id", o.AreaID)
	d.Set("area_type", o.AreaType)
	d.Set("summaries_enabled", o.SummariesEnabled)
	d.Set("suppress_area_range", o.SuppressAreaRange)
	d.Set("suppress_area_range_nssa", o.SuppressAreaRangeNSSA)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceOSPFAreaUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFArea{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	AreaID := d.Get("area_id").(int)
	o.AreaID = &AreaID

	if attr, ok := d.GetOk("redistribute_external_enabled"); ok {
		RedistributeExternalEnabled := attr.(bool)
		o.RedistributeExternalEnabled = &RedistributeExternalEnabled
	}
	if attr, ok := d.GetOk("default_metric"); ok {
		DefaultMetric := attr.(int)
		o.DefaultMetric = &DefaultMetric
	}
	if attr, ok := d.GetOk("default_originate_option"); ok {
		o.DefaultOriginateOption = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("aggregate_area_range"); ok {
		o.AggregateAreaRange = attr.([]interface{})
	}
	if attr, ok := d.GetOk("aggregate_area_range_nssa"); ok {
		o.AggregateAreaRangeNSSA = attr.([]interface{})
	}
	if attr, ok := d.GetOk("area_type"); ok {
		o.AreaType = attr.(string)
	}
	if attr, ok := d.GetOk("summaries_enabled"); ok {
		SummariesEnabled := attr.(bool)
		o.SummariesEnabled = &SummariesEnabled
	}
	if attr, ok := d.GetOk("suppress_area_range"); ok {
		o.SuppressAreaRange = attr.([]interface{})
	}
	if attr, ok := d.GetOk("suppress_area_range_nssa"); ok {
		o.SuppressAreaRangeNSSA = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceOSPFAreaDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFArea{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
