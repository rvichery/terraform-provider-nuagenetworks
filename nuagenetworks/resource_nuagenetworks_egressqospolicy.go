package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceEgressQOSPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceEgressQOSPolicyCreate,
		Read:   resourceEgressQOSPolicyRead,
		Update: resourceEgressQOSPolicyUpdate,
		Delete: resourceEgressQOSPolicyDelete,
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
			"parent_queue_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_service_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_egress_qos_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_cos_remarking_policy_table_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_dscp_remarking_policy_table_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue1_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue1_forwarding_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue2_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue2_forwarding_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue3_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue3_forwarding_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue4_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue4_forwarding_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEgressQOSPolicyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EgressQOSPolicy object
	o := &vspk.EgressQOSPolicy{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("parent_queue_associated_rate_limiter_id"); ok {
		o.ParentQueueAssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("default_service_class"); ok {
		o.DefaultServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_egress_qos_id"); ok {
		o.AssocEgressQosId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_cos_remarking_policy_table_id"); ok {
		o.AssociatedCOSRemarkingPolicyTableID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_remarking_policy_table_id"); ok {
		o.AssociatedDSCPRemarkingPolicyTableID = attr.(string)
	}
	if attr, ok := d.GetOk("queue1_associated_rate_limiter_id"); ok {
		o.Queue1AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue1_forwarding_classes"); ok {
		o.Queue1ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue2_associated_rate_limiter_id"); ok {
		o.Queue2AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue2_forwarding_classes"); ok {
		o.Queue2ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue3_associated_rate_limiter_id"); ok {
		o.Queue3AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue3_forwarding_classes"); ok {
		o.Queue3ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue4_associated_rate_limiter_id"); ok {
		o.Queue4AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue4_forwarding_classes"); ok {
		o.Queue4ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateEgressQOSPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateEgressQOSPolicy(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceEgressQOSPolicyRead(d, m)
}

func resourceEgressQOSPolicyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressQOSPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("parent_queue_associated_rate_limiter_id", o.ParentQueueAssociatedRateLimiterID)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("default_service_class", o.DefaultServiceClass)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("assoc_egress_qos_id", o.AssocEgressQosId)
	d.Set("associated_cos_remarking_policy_table_id", o.AssociatedCOSRemarkingPolicyTableID)
	d.Set("associated_dscp_remarking_policy_table_id", o.AssociatedDSCPRemarkingPolicyTableID)
	d.Set("queue1_associated_rate_limiter_id", o.Queue1AssociatedRateLimiterID)
	d.Set("queue1_forwarding_classes", o.Queue1ForwardingClasses)
	d.Set("queue2_associated_rate_limiter_id", o.Queue2AssociatedRateLimiterID)
	d.Set("queue2_forwarding_classes", o.Queue2ForwardingClasses)
	d.Set("queue3_associated_rate_limiter_id", o.Queue3AssociatedRateLimiterID)
	d.Set("queue3_forwarding_classes", o.Queue3ForwardingClasses)
	d.Set("queue4_associated_rate_limiter_id", o.Queue4AssociatedRateLimiterID)
	d.Set("queue4_forwarding_classes", o.Queue4ForwardingClasses)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceEgressQOSPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressQOSPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("parent_queue_associated_rate_limiter_id"); ok {
		o.ParentQueueAssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("default_service_class"); ok {
		o.DefaultServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_egress_qos_id"); ok {
		o.AssocEgressQosId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_cos_remarking_policy_table_id"); ok {
		o.AssociatedCOSRemarkingPolicyTableID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_remarking_policy_table_id"); ok {
		o.AssociatedDSCPRemarkingPolicyTableID = attr.(string)
	}
	if attr, ok := d.GetOk("queue1_associated_rate_limiter_id"); ok {
		o.Queue1AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue1_forwarding_classes"); ok {
		o.Queue1ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue2_associated_rate_limiter_id"); ok {
		o.Queue2AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue2_forwarding_classes"); ok {
		o.Queue2ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue3_associated_rate_limiter_id"); ok {
		o.Queue3AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue3_forwarding_classes"); ok {
		o.Queue3ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("queue4_associated_rate_limiter_id"); ok {
		o.Queue4AssociatedRateLimiterID = attr.(string)
	}
	if attr, ok := d.GetOk("queue4_forwarding_classes"); ok {
		o.Queue4ForwardingClasses = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEgressQOSPolicyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressQOSPolicy{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
