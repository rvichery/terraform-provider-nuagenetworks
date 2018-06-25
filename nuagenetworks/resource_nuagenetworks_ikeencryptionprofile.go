package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceIKEEncryptionprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIKEEncryptionprofileCreate,
		Read:   resourceIKEEncryptionprofileRead,
		Update: resourceIKEEncryptionprofileUpdate,
		Delete: resourceIKEEncryptionprofileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"dpd_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "REPLY_ONLY",
			},
			"dpd_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ipsec_authentication_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HMAC_SHA256",
			},
			"ipsec_dont_fragment": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipsec_enable_pfs": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipsec_encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES256",
			},
			"ipsec_pre_fragment": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipsec_sa_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
			},
			"ipsec_sa_replay_window_size": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WINDOW_SIZE_32",
			},
			"isakmp_authentication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PRE_SHARED_KEY",
			},
			"isakmp_diffie_helman_group_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GROUP_5_1536_BIT_DH",
			},
			"isakmp_encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES256",
			},
			"isakmp_encryption_key_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  28800,
			},
			"isakmp_hash_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SHA256",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sequence": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIKEEncryptionprofileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IKEEncryptionprofile object
	o := &vspk.IKEEncryptionprofile{}
	if attr, ok := d.GetOk("dpd_interval"); ok {
		DPDInterval := attr.(int)
		o.DPDInterval = &DPDInterval
	}
	if attr, ok := d.GetOk("dpd_mode"); ok {
		o.DPDMode = attr.(string)
	}
	if attr, ok := d.GetOk("dpd_timeout"); ok {
		DPDTimeout := attr.(int)
		o.DPDTimeout = &DPDTimeout
	}
	if attr, ok := d.GetOk("ipsec_authentication_algorithm"); ok {
		o.IPsecAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_dont_fragment"); ok {
		IPsecDontFragment := attr.(bool)
		o.IPsecDontFragment = &IPsecDontFragment
	}
	if attr, ok := d.GetOk("ipsec_enable_pfs"); ok {
		IPsecEnablePFS := attr.(bool)
		o.IPsecEnablePFS = &IPsecEnablePFS
	}
	if attr, ok := d.GetOk("ipsec_encryption_algorithm"); ok {
		o.IPsecEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_pre_fragment"); ok {
		IPsecPreFragment := attr.(bool)
		o.IPsecPreFragment = &IPsecPreFragment
	}
	if attr, ok := d.GetOk("ipsec_sa_lifetime"); ok {
		IPsecSALifetime := attr.(int)
		o.IPsecSALifetime = &IPsecSALifetime
	}
	if attr, ok := d.GetOk("ipsec_sa_replay_window_size"); ok {
		o.IPsecSAReplayWindowSize = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_authentication_mode"); ok {
		o.ISAKMPAuthenticationMode = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_diffie_helman_group_identifier"); ok {
		o.ISAKMPDiffieHelmanGroupIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_encryption_algorithm"); ok {
		o.ISAKMPEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_encryption_key_lifetime"); ok {
		ISAKMPEncryptionKeyLifetime := attr.(int)
		o.ISAKMPEncryptionKeyLifetime = &ISAKMPEncryptionKeyLifetime
	}
	if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
		o.ISAKMPHashAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		Sequence := attr.(int)
		o.Sequence = &Sequence
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateIKEEncryptionprofile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceIKEEncryptionprofileRead(d, m)
}

func resourceIKEEncryptionprofileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEEncryptionprofile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dpd_interval", o.DPDInterval)
	d.Set("dpd_mode", o.DPDMode)
	d.Set("dpd_timeout", o.DPDTimeout)
	d.Set("ipsec_authentication_algorithm", o.IPsecAuthenticationAlgorithm)
	d.Set("ipsec_dont_fragment", o.IPsecDontFragment)
	d.Set("ipsec_enable_pfs", o.IPsecEnablePFS)
	d.Set("ipsec_encryption_algorithm", o.IPsecEncryptionAlgorithm)
	d.Set("ipsec_pre_fragment", o.IPsecPreFragment)
	d.Set("ipsec_sa_lifetime", o.IPsecSALifetime)
	d.Set("ipsec_sa_replay_window_size", o.IPsecSAReplayWindowSize)
	d.Set("isakmp_authentication_mode", o.ISAKMPAuthenticationMode)
	d.Set("isakmp_diffie_helman_group_identifier", o.ISAKMPDiffieHelmanGroupIdentifier)
	d.Set("isakmp_encryption_algorithm", o.ISAKMPEncryptionAlgorithm)
	d.Set("isakmp_encryption_key_lifetime", o.ISAKMPEncryptionKeyLifetime)
	d.Set("isakmp_hash_algorithm", o.ISAKMPHashAlgorithm)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("sequence", o.Sequence)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIKEEncryptionprofileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEEncryptionprofile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("dpd_interval"); ok {
		DPDInterval := attr.(int)
		o.DPDInterval = &DPDInterval
	}
	if attr, ok := d.GetOk("dpd_mode"); ok {
		o.DPDMode = attr.(string)
	}
	if attr, ok := d.GetOk("dpd_timeout"); ok {
		DPDTimeout := attr.(int)
		o.DPDTimeout = &DPDTimeout
	}
	if attr, ok := d.GetOk("ipsec_authentication_algorithm"); ok {
		o.IPsecAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_dont_fragment"); ok {
		IPsecDontFragment := attr.(bool)
		o.IPsecDontFragment = &IPsecDontFragment
	}
	if attr, ok := d.GetOk("ipsec_enable_pfs"); ok {
		IPsecEnablePFS := attr.(bool)
		o.IPsecEnablePFS = &IPsecEnablePFS
	}
	if attr, ok := d.GetOk("ipsec_encryption_algorithm"); ok {
		o.IPsecEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_pre_fragment"); ok {
		IPsecPreFragment := attr.(bool)
		o.IPsecPreFragment = &IPsecPreFragment
	}
	if attr, ok := d.GetOk("ipsec_sa_lifetime"); ok {
		IPsecSALifetime := attr.(int)
		o.IPsecSALifetime = &IPsecSALifetime
	}
	if attr, ok := d.GetOk("ipsec_sa_replay_window_size"); ok {
		o.IPsecSAReplayWindowSize = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_authentication_mode"); ok {
		o.ISAKMPAuthenticationMode = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_diffie_helman_group_identifier"); ok {
		o.ISAKMPDiffieHelmanGroupIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_encryption_algorithm"); ok {
		o.ISAKMPEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("isakmp_encryption_key_lifetime"); ok {
		ISAKMPEncryptionKeyLifetime := attr.(int)
		o.ISAKMPEncryptionKeyLifetime = &ISAKMPEncryptionKeyLifetime
	}
	if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
		o.ISAKMPHashAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		Sequence := attr.(int)
		o.Sequence = &Sequence
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIKEEncryptionprofileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEEncryptionprofile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
