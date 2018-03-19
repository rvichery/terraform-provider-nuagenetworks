package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceCOSRemarkingPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceCOSRemarkingPolicyRead,
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
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "forwarding_class": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_cos_remarking_policy_table": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceCOSRemarkingPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredCOSRemarkingPolicies := vspk.COSRemarkingPoliciesList{}
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
    parent := &vspk.COSRemarkingPolicyTable{ID: d.Get("parent_cos_remarking_policy_table").(string)}
    filteredCOSRemarkingPolicies, err = parent.COSRemarkingPolicies(fetchFilter)
    if err != nil {
        return err
    }

    COSRemarkingPolicy := &vspk.COSRemarkingPolicy{}

    if len(filteredCOSRemarkingPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredCOSRemarkingPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    COSRemarkingPolicy = filteredCOSRemarkingPolicies[0]

    d.Set("dscp", COSRemarkingPolicy.DSCP)
    d.Set("last_updated_by", COSRemarkingPolicy.LastUpdatedBy)
    d.Set("entity_scope", COSRemarkingPolicy.EntityScope)
    d.Set("forwarding_class", COSRemarkingPolicy.ForwardingClass)
    d.Set("external_id", COSRemarkingPolicy.ExternalID)
    
    d.Set("id", COSRemarkingPolicy.Identifier())
    d.Set("parent_id", COSRemarkingPolicy.ParentID)
    d.Set("parent_type", COSRemarkingPolicy.ParentType)
    d.Set("owner", COSRemarkingPolicy.Owner)

    d.SetId(COSRemarkingPolicy.Identifier())
    
    return nil
}