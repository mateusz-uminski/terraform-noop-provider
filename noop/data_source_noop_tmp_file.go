package noop

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNoopTmpFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNoopTmpFileRead,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNoopTmpFileRead(d *schema.ResourceData, m interface{}) error {
	dir := m.(*Config).Dir

	filename := d.Get("filename").(string)

	filePath := fmt.Sprintf("%s/%s", dir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", filePath)
		}
		return fmt.Errorf("failed to read the file: %w", err)
	}

	d.SetId(filePath)
	d.Set("content", string(content)) // nolint:errcheck

	return nil
}
