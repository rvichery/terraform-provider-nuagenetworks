package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceNSGRoutingPolicyBinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSGRoutingPolicyBindingRead,
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
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_export_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_import_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_policy_object_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"export_to_overlay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceNSGRoutingPolicyBindingRead(d *schema.ResourceData, m interface{}) error {
	filteredNSGRoutingPolicyBindings := vspk.NSGRoutingPolicyBindingsList{}
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
	filteredNSGRoutingPolicyBindings, err = parent.NSGRoutingPolicyBindings(fetchFilter)
	if err != nil {
		return err
	}

	NSGRoutingPolicyBinding := &vspk.NSGRoutingPolicyBinding{}

	if len(filteredNSGRoutingPolicyBindings) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSGRoutingPolicyBindings) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		NSGRoutingPolicyBinding = filteredNSGRoutingPolicyBindings[0]
	}

	d.Set("name", NSGRoutingPolicyBinding.Name)
	d.Set("last_updated_by", NSGRoutingPolicyBinding.LastUpdatedBy)
	d.Set("description", NSGRoutingPolicyBinding.Description)
	d.Set("entity_scope", NSGRoutingPolicyBinding.EntityScope)
	d.Set("associated_export_routing_policy_id", NSGRoutingPolicyBinding.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", NSGRoutingPolicyBinding.AssociatedImportRoutingPolicyID)
	d.Set("associated_policy_object_group_id", NSGRoutingPolicyBinding.AssociatedPolicyObjectGroupID)
	d.Set("export_to_overlay", NSGRoutingPolicyBinding.ExportToOverlay)
	d.Set("external_id", NSGRoutingPolicyBinding.ExternalID)

	d.Set("id", NSGRoutingPolicyBinding.Identifier())
	d.Set("parent_id", NSGRoutingPolicyBinding.ParentID)
	d.Set("parent_type", NSGRoutingPolicyBinding.ParentType)
	d.Set("owner", NSGRoutingPolicyBinding.Owner)

	d.SetId(NSGRoutingPolicyBinding.Identifier())

	return nil
}
