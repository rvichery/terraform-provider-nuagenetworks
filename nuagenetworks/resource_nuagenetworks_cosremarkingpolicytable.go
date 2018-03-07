package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceCOSRemarkingPolicyTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceCOSRemarkingPolicyTableCreate,
		Read:   resourceCOSRemarkingPolicyTableRead,
		Update: resourceCOSRemarkingPolicyTableUpdate,
		Delete: resourceCOSRemarkingPolicyTableDelete,
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
				Optional: true,
			},
		},
	}
}

func resourceCOSRemarkingPolicyTableCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize COSRemarkingPolicyTable object
	o := &vspk.COSRemarkingPolicyTable{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateCOSRemarkingPolicyTable(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateCOSRemarkingPolicyTable(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceCOSRemarkingPolicyTableRead(d, m)
}

func resourceCOSRemarkingPolicyTableRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.COSRemarkingPolicyTable{
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

func resourceCOSRemarkingPolicyTableUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.COSRemarkingPolicyTable{
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

func resourceCOSRemarkingPolicyTableDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.COSRemarkingPolicyTable{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
