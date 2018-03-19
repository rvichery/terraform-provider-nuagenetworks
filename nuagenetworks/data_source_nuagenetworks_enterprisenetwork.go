package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEnterpriseNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnterpriseNetworkRead,
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
			"ip_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": {
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
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_network_macro_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_network_macro_group"},
			},
		},
	}
}

func dataSourceEnterpriseNetworkRead(d *schema.ResourceData, m interface{}) error {
	filteredEnterpriseNetworks := vspk.EnterpriseNetworksList{}
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
	if attr, ok := d.GetOk("parent_network_macro_group"); ok {
		parent := &vspk.NetworkMacroGroup{ID: attr.(string)}
		filteredEnterpriseNetworks, err = parent.EnterpriseNetworks(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredEnterpriseNetworks, err = parent.EnterpriseNetworks(fetchFilter)
		if err != nil {
			return err
		}
	}

	EnterpriseNetwork := &vspk.EnterpriseNetwork{}

	if len(filteredEnterpriseNetworks) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEnterpriseNetworks) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EnterpriseNetwork = filteredEnterpriseNetworks[0]

	d.Set("ip_type", EnterpriseNetwork.IPType)
	d.Set("ipv6_address", EnterpriseNetwork.IPv6Address)
	d.Set("name", EnterpriseNetwork.Name)
	d.Set("last_updated_by", EnterpriseNetwork.LastUpdatedBy)
	d.Set("address", EnterpriseNetwork.Address)
	d.Set("netmask", EnterpriseNetwork.Netmask)
	d.Set("entity_scope", EnterpriseNetwork.EntityScope)
	d.Set("external_id", EnterpriseNetwork.ExternalID)

	d.Set("id", EnterpriseNetwork.Identifier())
	d.Set("parent_id", EnterpriseNetwork.ParentID)
	d.Set("parent_type", EnterpriseNetwork.ParentType)
	d.Set("owner", EnterpriseNetwork.Owner)

	d.SetId(EnterpriseNetwork.Identifier())

	return nil
}
