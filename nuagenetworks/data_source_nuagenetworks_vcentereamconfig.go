package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVCenterEAMConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVCenterEAMConfigRead,
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
			"eam_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eam_server_port_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"eam_server_port_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vib_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ovf_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extension_key": {
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

func dataSourceVCenterEAMConfigRead(d *schema.ResourceData, m interface{}) error {
	filteredVCenterEAMConfigs := vspk.VCenterEAMConfigsList{}
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
	filteredVCenterEAMConfigs, err = parent.VCenterEAMConfigs(fetchFilter)
	if err != nil {
		return err
	}

	VCenterEAMConfig := &vspk.VCenterEAMConfig{}

	if len(filteredVCenterEAMConfigs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVCenterEAMConfigs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VCenterEAMConfig = filteredVCenterEAMConfigs[0]

	d.Set("eam_server_ip", VCenterEAMConfig.EamServerIP)
	d.Set("eam_server_port_number", VCenterEAMConfig.EamServerPortNumber)
	d.Set("eam_server_port_type", VCenterEAMConfig.EamServerPortType)
	d.Set("last_updated_by", VCenterEAMConfig.LastUpdatedBy)
	d.Set("vib_url", VCenterEAMConfig.VibURL)
	d.Set("entity_scope", VCenterEAMConfig.EntityScope)
	d.Set("ovf_url", VCenterEAMConfig.OvfURL)
	d.Set("extension_key", VCenterEAMConfig.ExtensionKey)
	d.Set("external_id", VCenterEAMConfig.ExternalID)

	d.Set("id", VCenterEAMConfig.Identifier())
	d.Set("parent_id", VCenterEAMConfig.ParentID)
	d.Set("parent_type", VCenterEAMConfig.ParentType)
	d.Set("owner", VCenterEAMConfig.Owner)

	d.SetId(VCenterEAMConfig.Identifier())

	return nil
}
