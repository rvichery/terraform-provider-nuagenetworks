package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePortMapping() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePortMappingRead,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vport": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePortMappingRead(d *schema.ResourceData, m interface{}) error {
	filteredPortMappings := vspk.PortMappingsList{}
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
	parent := &vspk.VPort{ID: d.Get("parent_vport").(string)}
	filteredPortMappings, err = parent.PortMappings(fetchFilter)
	if err != nil {
		return err
	}

	PortMapping := &vspk.PortMapping{}

	if len(filteredPortMappings) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPortMappings) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PortMapping = filteredPortMappings[0]

	d.Set("last_updated_by", PortMapping.LastUpdatedBy)
	d.Set("entity_scope", PortMapping.EntityScope)
	d.Set("private_port", PortMapping.PrivatePort)
	d.Set("public_port", PortMapping.PublicPort)
	d.Set("external_id", PortMapping.ExternalID)

	d.Set("id", PortMapping.Identifier())
	d.Set("parent_id", PortMapping.ParentID)
	d.Set("parent_type", PortMapping.ParentType)
	d.Set("owner", PortMapping.Owner)

	d.SetId(PortMapping.Identifier())

	return nil
}
