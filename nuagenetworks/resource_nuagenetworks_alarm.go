package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
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
			"target_object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acknowledged": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_condition": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"number_of_occurances": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_tca": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ike_gateway_connection"},
			},
			"parent_ike_gateway_connection": {
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
		Acknowledged := attr.(bool)
		o.Acknowledged = &Acknowledged
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
		Timestamp := attr.(int)
		o.Timestamp = &Timestamp
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("error_condition"); ok {
		ErrorCondition := attr.(int)
		o.ErrorCondition = &ErrorCondition
	}
	if attr, ok := d.GetOk("number_of_occurances"); ok {
		NumberOfOccurances := attr.(int)
		o.NumberOfOccurances = &NumberOfOccurances
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
		Acknowledged := attr.(bool)
		o.Acknowledged = &Acknowledged
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
		Timestamp := attr.(int)
		o.Timestamp = &Timestamp
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("error_condition"); ok {
		ErrorCondition := attr.(int)
		o.ErrorCondition = &ErrorCondition
	}
	if attr, ok := d.GetOk("number_of_occurances"); ok {
		NumberOfOccurances := attr.(int)
		o.NumberOfOccurances = &NumberOfOccurances
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
