package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePSPATMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePSPATMapRead,
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
			"reserved_spatips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"associated_spat_sources_pool_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_psnat_pool": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePSPATMapRead(d *schema.ResourceData, m interface{}) error {
	filteredPSPATMaps := vspk.PSPATMapsList{}
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
	parent := &vspk.PSNATPool{ID: d.Get("parent_psnat_pool").(string)}
	filteredPSPATMaps, err = parent.PSPATMaps(fetchFilter)
	if err != nil {
		return err
	}

	PSPATMap := &vspk.PSPATMap{}

	if len(filteredPSPATMaps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPSPATMaps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PSPATMap = filteredPSPATMaps[0]

	d.Set("name", PSPATMap.Name)
	d.Set("reserved_spatips", PSPATMap.ReservedSPATIPs)
	d.Set("associated_spat_sources_pool_id", PSPATMap.AssociatedSPATSourcesPoolID)

	d.Set("id", PSPATMap.Identifier())
	d.Set("parent_id", PSPATMap.ParentID)
	d.Set("parent_type", PSPATMap.ParentType)
	d.Set("owner", PSPATMap.Owner)

	d.SetId(PSPATMap.Identifier())

	return nil
}
