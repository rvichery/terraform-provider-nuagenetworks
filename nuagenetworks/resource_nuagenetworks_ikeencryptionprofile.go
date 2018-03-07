package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"dpd_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "REPLY_ONLY",
			},
			"dpd_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ipsec_authentication_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HMAC_SHA256",
			},
			"ipsec_dont_fragment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipsec_enable_pfs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipsec_encryption_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES256",
			},
			"ipsec_pre_fragment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipsec_sa_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
			},
			"ipsec_sa_replay_window_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WINDOW_SIZE_32",
			},
			"isakmp_authentication_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PRE_SHARED_KEY",
			},
			"isakmp_diffie_helman_group_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GROUP_5_1536_BIT_DH",
			},
			"isakmp_encryption_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES256",
			},
			"isakmp_encryption_key_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  28800,
			},
			"isakmp_hash_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SHA256",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
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
		o.DPDInterval = attr.(int)
	}
	if attr, ok := d.GetOk("dpd_mode"); ok {
		o.DPDMode = attr.(string)
	}
	if attr, ok := d.GetOk("dpd_timeout"); ok {
		o.DPDTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("ipsec_authentication_algorithm"); ok {
		o.IPsecAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_dont_fragment"); ok {
		o.IPsecDontFragment = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_enable_pfs"); ok {
		o.IPsecEnablePFS = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_encryption_algorithm"); ok {
		o.IPsecEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_pre_fragment"); ok {
		o.IPsecPreFragment = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_sa_lifetime"); ok {
		o.IPsecSALifetime = attr.(int)
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
		o.ISAKMPEncryptionKeyLifetime = attr.(int)
	}
	if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
		o.ISAKMPHashAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		o.Sequence = attr.(int)
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
		o.DPDInterval = attr.(int)
	}
	if attr, ok := d.GetOk("dpd_mode"); ok {
		o.DPDMode = attr.(string)
	}
	if attr, ok := d.GetOk("dpd_timeout"); ok {
		o.DPDTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("ipsec_authentication_algorithm"); ok {
		o.IPsecAuthenticationAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_dont_fragment"); ok {
		o.IPsecDontFragment = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_enable_pfs"); ok {
		o.IPsecEnablePFS = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_encryption_algorithm"); ok {
		o.IPsecEncryptionAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("ipsec_pre_fragment"); ok {
		o.IPsecPreFragment = attr.(bool)
	}
	if attr, ok := d.GetOk("ipsec_sa_lifetime"); ok {
		o.IPsecSALifetime = attr.(int)
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
		o.ISAKMPEncryptionKeyLifetime = attr.(int)
	}
	if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
		o.ISAKMPHashAlgorithm = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("sequence"); ok {
		o.Sequence = attr.(int)
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
