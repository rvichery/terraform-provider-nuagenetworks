package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceCommandCreate,
		Read:   resourceCommandRead,
		Update: resourceCommandUpdate,
		Delete: resourceCommandDelete,
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
				Optional: true,
				Computed: true,
			},
			"detailed_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detailed_status_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"command_information": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_param_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_command": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNSPECIFIED",
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_ns_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCommandCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Command object
	o := &vspk.Command{
		Command: d.Get("command").(string),
		Summary: d.Get("summary").(string),
	}
	if attr, ok := d.GetOk("associated_param"); ok {
		o.AssociatedParam = attr.(string)
	}
	if attr, ok := d.GetOk("associated_param_type"); ok {
		o.AssociatedParamType = attr.(string)
	}
	if attr, ok := d.GetOk("override"); ok {
		o.Override = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	err := parent.CreateCommand(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceCommandRead(d, m)
}

func resourceCommandRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Command{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("detailed_status", o.DetailedStatus)
	d.Set("detailed_status_code", o.DetailedStatusCode)
	d.Set("entity_scope", o.EntityScope)
	d.Set("command", o.Command)
	d.Set("command_information", o.CommandInformation)
	d.Set("associated_param", o.AssociatedParam)
	d.Set("associated_param_type", o.AssociatedParamType)
	d.Set("status", o.Status)
	d.Set("full_command", o.FullCommand)
	d.Set("summary", o.Summary)
	d.Set("override", o.Override)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceCommandUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Command{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Command = d.Get("command").(string)
	o.Summary = d.Get("summary").(string)

	if attr, ok := d.GetOk("associated_param"); ok {
		o.AssociatedParam = attr.(string)
	}
	if attr, ok := d.GetOk("associated_param_type"); ok {
		o.AssociatedParamType = attr.(string)
	}
	if attr, ok := d.GetOk("override"); ok {
		o.Override = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceCommandDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Command{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
