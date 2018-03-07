package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceOSPFInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceOSPFInstanceCreate,
		Read:   resourceOSPFInstanceRead,
		Update: resourceOSPFInstanceUpdate,
		Delete: resourceOSPFInstanceDelete,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"associated_export_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_import_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"super_backbone_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"export_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"export_to_overlay": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_preference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  150,
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOSPFInstanceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize OSPFInstance object
	o := &vspk.OSPFInstance{
		Name:        d.Get("name").(string),
		ExportLimit: d.Get("export_limit").(int),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("preference"); ok {
		o.Preference = attr.(int)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("super_backbone_enabled"); ok {
		o.SuperBackboneEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("export_to_overlay"); ok {
		o.ExportToOverlay = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("external_preference"); ok {
		o.ExternalPreference = attr.(int)
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateOSPFInstance(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceOSPFInstanceRead(d, m)
}

func resourceOSPFInstanceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInstance{
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
	d.Set("preference", o.Preference)
	d.Set("associated_export_routing_policy_id", o.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", o.AssociatedImportRoutingPolicyID)
	d.Set("super_backbone_enabled", o.SuperBackboneEnabled)
	d.Set("export_limit", o.ExportLimit)
	d.Set("export_to_overlay", o.ExportToOverlay)
	d.Set("external_id", o.ExternalID)
	d.Set("external_preference", o.ExternalPreference)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceOSPFInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInstance{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.ExportLimit = d.Get("export_limit").(int)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("preference"); ok {
		o.Preference = attr.(int)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("super_backbone_enabled"); ok {
		o.SuperBackboneEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("export_to_overlay"); ok {
		o.ExportToOverlay = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("external_preference"); ok {
		o.ExternalPreference = attr.(int)
	}

	o.Save()

	return nil
}

func resourceOSPFInstanceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.OSPFInstance{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
