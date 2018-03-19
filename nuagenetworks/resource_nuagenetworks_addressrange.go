package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceAddressRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceAddressRangeCreate,
		Read:   resourceAddressRangeRead,
		Update: resourceAddressRangeUpdate,
		Delete: resourceAddressRangeDelete,
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
			"dhcp_pool_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_address": {
				Type:     schema.TypeString,
				Required: true,
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
			"parent_shared_network_resource": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet_template", "parent_subnet", "parent_l2_domain_template"},
			},
			"parent_subnet_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet", "parent_l2_domain_template"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet_template", "parent_l2_domain_template"},
			},
			"parent_l2_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet_template", "parent_subnet"},
			},
		},
	}
}

func resourceAddressRangeCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize AddressRange object
	o := &vspk.AddressRange{
		MaxAddress: d.Get("max_address").(string),
		MinAddress: d.Get("min_address").(string),
	}
	if attr, ok := d.GetOk("dhcp_pool_type"); ok {
		o.DHCPPoolType = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		err := parent.CreateAddressRange(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet_template"); ok {
		parent := &vspk.SubnetTemplate{ID: attr.(string)}
		err := parent.CreateAddressRange(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateAddressRange(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		err := parent.CreateAddressRange(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceAddressRangeRead(d, m)
}

func resourceAddressRangeRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.AddressRange{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dhcp_pool_type", o.DHCPPoolType)
	d.Set("ip_type", o.IPType)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("max_address", o.MaxAddress)
	d.Set("min_address", o.MinAddress)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceAddressRangeUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.AddressRange{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.MaxAddress = d.Get("max_address").(string)
	o.MinAddress = d.Get("min_address").(string)

	if attr, ok := d.GetOk("dhcp_pool_type"); ok {
		o.DHCPPoolType = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceAddressRangeDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.AddressRange{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
