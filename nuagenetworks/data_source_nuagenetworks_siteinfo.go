package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSiteInfo() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSiteInfoRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "site_identifier": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "xmpp_domain": &schema.Schema{
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
        },
    }
}


func dataSourceSiteInfoRead(d *schema.ResourceData, m interface{}) error {
    filteredSiteInfos := vspk.SiteInfosList{}
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
    parent := m.(*vspk.Me)
    filteredSiteInfos, err = parent.SiteInfos(fetchFilter)
    if err != nil {
        return err
    }

    SiteInfo := &vspk.SiteInfo{}

    if len(filteredSiteInfos) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSiteInfos) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SiteInfo = filteredSiteInfos[0]

    d.Set("name", SiteInfo.Name)
    d.Set("last_updated_by", SiteInfo.LastUpdatedBy)
    d.Set("address", SiteInfo.Address)
    d.Set("description", SiteInfo.Description)
    d.Set("site_identifier", SiteInfo.SiteIdentifier)
    d.Set("xmpp_domain", SiteInfo.XmppDomain)
    d.Set("entity_scope", SiteInfo.EntityScope)
    d.Set("external_id", SiteInfo.ExternalID)
    
    d.Set("id", SiteInfo.Identifier())
    d.Set("parent_id", SiteInfo.ParentID)
    d.Set("parent_type", SiteInfo.ParentType)
    d.Set("owner", SiteInfo.Owner)

    d.SetId(SiteInfo.Identifier())
    
    return nil
}