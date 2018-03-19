package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePATIPEntry() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePATIPEntryRead,
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
			"pat_centralized": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": {
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
			"associated_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_shared_network_resource": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource"},
			},
		},
	}
}

func dataSourcePATIPEntryRead(d *schema.ResourceData, m interface{}) error {
	filteredPATIPEntries := vspk.PATIPEntriesList{}
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
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		filteredPATIPEntries, err = parent.PATIPEntries(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredPATIPEntries, err = parent.PATIPEntries(fetchFilter)
		if err != nil {
			return err
		}
	}

	PATIPEntry := &vspk.PATIPEntry{}

	if len(filteredPATIPEntries) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPATIPEntries) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PATIPEntry = filteredPATIPEntries[0]

	d.Set("pat_centralized", PATIPEntry.PATCentralized)
	d.Set("ip_address", PATIPEntry.IPAddress)
	d.Set("ip_type", PATIPEntry.IPType)
	d.Set("last_updated_by", PATIPEntry.LastUpdatedBy)
	d.Set("entity_scope", PATIPEntry.EntityScope)
	d.Set("associated_domain_id", PATIPEntry.AssociatedDomainID)
	d.Set("external_id", PATIPEntry.ExternalID)
	d.Set("hypervisor_id", PATIPEntry.HypervisorID)

	d.Set("id", PATIPEntry.Identifier())
	d.Set("parent_id", PATIPEntry.ParentID)
	d.Set("parent_type", PATIPEntry.ParentType)
	d.Set("owner", PATIPEntry.Owner)

	d.SetId(PATIPEntry.Identifier())

	return nil
}
