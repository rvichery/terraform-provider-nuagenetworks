package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceZFBRequest() *schema.Resource {
	return &schema.Resource{
		Create: resourceZFBRequestCreate,
		Read:   resourceZFBRequestRead,
		Update: resourceZFBRequestUpdate,
		Delete: resourceZFBRequestDelete,
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
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zfb_approval_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zfb_bootstrap_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"zfb_info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zfb_request_retry_timer": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sku": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_connected_time": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_enterprise_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"status_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZFBRequestCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ZFBRequest object
	o := &vspk.ZFBRequest{}
	if attr, ok := d.GetOk("mac_address"); ok {
		o.MACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_approval_status"); ok {
		o.ZFBApprovalStatus = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_bootstrap_enabled"); ok {
		ZFBBootstrapEnabled := attr.(bool)
		o.ZFBBootstrapEnabled = &ZFBBootstrapEnabled
	}
	if attr, ok := d.GetOk("zfb_info"); ok {
		o.ZFBInfo = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_request_retry_timer"); ok {
		ZFBRequestRetryTimer := attr.(int)
		o.ZFBRequestRetryTimer = &ZFBRequestRetryTimer
	}
	if attr, ok := d.GetOk("sku"); ok {
		o.SKU = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_type"); ok {
		o.CPUType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_version"); ok {
		o.NSGVersion = attr.(string)
	}
	if attr, ok := d.GetOk("uuid"); ok {
		o.UUID = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("hostname"); ok {
		o.Hostname = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_name"); ok {
		o.AssociatedEnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_id"); ok {
		o.AssociatedNSGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_name"); ok {
		o.AssociatedNSGatewayName = attr.(string)
	}
	if attr, ok := d.GetOk("status_string"); ok {
		o.StatusString = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateZFBRequest(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateZFBRequest(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceZFBRequestRead(d, m)
}

func resourceZFBRequestRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBRequest{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac_address", o.MACAddress)
	d.Set("zfb_approval_status", o.ZFBApprovalStatus)
	d.Set("zfb_bootstrap_enabled", o.ZFBBootstrapEnabled)
	d.Set("zfb_info", o.ZFBInfo)
	d.Set("zfb_request_retry_timer", o.ZFBRequestRetryTimer)
	d.Set("sku", o.SKU)
	d.Set("ip_address", o.IPAddress)
	d.Set("cpu_type", o.CPUType)
	d.Set("nsg_version", o.NSGVersion)
	d.Set("uuid", o.UUID)
	d.Set("family", o.Family)
	d.Set("last_connected_time", o.LastConnectedTime)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("serial_number", o.SerialNumber)
	d.Set("entity_scope", o.EntityScope)
	d.Set("hostname", o.Hostname)
	d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
	d.Set("associated_enterprise_name", o.AssociatedEnterpriseName)
	d.Set("associated_ns_gateway_id", o.AssociatedNSGatewayID)
	d.Set("associated_ns_gateway_name", o.AssociatedNSGatewayName)
	d.Set("status_string", o.StatusString)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceZFBRequestUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBRequest{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("mac_address"); ok {
		o.MACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_approval_status"); ok {
		o.ZFBApprovalStatus = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_bootstrap_enabled"); ok {
		ZFBBootstrapEnabled := attr.(bool)
		o.ZFBBootstrapEnabled = &ZFBBootstrapEnabled
	}
	if attr, ok := d.GetOk("zfb_info"); ok {
		o.ZFBInfo = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_request_retry_timer"); ok {
		ZFBRequestRetryTimer := attr.(int)
		o.ZFBRequestRetryTimer = &ZFBRequestRetryTimer
	}
	if attr, ok := d.GetOk("sku"); ok {
		o.SKU = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_type"); ok {
		o.CPUType = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_version"); ok {
		o.NSGVersion = attr.(string)
	}
	if attr, ok := d.GetOk("uuid"); ok {
		o.UUID = attr.(string)
	}
	if attr, ok := d.GetOk("family"); ok {
		o.Family = attr.(string)
	}
	if attr, ok := d.GetOk("serial_number"); ok {
		o.SerialNumber = attr.(string)
	}
	if attr, ok := d.GetOk("hostname"); ok {
		o.Hostname = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_name"); ok {
		o.AssociatedEnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_id"); ok {
		o.AssociatedNSGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ns_gateway_name"); ok {
		o.AssociatedNSGatewayName = attr.(string)
	}
	if attr, ok := d.GetOk("status_string"); ok {
		o.StatusString = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceZFBRequestDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBRequest{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
