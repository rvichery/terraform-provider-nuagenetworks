package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourcePolicyEntry() *schema.Resource {
    return &schema.Resource{
        Create: resourcePolicyEntryCreate,
        Read:   resourcePolicyEntryRead,
        Update: resourcePolicyEntryUpdate,
        Delete: resourcePolicyEntryDelete,

        Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "match_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "actions": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_policy_statement": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourcePolicyEntryCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PolicyEntry object
    o := &vspk.PolicyEntry{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("match_criteria"); ok {
        o.MatchCriteria = attr.(interface{})
    }
    if attr, ok := d.GetOk("actions"); ok {
        o.Actions = attr.(interface{})
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    parent := &vspk.PolicyStatement{ID: d.Get("parent_policy_statement").(string)}
    err := parent.CreatePolicyEntry(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePolicyEntryRead(d, m)
}

func resourcePolicyEntryRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyEntry{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
		d.SetId("")
		return nil
    }

    d.Set("name", o.Name)
    d.Set("match_criteria", o.MatchCriteria)
    d.Set("actions", o.Actions)
    d.Set("description", o.Description)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePolicyEntryUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyEntry{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("match_criteria"); ok {
        o.MatchCriteria = attr.(interface{})
    }
    if attr, ok := d.GetOk("actions"); ok {
        o.Actions = attr.(interface{})
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePolicyEntryDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyEntry{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}