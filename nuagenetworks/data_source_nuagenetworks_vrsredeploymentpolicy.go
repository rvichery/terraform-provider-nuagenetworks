package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVRSRedeploymentpolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVRSRedeploymentpolicyRead,
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
            "al_ubr0_status_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "cpu_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "cpu_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "vrs_corrective_action_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vrs_process_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vrsvsc_status_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redeployment_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "memory_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "memory_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "deployment_count_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "jesxmon_process_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "log_disk_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "log_disk_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "root_disk_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "root_disk_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
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


func dataSourceVRSRedeploymentpolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredVRSRedeploymentpolicies := vspk.VRSRedeploymentpoliciesList{}
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
        filteredVRSRedeploymentpolicies, err = parent.VRSRedeploymentpolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        filteredVRSRedeploymentpolicies, err = parent.VRSRedeploymentpolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter"); ok {
        parent := &vspk.VCenter{ID: attr.(string)}
        filteredVRSRedeploymentpolicies, err = parent.VRSRedeploymentpolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_vrs_config"); ok {
        parent := &vspk.VCenterVRSConfig{ID: attr.(string)}
        filteredVRSRedeploymentpolicies, err = parent.VRSRedeploymentpolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
        parent := &vspk.VCenterHypervisor{ID: attr.(string)}
        filteredVRSRedeploymentpolicies, err = parent.VRSRedeploymentpolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    VRSRedeploymentpolicy := &vspk.VRSRedeploymentpolicy{}

    if len(filteredVRSRedeploymentpolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVRSRedeploymentpolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VRSRedeploymentpolicy = filteredVRSRedeploymentpolicies[0]

    d.Set("al_ubr0_status_redeployment_enabled", VRSRedeploymentpolicy.ALUbr0StatusRedeploymentEnabled)
    d.Set("cpu_utilization_redeployment_enabled", VRSRedeploymentpolicy.CPUUtilizationRedeploymentEnabled)
    d.Set("cpu_utilization_threshold", VRSRedeploymentpolicy.CPUUtilizationThreshold)
    d.Set("vrs_corrective_action_delay", VRSRedeploymentpolicy.VRSCorrectiveActionDelay)
    d.Set("vrs_process_redeployment_enabled", VRSRedeploymentpolicy.VRSProcessRedeploymentEnabled)
    d.Set("vrsvsc_status_redeployment_enabled", VRSRedeploymentpolicy.VRSVSCStatusRedeploymentEnabled)
    d.Set("last_updated_by", VRSRedeploymentpolicy.LastUpdatedBy)
    d.Set("redeployment_delay", VRSRedeploymentpolicy.RedeploymentDelay)
    d.Set("memory_utilization_redeployment_enabled", VRSRedeploymentpolicy.MemoryUtilizationRedeploymentEnabled)
    d.Set("memory_utilization_threshold", VRSRedeploymentpolicy.MemoryUtilizationThreshold)
    d.Set("deployment_count_threshold", VRSRedeploymentpolicy.DeploymentCountThreshold)
    d.Set("jesxmon_process_redeployment_enabled", VRSRedeploymentpolicy.JesxmonProcessRedeploymentEnabled)
    d.Set("entity_scope", VRSRedeploymentpolicy.EntityScope)
    d.Set("log_disk_utilization_redeployment_enabled", VRSRedeploymentpolicy.LogDiskUtilizationRedeploymentEnabled)
    d.Set("log_disk_utilization_threshold", VRSRedeploymentpolicy.LogDiskUtilizationThreshold)
    d.Set("root_disk_utilization_redeployment_enabled", VRSRedeploymentpolicy.RootDiskUtilizationRedeploymentEnabled)
    d.Set("root_disk_utilization_threshold", VRSRedeploymentpolicy.RootDiskUtilizationThreshold)
    d.Set("external_id", VRSRedeploymentpolicy.ExternalID)
    
    d.Set("id", VRSRedeploymentpolicy.Identifier())
    d.Set("parent_id", VRSRedeploymentpolicy.ParentID)
    d.Set("parent_type", VRSRedeploymentpolicy.ParentType)
    d.Set("owner", VRSRedeploymentpolicy.Owner)

    d.SetId(VRSRedeploymentpolicy.Identifier())
    
    return nil
}