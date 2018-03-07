package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceL2Domain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceL2DomainRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_managed": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_collection_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vn_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_change_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_distinguisher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_preference": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_global_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_shared_network_resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_underlay_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stretched": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_ipv6_address": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_l2_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain_template"},
			},
		},
	}
}

func dataSourceL2DomainRead(d *schema.ResourceData, m interface{}) error {
	filteredL2Domains := vspk.L2DomainsList{}
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
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredL2Domains, err = parent.L2Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredL2Domains, err = parent.L2Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredL2Domains, err = parent.L2Domains(fetchFilter)
		if err != nil {
			return err
		}
	}

	L2Domain := &vspk.L2Domain{}

	if len(filteredL2Domains) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredL2Domains) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		L2Domain = filteredL2Domains[0]
	}

	d.Set("dhcp_managed", L2Domain.DHCPManaged)
	d.Set("dpi", L2Domain.DPI)
	d.Set("ip_type", L2Domain.IPType)
	d.Set("ipv6_address", L2Domain.IPv6Address)
	d.Set("ipv6_gateway", L2Domain.IPv6Gateway)
	d.Set("maintenance_mode", L2Domain.MaintenanceMode)
	d.Set("name", L2Domain.Name)
	d.Set("last_updated_by", L2Domain.LastUpdatedBy)
	d.Set("gateway", L2Domain.Gateway)
	d.Set("gateway_mac_address", L2Domain.GatewayMACAddress)
	d.Set("address", L2Domain.Address)
	d.Set("template_id", L2Domain.TemplateID)
	d.Set("service_id", L2Domain.ServiceID)
	d.Set("description", L2Domain.Description)
	d.Set("netmask", L2Domain.Netmask)
	d.Set("flow_collection_enabled", L2Domain.FlowCollectionEnabled)
	d.Set("vn_id", L2Domain.VnId)
	d.Set("encryption", L2Domain.Encryption)
	d.Set("entity_scope", L2Domain.EntityScope)
	d.Set("entity_state", L2Domain.EntityState)
	d.Set("policy_change_status", L2Domain.PolicyChangeStatus)
	d.Set("route_distinguisher", L2Domain.RouteDistinguisher)
	d.Set("route_target", L2Domain.RouteTarget)
	d.Set("uplink_preference", L2Domain.UplinkPreference)
	d.Set("use_global_mac", L2Domain.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", L2Domain.AssociatedMulticastChannelMapID)
	d.Set("associated_shared_network_resource_id", L2Domain.AssociatedSharedNetworkResourceID)
	d.Set("associated_underlay_id", L2Domain.AssociatedUnderlayID)
	d.Set("stretched", L2Domain.Stretched)
	d.Set("multicast", L2Domain.Multicast)
	d.Set("external_id", L2Domain.ExternalID)
	d.Set("dynamic_ipv6_address", L2Domain.DynamicIpv6Address)

	d.Set("id", L2Domain.Identifier())
	d.Set("parent_id", L2Domain.ParentID)
	d.Set("parent_type", L2Domain.ParentType)
	d.Set("owner", L2Domain.Owner)

	d.SetId(L2Domain.Identifier())

	return nil
}
