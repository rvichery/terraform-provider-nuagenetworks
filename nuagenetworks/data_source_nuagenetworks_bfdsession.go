package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceBFDSession() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBFDSessionRead,
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
			"bfd_destination_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd_multiplier": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_hop_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_uplink_connection": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_br_connection"},
			},
			"parent_br_connection": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_uplink_connection"},
			},
		},
	}
}

func dataSourceBFDSessionRead(d *schema.ResourceData, m interface{}) error {
	filteredBFDSessions := vspk.BFDSessionsList{}
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
	if attr, ok := d.GetOk("parent_uplink_connection"); ok {
		parent := &vspk.UplinkConnection{ID: attr.(string)}
		filteredBFDSessions, err = parent.BFDSessions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_br_connection"); ok {
		parent := &vspk.BRConnection{ID: attr.(string)}
		filteredBFDSessions, err = parent.BFDSessions(fetchFilter)
		if err != nil {
			return err
		}
	}

	BFDSession := &vspk.BFDSession{}

	if len(filteredBFDSessions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredBFDSessions) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		BFDSession = filteredBFDSessions[0]
	}

	d.Set("bfd_destination_ip", BFDSession.BFDDestinationIP)
	d.Set("bfd_multiplier", BFDSession.BFDMultiplier)
	d.Set("bfd_timer", BFDSession.BFDTimer)
	d.Set("last_updated_by", BFDSession.LastUpdatedBy)
	d.Set("entity_scope", BFDSession.EntityScope)
	d.Set("multi_hop_enabled", BFDSession.MultiHopEnabled)
	d.Set("external_id", BFDSession.ExternalID)

	d.Set("id", BFDSession.Identifier())
	d.Set("parent_id", BFDSession.ParentID)
	d.Set("parent_type", BFDSession.ParentType)
	d.Set("owner", BFDSession.Owner)

	d.SetId(BFDSession.Identifier())

	return nil
}
