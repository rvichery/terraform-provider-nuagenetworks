package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceIKEGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceIKEGatewayCreate,
		Read:   resourceIKEGatewayRead,
		Update: resourceIKEGatewayUpdate,
		Delete: resourceIKEGatewayDelete,
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
			"ike_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V2",
			},
			"ik_ev1_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIKEGatewayCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IKEGateway object
	o := &vspk.IKEGateway{}
	if attr, ok := d.GetOk("ike_version"); ok {
		o.IKEVersion = attr.(string)
	}
	if attr, ok := d.GetOk("ik_ev1_mode"); ok {
		o.IKEv1Mode = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateIKEGateway(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("ikegatewayconfigs"); ok {
		o.AssignIKEGatewayConfigs(attr.(vspk.IKEGatewayConfigsList))
	}
	return resourceIKEGatewayRead(d, m)
}

func resourceIKEGatewayRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ike_version", o.IKEVersion)
	d.Set("ik_ev1_mode", o.IKEv1Mode)
	d.Set("ip_address", o.IPAddress)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIKEGatewayUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("ike_version"); ok {
		o.IKEVersion = attr.(string)
	}
	if attr, ok := d.GetOk("ik_ev1_mode"); ok {
		o.IKEv1Mode = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIKEGatewayDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGateway{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
