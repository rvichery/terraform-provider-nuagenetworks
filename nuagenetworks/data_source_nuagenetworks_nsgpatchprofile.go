package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceNSGPatchProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSGPatchProfileRead,
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
			"patch_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_url": {
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNSGPatchProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredNSGPatchProfiles := vspk.NSGPatchProfilesList{}
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
	parent := m.(*vspk.Me)
	filteredNSGPatchProfiles, err = parent.NSGPatchProfiles(fetchFilter)
	if err != nil {
		return err
	}

	NSGPatchProfile := &vspk.NSGPatchProfile{}

	if len(filteredNSGPatchProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSGPatchProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NSGPatchProfile = filteredNSGPatchProfiles[0]

	d.Set("name", NSGPatchProfile.Name)
	d.Set("last_updated_by", NSGPatchProfile.LastUpdatedBy)
	d.Set("patch_tag", NSGPatchProfile.PatchTag)
	d.Set("patch_url", NSGPatchProfile.PatchURL)
	d.Set("description", NSGPatchProfile.Description)
	d.Set("enterprise_id", NSGPatchProfile.EnterpriseID)
	d.Set("entity_scope", NSGPatchProfile.EntityScope)
	d.Set("external_id", NSGPatchProfile.ExternalID)

	d.Set("id", NSGPatchProfile.Identifier())
	d.Set("parent_id", NSGPatchProfile.ParentID)
	d.Set("parent_type", NSGPatchProfile.ParentType)
	d.Set("owner", NSGPatchProfile.Owner)

	d.SetId(NSGPatchProfile.Identifier())

	return nil
}
