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

type InlineResponse2003History struct {
	Id int32 `json:"id,omitempty"`
	Created *InlineResponse2003Created `json:"created,omitempty"`
	DocumentId int32 `json:"document_id,omitempty"`
	Document string `json:"document,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	// The status of the fax job
	Status string `json:"status,omitempty"`
}