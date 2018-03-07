package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceRedirectionTarget() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRedirectionTargetRead,
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
			"esi": &schema.Schema{
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
			"redundancy_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_point_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vm_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_bridge_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_host_interface"},
			},
			"parent_host_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceRedirectionTargetRead(d *schema.ResourceData, m interface{}) error {
	filteredRedirectionTargets := vspk.RedirectionTargetsList{}
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
	if attr, ok := d.GetOk("parent_container_interface"); ok {
		parent := &vspk.ContainerInterface{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm_interface"); ok {
		parent := &vspk.VMInterface{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredRedirectionTargets, err = parent.RedirectionTargets(fetchFilter)
		if err != nil {
			return err
		}
	}

	RedirectionTarget := &vspk.RedirectionTarget{}

	if len(filteredRedirectionTargets) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredRedirectionTargets) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		RedirectionTarget = filteredRedirectionTargets[0]
	}

	d.Set("esi", RedirectionTarget.ESI)
	d.Set("name", RedirectionTarget.Name)
	d.Set("last_updated_by", RedirectionTarget.LastUpdatedBy)
	d.Set("redundancy_enabled", RedirectionTarget.RedundancyEnabled)
	d.Set("template_id", RedirectionTarget.TemplateID)
	d.Set("description", RedirectionTarget.Description)
	d.Set("virtual_network_id", RedirectionTarget.VirtualNetworkID)
	d.Set("end_point_type", RedirectionTarget.EndPointType)
	d.Set("entity_scope", RedirectionTarget.EntityScope)
	d.Set("trigger_type", RedirectionTarget.TriggerType)
	d.Set("external_id", RedirectionTarget.ExternalID)

	d.Set("id", RedirectionTarget.Identifier())
	d.Set("parent_id", RedirectionTarget.ParentID)
	d.Set("parent_type", RedirectionTarget.ParentType)
	d.Set("owner", RedirectionTarget.Owner)

	d.SetId(RedirectionTarget.Identifier())

	return nil
}
