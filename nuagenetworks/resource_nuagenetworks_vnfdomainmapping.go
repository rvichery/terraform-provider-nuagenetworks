package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVNFDomainMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceVNFDomainMappingCreate,
		Read:   resourceVNFDomainMappingRead,
		Update: resourceVNFDomainMappingUpdate,
		Delete: resourceVNFDomainMappingDelete,
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
			"segmentation_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"segmentation_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VLAN",
			},
			"associated_ns_gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_ns_gateway_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVNFDomainMappingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VNFDomainMapping object
	o := &vspk.VNFDomainMapping{}
	if attr, ok := d.GetOk("segmentation_id"); ok {
		SegmentationID := attr.(int)
		o.SegmentationID = &SegmentationID
	}
	if attr, ok := d.GetOk("segmentation_type"); ok {
		o.SegmentationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_id"); ok {
		o.AssociatedNSGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_name"); ok {
		o.AssociatedNSGatewayName = attr.(string)
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateVNFDomainMapping(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceVNFDomainMappingRead(d, m)
}

func resourceVNFDomainMappingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFDomainMapping{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("segmentation_id", o.SegmentationID)
	d.Set("segmentation_type", o.SegmentationType)
	d.Set("associated_ns_gateway_id", o.AssociatedNSGatewayID)
	d.Set("associated_ns_gateway_name", o.AssociatedNSGatewayName)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVNFDomainMappingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFDomainMapping{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("segmentation_id"); ok {
		SegmentationID := attr.(int)
		o.SegmentationID = &SegmentationID
	}
	if attr, ok := d.GetOk("segmentation_type"); ok {
		o.SegmentationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_id"); ok {
		o.AssociatedNSGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_name"); ok {
		o.AssociatedNSGatewayName = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVNFDomainMappingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNFDomainMapping{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
