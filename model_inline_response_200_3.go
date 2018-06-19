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

type InlineResponse2003 struct {
	// The status of the API request
	Status string `json:"status,omitempty"`
	// The fax jobs requested
	History []InlineResponse2003History `json:"history,omitempty"`
}
