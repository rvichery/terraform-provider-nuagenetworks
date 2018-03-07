package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNSGatewayTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSGatewayTemplateRead,
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
			"ssh_service": &schema.Schema{
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
			"personality": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"infrastructure_access_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"infrastructure_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_ssh_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNSGatewayTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredNSGatewayTemplates := vspk.NSGatewayTemplatesList{}
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
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredNSGatewayTemplates, err = parent.NSGatewayTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredNSGatewayTemplates, err = parent.NSGatewayTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	NSGatewayTemplate := &vspk.NSGatewayTemplate{}

	if len(filteredNSGatewayTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSGatewayTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		NSGatewayTemplate = filteredNSGatewayTemplates[0]
	}

	d.Set("ssh_service", NSGatewayTemplate.SSHService)
	d.Set("name", NSGatewayTemplate.Name)
	d.Set("last_updated_by", NSGatewayTemplate.LastUpdatedBy)
	d.Set("personality", NSGatewayTemplate.Personality)
	d.Set("description", NSGatewayTemplate.Description)
	d.Set("infrastructure_access_profile_id", NSGatewayTemplate.InfrastructureAccessProfileID)
	d.Set("infrastructure_profile_id", NSGatewayTemplate.InfrastructureProfileID)
	d.Set("instance_ssh_override", NSGatewayTemplate.InstanceSSHOverride)
	d.Set("enterprise_id", NSGatewayTemplate.EnterpriseID)
	d.Set("entity_scope", NSGatewayTemplate.EntityScope)
	d.Set("external_id", NSGatewayTemplate.ExternalID)

	d.Set("id", NSGatewayTemplate.Identifier())
	d.Set("parent_id", NSGatewayTemplate.ParentID)
	d.Set("parent_type", NSGatewayTemplate.ParentType)
	d.Set("owner", NSGatewayTemplate.Owner)

	d.SetId(NSGatewayTemplate.Identifier())

	return nil
}
