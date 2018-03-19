package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceSPATSourcesPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceSPATSourcesPoolCreate,
		Read:   resourceSPATSourcesPoolRead,
		Update: resourceSPATSourcesPoolUpdate,
		Delete: resourceSPATSourcesPoolDelete,
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
				Default:  "IPV4",
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSPATSourcesPoolCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize SPATSourcesPool object
	o := &vspk.SPATSourcesPool{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("address_list"); ok {
		o.AddressList = attr.([]interface{})
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateSPATSourcesPool(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceSPATSourcesPoolRead(d, m)
}

func resourceSPATSourcesPoolRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SPATSourcesPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("family", o.Family)
	d.Set("address_list", o.AddressList)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceSPATSourcesPoolUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SPATSourcesPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("address_list"); ok {
		o.AddressList = attr.([]interface{})
	}

	o.Save()

	return nil
}

func resourceSPATSourcesPoolDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SPATSourcesPool{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
