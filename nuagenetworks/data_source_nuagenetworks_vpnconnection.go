package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVPNConnection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVPNConnectionRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_wan_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_shared_network_resource": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_l2_domain"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_domain"},
			},
		},
	}
}

func dataSourceVPNConnectionRead(d *schema.ResourceData, m interface{}) error {
	filteredVPNConnections := vspk.VPNConnectionsList{}
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
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		filteredVPNConnections, err = parent.VPNConnections(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredVPNConnections, err = parent.VPNConnections(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredVPNConnections, err = parent.VPNConnections(fetchFilter)
		if err != nil {
			return err
		}
	}

	VPNConnection := &vspk.VPNConnection{}

	if len(filteredVPNConnections) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVPNConnections) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		VPNConnection = filteredVPNConnections[0]
	}

	d.Set("name", VPNConnection.Name)
	d.Set("last_updated_by", VPNConnection.LastUpdatedBy)
	d.Set("description", VPNConnection.Description)
	d.Set("entity_scope", VPNConnection.EntityScope)
	d.Set("associated_wan_service_id", VPNConnection.AssociatedWANServiceID)
	d.Set("external_id", VPNConnection.ExternalID)

	d.Set("id", VPNConnection.Identifier())
	d.Set("parent_id", VPNConnection.ParentID)
	d.Set("parent_type", VPNConnection.ParentType)
	d.Set("owner", VPNConnection.Owner)

	d.SetId(VPNConnection.Identifier())

	return nil
}
