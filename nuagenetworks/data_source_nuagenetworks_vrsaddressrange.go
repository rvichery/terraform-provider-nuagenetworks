package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVRSAddressRange() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVRSAddressRangeRead,
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
            "max_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "min_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_vrs_config": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_hypervisor"},
            },
            "parent_vcenter_hypervisor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config"},
            },
        },
    }
}


func dataSourceVRSAddressRangeRead(d *schema.ResourceData, m interface{}) error {
    filteredVRSAddressRanges := vspk.VRSAddressRangesList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        filteredVRSAddressRanges, err = parent.VRSAddressRanges(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        filteredVRSAddressRanges, err = parent.VRSAddressRanges(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter"); ok {
        parent := &vspk.VCenter{ID: attr.(string)}
        filteredVRSAddressRanges, err = parent.VRSAddressRanges(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_vrs_config"); ok {
        parent := &vspk.VCenterVRSConfig{ID: attr.(string)}
        filteredVRSAddressRanges, err = parent.VRSAddressRanges(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
        parent := &vspk.VCenterHypervisor{ID: attr.(string)}
        filteredVRSAddressRanges, err = parent.VRSAddressRanges(fetchFilter)
        if err != nil {
            return err
        }
    }

    VRSAddressRange := &vspk.VRSAddressRange{}

    if len(filteredVRSAddressRanges) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVRSAddressRanges) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VRSAddressRange = filteredVRSAddressRanges[0]

    d.Set("last_updated_by", VRSAddressRange.LastUpdatedBy)
    d.Set("max_address", VRSAddressRange.MaxAddress)
    d.Set("min_address", VRSAddressRange.MinAddress)
    d.Set("entity_scope", VRSAddressRange.EntityScope)
    d.Set("external_id", VRSAddressRange.ExternalID)
    
    d.Set("id", VRSAddressRange.Identifier())
    d.Set("parent_id", VRSAddressRange.ParentID)
    d.Set("parent_type", VRSAddressRange.ParentType)
    d.Set("owner", VRSAddressRange.Owner)

    d.SetId(VRSAddressRange.Identifier())
    
    return nil
}