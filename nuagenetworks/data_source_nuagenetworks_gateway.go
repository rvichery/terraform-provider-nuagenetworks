package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGatewayRead,
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
			"redundancy_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pending": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_gateway_vlanvnid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vtep": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_disc_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_redundancy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group"},
			},
		},
	}
}

func dataSourceGatewayRead(d *schema.ResourceData, m interface{}) error {
	filteredGateways := vspk.GatewaysList{}
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
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		filteredGateways, err = parent.Gateways(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredGateways, err = parent.Gateways(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredGateways, err = parent.Gateways(fetchFilter)
		if err != nil {
			return err
		}
	}

	Gateway := &vspk.Gateway{}

	if len(filteredGateways) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Gateway = filteredGateways[0]

	d.Set("name", Gateway.Name)
	d.Set("last_updated_by", Gateway.LastUpdatedBy)
	d.Set("redundancy_group_id", Gateway.RedundancyGroupID)
	d.Set("peer", Gateway.Peer)
	d.Set("template_id", Gateway.TemplateID)
	d.Set("pending", Gateway.Pending)
	d.Set("permitted_action", Gateway.PermittedAction)
	d.Set("personality", Gateway.Personality)
	d.Set("description", Gateway.Description)
	d.Set("enterprise_id", Gateway.EnterpriseID)
	d.Set("entity_scope", Gateway.EntityScope)
	d.Set("use_gateway_vlanvnid", Gateway.UseGatewayVLANVNID)
	d.Set("vtep", Gateway.Vtep)
	d.Set("auto_disc_gateway_id", Gateway.AutoDiscGatewayID)
	d.Set("external_id", Gateway.ExternalID)
	d.Set("system_id", Gateway.SystemID)

	d.Set("id", Gateway.Identifier())
	d.Set("parent_id", Gateway.ParentID)
	d.Set("parent_type", Gateway.ParentType)
	d.Set("owner", Gateway.Owner)

	d.SetId(Gateway.Identifier())

	return nil
}
