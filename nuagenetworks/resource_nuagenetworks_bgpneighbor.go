package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceBGPNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceBGPNeighborCreate,
		Read:   resourceBGPNeighborRead,
		Update: resourceBGPNeighborUpdate,
		Delete: resourceBGPNeighborDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dampening_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"peer_as": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"peer_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_export_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_import_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet", "parent_vlan"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_vlan"},
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_subnet"},
			},
		},
	}
}

func resourceBGPNeighborCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize BGPNeighbor object
	PeerAS := d.Get("peer_as").(int)
	o := &vspk.BGPNeighbor{
		Name:   d.Get("name").(string),
		PeerAS: &PeerAS,
	}
	if attr, ok := d.GetOk("bfd_enabled"); ok {
		BFDEnabled := attr.(bool)
		o.BFDEnabled = &BFDEnabled
	}
	if attr, ok := d.GetOk("dampening_enabled"); ok {
		DampeningEnabled := attr.(bool)
		o.DampeningEnabled = &DampeningEnabled
	}
	if attr, ok := d.GetOk("peer_ip"); ok {
		o.PeerIP = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("session"); ok {
		o.Session = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateBGPNeighbor(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateBGPNeighbor(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		err := parent.CreateBGPNeighbor(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceBGPNeighborRead(d, m)
}

func resourceBGPNeighborRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPNeighbor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("bfd_enabled", o.BFDEnabled)
	d.Set("name", o.Name)
	d.Set("dampening_enabled", o.DampeningEnabled)
	d.Set("peer_as", o.PeerAS)
	d.Set("peer_ip", o.PeerIP)
	d.Set("description", o.Description)
	d.Set("session", o.Session)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_export_routing_policy_id", o.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", o.AssociatedImportRoutingPolicyID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceBGPNeighborUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPNeighbor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	PeerAS := d.Get("peer_as").(int)
	o.PeerAS = &PeerAS

	if attr, ok := d.GetOk("bfd_enabled"); ok {
		BFDEnabled := attr.(bool)
		o.BFDEnabled = &BFDEnabled
	}
	if attr, ok := d.GetOk("dampening_enabled"); ok {
		DampeningEnabled := attr.(bool)
		o.DampeningEnabled = &DampeningEnabled
	}
	if attr, ok := d.GetOk("peer_ip"); ok {
		o.PeerIP = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("session"); ok {
		o.Session = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceBGPNeighborDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPNeighbor{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
