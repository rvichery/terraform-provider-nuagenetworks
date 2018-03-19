package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceBGPNeighbor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBGPNeighborRead,
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
			"bfd_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dampening_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"peer_as": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"peer_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"session": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_export_routing_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_import_routing_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet", "parent_vlan"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_vlan"},
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_subnet"},
			},
		},
	}
}

func dataSourceBGPNeighborRead(d *schema.ResourceData, m interface{}) error {
	filteredBGPNeighbors := vspk.BGPNeighborsList{}
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
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
		if err != nil {
			return err
		}
	}

	BGPNeighbor := &vspk.BGPNeighbor{}

	if len(filteredBGPNeighbors) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredBGPNeighbors) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	BGPNeighbor = filteredBGPNeighbors[0]

	d.Set("bfd_enabled", BGPNeighbor.BFDEnabled)
	d.Set("name", BGPNeighbor.Name)
	d.Set("dampening_enabled", BGPNeighbor.DampeningEnabled)
	d.Set("peer_as", BGPNeighbor.PeerAS)
	d.Set("peer_ip", BGPNeighbor.PeerIP)
	d.Set("description", BGPNeighbor.Description)
	d.Set("session", BGPNeighbor.Session)
	d.Set("entity_scope", BGPNeighbor.EntityScope)
	d.Set("associated_export_routing_policy_id", BGPNeighbor.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", BGPNeighbor.AssociatedImportRoutingPolicyID)
	d.Set("external_id", BGPNeighbor.ExternalID)

	d.Set("id", BGPNeighbor.Identifier())
	d.Set("parent_id", BGPNeighbor.ParentID)
	d.Set("parent_type", BGPNeighbor.ParentType)
	d.Set("owner", BGPNeighbor.Owner)

	d.SetId(BGPNeighbor.Identifier())

	return nil
}
