package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceGroupKeyEncryptionProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGroupKeyEncryptionProfileRead,
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
			"sek_generation_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sek_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sek_payload_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sek_payload_encryption_bc_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sek_payload_encryption_key_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sek_payload_signing_algorithm": {
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
			"seed_generation_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seed_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seed_payload_authentication_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_authentication_bc_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_authentication_key_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seed_payload_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_encryption_bc_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_payload_encryption_key_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seed_payload_signing_algorithm": {
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
			"traffic_authentication_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_encryption_key_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGroupKeyEncryptionProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredGroupKeyEncryptionProfiles := vspk.GroupKeyEncryptionProfilesList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredGroupKeyEncryptionProfiles, err = parent.GroupKeyEncryptionProfiles(fetchFilter)
	if err != nil {
		return err
	}

	GroupKeyEncryptionProfile := &vspk.GroupKeyEncryptionProfile{}

	if len(filteredGroupKeyEncryptionProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredGroupKeyEncryptionProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	GroupKeyEncryptionProfile = filteredGroupKeyEncryptionProfiles[0]

	d.Set("sek_generation_interval", GroupKeyEncryptionProfile.SEKGenerationInterval)
	d.Set("sek_lifetime", GroupKeyEncryptionProfile.SEKLifetime)
	d.Set("sek_payload_encryption_algorithm", GroupKeyEncryptionProfile.SEKPayloadEncryptionAlgorithm)
	d.Set("sek_payload_encryption_bc_algorithm", GroupKeyEncryptionProfile.SEKPayloadEncryptionBCAlgorithm)
	d.Set("sek_payload_encryption_key_length", GroupKeyEncryptionProfile.SEKPayloadEncryptionKeyLength)
	d.Set("sek_payload_signing_algorithm", GroupKeyEncryptionProfile.SEKPayloadSigningAlgorithm)
	d.Set("name", GroupKeyEncryptionProfile.Name)
	d.Set("last_updated_by", GroupKeyEncryptionProfile.LastUpdatedBy)
	d.Set("seed_generation_interval", GroupKeyEncryptionProfile.SeedGenerationInterval)
	d.Set("seed_lifetime", GroupKeyEncryptionProfile.SeedLifetime)
	d.Set("seed_payload_authentication_algorithm", GroupKeyEncryptionProfile.SeedPayloadAuthenticationAlgorithm)
	d.Set("seed_payload_authentication_bc_algorithm", GroupKeyEncryptionProfile.SeedPayloadAuthenticationBCAlgorithm)
	d.Set("seed_payload_authentication_key_length", GroupKeyEncryptionProfile.SeedPayloadAuthenticationKeyLength)
	d.Set("seed_payload_encryption_algorithm", GroupKeyEncryptionProfile.SeedPayloadEncryptionAlgorithm)
	d.Set("seed_payload_encryption_bc_algorithm", GroupKeyEncryptionProfile.SeedPayloadEncryptionBCAlgorithm)
	d.Set("seed_payload_encryption_key_length", GroupKeyEncryptionProfile.SeedPayloadEncryptionKeyLength)
	d.Set("seed_payload_signing_algorithm", GroupKeyEncryptionProfile.SeedPayloadSigningAlgorithm)
	d.Set("description", GroupKeyEncryptionProfile.Description)
	d.Set("entity_scope", GroupKeyEncryptionProfile.EntityScope)
	d.Set("traffic_authentication_algorithm", GroupKeyEncryptionProfile.TrafficAuthenticationAlgorithm)
	d.Set("traffic_encryption_algorithm", GroupKeyEncryptionProfile.TrafficEncryptionAlgorithm)
	d.Set("traffic_encryption_key_lifetime", GroupKeyEncryptionProfile.TrafficEncryptionKeyLifetime)
	d.Set("associated_enterprise_id", GroupKeyEncryptionProfile.AssociatedEnterpriseID)
	d.Set("external_id", GroupKeyEncryptionProfile.ExternalID)

	d.Set("id", GroupKeyEncryptionProfile.Identifier())
	d.Set("parent_id", GroupKeyEncryptionProfile.ParentID)
	d.Set("parent_type", GroupKeyEncryptionProfile.ParentType)
	d.Set("owner", GroupKeyEncryptionProfile.Owner)

	d.SetId(GroupKeyEncryptionProfile.Identifier())

	return nil
}
