package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceL4Service() *schema.Resource {
	return &schema.Resource{
		Create: resourceL4ServiceCreate,
		Read:   resourceL4ServiceRead,
		Update: resourceL4ServiceUpdate,
		Delete: resourceL4ServiceDelete,

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
			"default_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"ports": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceL4ServiceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize L4Service object
	o := &vspk.L4Service{
		Name:     d.Get("name").(string),
		Ports:    d.Get("ports").(string),
		Protocol: d.Get("protocol").(string),
	}
	if attr, ok := d.GetOk("default_service"); ok {
		o.DefaultService = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateL4Service(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateL4Service(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceL4ServiceRead(d, m)
}

func resourceL4ServiceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4Service{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("default_service", o.DefaultService)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("ports", o.Ports)
	d.Set("protocol", o.Protocol)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceL4ServiceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4Service{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Ports = d.Get("ports").(string)
	o.Protocol = d.Get("protocol").(string)

	if attr, ok := d.GetOk("default_service"); ok {
		o.DefaultService = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceL4ServiceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L4Service{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
