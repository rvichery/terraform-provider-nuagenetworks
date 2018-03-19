package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEnterprisePermission() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnterprisePermissionRead,
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
			"parent_shared_network_resource": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_redundancy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_vsg_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_wan_service": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_gateway", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_ns_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_ns_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_ns_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_patnat_pool"},
			},
			"parent_patnat_pool": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_redundancy_group", "parent_vsg_redundant_port", "parent_wan_service", "parent_port", "parent_subnet", "parent_vlan", "parent_gateway", "parent_ns_gateway", "parent_ns_port"},
			},
		},
	}
}

func dataSourceEnterprisePermissionRead(d *schema.ResourceData, m interface{}) error {
	filteredEnterprisePermissions := vspk.EnterprisePermissionsList{}
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
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_wan_service"); ok {
		parent := &vspk.WANService{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_patnat_pool"); ok {
		parent := &vspk.PATNATPool{ID: attr.(string)}
		filteredEnterprisePermissions, err = parent.EnterprisePermissions(fetchFilter)
		if err != nil {
			return err
		}
	}

	EnterprisePermission := &vspk.EnterprisePermission{}

	if len(filteredEnterprisePermissions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEnterprisePermissions) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EnterprisePermission = filteredEnterprisePermissions[0]

	d.Set("name", EnterprisePermission.Name)
	d.Set("last_updated_by", EnterprisePermission.LastUpdatedBy)
	d.Set("permitted_action", EnterprisePermission.PermittedAction)
	d.Set("permitted_entity_description", EnterprisePermission.PermittedEntityDescription)
	d.Set("permitted_entity_id", EnterprisePermission.PermittedEntityID)
	d.Set("permitted_entity_name", EnterprisePermission.PermittedEntityName)
	d.Set("permitted_entity_type", EnterprisePermission.PermittedEntityType)
	d.Set("entity_scope", EnterprisePermission.EntityScope)
	d.Set("external_id", EnterprisePermission.ExternalID)

	d.Set("id", EnterprisePermission.Identifier())
	d.Set("parent_id", EnterprisePermission.ParentID)
	d.Set("parent_type", EnterprisePermission.ParentType)
	d.Set("owner", EnterprisePermission.Owner)

	d.SetId(EnterprisePermission.Identifier())

	return nil
}
