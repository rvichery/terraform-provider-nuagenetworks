package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerCreate,
		Read:   resourceContainerRead,
		Update: resourceContainerUpdate,
		Delete: resourceContainerDelete,
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
			"l2_domain_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vrsid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reason_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_expiry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delete_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resync_info": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"resync_info_raw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"site_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"zone_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"orchestration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceContainerCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Container object
	o := &vspk.Container{
		UUID:            d.Get("uuid").(string),
		Name:            d.Get("name").(string),
		OrchestrationID: d.Get("orchestration_id").(string),
	}
	if attr, ok := d.GetOk("l2_domain_ids"); ok {
		o.L2DomainIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("vrsid"); ok {
		o.VRSID = attr.(string)
	}
	if attr, ok := d.GetOk("reason_type"); ok {
		o.ReasonType = attr.(string)
	}
	if attr, ok := d.GetOk("delete_expiry"); ok {
		DeleteExpiry := attr.(int)
		o.DeleteExpiry = &DeleteExpiry
	}
	if attr, ok := d.GetOk("delete_mode"); ok {
		o.DeleteMode = attr.(string)
	}
	if attr, ok := d.GetOk("resync_info"); ok {
		o.ResyncInfo = attr.(interface{})
	}
	if attr, ok := d.GetOk("site_identifier"); ok {
		o.SiteIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("image_id"); ok {
		o.ImageID = attr.(string)
	}
	if attr, ok := d.GetOk("image_name"); ok {
		o.ImageName = attr.(string)
	}
	if attr, ok := d.GetOk("interfaces"); ok {
		o.Interfaces = attr.([]interface{})
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_name"); ok {
		o.EnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("domain_ids"); ok {
		o.DomainIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("zone_ids"); ok {
		o.ZoneIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("user_id"); ok {
		o.UserID = attr.(string)
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("subnet_ids"); ok {
		o.SubnetIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("hypervisor_ip"); ok {
		o.HypervisorIP = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateContainer(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceContainerRead(d, m)
}

func resourceContainerRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Container{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("l2_domain_ids", o.L2DomainIDs)
	d.Set("vrsid", o.VRSID)
	d.Set("uuid", o.UUID)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("reason_type", o.ReasonType)
	d.Set("delete_expiry", o.DeleteExpiry)
	d.Set("delete_mode", o.DeleteMode)
	if v, ok := o.ResyncInfo.(string); ok {
		raw := make(map[string]string)
		raw["raw"] = v
		d.Set("resync_info_raw", raw)
	} else {
		d.Set("resync_info", o.ResyncInfo)
	}
	d.Set("site_identifier", o.SiteIdentifier)
	d.Set("image_id", o.ImageID)
	d.Set("image_name", o.ImageName)
	d.Set("interfaces", o.Interfaces)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("enterprise_name", o.EnterpriseName)
	d.Set("entity_scope", o.EntityScope)
	d.Set("domain_ids", o.DomainIDs)
	d.Set("zone_ids", o.ZoneIDs)
	d.Set("orchestration_id", o.OrchestrationID)
	d.Set("user_id", o.UserID)
	d.Set("user_name", o.UserName)
	d.Set("status", o.Status)
	d.Set("subnet_ids", o.SubnetIDs)
	d.Set("external_id", o.ExternalID)
	d.Set("hypervisor_ip", o.HypervisorIP)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceContainerUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Container{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.UUID = d.Get("uuid").(string)
	o.Name = d.Get("name").(string)
	o.OrchestrationID = d.Get("orchestration_id").(string)

	if attr, ok := d.GetOk("l2_domain_ids"); ok {
		o.L2DomainIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("vrsid"); ok {
		o.VRSID = attr.(string)
	}
	if attr, ok := d.GetOk("reason_type"); ok {
		o.ReasonType = attr.(string)
	}
	if attr, ok := d.GetOk("delete_expiry"); ok {
		DeleteExpiry := attr.(int)
		o.DeleteExpiry = &DeleteExpiry
	}
	if attr, ok := d.GetOk("delete_mode"); ok {
		o.DeleteMode = attr.(string)
	}
	if attr, ok := d.GetOk("resync_info"); ok {
		o.ResyncInfo = attr.(interface{})
	}
	if attr, ok := d.GetOk("site_identifier"); ok {
		o.SiteIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("image_id"); ok {
		o.ImageID = attr.(string)
	}
	if attr, ok := d.GetOk("image_name"); ok {
		o.ImageName = attr.(string)
	}
	if attr, ok := d.GetOk("interfaces"); ok {
		o.Interfaces = attr.([]interface{})
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_name"); ok {
		o.EnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("domain_ids"); ok {
		o.DomainIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("zone_ids"); ok {
		o.ZoneIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("user_id"); ok {
		o.UserID = attr.(string)
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("subnet_ids"); ok {
		o.SubnetIDs = attr.([]interface{})
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("hypervisor_ip"); ok {
		o.HypervisorIP = attr.(string)
	}

	o.Save()

	return nil
}

func resourceContainerDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Container{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
