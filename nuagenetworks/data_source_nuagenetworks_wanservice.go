package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceWANService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWANServiceRead,
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
			"wan_service_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"irb_enabled": {
				Type:     schema.TypeBool,
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
			"service_policy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vn_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"orphan": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_user_mnemonic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_mnemonic": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_vpn_connect_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_route_target": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_redundancy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_auto_discovered_gateway", "parent_gateway"},
			},
			"parent_auto_discovered_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_gateway"},
			},
			"parent_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_auto_discovered_gateway"},
			},
		},
	}
}

func dataSourceWANServiceRead(d *schema.ResourceData, m interface{}) error {
	filteredWANServices := vspk.WANServicesList{}
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
		filteredWANServices, err = parent.WANServices(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_auto_discovered_gateway"); ok {
		parent := &vspk.AutoDiscoveredGateway{ID: attr.(string)}
		filteredWANServices, err = parent.WANServices(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredWANServices, err = parent.WANServices(fetchFilter)
		if err != nil {
			return err
		}
	}

	WANService := &vspk.WANService{}

	if len(filteredWANServices) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredWANServices) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	WANService = filteredWANServices[0]

	d.Set("wan_service_identifier", WANService.WANServiceIdentifier)
	d.Set("irb_enabled", WANService.IRBEnabled)
	d.Set("name", WANService.Name)
	d.Set("last_updated_by", WANService.LastUpdatedBy)
	d.Set("permitted_action", WANService.PermittedAction)
	d.Set("service_policy", WANService.ServicePolicy)
	d.Set("service_type", WANService.ServiceType)
	d.Set("description", WANService.Description)
	d.Set("vn_id", WANService.VnId)
	d.Set("enterprise_name", WANService.EnterpriseName)
	d.Set("entity_scope", WANService.EntityScope)
	d.Set("domain_name", WANService.DomainName)
	d.Set("config_type", WANService.ConfigType)
	d.Set("orphan", WANService.Orphan)
	d.Set("use_user_mnemonic", WANService.UseUserMnemonic)
	d.Set("user_mnemonic", WANService.UserMnemonic)
	d.Set("associated_domain_id", WANService.AssociatedDomainID)
	d.Set("associated_vpn_connect_id", WANService.AssociatedVPNConnectID)
	d.Set("tunnel_type", WANService.TunnelType)
	d.Set("external_id", WANService.ExternalID)
	d.Set("external_route_target", WANService.ExternalRouteTarget)

	d.Set("id", WANService.Identifier())
	d.Set("parent_id", WANService.ParentID)
	d.Set("parent_type", WANService.ParentType)
	d.Set("owner", WANService.Owner)

	d.SetId(WANService.Identifier())

	return nil
}
