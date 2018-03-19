package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceKeyServerMonitorEncryptedSeed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeyServerMonitorEncryptedSeedRead,
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
			"sek_creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_server_certificate_serial_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enterprise_secured_data_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_key_server_monitor_sek_creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_key_server_monitor_sekid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_key_server_monitor_seed_creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_key_server_monitor_seed_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_key_server_monitor_seed": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_key_server_monitor"},
			},
			"parent_key_server_monitor": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_key_server_monitor_seed"},
			},
		},
	}
}

func dataSourceKeyServerMonitorEncryptedSeedRead(d *schema.ResourceData, m interface{}) error {
	filteredKeyServerMonitorEncryptedSeeds := vspk.KeyServerMonitorEncryptedSeedsList{}
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
	if attr, ok := d.GetOk("parent_key_server_monitor_seed"); ok {
		parent := &vspk.KeyServerMonitorSeed{ID: attr.(string)}
		filteredKeyServerMonitorEncryptedSeeds, err = parent.KeyServerMonitorEncryptedSeeds(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_key_server_monitor"); ok {
		parent := &vspk.KeyServerMonitor{ID: attr.(string)}
		filteredKeyServerMonitorEncryptedSeeds, err = parent.KeyServerMonitorEncryptedSeeds(fetchFilter)
		if err != nil {
			return err
		}
	}

	KeyServerMonitorEncryptedSeed := &vspk.KeyServerMonitorEncryptedSeed{}

	if len(filteredKeyServerMonitorEncryptedSeeds) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredKeyServerMonitorEncryptedSeeds) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	KeyServerMonitorEncryptedSeed = filteredKeyServerMonitorEncryptedSeeds[0]

	d.Set("sek_creation_time", KeyServerMonitorEncryptedSeed.SEKCreationTime)
	d.Set("last_updated_by", KeyServerMonitorEncryptedSeed.LastUpdatedBy)
	d.Set("key_server_certificate_serial_number", KeyServerMonitorEncryptedSeed.KeyServerCertificateSerialNumber)
	d.Set("enterprise_secured_data_id", KeyServerMonitorEncryptedSeed.EnterpriseSecuredDataID)
	d.Set("entity_scope", KeyServerMonitorEncryptedSeed.EntityScope)
	d.Set("associated_key_server_monitor_sek_creation_time", KeyServerMonitorEncryptedSeed.AssociatedKeyServerMonitorSEKCreationTime)
	d.Set("associated_key_server_monitor_sekid", KeyServerMonitorEncryptedSeed.AssociatedKeyServerMonitorSEKID)
	d.Set("associated_key_server_monitor_seed_creation_time", KeyServerMonitorEncryptedSeed.AssociatedKeyServerMonitorSeedCreationTime)
	d.Set("associated_key_server_monitor_seed_id", KeyServerMonitorEncryptedSeed.AssociatedKeyServerMonitorSeedID)
	d.Set("external_id", KeyServerMonitorEncryptedSeed.ExternalID)

	d.Set("id", KeyServerMonitorEncryptedSeed.Identifier())
	d.Set("parent_id", KeyServerMonitorEncryptedSeed.ParentID)
	d.Set("parent_type", KeyServerMonitorEncryptedSeed.ParentType)
	d.Set("owner", KeyServerMonitorEncryptedSeed.Owner)

	d.SetId(KeyServerMonitorEncryptedSeed.Identifier())

	return nil
}
