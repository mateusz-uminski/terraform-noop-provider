package noop

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func diagErrorf(format string, a ...interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf(format, a...))
}
