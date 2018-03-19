package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceStaticRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceStaticRouteCreate,
		Read:   resourceStaticRouteRead,
		Update: resourceStaticRouteUpdate,
		Delete: resourceStaticRouteDelete,
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
			"bfd_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
			},
			"next_hop_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_distinguisher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_shared_network_resource": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource"},
			},
		},
	}
}

func resourceStaticRouteCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize StaticRoute object
	o := &vspk.StaticRoute{
		Address:   d.Get("address").(string),
		Netmask:   d.Get("netmask").(string),
		NextHopIp: d.Get("next_hop_ip").(string),
	}
	if attr, ok := d.GetOk("bfd_enabled"); ok {
		o.BFDEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("associated_subnet_id"); ok {
		o.AssociatedSubnetID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		err := parent.CreateStaticRoute(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateStaticRoute(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceStaticRouteRead(d, m)
}

func resourceStaticRouteRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StaticRoute{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("bfd_enabled", o.BFDEnabled)
	d.Set("ip_type", o.IPType)
	d.Set("ipv6_address", o.IPv6Address)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("address", o.Address)
	d.Set("netmask", o.Netmask)
	d.Set("next_hop_ip", o.NextHopIp)
	d.Set("entity_scope", o.EntityScope)
	d.Set("route_distinguisher", o.RouteDistinguisher)
	d.Set("associated_subnet_id", o.AssociatedSubnetID)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceStaticRouteUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StaticRoute{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Address = d.Get("address").(string)
	o.Netmask = d.Get("netmask").(string)
	o.NextHopIp = d.Get("next_hop_ip").(string)

	if attr, ok := d.GetOk("bfd_enabled"); ok {
		o.BFDEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("associated_subnet_id"); ok {
		o.AssociatedSubnetID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}

	o.Save()

	return nil
}

func resourceStaticRouteDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StaticRoute{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
