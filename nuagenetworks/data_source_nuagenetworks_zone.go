package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceZoneRead,
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
			"dpi": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": {
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
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": {
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
			"flow_collection_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_zone": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"number_of_hosts_in_subnets": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_ipv6_address": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceZoneRead(d *schema.ResourceData, m interface{}) error {
	filteredZones := vspk.ZonesList{}
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
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredZones, err = parent.Zones(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredZones, err = parent.Zones(fetchFilter)
		if err != nil {
			return err
		}
	}

	Zone := &vspk.Zone{}

	if len(filteredZones) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredZones) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Zone = filteredZones[0]

	d.Set("dpi", Zone.DPI)
	d.Set("ip_type", Zone.IPType)
	d.Set("ipv6_address", Zone.IPv6Address)
	d.Set("maintenance_mode", Zone.MaintenanceMode)
	d.Set("name", Zone.Name)
	d.Set("last_updated_by", Zone.LastUpdatedBy)
	d.Set("address", Zone.Address)
	d.Set("template_id", Zone.TemplateID)
	d.Set("description", Zone.Description)
	d.Set("netmask", Zone.Netmask)
	d.Set("flow_collection_enabled", Zone.FlowCollectionEnabled)
	d.Set("encryption", Zone.Encryption)
	d.Set("entity_scope", Zone.EntityScope)
	d.Set("policy_group_id", Zone.PolicyGroupID)
	d.Set("associated_multicast_channel_map_id", Zone.AssociatedMulticastChannelMapID)
	d.Set("public_zone", Zone.PublicZone)
	d.Set("multicast", Zone.Multicast)
	d.Set("number_of_hosts_in_subnets", Zone.NumberOfHostsInSubnets)
	d.Set("external_id", Zone.ExternalID)
	d.Set("dynamic_ipv6_address", Zone.DynamicIpv6Address)

	d.Set("id", Zone.Identifier())
	d.Set("parent_id", Zone.ParentID)
	d.Set("parent_type", Zone.ParentType)
	d.Set("owner", Zone.Owner)

	d.SetId(Zone.Identifier())

	return nil
}
