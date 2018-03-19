package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVPortRead,
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
			"vlanid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dpi": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_attached_interfaces": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mac_move_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"address_spoofing": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segmentation_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"segmentation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_collection_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operational_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trunk_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_floating_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ssid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_send_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_trunk_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sub_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_nic_vport_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_overlay_mirror_destination": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_redirection_target": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_multi_nic_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_vrs": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_trunk", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_trunk": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_l2_domain", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_floating_ip", "parent_policy_group"},
			},
			"parent_floating_ip": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_policy_group"},
			},
			"parent_policy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_overlay_mirror_destination", "parent_zone", "parent_redirection_target", "parent_domain", "parent_subnet", "parent_multi_nic_vport", "parent_vrs", "parent_trunk", "parent_l2_domain", "parent_floating_ip"},
			},
		},
	}
}

func dataSourceVPortRead(d *schema.ResourceData, m interface{}) error {
	filteredVPorts := vspk.VPortsList{}
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
	if attr, ok := d.GetOk("parent_overlay_mirror_destination"); ok {
		parent := &vspk.OverlayMirrorDestination{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_redirection_target"); ok {
		parent := &vspk.RedirectionTarget{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_multi_nic_vport"); ok {
		parent := &vspk.MultiNICVPort{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vrs"); ok {
		parent := &vspk.VRS{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_trunk"); ok {
		parent := &vspk.Trunk{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_floating_ip"); ok {
		parent := &vspk.FloatingIp{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_policy_group"); ok {
		parent := &vspk.PolicyGroup{ID: attr.(string)}
		filteredVPorts, err = parent.VPorts(fetchFilter)
		if err != nil {
			return err
		}
	}

	VPort := &vspk.VPort{}

	if len(filteredVPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VPort = filteredVPorts[0]

	d.Set("vlanid", VPort.VLANID)
	d.Set("dpi", VPort.DPI)
	d.Set("name", VPort.Name)
	d.Set("has_attached_interfaces", VPort.HasAttachedInterfaces)
	d.Set("last_updated_by", VPort.LastUpdatedBy)
	d.Set("gateway_mac_move_role", VPort.GatewayMACMoveRole)
	d.Set("active", VPort.Active)
	d.Set("address_spoofing", VPort.AddressSpoofing)
	d.Set("segmentation_id", VPort.SegmentationID)
	d.Set("segmentation_type", VPort.SegmentationType)
	d.Set("description", VPort.Description)
	d.Set("flow_collection_enabled", VPort.FlowCollectionEnabled)
	d.Set("entity_scope", VPort.EntityScope)
	d.Set("domain_id", VPort.DomainID)
	d.Set("zone_id", VPort.ZoneID)
	d.Set("operational_state", VPort.OperationalState)
	d.Set("trunk_role", VPort.TrunkRole)
	d.Set("associated_floating_ip_id", VPort.AssociatedFloatingIPID)
	d.Set("associated_multicast_channel_map_id", VPort.AssociatedMulticastChannelMapID)
	d.Set("associated_ssid", VPort.AssociatedSSID)
	d.Set("associated_send_multicast_channel_map_id", VPort.AssociatedSendMulticastChannelMapID)
	d.Set("associated_trunk_id", VPort.AssociatedTrunkID)
	d.Set("sub_type", VPort.SubType)
	d.Set("multi_nic_vport_id", VPort.MultiNICVPortID)
	d.Set("multicast", VPort.Multicast)
	d.Set("external_id", VPort.ExternalID)
	d.Set("type", VPort.Type)
	d.Set("system_type", VPort.SystemType)

	d.Set("id", VPort.Identifier())
	d.Set("parent_id", VPort.ParentID)
	d.Set("parent_type", VPort.ParentType)
	d.Set("owner", VPort.Owner)

	d.SetId(VPort.Identifier())

	return nil
}
