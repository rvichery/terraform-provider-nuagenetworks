package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVRSRedeploymentpolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVRSRedeploymentpolicyCreate,
		Read:   resourceVRSRedeploymentpolicyRead,
		Update: resourceVRSRedeploymentpolicyUpdate,
		Delete: resourceVRSRedeploymentpolicyDelete,
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
			"al_ubr0_status_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cpu_utilization_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cpu_utilization_threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"vrs_corrective_action_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vrs_process_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vrsvsc_status_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redeployment_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_utilization_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"memory_utilization_threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"deployment_count_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jesxmon_process_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_disk_utilization_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"log_disk_utilization_threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"root_disk_utilization_redeployment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"root_disk_utilization_threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_vcenter_cluster": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
			},
			"parent_vcenter_data_center": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
			},
			"parent_vcenter": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter_vrs_config", "parent_vcenter_hypervisor"},
			},
			"parent_vcenter_vrs_config": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_hypervisor"},
			},
			"parent_vcenter_hypervisor": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_cluster", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config"},
			},
		},
	}
}

func resourceVRSRedeploymentpolicyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VRSRedeploymentpolicy object
	o := &vspk.VRSRedeploymentpolicy{}
	if attr, ok := d.GetOk("al_ubr0_status_redeployment_enabled"); ok {
		ALUbr0StatusRedeploymentEnabled := attr.(bool)
		o.ALUbr0StatusRedeploymentEnabled = &ALUbr0StatusRedeploymentEnabled
	}
	if attr, ok := d.GetOk("cpu_utilization_redeployment_enabled"); ok {
		CPUUtilizationRedeploymentEnabled := attr.(bool)
		o.CPUUtilizationRedeploymentEnabled = &CPUUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("cpu_utilization_threshold"); ok {
		o.CPUUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("vrs_corrective_action_delay"); ok {
		VRSCorrectiveActionDelay := attr.(int)
		o.VRSCorrectiveActionDelay = &VRSCorrectiveActionDelay
	}
	if attr, ok := d.GetOk("vrs_process_redeployment_enabled"); ok {
		VRSProcessRedeploymentEnabled := attr.(bool)
		o.VRSProcessRedeploymentEnabled = &VRSProcessRedeploymentEnabled
	}
	if attr, ok := d.GetOk("vrsvsc_status_redeployment_enabled"); ok {
		VRSVSCStatusRedeploymentEnabled := attr.(bool)
		o.VRSVSCStatusRedeploymentEnabled = &VRSVSCStatusRedeploymentEnabled
	}
	if attr, ok := d.GetOk("redeployment_delay"); ok {
		RedeploymentDelay := attr.(int)
		o.RedeploymentDelay = &RedeploymentDelay
	}
	if attr, ok := d.GetOk("memory_utilization_redeployment_enabled"); ok {
		MemoryUtilizationRedeploymentEnabled := attr.(bool)
		o.MemoryUtilizationRedeploymentEnabled = &MemoryUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("memory_utilization_threshold"); ok {
		o.MemoryUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("deployment_count_threshold"); ok {
		DeploymentCountThreshold := attr.(int)
		o.DeploymentCountThreshold = &DeploymentCountThreshold
	}
	if attr, ok := d.GetOk("jesxmon_process_redeployment_enabled"); ok {
		JesxmonProcessRedeploymentEnabled := attr.(bool)
		o.JesxmonProcessRedeploymentEnabled = &JesxmonProcessRedeploymentEnabled
	}
	if attr, ok := d.GetOk("log_disk_utilization_redeployment_enabled"); ok {
		LogDiskUtilizationRedeploymentEnabled := attr.(bool)
		o.LogDiskUtilizationRedeploymentEnabled = &LogDiskUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("log_disk_utilization_threshold"); ok {
		o.LogDiskUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("root_disk_utilization_redeployment_enabled"); ok {
		RootDiskUtilizationRedeploymentEnabled := attr.(bool)
		o.RootDiskUtilizationRedeploymentEnabled = &RootDiskUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("root_disk_utilization_threshold"); ok {
		o.RootDiskUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
		parent := &vspk.VCenterCluster{ID: attr.(string)}
		err := parent.CreateVRSRedeploymentpolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
		parent := &vspk.VCenterDataCenter{ID: attr.(string)}
		err := parent.CreateVRSRedeploymentpolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vcenter"); ok {
		parent := &vspk.VCenter{ID: attr.(string)}
		err := parent.CreateVRSRedeploymentpolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vcenter_vrs_config"); ok {
		parent := &vspk.VCenterVRSConfig{ID: attr.(string)}
		err := parent.CreateVRSRedeploymentpolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
		parent := &vspk.VCenterHypervisor{ID: attr.(string)}
		err := parent.CreateVRSRedeploymentpolicy(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceVRSRedeploymentpolicyRead(d, m)
}

func resourceVRSRedeploymentpolicyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VRSRedeploymentpolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("al_ubr0_status_redeployment_enabled", o.ALUbr0StatusRedeploymentEnabled)
	d.Set("cpu_utilization_redeployment_enabled", o.CPUUtilizationRedeploymentEnabled)
	d.Set("cpu_utilization_threshold", o.CPUUtilizationThreshold)
	d.Set("vrs_corrective_action_delay", o.VRSCorrectiveActionDelay)
	d.Set("vrs_process_redeployment_enabled", o.VRSProcessRedeploymentEnabled)
	d.Set("vrsvsc_status_redeployment_enabled", o.VRSVSCStatusRedeploymentEnabled)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("redeployment_delay", o.RedeploymentDelay)
	d.Set("memory_utilization_redeployment_enabled", o.MemoryUtilizationRedeploymentEnabled)
	d.Set("memory_utilization_threshold", o.MemoryUtilizationThreshold)
	d.Set("deployment_count_threshold", o.DeploymentCountThreshold)
	d.Set("jesxmon_process_redeployment_enabled", o.JesxmonProcessRedeploymentEnabled)
	d.Set("entity_scope", o.EntityScope)
	d.Set("log_disk_utilization_redeployment_enabled", o.LogDiskUtilizationRedeploymentEnabled)
	d.Set("log_disk_utilization_threshold", o.LogDiskUtilizationThreshold)
	d.Set("root_disk_utilization_redeployment_enabled", o.RootDiskUtilizationRedeploymentEnabled)
	d.Set("root_disk_utilization_threshold", o.RootDiskUtilizationThreshold)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVRSRedeploymentpolicyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VRSRedeploymentpolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("al_ubr0_status_redeployment_enabled"); ok {
		ALUbr0StatusRedeploymentEnabled := attr.(bool)
		o.ALUbr0StatusRedeploymentEnabled = &ALUbr0StatusRedeploymentEnabled
	}
	if attr, ok := d.GetOk("cpu_utilization_redeployment_enabled"); ok {
		CPUUtilizationRedeploymentEnabled := attr.(bool)
		o.CPUUtilizationRedeploymentEnabled = &CPUUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("cpu_utilization_threshold"); ok {
		o.CPUUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("vrs_corrective_action_delay"); ok {
		VRSCorrectiveActionDelay := attr.(int)
		o.VRSCorrectiveActionDelay = &VRSCorrectiveActionDelay
	}
	if attr, ok := d.GetOk("vrs_process_redeployment_enabled"); ok {
		VRSProcessRedeploymentEnabled := attr.(bool)
		o.VRSProcessRedeploymentEnabled = &VRSProcessRedeploymentEnabled
	}
	if attr, ok := d.GetOk("vrsvsc_status_redeployment_enabled"); ok {
		VRSVSCStatusRedeploymentEnabled := attr.(bool)
		o.VRSVSCStatusRedeploymentEnabled = &VRSVSCStatusRedeploymentEnabled
	}
	if attr, ok := d.GetOk("redeployment_delay"); ok {
		RedeploymentDelay := attr.(int)
		o.RedeploymentDelay = &RedeploymentDelay
	}
	if attr, ok := d.GetOk("memory_utilization_redeployment_enabled"); ok {
		MemoryUtilizationRedeploymentEnabled := attr.(bool)
		o.MemoryUtilizationRedeploymentEnabled = &MemoryUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("memory_utilization_threshold"); ok {
		o.MemoryUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("deployment_count_threshold"); ok {
		DeploymentCountThreshold := attr.(int)
		o.DeploymentCountThreshold = &DeploymentCountThreshold
	}
	if attr, ok := d.GetOk("jesxmon_process_redeployment_enabled"); ok {
		JesxmonProcessRedeploymentEnabled := attr.(bool)
		o.JesxmonProcessRedeploymentEnabled = &JesxmonProcessRedeploymentEnabled
	}
	if attr, ok := d.GetOk("log_disk_utilization_redeployment_enabled"); ok {
		LogDiskUtilizationRedeploymentEnabled := attr.(bool)
		o.LogDiskUtilizationRedeploymentEnabled = &LogDiskUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("log_disk_utilization_threshold"); ok {
		o.LogDiskUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("root_disk_utilization_redeployment_enabled"); ok {
		RootDiskUtilizationRedeploymentEnabled := attr.(bool)
		o.RootDiskUtilizationRedeploymentEnabled = &RootDiskUtilizationRedeploymentEnabled
	}
	if attr, ok := d.GetOk("root_disk_utilization_threshold"); ok {
		o.RootDiskUtilizationThreshold = attr.(float64)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVRSRedeploymentpolicyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VRSRedeploymentpolicy{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
