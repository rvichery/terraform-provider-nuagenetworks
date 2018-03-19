package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceBRConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceBRConnectionCreate,
		Read:   resourceBRConnectionRead,
		Update: resourceBRConnectionUpdate,
		Delete: resourceBRConnectionDelete,
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
			"dns_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"advertisement_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uplink_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vlan_template"},
			},
			"parent_vlan_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vlan"},
			},
		},
	}
}

func resourceBRConnectionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize BRConnection object
	o := &vspk.BRConnection{}
	if attr, ok := d.GetOk("dns_address"); ok {
		o.DNSAddress = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertisement_criteria"); ok {
		o.AdvertisementCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("mode"); ok {
		o.Mode = attr.(string)
	}
	if attr, ok := d.GetOk("uplink_id"); ok {
		o.UplinkID = attr.(int)
	}
	if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		err := parent.CreateBRConnection(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vlan_template"); ok {
		parent := &vspk.VLANTemplate{ID: attr.(string)}
		err := parent.CreateBRConnection(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceBRConnectionRead(d, m)
}

func resourceBRConnectionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BRConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dns_address", o.DNSAddress)
	d.Set("gateway", o.Gateway)
	d.Set("address", o.Address)
	d.Set("advertisement_criteria", o.AdvertisementCriteria)
	d.Set("netmask", o.Netmask)
	d.Set("mode", o.Mode)
	d.Set("uplink_id", o.UplinkID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceBRConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BRConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("dns_address"); ok {
		o.DNSAddress = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertisement_criteria"); ok {
		o.AdvertisementCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("mode"); ok {
		o.Mode = attr.(string)
	}
	if attr, ok := d.GetOk("uplink_id"); ok {
		o.UplinkID = attr.(int)
	}

	o.Save()

	return nil
}

func resourceBRConnectionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BRConnection{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
