package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourcePTranslationMap() *schema.Resource {
	return &schema.Resource{
		Create: resourcePTranslationMapCreate,
		Read:   resourcePTranslationMapRead,
		Update: resourcePTranslationMapUpdate,
		Delete: resourcePTranslationMapDelete,

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
			"spat_source_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mapping_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_alias_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_psnat_pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePTranslationMapCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PTranslationMap object
	o := &vspk.PTranslationMap{
		MappingType:     d.Get("mapping_type").(string),
		ProviderAliasIP: d.Get("provider_alias_ip").(string),
		ProviderIP:      d.Get("provider_ip").(string),
	}
	if attr, ok := d.GetOk("spat_source_list"); ok {
		o.SPATSourceList = attr.([]interface{})
	}
	parent := &vspk.PSNATPool{ID: d.Get("parent_psnat_pool").(string)}
	err := parent.CreatePTranslationMap(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourcePTranslationMapRead(d, m)
}

func resourcePTranslationMapRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PTranslationMap{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("spat_source_list", o.SPATSourceList)
	d.Set("mapping_type", o.MappingType)
	d.Set("provider_alias_ip", o.ProviderAliasIP)
	d.Set("provider_ip", o.ProviderIP)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePTranslationMapUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PTranslationMap{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.MappingType = d.Get("mapping_type").(string)
	o.ProviderAliasIP = d.Get("provider_alias_ip").(string)
	o.ProviderIP = d.Get("provider_ip").(string)

	if attr, ok := d.GetOk("spat_source_list"); ok {
		o.SPATSourceList = attr.([]interface{})
	}

	o.Save()

	return nil
}

func resourcePTranslationMapDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PTranslationMap{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
