package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDSCPForwardingClassTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceDSCPForwardingClassTableCreate,
		Read:   resourceDSCPForwardingClassTableRead,
		Update: resourceDSCPForwardingClassTableUpdate,
		Delete: resourceDSCPForwardingClassTableDelete,
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
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDSCPForwardingClassTableCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DSCPForwardingClassTable object
	o := &vspk.DSCPForwardingClassTable{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateDSCPForwardingClassTable(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceDSCPForwardingClassTableRead(d, m)
}

func resourceDSCPForwardingClassTableRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPForwardingClassTable{
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

func resourceDSCPForwardingClassTableUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPForwardingClassTable{
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

func resourceDSCPForwardingClassTableDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DSCPForwardingClassTable{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
