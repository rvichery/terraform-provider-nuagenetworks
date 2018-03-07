package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceNATMapEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceNATMapEntryCreate,
		Read:   resourceNATMapEntryRead,
		Update: resourceNATMapEntryUpdate,
		Delete: resourceNATMapEntryDelete,
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
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"associated_patnat_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_patnat_pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNATMapEntryCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NATMapEntry object
	o := &vspk.NATMapEntry{
		PrivateIP: d.Get("private_ip").(string),
		PublicIP:  d.Get("public_ip").(string),
		Type:      d.Get("type").(string),
	}
	if attr, ok := d.GetOk("associated_patnat_pool_id"); ok {
		o.AssociatedPATNATPoolID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.PATNATPool{ID: d.Get("parent_patnat_pool").(string)}
	err := parent.CreateNATMapEntry(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceNATMapEntryRead(d, m)
}

func resourceNATMapEntryRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NATMapEntry{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("private_ip", o.PrivateIP)
	d.Set("associated_patnat_pool_id", o.AssociatedPATNATPoolID)
	d.Set("public_ip", o.PublicIP)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNATMapEntryUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NATMapEntry{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.PrivateIP = d.Get("private_ip").(string)
	o.PublicIP = d.Get("public_ip").(string)
	o.Type = d.Get("type").(string)

	if attr, ok := d.GetOk("associated_patnat_pool_id"); ok {
		o.AssociatedPATNATPoolID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNATMapEntryDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NATMapEntry{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
