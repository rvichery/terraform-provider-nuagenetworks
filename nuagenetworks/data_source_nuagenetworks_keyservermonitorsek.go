package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceKeyServerMonitorSEK() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeyServerMonitorSEKRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_authentication_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_encryption_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_key_server_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceKeyServerMonitorSEKRead(d *schema.ResourceData, m interface{}) error {
	filteredKeyServerMonitorSEKs := vspk.KeyServerMonitorSEKsList{}
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
	parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
	filteredKeyServerMonitorSEKs, err = parent.KeyServerMonitorSEKs(fetchFilter)
	if err != nil {
		return err
	}

	KeyServerMonitorSEK := &vspk.KeyServerMonitorSEK{}

	if len(filteredKeyServerMonitorSEKs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredKeyServerMonitorSEKs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		KeyServerMonitorSEK = filteredKeyServerMonitorSEKs[0]
	}

	d.Set("last_updated_by", KeyServerMonitorSEK.LastUpdatedBy)
	d.Set("seed_payload_authentication_algorithm", KeyServerMonitorSEK.SeedPayloadAuthenticationAlgorithm)
	d.Set("seed_payload_encryption_algorithm", KeyServerMonitorSEK.SeedPayloadEncryptionAlgorithm)
	d.Set("lifetime", KeyServerMonitorSEK.Lifetime)
	d.Set("entity_scope", KeyServerMonitorSEK.EntityScope)
	d.Set("creation_time", KeyServerMonitorSEK.CreationTime)
	d.Set("start_time", KeyServerMonitorSEK.StartTime)
	d.Set("external_id", KeyServerMonitorSEK.ExternalID)

	d.Set("id", KeyServerMonitorSEK.Identifier())
	d.Set("parent_id", KeyServerMonitorSEK.ParentID)
	d.Set("parent_type", KeyServerMonitorSEK.ParentType)
	d.Set("owner", KeyServerMonitorSEK.Owner)

	d.SetId(KeyServerMonitorSEK.Identifier())

	return nil
}
