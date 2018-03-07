package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceAlarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmCreate,
		Read:   resourceAlarmRead,
		Update: resourceAlarmUpdate,
		Delete: resourceAlarmDelete,
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
			"target_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acknowledged": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_condition": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"number_of_occurances": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_tca": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ike_gateway_connection"},
			},
			"parent_ike_gateway_connection": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_tca"},
			},
		},
	}
}

func resourceAlarmCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Alarm object
	o := &vspk.Alarm{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("target_object"); ok {
		o.TargetObject = attr.(string)
	}
	if attr, ok := d.GetOk("acknowledged"); ok {
		o.Acknowledged = attr.(bool)
	}
	if attr, ok := d.GetOk("reason"); ok {
		o.Reason = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("severity"); ok {
		o.Severity = attr.(string)
	}
	if attr, ok := d.GetOk("timestamp"); ok {
		o.Timestamp = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("error_condition"); ok {
		o.ErrorCondition = attr.(int)
	}
	if attr, ok := d.GetOk("number_of_occurances"); ok {
		o.NumberOfOccurances = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_tca"); ok {
		parent := &vspk.TCA{ID: attr.(string)}
		err := parent.CreateAlarm(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ike_gateway_connection"); ok {
		parent := &vspk.IKEGatewayConnection{ID: attr.(string)}
		err := parent.CreateAlarm(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceAlarmRead(d, m)
}

func resourceAlarmRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Alarm{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("target_object", o.TargetObject)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("acknowledged", o.Acknowledged)
	d.Set("reason", o.Reason)
	d.Set("description", o.Description)
	d.Set("severity", o.Severity)
	d.Set("timestamp", o.Timestamp)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("error_condition", o.ErrorCondition)
	d.Set("number_of_occurances", o.NumberOfOccurances)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceAlarmUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Alarm{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("target_object"); ok {
		o.TargetObject = attr.(string)
	}
	if attr, ok := d.GetOk("acknowledged"); ok {
		o.Acknowledged = attr.(bool)
	}
	if attr, ok := d.GetOk("reason"); ok {
		o.Reason = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("severity"); ok {
		o.Severity = attr.(string)
	}
	if attr, ok := d.GetOk("timestamp"); ok {
		o.Timestamp = attr.(int)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("error_condition"); ok {
		o.ErrorCondition = attr.(int)
	}
	if attr, ok := d.GetOk("number_of_occurances"); ok {
		o.NumberOfOccurances = attr.(int)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceAlarmDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Alarm{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
