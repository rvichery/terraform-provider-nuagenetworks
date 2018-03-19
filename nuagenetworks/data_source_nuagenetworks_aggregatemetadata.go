package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAggregateMetadata() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAggregateMetadataRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata_tag_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "network_notification_disabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "blob": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "global_metadata": &schema.Schema{
                Type:     schema.TypeBool,
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
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceAggregateMetadataRead(d *schema.ResourceData, m interface{}) error {
    filteredAggregateMetadatas := vspk.AggregateMetadatasList{}
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
    parent := &vspk.VPort{ID: d.Get("parent_vport").(string)}
    filteredAggregateMetadatas, err = parent.AggregateMetadatas(fetchFilter)
    if err != nil {
        return err
    }

    AggregateMetadata := &vspk.AggregateMetadata{}

    if len(filteredAggregateMetadatas) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAggregateMetadatas) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AggregateMetadata = filteredAggregateMetadatas[0]

    d.Set("name", AggregateMetadata.Name)
    d.Set("description", AggregateMetadata.Description)
    d.Set("metadata_tag_ids", AggregateMetadata.MetadataTagIDs)
    d.Set("network_notification_disabled", AggregateMetadata.NetworkNotificationDisabled)
    d.Set("blob", AggregateMetadata.Blob)
    d.Set("global_metadata", AggregateMetadata.GlobalMetadata)
    d.Set("entity_scope", AggregateMetadata.EntityScope)
    d.Set("external_id", AggregateMetadata.ExternalID)
    
    d.Set("id", AggregateMetadata.Identifier())
    d.Set("parent_id", AggregateMetadata.ParentID)
    d.Set("parent_type", AggregateMetadata.ParentType)
    d.Set("owner", AggregateMetadata.Owner)

    d.SetId(AggregateMetadata.Identifier())
    
    return nil
}