package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceL4ServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceL4ServiceGroupCreate,
		Read:   resourceL4ServiceGroupRead,
		Update: resourceL4ServiceGroupUpdate,
		Delete: resourceL4ServiceGroupDelete,
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
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
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

func resourceL4ServiceGroupCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize L4ServiceGroup object
	o := &vspk.L4ServiceGroup{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateL4ServiceGroup(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("l4services"); ok {
		o.AssignL4Services(attr.(vspk.L4ServicesList))
	}
	return resourceL4ServiceGroupRead(d, m)
}

func resourceL4ServiceGroupRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4ServiceGroup{
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
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceL4ServiceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4ServiceGroup{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceL4ServiceGroupDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4ServiceGroup{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
