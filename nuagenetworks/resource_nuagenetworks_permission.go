package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourcePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourcePermissionCreate,
		Read:   resourcePermissionRead,
		Update: resourcePermissionUpdate,
		Delete: resourcePermissionDelete,
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permitted_entity_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_entity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permitted_entity_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permitted_entity_type": {
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
			"parent_redundancy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_vsg_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_wan_service": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_l2_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_domain_template"},
			},
			"parent_ns_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_port", "parent_domain_template"},
			},
			"parent_ns_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_zone", "parent_domain", "parent_wan_service", "parent_port", "parent_l2_domain_template", "parent_vlan", "parent_gateway", "parent_l2_domain", "parent_ns_gateway", "parent_ns_port"},
			},
		},
	}
}

func resourcePermissionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Permission object
	o := &vspk.Permission{
		PermittedAction:   d.Get("permitted_action").(string),
		PermittedEntityID: d.Get("permitted_entity_id").(string),
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_description"); ok {
		o.PermittedEntityDescription = attr.(string)
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
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_wan_service"); ok {
		parent := &vspk.WANService{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		err := parent.CreatePermission(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourcePermissionRead(d, m)
}

func resourcePermissionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Permission{
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

func resourcePermissionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Permission{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.PermittedAction = d.Get("permitted_action").(string)
	o.PermittedEntityID = d.Get("permitted_entity_id").(string)

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("permitted_entity_description"); ok {
		o.PermittedEntityDescription = attr.(string)
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

func resourcePermissionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Permission{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
