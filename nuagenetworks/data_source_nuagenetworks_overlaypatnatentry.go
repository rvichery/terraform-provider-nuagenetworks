package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayPATNATEntry() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayPATNATEntryRead,
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
            "nat_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "private_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_link_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "public_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_overlay_address_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceOverlayPATNATEntryRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayPATNATEntries := vspk.OverlayPATNATEntriesList{}
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
    parent := &vspk.OverlayAddressPool{ID: d.Get("parent_overlay_address_pool").(string)}
    filteredOverlayPATNATEntries, err = parent.OverlayPATNATEntries(fetchFilter)
    if err != nil {
        return err
    }

    OverlayPATNATEntry := &vspk.OverlayPATNATEntry{}

    if len(filteredOverlayPATNATEntries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayPATNATEntries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayPATNATEntry = filteredOverlayPATNATEntries[0]

    d.Set("nat_enabled", OverlayPATNATEntry.NATEnabled)
    d.Set("private_ip", OverlayPATNATEntry.PrivateIP)
    d.Set("associated_domain_id", OverlayPATNATEntry.AssociatedDomainID)
    d.Set("associated_link_id", OverlayPATNATEntry.AssociatedLinkID)
    d.Set("public_ip", OverlayPATNATEntry.PublicIP)
    
    d.Set("id", OverlayPATNATEntry.Identifier())
    d.Set("parent_id", OverlayPATNATEntry.ParentID)
    d.Set("parent_type", OverlayPATNATEntry.ParentType)
    d.Set("owner", OverlayPATNATEntry.Owner)

    d.SetId(OverlayPATNATEntry.Identifier())
    
    return nil
}