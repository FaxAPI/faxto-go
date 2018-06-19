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

type InlineResponse2001 struct {
	// The status of the API request
	Status string `json:"status,omitempty"`
	// The cost of fax for the specified fax number
	Cost float32 `json:"cost,omitempty"`
}
