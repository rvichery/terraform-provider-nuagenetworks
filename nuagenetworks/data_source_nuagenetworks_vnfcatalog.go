package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVNFCatalog() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNFCatalogRead,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVNFCatalogRead(d *schema.ResourceData, m interface{}) error {
	filteredVNFCatalogs := vspk.VNFCatalogsList{}
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
	filteredVNFCatalogs, err = parent.VNFCatalogs(fetchFilter)
	if err != nil {
		return err
	}

	VNFCatalog := &vspk.VNFCatalog{}

	if len(filteredVNFCatalogs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVNFCatalogs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		VNFCatalog = filteredVNFCatalogs[0]
	}

	d.Set("name", VNFCatalog.Name)
	d.Set("description", VNFCatalog.Description)

	d.Set("id", VNFCatalog.Identifier())
	d.Set("parent_id", VNFCatalog.ParentID)
	d.Set("parent_type", VNFCatalog.ParentType)
	d.Set("owner", VNFCatalog.Owner)

	d.SetId(VNFCatalog.Identifier())

	return nil
}
