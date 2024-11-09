package noop

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNoopTmpFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceNoopTmpFileCreate,
		Read:   resourceNoopTmpFileRead,
		Update: resourceNoopTmpFileUpdate,
		Delete: resourceNoopTmpFileDelete,
		Importer: &schema.ResourceImporter{
			State: resourceNoopTmpFileImport,
		},

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNoopTmpFileCreate(d *schema.ResourceData, m interface{}) error {
	dir := m.(*Config).Dir

	filename := d.Get("filename").(string)
	content := d.Get("content").(string)

	filePath := fmt.Sprintf("%s/%s", dir, filename)

	if _, err := os.Stat(filePath); err == nil {
		return errors.New("the file already exists")
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to create the file: %w", err)
	}

	d.SetId(filePath)
	return resourceNoopTmpFileRead(d, m)
}

func resourceNoopTmpFileRead(d *schema.ResourceData, m interface{}) error {
	filePath := d.Id()

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to describe the file: %w", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}
	d.Set("content", string(content)) // nolint:errcheck

	return nil
}

func resourceNoopTmpFileUpdate(d *schema.ResourceData, m interface{}) error {
	filePath := d.Id()
	content := d.Get("content").(string)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("faile to update the file: %w", err)
	}

	return resourceNoopTmpFileRead(d, m)
}

func resourceNoopTmpFileDelete(d *schema.ResourceData, _ interface{}) error {
	filePath := d.Id()

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete the file: %w", err)
	}

	d.SetId("")
	return nil
}

func resourceNoopTmpFileImport(d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	filePath := d.Id()
	d.SetId(filePath)

	return []*schema.ResourceData{d}, nil
}
