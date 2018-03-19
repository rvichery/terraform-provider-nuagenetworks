package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceCTranslationMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCTranslationMapRead,
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
			"mapping_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_alias_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_csnat_pool": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCTranslationMapRead(d *schema.ResourceData, m interface{}) error {
	filteredCTranslationMaps := vspk.CTranslationMapsList{}
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
	parent := &vspk.CSNATPool{ID: d.Get("parent_csnat_pool").(string)}
	filteredCTranslationMaps, err = parent.CTranslationMaps(fetchFilter)
	if err != nil {
		return err
	}

	CTranslationMap := &vspk.CTranslationMap{}

	if len(filteredCTranslationMaps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCTranslationMaps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	CTranslationMap = filteredCTranslationMaps[0]

	d.Set("mapping_type", CTranslationMap.MappingType)
	d.Set("customer_alias_ip", CTranslationMap.CustomerAliasIP)
	d.Set("customer_ip", CTranslationMap.CustomerIP)

	d.Set("id", CTranslationMap.Identifier())
	d.Set("parent_id", CTranslationMap.ParentID)
	d.Set("parent_type", CTranslationMap.ParentType)
	d.Set("owner", CTranslationMap.Owner)

	d.SetId(CTranslationMap.Identifier())

	return nil
}
