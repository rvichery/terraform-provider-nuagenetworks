package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIKEEncryptionprofile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKEEncryptionprofileRead,
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
			"dpd_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dpd_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dpd_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_authentication_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_dont_fragment": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipsec_enable_pfs": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipsec_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_pre_fragment": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipsec_sa_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_sa_replay_window_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isakmp_authentication_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isakmp_diffie_helman_group_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isakmp_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isakmp_encryption_key_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"isakmp_hash_algorithm": {
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
			"sequence": {
				Type:     schema.TypeInt,
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

func dataSourceIKEEncryptionprofileRead(d *schema.ResourceData, m interface{}) error {
	filteredIKEEncryptionprofiles := vspk.IKEEncryptionprofilesList{}
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
	filteredIKEEncryptionprofiles, err = parent.IKEEncryptionprofiles(fetchFilter)
	if err != nil {
		return err
	}

	IKEEncryptionprofile := &vspk.IKEEncryptionprofile{}

	if len(filteredIKEEncryptionprofiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKEEncryptionprofiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IKEEncryptionprofile = filteredIKEEncryptionprofiles[0]

	d.Set("dpd_interval", IKEEncryptionprofile.DPDInterval)
	d.Set("dpd_mode", IKEEncryptionprofile.DPDMode)
	d.Set("dpd_timeout", IKEEncryptionprofile.DPDTimeout)
	d.Set("ipsec_authentication_algorithm", IKEEncryptionprofile.IPsecAuthenticationAlgorithm)
	d.Set("ipsec_dont_fragment", IKEEncryptionprofile.IPsecDontFragment)
	d.Set("ipsec_enable_pfs", IKEEncryptionprofile.IPsecEnablePFS)
	d.Set("ipsec_encryption_algorithm", IKEEncryptionprofile.IPsecEncryptionAlgorithm)
	d.Set("ipsec_pre_fragment", IKEEncryptionprofile.IPsecPreFragment)
	d.Set("ipsec_sa_lifetime", IKEEncryptionprofile.IPsecSALifetime)
	d.Set("ipsec_sa_replay_window_size", IKEEncryptionprofile.IPsecSAReplayWindowSize)
	d.Set("isakmp_authentication_mode", IKEEncryptionprofile.ISAKMPAuthenticationMode)
	d.Set("isakmp_diffie_helman_group_identifier", IKEEncryptionprofile.ISAKMPDiffieHelmanGroupIdentifier)
	d.Set("isakmp_encryption_algorithm", IKEEncryptionprofile.ISAKMPEncryptionAlgorithm)
	d.Set("isakmp_encryption_key_lifetime", IKEEncryptionprofile.ISAKMPEncryptionKeyLifetime)
	d.Set("isakmp_hash_algorithm", IKEEncryptionprofile.ISAKMPHashAlgorithm)
	d.Set("name", IKEEncryptionprofile.Name)
	d.Set("last_updated_by", IKEEncryptionprofile.LastUpdatedBy)
	d.Set("sequence", IKEEncryptionprofile.Sequence)
	d.Set("description", IKEEncryptionprofile.Description)
	d.Set("entity_scope", IKEEncryptionprofile.EntityScope)
	d.Set("associated_enterprise_id", IKEEncryptionprofile.AssociatedEnterpriseID)
	d.Set("external_id", IKEEncryptionprofile.ExternalID)

	d.Set("id", IKEEncryptionprofile.Identifier())
	d.Set("parent_id", IKEEncryptionprofile.ParentID)
	d.Set("parent_type", IKEEncryptionprofile.ParentType)
	d.Set("owner", IKEEncryptionprofile.Owner)

	d.SetId(IKEEncryptionprofile.Identifier())

	return nil
}
