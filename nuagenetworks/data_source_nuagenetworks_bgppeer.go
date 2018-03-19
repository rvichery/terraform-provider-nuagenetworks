package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceBGPPeer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBGPPeerRead,
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
			"last_state_change": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_hsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vsc"},
			},
			"parent_vsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_hsc"},
			},
		},
	}
}

func dataSourceBGPPeerRead(d *schema.ResourceData, m interface{}) error {
	filteredBGPPeers := vspk.BGPPeersList{}
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
	if attr, ok := d.GetOk("parent_hsc"); ok {
		parent := &vspk.HSC{ID: attr.(string)}
		filteredBGPPeers, err = parent.BGPPeers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsc"); ok {
		parent := &vspk.VSC{ID: attr.(string)}
		filteredBGPPeers, err = parent.BGPPeers(fetchFilter)
		if err != nil {
			return err
		}
	}

	BGPPeer := &vspk.BGPPeer{}

	if len(filteredBGPPeers) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredBGPPeers) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	BGPPeer = filteredBGPPeers[0]

	d.Set("last_state_change", BGPPeer.LastStateChange)
	d.Set("address", BGPPeer.Address)
	d.Set("entity_scope", BGPPeer.EntityScope)
	d.Set("status", BGPPeer.Status)
	d.Set("external_id", BGPPeer.ExternalID)

	d.Set("id", BGPPeer.Identifier())
	d.Set("parent_id", BGPPeer.ParentID)
	d.Set("parent_type", BGPPeer.ParentType)
	d.Set("owner", BGPPeer.Owner)

	d.SetId(BGPPeer.Identifier())

	return nil
}
