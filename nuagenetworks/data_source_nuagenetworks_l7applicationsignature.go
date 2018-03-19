package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceL7applicationsignature() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceL7applicationsignatureRead,
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
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"readonly": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deprecated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"deprecated_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dictionary_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"signature_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"risk": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"plugin_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_flags": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"productivity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"guidstring": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceL7applicationsignatureRead(d *schema.ResourceData, m interface{}) error {
	filteredL7applicationsignatures := vspk.L7applicationsignaturesList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredL7applicationsignatures, err = parent.L7applicationsignatures(fetchFilter)
	if err != nil {
		return err
	}

	L7applicationsignature := &vspk.L7applicationsignature{}

	if len(filteredL7applicationsignatures) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredL7applicationsignatures) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	L7applicationsignature = filteredL7applicationsignatures[0]

	d.Set("name", L7applicationsignature.Name)
	d.Set("category", L7applicationsignature.Category)
	d.Set("readonly", L7applicationsignature.Readonly)
	d.Set("reference", L7applicationsignature.Reference)
	d.Set("deprecated", L7applicationsignature.Deprecated)
	d.Set("deprecated_version", L7applicationsignature.DeprecatedVersion)
	d.Set("description", L7applicationsignature.Description)
	d.Set("dictionary_version", L7applicationsignature.DictionaryVersion)
	d.Set("signature_index", L7applicationsignature.SignatureIndex)
	d.Set("risk", L7applicationsignature.Risk)
	d.Set("plugin_name", L7applicationsignature.PluginName)
	d.Set("software_flags", L7applicationsignature.SoftwareFlags)
	d.Set("productivity", L7applicationsignature.Productivity)
	d.Set("guidstring", L7applicationsignature.Guidstring)

	d.Set("id", L7applicationsignature.Identifier())
	d.Set("parent_id", L7applicationsignature.ParentID)
	d.Set("parent_type", L7applicationsignature.ParentType)
	d.Set("owner", L7applicationsignature.Owner)

	d.SetId(L7applicationsignature.Identifier())

	return nil
}
