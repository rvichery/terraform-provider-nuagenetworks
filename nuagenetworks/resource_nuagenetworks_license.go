package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceLicense() *schema.Resource {
	return &schema.Resource{
		Create: resourceLicenseCreate,
		Read:   resourceLicenseRead,
		Update: resourceLicenseUpdate,
		Delete: resourceLicenseDelete,
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
			"major_release": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_supported_versions": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_encryption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_entities": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minor_release": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowed_avrsgs_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_avrss_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_cpes_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_nics_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_vms_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_vrsgs_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_vrss_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"unique_license_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_cluster_license": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"street": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customer_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expiration_date": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"expiry_timestamp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLicenseCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize License object
	o := &vspk.License{
		License: d.Get("license").(string),
	}
	if attr, ok := d.GetOk("major_release"); ok {
		MajorRelease := attr.(int)
		o.MajorRelease = &MajorRelease
	}
	if attr, ok := d.GetOk("additional_supported_versions"); ok {
		o.AdditionalSupportedVersions = attr.(string)
	}
	if attr, ok := d.GetOk("phone"); ok {
		o.Phone = attr.(string)
	}
	if attr, ok := d.GetOk("license_encryption"); ok {
		o.LicenseEncryption = attr.(string)
	}
	if attr, ok := d.GetOk("license_entities"); ok {
		o.LicenseEntities = attr.(string)
	}
	if attr, ok := d.GetOk("license_id"); ok {
		LicenseID := attr.(int)
		o.LicenseID = &LicenseID
	}
	if attr, ok := d.GetOk("license_type"); ok {
		o.LicenseType = attr.(string)
	}
	if attr, ok := d.GetOk("minor_release"); ok {
		MinorRelease := attr.(int)
		o.MinorRelease = &MinorRelease
	}
	if attr, ok := d.GetOk("zip"); ok {
		o.Zip = attr.(string)
	}
	if attr, ok := d.GetOk("city"); ok {
		o.City = attr.(string)
	}
	if attr, ok := d.GetOk("allowed_avrsgs_count"); ok {
		AllowedAVRSGsCount := attr.(int)
		o.AllowedAVRSGsCount = &AllowedAVRSGsCount
	}
	if attr, ok := d.GetOk("allowed_avrss_count"); ok {
		AllowedAVRSsCount := attr.(int)
		o.AllowedAVRSsCount = &AllowedAVRSsCount
	}
	if attr, ok := d.GetOk("allowed_cpes_count"); ok {
		AllowedCPEsCount := attr.(int)
		o.AllowedCPEsCount = &AllowedCPEsCount
	}
	if attr, ok := d.GetOk("allowed_nics_count"); ok {
		AllowedNICsCount := attr.(int)
		o.AllowedNICsCount = &AllowedNICsCount
	}
	if attr, ok := d.GetOk("allowed_vms_count"); ok {
		AllowedVMsCount := attr.(int)
		o.AllowedVMsCount = &AllowedVMsCount
	}
	if attr, ok := d.GetOk("allowed_vrsgs_count"); ok {
		AllowedVRSGsCount := attr.(int)
		o.AllowedVRSGsCount = &AllowedVRSGsCount
	}
	if attr, ok := d.GetOk("allowed_vrss_count"); ok {
		AllowedVRSsCount := attr.(int)
		o.AllowedVRSsCount = &AllowedVRSsCount
	}
	if attr, ok := d.GetOk("email"); ok {
		o.Email = attr.(string)
	}
	if attr, ok := d.GetOk("encryption_mode"); ok {
		EncryptionMode := attr.(bool)
		o.EncryptionMode = &EncryptionMode
	}
	if attr, ok := d.GetOk("unique_license_identifier"); ok {
		o.UniqueLicenseIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("company"); ok {
		o.Company = attr.(string)
	}
	if attr, ok := d.GetOk("country"); ok {
		o.Country = attr.(string)
	}
	if attr, ok := d.GetOk("product_version"); ok {
		o.ProductVersion = attr.(string)
	}
	if attr, ok := d.GetOk("is_cluster_license"); ok {
		IsClusterLicense := attr.(bool)
		o.IsClusterLicense = &IsClusterLicense
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("state"); ok {
		o.State = attr.(string)
	}
	if attr, ok := d.GetOk("street"); ok {
		o.Street = attr.(string)
	}
	if attr, ok := d.GetOk("customer_key"); ok {
		o.CustomerKey = attr.(string)
	}
	if attr, ok := d.GetOk("expiration_date"); ok {
		o.ExpirationDate = attr.(float64)
	}
	if attr, ok := d.GetOk("expiry_timestamp"); ok {
		ExpiryTimestamp := attr.(int)
		o.ExpiryTimestamp = &ExpiryTimestamp
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateLicense(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceLicenseRead(d, m)
}

func resourceLicenseRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.License{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("major_release", o.MajorRelease)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("additional_supported_versions", o.AdditionalSupportedVersions)
	d.Set("phone", o.Phone)
	d.Set("license", o.License)
	d.Set("license_encryption", o.LicenseEncryption)
	d.Set("license_entities", o.LicenseEntities)
	d.Set("license_id", o.LicenseID)
	d.Set("license_type", o.LicenseType)
	d.Set("minor_release", o.MinorRelease)
	d.Set("zip", o.Zip)
	d.Set("city", o.City)
	d.Set("allowed_avrsgs_count", o.AllowedAVRSGsCount)
	d.Set("allowed_avrss_count", o.AllowedAVRSsCount)
	d.Set("allowed_cpes_count", o.AllowedCPEsCount)
	d.Set("allowed_nics_count", o.AllowedNICsCount)
	d.Set("allowed_vms_count", o.AllowedVMsCount)
	d.Set("allowed_vrsgs_count", o.AllowedVRSGsCount)
	d.Set("allowed_vrss_count", o.AllowedVRSsCount)
	d.Set("email", o.Email)
	d.Set("encryption_mode", o.EncryptionMode)
	d.Set("unique_license_identifier", o.UniqueLicenseIdentifier)
	d.Set("entity_scope", o.EntityScope)
	d.Set("company", o.Company)
	d.Set("country", o.Country)
	d.Set("product_version", o.ProductVersion)

	d.Set("is_cluster_license", o.IsClusterLicense)
	d.Set("user_name", o.UserName)
	d.Set("state", o.State)
	d.Set("street", o.Street)
	d.Set("customer_key", o.CustomerKey)
	d.Set("expiration_date", o.ExpirationDate)
	d.Set("expiry_timestamp", o.ExpiryTimestamp)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceLicenseUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.License{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.License = d.Get("license").(string)

	if attr, ok := d.GetOk("major_release"); ok {
		MajorRelease := attr.(int)
		o.MajorRelease = &MajorRelease
	}
	if attr, ok := d.GetOk("additional_supported_versions"); ok {
		o.AdditionalSupportedVersions = attr.(string)
	}
	if attr, ok := d.GetOk("phone"); ok {
		o.Phone = attr.(string)
	}
	if attr, ok := d.GetOk("license_encryption"); ok {
		o.LicenseEncryption = attr.(string)
	}
	if attr, ok := d.GetOk("license_entities"); ok {
		o.LicenseEntities = attr.(string)
	}
	if attr, ok := d.GetOk("license_id"); ok {
		LicenseID := attr.(int)
		o.LicenseID = &LicenseID
	}
	if attr, ok := d.GetOk("license_type"); ok {
		o.LicenseType = attr.(string)
	}
	if attr, ok := d.GetOk("minor_release"); ok {
		MinorRelease := attr.(int)
		o.MinorRelease = &MinorRelease
	}
	if attr, ok := d.GetOk("zip"); ok {
		o.Zip = attr.(string)
	}
	if attr, ok := d.GetOk("city"); ok {
		o.City = attr.(string)
	}
	if attr, ok := d.GetOk("allowed_avrsgs_count"); ok {
		AllowedAVRSGsCount := attr.(int)
		o.AllowedAVRSGsCount = &AllowedAVRSGsCount
	}
	if attr, ok := d.GetOk("allowed_avrss_count"); ok {
		AllowedAVRSsCount := attr.(int)
		o.AllowedAVRSsCount = &AllowedAVRSsCount
	}
	if attr, ok := d.GetOk("allowed_cpes_count"); ok {
		AllowedCPEsCount := attr.(int)
		o.AllowedCPEsCount = &AllowedCPEsCount
	}
	if attr, ok := d.GetOk("allowed_nics_count"); ok {
		AllowedNICsCount := attr.(int)
		o.AllowedNICsCount = &AllowedNICsCount
	}
	if attr, ok := d.GetOk("allowed_vms_count"); ok {
		AllowedVMsCount := attr.(int)
		o.AllowedVMsCount = &AllowedVMsCount
	}
	if attr, ok := d.GetOk("allowed_vrsgs_count"); ok {
		AllowedVRSGsCount := attr.(int)
		o.AllowedVRSGsCount = &AllowedVRSGsCount
	}
	if attr, ok := d.GetOk("allowed_vrss_count"); ok {
		AllowedVRSsCount := attr.(int)
		o.AllowedVRSsCount = &AllowedVRSsCount
	}
	if attr, ok := d.GetOk("email"); ok {
		o.Email = attr.(string)
	}
	if attr, ok := d.GetOk("encryption_mode"); ok {
		EncryptionMode := attr.(bool)
		o.EncryptionMode = &EncryptionMode
	}
	if attr, ok := d.GetOk("unique_license_identifier"); ok {
		o.UniqueLicenseIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("company"); ok {
		o.Company = attr.(string)
	}
	if attr, ok := d.GetOk("country"); ok {
		o.Country = attr.(string)
	}
	if attr, ok := d.GetOk("product_version"); ok {
		o.ProductVersion = attr.(string)
	}
	if attr, ok := d.GetOk("is_cluster_license"); ok {
		IsClusterLicense := attr.(bool)
		o.IsClusterLicense = &IsClusterLicense
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("state"); ok {
		o.State = attr.(string)
	}
	if attr, ok := d.GetOk("street"); ok {
		o.Street = attr.(string)
	}
	if attr, ok := d.GetOk("customer_key"); ok {
		o.CustomerKey = attr.(string)
	}
	if attr, ok := d.GetOk("expiration_date"); ok {
		o.ExpirationDate = attr.(float64)
	}
	if attr, ok := d.GetOk("expiry_timestamp"); ok {
		ExpiryTimestamp := attr.(int)
		o.ExpiryTimestamp = &ExpiryTimestamp
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceLicenseDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.License{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
