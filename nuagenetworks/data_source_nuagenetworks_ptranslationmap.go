package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePTranslationMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePTranslationMapRead,
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
			"spat_source_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mapping_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_alias_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_psnat_pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePTranslationMapRead(d *schema.ResourceData, m interface{}) error {
	filteredPTranslationMaps := vspk.PTranslationMapsList{}
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
	filteredPTranslationMaps, err = parent.PTranslationMaps(fetchFilter)
	if err != nil {
		return err
	}

	PTranslationMap := &vspk.PTranslationMap{}

	if len(filteredPTranslationMaps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPTranslationMaps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		PTranslationMap = filteredPTranslationMaps[0]
	}

	d.Set("spat_source_list", PTranslationMap.SPATSourceList)
	d.Set("mapping_type", PTranslationMap.MappingType)
	d.Set("provider_alias_ip", PTranslationMap.ProviderAliasIP)
	d.Set("provider_ip", PTranslationMap.ProviderIP)

	d.Set("id", PTranslationMap.Identifier())
	d.Set("parent_id", PTranslationMap.ParentID)
	d.Set("parent_type", PTranslationMap.ParentType)
	d.Set("owner", PTranslationMap.Owner)

	d.SetId(PTranslationMap.Identifier())

	return nil
}
