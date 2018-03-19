package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNSPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSPortRead,
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
			"nat_traversal": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_range": {
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
			"template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"physical_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_nat_probes": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"speed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_through_ubr_only": {
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
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_redundant_port_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_auto_discovered_gateway", "parent_ns_gateway"},
			},
			"parent_auto_discovered_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundant_port", "parent_ns_gateway"},
			},
			"parent_ns_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundant_port", "parent_auto_discovered_gateway"},
			},
		},
	}
}

func dataSourceNSPortRead(d *schema.ResourceData, m interface{}) error {
	filteredNSPorts := vspk.NSPortsList{}
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
	if attr, ok := d.GetOk("parent_redundant_port"); ok {
		parent := &vspk.RedundantPort{ID: attr.(string)}
		filteredNSPorts, err = parent.NSPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_auto_discovered_gateway"); ok {
		parent := &vspk.AutoDiscoveredGateway{ID: attr.(string)}
		filteredNSPorts, err = parent.NSPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredNSPorts, err = parent.NSPorts(fetchFilter)
		if err != nil {
			return err
		}
	}

	NSPort := &vspk.NSPort{}

	if len(filteredNSPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NSPort = filteredNSPorts[0]

	d.Set("nat_traversal", NSPort.NATTraversal)
	d.Set("vlan_range", NSPort.VLANRange)
	d.Set("name", NSPort.Name)
	d.Set("last_updated_by", NSPort.LastUpdatedBy)
	d.Set("template_id", NSPort.TemplateID)
	d.Set("permitted_action", NSPort.PermittedAction)
	d.Set("description", NSPort.Description)
	d.Set("physical_name", NSPort.PhysicalName)
	d.Set("enable_nat_probes", NSPort.EnableNATProbes)
	d.Set("entity_scope", NSPort.EntityScope)
	d.Set("port_type", NSPort.PortType)
	d.Set("speed", NSPort.Speed)
	d.Set("traffic_through_ubr_only", NSPort.TrafficThroughUBROnly)
	d.Set("use_user_mnemonic", NSPort.UseUserMnemonic)
	d.Set("user_mnemonic", NSPort.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", NSPort.AssociatedEgressQOSPolicyID)
	d.Set("associated_redundant_port_id", NSPort.AssociatedRedundantPortID)
	d.Set("status", NSPort.Status)
	d.Set("mtu", NSPort.Mtu)
	d.Set("external_id", NSPort.ExternalID)

	d.Set("id", NSPort.Identifier())
	d.Set("parent_id", NSPort.ParentID)
	d.Set("parent_type", NSPort.ParentType)
	d.Set("owner", NSPort.Owner)

	d.SetId(NSPort.Identifier())

	return nil
}
