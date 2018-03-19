package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLicenseStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseStatusRead,
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
			"accumulate_licenses_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"total_licensed_avrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_avrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_gateways_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_nics_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_nsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_avrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_avrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_nics_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_nsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_vms_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_vrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_used_vrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_vms_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_vrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_licensed_vrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_used_gateways_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceLicenseStatusRead(d *schema.ResourceData, m interface{}) error {
	filteredLicenseStatus := vspk.LicenseStatusList{}
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
	filteredLicenseStatus, err = parent.LicenseStatus(fetchFilter)
	if err != nil {
		return err
	}

	LicenseStatus := &vspk.LicenseStatus{}

	if len(filteredLicenseStatus) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLicenseStatus) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	LicenseStatus = filteredLicenseStatus[0]

	d.Set("accumulate_licenses_enabled", LicenseStatus.AccumulateLicensesEnabled)
	d.Set("total_licensed_avrsgs_count", LicenseStatus.TotalLicensedAVRSGsCount)
	d.Set("total_licensed_avrss_count", LicenseStatus.TotalLicensedAVRSsCount)
	d.Set("total_licensed_gateways_count", LicenseStatus.TotalLicensedGatewaysCount)
	d.Set("total_licensed_nics_count", LicenseStatus.TotalLicensedNICsCount)
	d.Set("total_licensed_nsgs_count", LicenseStatus.TotalLicensedNSGsCount)
	d.Set("total_licensed_used_avrsgs_count", LicenseStatus.TotalLicensedUsedAVRSGsCount)
	d.Set("total_licensed_used_avrss_count", LicenseStatus.TotalLicensedUsedAVRSsCount)
	d.Set("total_licensed_used_nics_count", LicenseStatus.TotalLicensedUsedNICsCount)
	d.Set("total_licensed_used_nsgs_count", LicenseStatus.TotalLicensedUsedNSGsCount)
	d.Set("total_licensed_used_vms_count", LicenseStatus.TotalLicensedUsedVMsCount)
	d.Set("total_licensed_used_vrsgs_count", LicenseStatus.TotalLicensedUsedVRSGsCount)
	d.Set("total_licensed_used_vrss_count", LicenseStatus.TotalLicensedUsedVRSsCount)
	d.Set("total_licensed_vms_count", LicenseStatus.TotalLicensedVMsCount)
	d.Set("total_licensed_vrsgs_count", LicenseStatus.TotalLicensedVRSGsCount)
	d.Set("total_licensed_vrss_count", LicenseStatus.TotalLicensedVRSsCount)
	d.Set("total_used_gateways_count", LicenseStatus.TotalUsedGatewaysCount)

	d.Set("id", LicenseStatus.Identifier())
	d.Set("parent_id", LicenseStatus.ParentID)
	d.Set("parent_type", LicenseStatus.ParentType)
	d.Set("owner", LicenseStatus.Owner)

	d.SetId(LicenseStatus.Identifier())

	return nil
}
