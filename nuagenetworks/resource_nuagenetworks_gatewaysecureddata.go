package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceGatewaySecuredData() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatewaySecuredDataCreate,
		Read:   resourceGatewaySecuredDataRead,
		Update: resourceGatewaySecuredDataUpdate,
		Delete: resourceGatewaySecuredDataDelete,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_cert_serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"keyserver_cert_serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signed_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_gateway_security": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGatewaySecuredDataCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize GatewaySecuredData object
	o := &vspk.GatewaySecuredData{}
	if attr, ok := d.GetOk("data"); ok {
		o.Data = attr.(string)
	}
	if attr, ok := d.GetOk("gateway_cert_serial_number"); ok {
		o.GatewayCertSerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("keyserver_cert_serial_number"); ok {
		o.KeyserverCertSerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("signed_data"); ok {
		o.SignedData = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.GatewaySecurity{ID: d.Get("parent_gateway_security").(string)}
	err := parent.CreateGatewaySecuredData(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceGatewaySecuredDataRead(d, m)
}

func resourceGatewaySecuredDataRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.GatewaySecuredData{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("data", o.Data)
	d.Set("gateway_cert_serial_number", o.GatewayCertSerialNumber)
	d.Set("keyserver_cert_serial_number", o.KeyserverCertSerialNumber)
	d.Set("signed_data", o.SignedData)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceGatewaySecuredDataUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.GatewaySecuredData{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("data"); ok {
		o.Data = attr.(string)
	}
	if attr, ok := d.GetOk("gateway_cert_serial_number"); ok {
		o.GatewayCertSerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("keyserver_cert_serial_number"); ok {
		o.KeyserverCertSerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("signed_data"); ok {
		o.SignedData = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceGatewaySecuredDataDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.GatewaySecuredData{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
