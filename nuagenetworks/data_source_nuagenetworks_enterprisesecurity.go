package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEnterpriseSecurity() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEnterpriseSecurityRead,
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
            "gateway_security_revision": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "revision": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
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
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceEnterpriseSecurityRead(d *schema.ResourceData, m interface{}) error {
    filteredEnterpriseSecurities := vspk.EnterpriseSecuritiesList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredEnterpriseSecurities, err = parent.EnterpriseSecurities(fetchFilter)
    if err != nil {
        return err
    }

    EnterpriseSecurity := &vspk.EnterpriseSecurity{}

    if len(filteredEnterpriseSecurities) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEnterpriseSecurities) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EnterpriseSecurity = filteredEnterpriseSecurities[0]

    d.Set("last_updated_by", EnterpriseSecurity.LastUpdatedBy)
    d.Set("gateway_security_revision", EnterpriseSecurity.GatewaySecurityRevision)
    d.Set("revision", EnterpriseSecurity.Revision)
    d.Set("enterprise_id", EnterpriseSecurity.EnterpriseID)
    d.Set("entity_scope", EnterpriseSecurity.EntityScope)
    d.Set("external_id", EnterpriseSecurity.ExternalID)
    
    d.Set("id", EnterpriseSecurity.Identifier())
    d.Set("parent_id", EnterpriseSecurity.ParentID)
    d.Set("parent_type", EnterpriseSecurity.ParentType)
    d.Set("owner", EnterpriseSecurity.Owner)

    d.SetId(EnterpriseSecurity.Identifier())
    
    return nil
}