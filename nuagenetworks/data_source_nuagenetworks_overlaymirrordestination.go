package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayMirrorDestination() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayMirrorDestinationRead,
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
            "esi": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redundancy_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "virtual_network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_point_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "trigger_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceOverlayMirrorDestinationRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayMirrorDestinations := vspk.OverlayMirrorDestinationsList{}
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
    parent := &vspk.L2Domain{ID: d.Get("parent_l2_domain").(string)}
    filteredOverlayMirrorDestinations, err = parent.OverlayMirrorDestinations(fetchFilter)
    if err != nil {
        return err
    }

    OverlayMirrorDestination := &vspk.OverlayMirrorDestination{}

    if len(filteredOverlayMirrorDestinations) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayMirrorDestinations) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayMirrorDestination = filteredOverlayMirrorDestinations[0]

    d.Set("esi", OverlayMirrorDestination.ESI)
    d.Set("name", OverlayMirrorDestination.Name)
    d.Set("last_updated_by", OverlayMirrorDestination.LastUpdatedBy)
    d.Set("redundancy_enabled", OverlayMirrorDestination.RedundancyEnabled)
    d.Set("template_id", OverlayMirrorDestination.TemplateID)
    d.Set("description", OverlayMirrorDestination.Description)
    d.Set("virtual_network_id", OverlayMirrorDestination.VirtualNetworkID)
    d.Set("end_point_type", OverlayMirrorDestination.EndPointType)
    d.Set("entity_scope", OverlayMirrorDestination.EntityScope)
    d.Set("trigger_type", OverlayMirrorDestination.TriggerType)
    d.Set("external_id", OverlayMirrorDestination.ExternalID)
    
    d.Set("id", OverlayMirrorDestination.Identifier())
    d.Set("parent_id", OverlayMirrorDestination.ParentID)
    d.Set("parent_type", OverlayMirrorDestination.ParentType)
    d.Set("owner", OverlayMirrorDestination.Owner)

    d.SetId(OverlayMirrorDestination.Identifier())
    
    return nil
}