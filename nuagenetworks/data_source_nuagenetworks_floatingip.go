package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceFloatingIp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFloatingIpRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_control": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assigned": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"assigned_to_object_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_shared_network_resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceFloatingIpRead(d *schema.ResourceData, m interface{}) error {
	filteredFloatingIps := vspk.FloatingIpsList{}
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
		filteredFloatingIps, err = parent.FloatingIps(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredFloatingIps, err = parent.FloatingIps(fetchFilter)
		if err != nil {
			return err
		}
	}

	FloatingIp := &vspk.FloatingIp{}

	if len(filteredFloatingIps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredFloatingIps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		FloatingIp = filteredFloatingIps[0]
	}

	d.Set("last_updated_by", FloatingIp.LastUpdatedBy)
	d.Set("access_control", FloatingIp.AccessControl)
	d.Set("address", FloatingIp.Address)
	d.Set("entity_scope", FloatingIp.EntityScope)
	d.Set("assigned", FloatingIp.Assigned)
	d.Set("assigned_to_object_type", FloatingIp.AssignedToObjectType)
	d.Set("associated_shared_network_resource_id", FloatingIp.AssociatedSharedNetworkResourceID)
	d.Set("external_id", FloatingIp.ExternalID)

	d.Set("id", FloatingIp.Identifier())
	d.Set("parent_id", FloatingIp.ParentID)
	d.Set("parent_type", FloatingIp.ParentType)
	d.Set("owner", FloatingIp.Owner)

	d.SetId(FloatingIp.Identifier())

	return nil
}
