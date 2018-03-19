package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVSP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVSPRead,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVSPRead(d *schema.ResourceData, m interface{}) error {
	filteredVSPs := vspk.VSPsList{}
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
	parent := m.(*vspk.Me)
	filteredVSPs, err = parent.VSPs(fetchFilter)
	if err != nil {
		return err
	}

	VSP := &vspk.VSP{}

	if len(filteredVSPs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVSPs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VSP = filteredVSPs[0]

	d.Set("name", VSP.Name)
	d.Set("last_updated_by", VSP.LastUpdatedBy)
	d.Set("description", VSP.Description)
	d.Set("entity_scope", VSP.EntityScope)
	d.Set("location", VSP.Location)
	d.Set("product_version", VSP.ProductVersion)
	d.Set("external_id", VSP.ExternalID)

	d.Set("id", VSP.Identifier())
	d.Set("parent_id", VSP.ParentID)
	d.Set("parent_type", VSP.ParentType)
	d.Set("owner", VSP.Owner)

	d.SetId(VSP.Identifier())

	return nil
}
