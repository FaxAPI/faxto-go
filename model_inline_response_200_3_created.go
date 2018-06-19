/*
 * Fax.to REST API client for Go
 *
 * This is Fax.to REST API client for Go.
 *
 * API version: 2.0.0
 * Contact: inquiries@fax.to
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type InlineResponse2003Created struct {
	Date time.Time `json:"date,omitempty"`
	TimezoneType int32 `json:"timezone_type,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}
