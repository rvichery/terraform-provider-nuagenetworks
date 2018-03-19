package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFiltersSchema() *schema.Schema {
    return &schema.Schema{
        Type:     schema.TypeSet,
        Optional: true,
        ForceNew: true,
        Elem: &schema.Resource{
            Schema: map[string]*schema.Schema{
                "key": {
                    Type:     schema.TypeString,
                    Required: true,
                },
                "operator": {
                    Type:     schema.TypeString,
                    Optional: true,
                    Default:  "==",
                },
                "value": {
                    Type:     schema.TypeString,
                    Required: true,
                },
            },
        },
    }
}