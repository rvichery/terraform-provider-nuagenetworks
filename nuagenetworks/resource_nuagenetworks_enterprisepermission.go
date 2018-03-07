package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceEnterprisePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnterprisePermissionCreate,
		Read:   resourceEnterprisePermissionRead,
		Update: resourceEnterprisePermissionUpdate,
		Delete: resourceEnterprisePermissionDelete,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permitted_action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"permitted_entity_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_entity_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_entity_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_shared_network_resource": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_redundancy_group": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_vsg_redundant_port": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_wan_service": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_port": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_vlan": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_gateway": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_ns_port": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_patnat_pool"},
			},
			"parent_patnat_pool": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port"},
			},
		},
	}
}

func resourceEnterprisePermissionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EnterprisePermission object
	o := &vspk.EnterprisePermission{
		PermittedAction: d.Get("permitted_action").(string),
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_description"); ok {
		o.PermittedEntityDescription = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_id"); ok {
		o.PermittedEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_name"); ok {
		o.PermittedEntityName = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_type"); ok {
		o.PermittedEntityType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_wan_service"); ok {
		parent := &vspk.WANService{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_patnat_pool"); ok {
		parent := &vspk.PATNATPool{ID: attr.(string)}
		err := parent.CreateEnterprisePermission(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceEnterprisePermissionRead(d, m)
}

func resourceEnterprisePermissionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterprisePermission{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("permitted_entity_description", o.PermittedEntityDescription)
	d.Set("permitted_entity_id", o.PermittedEntityID)
	d.Set("permitted_entity_name", o.PermittedEntityName)
	d.Set("permitted_entity_type", o.PermittedEntityType)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceEnterprisePermissionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterprisePermission{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.PermittedAction = d.Get("permitted_action").(string)

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_description"); ok {
		o.PermittedEntityDescription = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_id"); ok {
		o.PermittedEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_name"); ok {
		o.PermittedEntityName = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_type"); ok {
		o.PermittedEntityType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEnterprisePermissionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterprisePermission{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
