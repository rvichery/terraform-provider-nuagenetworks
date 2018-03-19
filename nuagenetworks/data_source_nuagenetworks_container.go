package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceContainer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContainerRead,
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
			"l2_domain_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vrsid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
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
			"reason_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_expiry": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"delete_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resync_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"zone_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"orchestration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_qos": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_vrs": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_egress_acl_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_user": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_enterprise", "parent_ingress_acl_template"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_ingress_acl_template"},
			},
			"parent_ingress_acl_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise"},
			},
		},
	}
}

func dataSourceContainerRead(d *schema.ResourceData, m interface{}) error {
	filteredContainers := vspk.ContainersList{}
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
	if attr, ok := d.GetOk("parent_qos"); ok {
		parent := &vspk.QOS{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vrs"); ok {
		parent := &vspk.VRS{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_egress_acl_template"); ok {
		parent := &vspk.EgressACLTemplate{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_user"); ok {
		parent := &vspk.User{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
		parent := &vspk.IngressACLTemplate{ID: attr.(string)}
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredContainers, err = parent.Containers(fetchFilter)
		if err != nil {
			return err
		}
	}

	Container := &vspk.Container{}

	if len(filteredContainers) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredContainers) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Container = filteredContainers[0]

	d.Set("l2_domain_ids", Container.L2DomainIDs)
	d.Set("vrsid", Container.VRSID)
	d.Set("uuid", Container.UUID)
	d.Set("name", Container.Name)
	d.Set("last_updated_by", Container.LastUpdatedBy)
	d.Set("reason_type", Container.ReasonType)
	d.Set("delete_expiry", Container.DeleteExpiry)
	d.Set("delete_mode", Container.DeleteMode)
	d.Set("resync_info", Container.ResyncInfo)
	d.Set("site_identifier", Container.SiteIdentifier)
	d.Set("image_id", Container.ImageID)
	d.Set("image_name", Container.ImageName)
	d.Set("interfaces", Container.Interfaces)
	d.Set("enterprise_id", Container.EnterpriseID)
	d.Set("enterprise_name", Container.EnterpriseName)
	d.Set("entity_scope", Container.EntityScope)
	d.Set("domain_ids", Container.DomainIDs)
	d.Set("zone_ids", Container.ZoneIDs)
	d.Set("orchestration_id", Container.OrchestrationID)
	d.Set("user_id", Container.UserID)
	d.Set("user_name", Container.UserName)
	d.Set("status", Container.Status)
	d.Set("subnet_ids", Container.SubnetIDs)
	d.Set("external_id", Container.ExternalID)
	d.Set("hypervisor_ip", Container.HypervisorIP)

	d.Set("id", Container.Identifier())
	d.Set("parent_id", Container.ParentID)
	d.Set("parent_type", Container.ParentType)
	d.Set("owner", Container.Owner)

	d.SetId(Container.Identifier())

	return nil
}
