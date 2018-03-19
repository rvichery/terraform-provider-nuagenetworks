package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGInfo() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGInfoRead,
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
            "mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_release_date": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sku": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tpm_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tpm_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cpu_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "patches": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "libraries": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "product_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ns_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNSGInfoRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGInfos := vspk.NSGInfosList{}
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
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    filteredNSGInfos, err = parent.NSGInfos(fetchFilter)
    if err != nil {
        return err
    }

    NSGInfo := &vspk.NSGInfo{}

    if len(filteredNSGInfos) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGInfos) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGInfo = filteredNSGInfos[0]

    d.Set("mac_address", NSGInfo.MACAddress)
    d.Set("bios_release_date", NSGInfo.BIOSReleaseDate)
    d.Set("bios_version", NSGInfo.BIOSVersion)
    d.Set("sku", NSGInfo.SKU)
    d.Set("tpm_status", NSGInfo.TPMStatus)
    d.Set("tpm_version", NSGInfo.TPMVersion)
    d.Set("cpu_type", NSGInfo.CPUType)
    d.Set("nsg_version", NSGInfo.NSGVersion)
    d.Set("uuid", NSGInfo.UUID)
    d.Set("family", NSGInfo.Family)
    d.Set("patches", NSGInfo.Patches)
    d.Set("serial_number", NSGInfo.SerialNumber)
    d.Set("libraries", NSGInfo.Libraries)
    d.Set("entity_scope", NSGInfo.EntityScope)
    d.Set("product_name", NSGInfo.ProductName)
    d.Set("associated_ns_gateway_id", NSGInfo.AssociatedNSGatewayID)
    d.Set("external_id", NSGInfo.ExternalID)
    
    d.Set("id", NSGInfo.Identifier())
    d.Set("parent_id", NSGInfo.ParentID)
    d.Set("parent_type", NSGInfo.ParentType)
    d.Set("owner", NSGInfo.Owner)

    d.SetId(NSGInfo.Identifier())
    
    return nil
}