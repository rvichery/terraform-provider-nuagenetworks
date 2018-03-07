package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomainRead,
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
			"pat_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ecmp_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bgp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dhcp_behavior": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_server_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_underlay": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"back_haul_route_distinguisher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_route_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"back_haul_subnet_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_subnet_mask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"back_haul_vnid": &schema.Schema{
				Type:     schema.TypeInt,
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
			"advertise_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"leaking_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_dhcp_server_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action": &schema.Schema{
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
			"dhcp_server_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"global_routing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"flow_collection_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"import_route_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"underlay_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"policy_change_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"domain_vlanid": &schema.Schema{
				Type:     schema.TypeInt,
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
			"associated_bgp_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_pat_mapper_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_shared_pat_mapper_id": &schema.Schema{
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
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"export_route_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_firewall_acl": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_domain_template", "parent_enterprise"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_firewall_acl", "parent_domain_template", "parent_enterprise"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_firewall_acl", "parent_domain", "parent_enterprise"},
			},
			"parent_enterprise": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_firewall_acl", "parent_domain", "parent_domain_template"},
			},
		},
	}
}

func dataSourceDomainRead(d *schema.ResourceData, m interface{}) error {
	filteredDomains := vspk.DomainsList{}
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
	if attr, ok := d.GetOk("parent_firewall_acl"); ok {
		parent := &vspk.FirewallAcl{ID: attr.(string)}
		filteredDomains, err = parent.Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredDomains, err = parent.Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredDomains, err = parent.Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredDomains, err = parent.Domains(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredDomains, err = parent.Domains(fetchFilter)
		if err != nil {
			return err
		}
	}

	Domain := &vspk.Domain{}

	if len(filteredDomains) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDomains) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Domain = filteredDomains[0]
	}

	d.Set("pat_enabled", Domain.PATEnabled)
	d.Set("ecmp_count", Domain.ECMPCount)
	d.Set("bgp_enabled", Domain.BGPEnabled)
	d.Set("dhcp_behavior", Domain.DHCPBehavior)
	d.Set("dhcp_server_address", Domain.DHCPServerAddress)
	d.Set("fip_underlay", Domain.FIPUnderlay)
	d.Set("dpi", Domain.DPI)
	d.Set("label_id", Domain.LabelID)
	d.Set("back_haul_route_distinguisher", Domain.BackHaulRouteDistinguisher)
	d.Set("back_haul_route_target", Domain.BackHaulRouteTarget)
	d.Set("back_haul_service_id", Domain.BackHaulServiceID)
	d.Set("back_haul_subnet_ip_address", Domain.BackHaulSubnetIPAddress)
	d.Set("back_haul_subnet_mask", Domain.BackHaulSubnetMask)
	d.Set("back_haul_vnid", Domain.BackHaulVNID)
	d.Set("maintenance_mode", Domain.MaintenanceMode)
	d.Set("name", Domain.Name)
	d.Set("last_updated_by", Domain.LastUpdatedBy)
	d.Set("advertise_criteria", Domain.AdvertiseCriteria)
	d.Set("leaking_enabled", Domain.LeakingEnabled)
	d.Set("secondary_dhcp_server_address", Domain.SecondaryDHCPServerAddress)
	d.Set("template_id", Domain.TemplateID)
	d.Set("permitted_action", Domain.PermittedAction)
	d.Set("service_id", Domain.ServiceID)
	d.Set("description", Domain.Description)
	d.Set("dhcp_server_addresses", Domain.DhcpServerAddresses)
	d.Set("global_routing_enabled", Domain.GlobalRoutingEnabled)
	d.Set("flow_collection_enabled", Domain.FlowCollectionEnabled)
	d.Set("import_route_target", Domain.ImportRouteTarget)
	d.Set("encryption", Domain.Encryption)
	d.Set("underlay_enabled", Domain.UnderlayEnabled)
	d.Set("entity_scope", Domain.EntityScope)
	d.Set("local_as", Domain.LocalAS)
	d.Set("policy_change_status", Domain.PolicyChangeStatus)
	d.Set("domain_id", Domain.DomainID)
	d.Set("domain_vlanid", Domain.DomainVLANID)
	d.Set("route_distinguisher", Domain.RouteDistinguisher)
	d.Set("route_target", Domain.RouteTarget)
	d.Set("uplink_preference", Domain.UplinkPreference)
	d.Set("associated_bgp_profile_id", Domain.AssociatedBGPProfileID)
	d.Set("associated_multicast_channel_map_id", Domain.AssociatedMulticastChannelMapID)
	d.Set("associated_pat_mapper_id", Domain.AssociatedPATMapperID)
	d.Set("associated_shared_pat_mapper_id", Domain.AssociatedSharedPATMapperID)
	d.Set("associated_underlay_id", Domain.AssociatedUnderlayID)
	d.Set("stretched", Domain.Stretched)
	d.Set("multicast", Domain.Multicast)
	d.Set("tunnel_type", Domain.TunnelType)
	d.Set("customer_id", Domain.CustomerID)
	d.Set("export_route_target", Domain.ExportRouteTarget)
	d.Set("external_id", Domain.ExternalID)

	d.Set("id", Domain.Identifier())
	d.Set("parent_id", Domain.ParentID)
	d.Set("parent_type", Domain.ParentType)
	d.Set("owner", Domain.Owner)

	d.SetId(Domain.Identifier())

	return nil
}
