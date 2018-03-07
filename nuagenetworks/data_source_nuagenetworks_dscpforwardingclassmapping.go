package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDSCPForwardingClassMapping() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDSCPForwardingClassMappingRead,
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
			"dscp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forwarding_class": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_dscp_forwarding_class_table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDSCPForwardingClassMappingRead(d *schema.ResourceData, m interface{}) error {
	filteredDSCPForwardingClassMappings := vspk.DSCPForwardingClassMappingsList{}
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
	parent := &vspk.DSCPForwardingClassTable{ID: d.Get("parent_dscp_forwarding_class_table").(string)}
	filteredDSCPForwardingClassMappings, err = parent.DSCPForwardingClassMappings(fetchFilter)
	if err != nil {
		return err
	}

	DSCPForwardingClassMapping := &vspk.DSCPForwardingClassMapping{}

	if len(filteredDSCPForwardingClassMappings) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDSCPForwardingClassMappings) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		DSCPForwardingClassMapping = filteredDSCPForwardingClassMappings[0]
	}

	d.Set("dscp", DSCPForwardingClassMapping.DSCP)
	d.Set("last_updated_by", DSCPForwardingClassMapping.LastUpdatedBy)
	d.Set("entity_scope", DSCPForwardingClassMapping.EntityScope)
	d.Set("forwarding_class", DSCPForwardingClassMapping.ForwardingClass)
	d.Set("external_id", DSCPForwardingClassMapping.ExternalID)

	d.Set("id", DSCPForwardingClassMapping.Identifier())
	d.Set("parent_id", DSCPForwardingClassMapping.ParentID)
	d.Set("parent_type", DSCPForwardingClassMapping.ParentType)
	d.Set("owner", DSCPForwardingClassMapping.Owner)

	d.SetId(DSCPForwardingClassMapping.Identifier())

	return nil
}
