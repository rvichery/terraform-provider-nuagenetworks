package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourcePolicyObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyObjectGroupCreate,
		Read:   resourcePolicyObjectGroupRead,
		Update: resourcePolicyObjectGroupUpdate,
		Delete: resourcePolicyObjectGroupDelete,
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
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePolicyObjectGroupCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PolicyObjectGroup object
	o := &vspk.PolicyObjectGroup{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreatePolicyObjectGroup(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("nsgateways"); ok {
		o.AssignNSGateways(attr.(vspk.NSGatewaysList))
	}
	return resourcePolicyObjectGroupRead(d, m)
}

func resourcePolicyObjectGroupRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyObjectGroup{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePolicyObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyObjectGroup{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}

	o.Save()

	return nil
}

func resourcePolicyObjectGroupDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyObjectGroup{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
