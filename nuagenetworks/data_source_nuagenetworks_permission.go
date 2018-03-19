package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePermission() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePermissionRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_entity_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
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

func dataSourcePermissionRead(d *schema.ResourceData, m interface{}) error {
	filteredPermissions := vspk.PermissionsList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_wan_service"); ok {
		parent := &vspk.WANService{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredPermissions, err = parent.Permissions(fetchFilter)
		if err != nil {
			return err
		}
	}

	Permission := &vspk.Permission{}

	if len(filteredPermissions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPermissions) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Permission = filteredPermissions[0]

	d.Set("name", Permission.Name)
	d.Set("last_updated_by", Permission.LastUpdatedBy)
	d.Set("permitted_action", Permission.PermittedAction)
	d.Set("permitted_entity_description", Permission.PermittedEntityDescription)
	d.Set("permitted_entity_id", Permission.PermittedEntityID)
	d.Set("permitted_entity_name", Permission.PermittedEntityName)
	d.Set("permitted_entity_type", Permission.PermittedEntityType)
	d.Set("entity_scope", Permission.EntityScope)
	d.Set("external_id", Permission.ExternalID)

	d.Set("id", Permission.Identifier())
	d.Set("parent_id", Permission.ParentID)
	d.Set("parent_type", Permission.ParentType)
	d.Set("owner", Permission.Owner)

	d.SetId(Permission.Identifier())

	return nil
}
