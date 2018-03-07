package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDSCPRemarkingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDSCPRemarkingPolicyCreate,
		Read:   resourceDSCPRemarkingPolicyRead,
		Update: resourceDSCPRemarkingPolicyUpdate,
		Delete: resourceDSCPRemarkingPolicyDelete,
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
			"dscp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forwarding_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_dscp_remarking_policy_table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDSCPRemarkingPolicyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DSCPRemarkingPolicy object
	o := &vspk.DSCPRemarkingPolicy{}
	if attr, ok := d.GetOk("dscp"); ok {
		o.DSCP = attr.(string)
	}
	if attr, ok := d.GetOk("forwarding_class"); ok {
		o.ForwardingClass = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.DSCPRemarkingPolicyTable{ID: d.Get("parent_dscp_remarking_policy_table").(string)}
	err := parent.CreateDSCPRemarkingPolicy(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceDSCPRemarkingPolicyRead(d, m)
}

func resourceDSCPRemarkingPolicyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPRemarkingPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dscp", o.DSCP)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("forwarding_class", o.ForwardingClass)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDSCPRemarkingPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPRemarkingPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("dscp"); ok {
		o.DSCP = attr.(string)
	}
	if attr, ok := d.GetOk("forwarding_class"); ok {
		o.ForwardingClass = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDSCPRemarkingPolicyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPRemarkingPolicy{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
