package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinkCreate,
		Read:   resourceLinkRead,
		Update: resourceLinkUpdate,
		Delete: resourceLinkDelete,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acceptance_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALL",
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_destination_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_destination_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_source_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_source_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_source_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLinkCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Link object
	o := &vspk.Link{}
	if attr, ok := d.GetOk("acceptance_criteria"); ok {
		o.AcceptanceCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_destination_id"); ok {
		o.AssociatedDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_destination_name"); ok {
		o.AssociatedDestinationName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_destination_type"); ok {
		o.AssociatedDestinationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_id"); ok {
		o.AssociatedSourceID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_name"); ok {
		o.AssociatedSourceName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_type"); ok {
		o.AssociatedSourceType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateLink(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceLinkRead(d, m)
}

func resourceLinkRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Link{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("acceptance_criteria", o.AcceptanceCriteria)
	d.Set("read_only", o.ReadOnly)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_destination_id", o.AssociatedDestinationID)
	d.Set("associated_destination_name", o.AssociatedDestinationName)
	d.Set("associated_destination_type", o.AssociatedDestinationType)
	d.Set("associated_source_id", o.AssociatedSourceID)
	d.Set("associated_source_name", o.AssociatedSourceName)
	d.Set("associated_source_type", o.AssociatedSourceType)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceLinkUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Link{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("acceptance_criteria"); ok {
		o.AcceptanceCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_destination_id"); ok {
		o.AssociatedDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_destination_name"); ok {
		o.AssociatedDestinationName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_destination_type"); ok {
		o.AssociatedDestinationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_id"); ok {
		o.AssociatedSourceID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_name"); ok {
		o.AssociatedSourceName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_source_type"); ok {
		o.AssociatedSourceType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}

	o.Save()

	return nil
}

func resourceLinkDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Link{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
