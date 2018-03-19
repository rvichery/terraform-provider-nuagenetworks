package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVLANTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVLANTemplateRead,
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
            "value": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "is_uplink": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "associated_connection_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ingress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_uplink_connection_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vsc_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "duc_vlan": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_port_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_port_template"},
            },
            "parent_port_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port_template"},
            },
        },
    }
}


func dataSourceVLANTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredVLANTemplates := vspk.VLANTemplatesList{}
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
    if attr, ok := d.GetOk("parent_ns_port_template"); ok {
        parent := &vspk.NSPortTemplate{ID: attr.(string)}
        filteredVLANTemplates, err = parent.VLANTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_port_template"); ok {
        parent := &vspk.PortTemplate{ID: attr.(string)}
        filteredVLANTemplates, err = parent.VLANTemplates(fetchFilter)
        if err != nil {
            return err
        }
    }

    VLANTemplate := &vspk.VLANTemplate{}

    if len(filteredVLANTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVLANTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VLANTemplate = filteredVLANTemplates[0]

    d.Set("value", VLANTemplate.Value)
    d.Set("last_updated_by", VLANTemplate.LastUpdatedBy)
    d.Set("description", VLANTemplate.Description)
    d.Set("entity_scope", VLANTemplate.EntityScope)
    d.Set("is_uplink", VLANTemplate.IsUplink)
    d.Set("associated_connection_type", VLANTemplate.AssociatedConnectionType)
    d.Set("associated_egress_qos_policy_id", VLANTemplate.AssociatedEgressQOSPolicyID)
    d.Set("associated_ingress_qos_policy_id", VLANTemplate.AssociatedIngressQOSPolicyID)
    d.Set("associated_uplink_connection_id", VLANTemplate.AssociatedUplinkConnectionID)
    d.Set("associated_vsc_profile_id", VLANTemplate.AssociatedVSCProfileID)
    d.Set("duc_vlan", VLANTemplate.DucVlan)
    d.Set("external_id", VLANTemplate.ExternalID)
    d.Set("type", VLANTemplate.Type)
    
    d.Set("id", VLANTemplate.Identifier())
    d.Set("parent_id", VLANTemplate.ParentID)
    d.Set("parent_type", VLANTemplate.ParentType)
    d.Set("owner", VLANTemplate.Owner)

    d.SetId(VLANTemplate.Identifier())
    
    return nil
}