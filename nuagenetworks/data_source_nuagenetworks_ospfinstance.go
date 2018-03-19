package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceOSPFInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOSPFInstanceRead,
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
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preference": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_export_routing_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_import_routing_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"super_backbone_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"export_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"export_to_overlay": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_preference": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceOSPFInstanceRead(d *schema.ResourceData, m interface{}) error {
	filteredOSPFInstances := vspk.OSPFInstancesList{}
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
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	filteredOSPFInstances, err = parent.OSPFInstances(fetchFilter)
	if err != nil {
		return err
	}

	OSPFInstance := &vspk.OSPFInstance{}

	if len(filteredOSPFInstances) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredOSPFInstances) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	OSPFInstance = filteredOSPFInstances[0]

	d.Set("name", OSPFInstance.Name)
	d.Set("last_updated_by", OSPFInstance.LastUpdatedBy)
	d.Set("description", OSPFInstance.Description)
	d.Set("entity_scope", OSPFInstance.EntityScope)
	d.Set("preference", OSPFInstance.Preference)
	d.Set("associated_export_routing_policy_id", OSPFInstance.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", OSPFInstance.AssociatedImportRoutingPolicyID)
	d.Set("super_backbone_enabled", OSPFInstance.SuperBackboneEnabled)
	d.Set("export_limit", OSPFInstance.ExportLimit)
	d.Set("export_to_overlay", OSPFInstance.ExportToOverlay)
	d.Set("external_id", OSPFInstance.ExternalID)
	d.Set("external_preference", OSPFInstance.ExternalPreference)

	d.Set("id", OSPFInstance.Identifier())
	d.Set("parent_id", OSPFInstance.ParentID)
	d.Set("parent_type", OSPFInstance.ParentType)
	d.Set("owner", OSPFInstance.Owner)

	d.SetId(OSPFInstance.Identifier())

	return nil
}
