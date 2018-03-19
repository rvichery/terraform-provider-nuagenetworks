package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceAvatar() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAvatarRead,
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_user": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_user"},
			},
		},
	}
}

func dataSourceAvatarRead(d *schema.ResourceData, m interface{}) error {
	filteredAvatars := vspk.AvatarsList{}
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
	if attr, ok := d.GetOk("parent_user"); ok {
		parent := &vspk.User{ID: attr.(string)}
		filteredAvatars, err = parent.Avatars(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredAvatars, err = parent.Avatars(fetchFilter)
		if err != nil {
			return err
		}
	}

	Avatar := &vspk.Avatar{}

	if len(filteredAvatars) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredAvatars) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Avatar = filteredAvatars[0]

	d.Set("last_updated_by", Avatar.LastUpdatedBy)
	d.Set("entity_scope", Avatar.EntityScope)
	d.Set("external_id", Avatar.ExternalID)
	d.Set("type", Avatar.Type)

	d.Set("id", Avatar.Identifier())
	d.Set("parent_id", Avatar.ParentID)
	d.Set("parent_type", Avatar.ParentType)
	d.Set("owner", Avatar.Owner)

	d.SetId(Avatar.Identifier())

	return nil
}
