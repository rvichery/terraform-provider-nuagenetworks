package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceNSGRoutingPolicyBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGRoutingPolicyBindingCreate,
		Read:   resourceNSGRoutingPolicyBindingRead,
		Update: resourceNSGRoutingPolicyBindingUpdate,
		Delete: resourceNSGRoutingPolicyBindingDelete,
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
				Optional: true,
				Computed: true,
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
			"associated_export_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_import_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_policy_object_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"export_to_overlay": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNSGRoutingPolicyBindingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NSGRoutingPolicyBinding object
	o := &vspk.NSGRoutingPolicyBinding{
		Name: d.Get("name").(string),
		AssociatedPolicyObjectGroupID: d.Get("associated_policy_object_group_id").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("export_to_overlay"); ok {
		o.ExportToOverlay = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateNSGRoutingPolicyBinding(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceNSGRoutingPolicyBindingRead(d, m)
}

func resourceNSGRoutingPolicyBindingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGRoutingPolicyBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_export_routing_policy_id", o.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", o.AssociatedImportRoutingPolicyID)
	d.Set("associated_policy_object_group_id", o.AssociatedPolicyObjectGroupID)
	d.Set("export_to_overlay", o.ExportToOverlay)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNSGRoutingPolicyBindingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGRoutingPolicyBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.AssociatedPolicyObjectGroupID = d.Get("associated_policy_object_group_id").(string)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("export_to_overlay"); ok {
		o.ExportToOverlay = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNSGRoutingPolicyBindingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NSGRoutingPolicyBinding{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
