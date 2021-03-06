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

type InlineResponse2008 struct {
	// The status of the API request
	Status string `json:"status,omitempty"`
	// File data
	Files []InlineResponse2008Files `json:"files,omitempty"`
}
