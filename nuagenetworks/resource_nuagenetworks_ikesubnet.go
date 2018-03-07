package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceIKESubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceIKESubnetCreate,
		Read:   resourceIKESubnetRead,
		Update: resourceIKESubnetUpdate,
		Delete: resourceIKESubnetDelete,
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
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_ike_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIKESubnetCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IKESubnet object
	o := &vspk.IKESubnet{}
	if attr, ok := d.GetOk("prefix"); ok {
		o.Prefix = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_id"); ok {
		o.AssociatedIKEGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.IKEGateway{ID: d.Get("parent_ike_gateway").(string)}
	err := parent.CreateIKESubnet(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceIKESubnetRead(d, m)
}

func resourceIKESubnetRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKESubnet{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("prefix", o.Prefix)
	d.Set("associated_ike_gateway_id", o.AssociatedIKEGatewayID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIKESubnetUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKESubnet{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("prefix"); ok {
		o.Prefix = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_id"); ok {
		o.AssociatedIKEGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIKESubnetDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKESubnet{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
