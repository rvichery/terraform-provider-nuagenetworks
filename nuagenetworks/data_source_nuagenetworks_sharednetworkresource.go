package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceSharedNetworkResource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSharedNetworkResourceRead,
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
			"ecmp_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dhcp_managed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"back_haul_route_distinguisher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_route_target": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_vnid": {
				Type:     schema.TypeInt,
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
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_restriction_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared_resource_parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vn_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"underlay": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_route_distinguisher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_route_target": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_gw_vlan_attachment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_interface_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_vport_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_global_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_pat_mapper_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_pat_allocation_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_pat_mapper": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_pat_mapper"},
			},
		},
	}
}

func dataSourceSharedNetworkResourceRead(d *schema.ResourceData, m interface{}) error {
	filteredSharedNetworkResources := vspk.SharedNetworkResourcesList{}
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
	if attr, ok := d.GetOk("parent_pat_mapper"); ok {
		parent := &vspk.PATMapper{ID: attr.(string)}
		filteredSharedNetworkResources, err = parent.SharedNetworkResources(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredSharedNetworkResources, err = parent.SharedNetworkResources(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredSharedNetworkResources, err = parent.SharedNetworkResources(fetchFilter)
		if err != nil {
			return err
		}
	}

	SharedNetworkResource := &vspk.SharedNetworkResource{}

	if len(filteredSharedNetworkResources) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredSharedNetworkResources) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	SharedNetworkResource = filteredSharedNetworkResources[0]

	d.Set("ecmp_count", SharedNetworkResource.ECMPCount)
	d.Set("dhcp_managed", SharedNetworkResource.DHCPManaged)
	d.Set("back_haul_route_distinguisher", SharedNetworkResource.BackHaulRouteDistinguisher)
	d.Set("back_haul_route_target", SharedNetworkResource.BackHaulRouteTarget)
	d.Set("back_haul_vnid", SharedNetworkResource.BackHaulVNID)
	d.Set("name", SharedNetworkResource.Name)
	d.Set("last_updated_by", SharedNetworkResource.LastUpdatedBy)
	d.Set("gateway", SharedNetworkResource.Gateway)
	d.Set("gateway_mac_address", SharedNetworkResource.GatewayMACAddress)
	d.Set("access_restriction_enabled", SharedNetworkResource.AccessRestrictionEnabled)
	d.Set("address", SharedNetworkResource.Address)
	d.Set("permitted_action_type", SharedNetworkResource.PermittedActionType)
	d.Set("description", SharedNetworkResource.Description)
	d.Set("netmask", SharedNetworkResource.Netmask)
	d.Set("shared_resource_parent_id", SharedNetworkResource.SharedResourceParentID)
	d.Set("vn_id", SharedNetworkResource.VnID)
	d.Set("underlay", SharedNetworkResource.Underlay)
	d.Set("entity_scope", SharedNetworkResource.EntityScope)
	d.Set("domain_route_distinguisher", SharedNetworkResource.DomainRouteDistinguisher)
	d.Set("domain_route_target", SharedNetworkResource.DomainRouteTarget)
	d.Set("uplink_gw_vlan_attachment_id", SharedNetworkResource.UplinkGWVlanAttachmentID)
	d.Set("uplink_interface_ip", SharedNetworkResource.UplinkInterfaceIP)
	d.Set("uplink_interface_mac", SharedNetworkResource.UplinkInterfaceMAC)
	d.Set("uplink_vport_name", SharedNetworkResource.UplinkVPortName)
	d.Set("use_global_mac", SharedNetworkResource.UseGlobalMAC)
	d.Set("associated_pat_mapper_id", SharedNetworkResource.AssociatedPATMapperID)
	d.Set("external_id", SharedNetworkResource.ExternalID)
	d.Set("dynamic_pat_allocation_enabled", SharedNetworkResource.DynamicPATAllocationEnabled)
	d.Set("type", SharedNetworkResource.Type)

	d.Set("id", SharedNetworkResource.Identifier())
	d.Set("parent_id", SharedNetworkResource.ParentID)
	d.Set("parent_type", SharedNetworkResource.ParentType)
	d.Set("owner", SharedNetworkResource.Owner)

	d.SetId(SharedNetworkResource.Identifier())

	return nil
}
