// Contains functions that don't really belong anywhere else.
package teamcity

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func optionalSuffixReturnSuppress(k, old, new string, d *schema.ResourceData) bool {
	return strings.TrimSuffix(old, "\n") == strings.TrimSuffix(new, "\n")
}
