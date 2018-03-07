package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePortTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePortTemplateRead,
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
			"vlan_range": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"physical_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_egress_qos_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_gateway_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePortTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredPortTemplates := vspk.PortTemplatesList{}
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
	parent := &vspk.GatewayTemplate{ID: d.Get("parent_gateway_template").(string)}
	filteredPortTemplates, err = parent.PortTemplates(fetchFilter)
	if err != nil {
		return err
	}

	PortTemplate := &vspk.PortTemplate{}

	if len(filteredPortTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPortTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		PortTemplate = filteredPortTemplates[0]
	}

	d.Set("vlan_range", PortTemplate.VLANRange)
	d.Set("name", PortTemplate.Name)
	d.Set("last_updated_by", PortTemplate.LastUpdatedBy)
	d.Set("description", PortTemplate.Description)
	d.Set("physical_name", PortTemplate.PhysicalName)
	d.Set("entity_scope", PortTemplate.EntityScope)
	d.Set("port_type", PortTemplate.PortType)
	d.Set("associated_egress_qos_policy_id", PortTemplate.AssociatedEgressQOSPolicyID)
	d.Set("external_id", PortTemplate.ExternalID)

	d.Set("id", PortTemplate.Identifier())
	d.Set("parent_id", PortTemplate.ParentID)
	d.Set("parent_type", PortTemplate.ParentType)
	d.Set("owner", PortTemplate.Owner)

	d.SetId(PortTemplate.Identifier())

	return nil
}
