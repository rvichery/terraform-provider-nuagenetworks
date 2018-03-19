package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceAutoDiscoverCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutoDiscoverClusterRead,
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assoc_vcenter_data_center_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vcenter_data_center": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAutoDiscoverClusterRead(d *schema.ResourceData, m interface{}) error {
	filteredAutoDiscoverClusters := vspk.AutoDiscoverClustersList{}
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
	parent := &vspk.VCenterDataCenter{ID: d.Get("parent_vcenter_data_center").(string)}
	filteredAutoDiscoverClusters, err = parent.AutoDiscoverClusters(fetchFilter)
	if err != nil {
		return err
	}

	AutoDiscoverCluster := &vspk.AutoDiscoverCluster{}

	if len(filteredAutoDiscoverClusters) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredAutoDiscoverClusters) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	AutoDiscoverCluster = filteredAutoDiscoverClusters[0]

	d.Set("name", AutoDiscoverCluster.Name)
	d.Set("managed_object_id", AutoDiscoverCluster.ManagedObjectID)
	d.Set("last_updated_by", AutoDiscoverCluster.LastUpdatedBy)
	d.Set("entity_scope", AutoDiscoverCluster.EntityScope)
	d.Set("assoc_vcenter_data_center_id", AutoDiscoverCluster.AssocVCenterDataCenterID)
	d.Set("external_id", AutoDiscoverCluster.ExternalID)

	d.Set("id", AutoDiscoverCluster.Identifier())
	d.Set("parent_id", AutoDiscoverCluster.ParentID)
	d.Set("parent_type", AutoDiscoverCluster.ParentType)
	d.Set("owner", AutoDiscoverCluster.Owner)

	d.SetId(AutoDiscoverCluster.Identifier())

	return nil
}
