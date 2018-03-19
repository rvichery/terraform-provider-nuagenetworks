package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceMultiCastChannelMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMultiCastChannelMapRead,
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_interface": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vm_interface", "parent_multi_cast_list", "parent_host_interface"},
			},
			"parent_vm_interface": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_multi_cast_list", "parent_host_interface"},
			},
			"parent_multi_cast_list": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_vm_interface", "parent_host_interface"},
			},
			"parent_host_interface": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_vm_interface", "parent_multi_cast_list"},
			},
		},
	}
}

func dataSourceMultiCastChannelMapRead(d *schema.ResourceData, m interface{}) error {
	filteredMultiCastChannelMaps := vspk.MultiCastChannelMapsList{}
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
		filteredMultiCastChannelMaps, err = parent.MultiCastChannelMaps(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm_interface"); ok {
		parent := &vspk.VMInterface{ID: attr.(string)}
		filteredMultiCastChannelMaps, err = parent.MultiCastChannelMaps(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_multi_cast_list"); ok {
		parent := &vspk.MultiCastList{ID: attr.(string)}
		filteredMultiCastChannelMaps, err = parent.MultiCastChannelMaps(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredMultiCastChannelMaps, err = parent.MultiCastChannelMaps(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredMultiCastChannelMaps, err = parent.MultiCastChannelMaps(fetchFilter)
		if err != nil {
			return err
		}
	}

	MultiCastChannelMap := &vspk.MultiCastChannelMap{}

	if len(filteredMultiCastChannelMaps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredMultiCastChannelMaps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	MultiCastChannelMap = filteredMultiCastChannelMaps[0]

	d.Set("name", MultiCastChannelMap.Name)
	d.Set("last_updated_by", MultiCastChannelMap.LastUpdatedBy)
	d.Set("description", MultiCastChannelMap.Description)
	d.Set("entity_scope", MultiCastChannelMap.EntityScope)
	d.Set("external_id", MultiCastChannelMap.ExternalID)

	d.Set("id", MultiCastChannelMap.Identifier())
	d.Set("parent_id", MultiCastChannelMap.ParentID)
	d.Set("parent_type", MultiCastChannelMap.ParentType)
	d.Set("owner", MultiCastChannelMap.Owner)

	d.SetId(MultiCastChannelMap.Identifier())

	return nil
}
