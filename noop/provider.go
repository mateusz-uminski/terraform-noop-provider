package noop

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/tmp",
				Description: "The directory where files will be created.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"noop_tmp_file": resourceNoopTmpFile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"noop_tmp_file": dataSourceNoopTmpFile(),
		},
	}

	p.ConfigureContextFunc = providerConfigure(p)

	return p
}

func providerConfigure(_ *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		dir := d.Get("dir").(string)

		if dir == "" {
			dir = "/tmp"
		}

		config := Config{
			Dir: dir,
		}

		meta, err := config.Meta()
		if err != nil {
			return nil, diagErrorf("something went wrong: %w", err)
		}

		return meta, nil
	}
}
