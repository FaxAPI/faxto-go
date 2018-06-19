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

type InlineResponse2009 struct {
	Status string `json:"status,omitempty"`
	UserId int32 `json:"user_id,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Id int32 `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	OrigFilename string `json:"orig_filename,omitempty"`
	PreviewFile string `json:"preview_file,omitempty"`
	PreviewImage string `json:"preview_image,omitempty"`
	PreviewInStorage int32 `json:"preview_in_storage,omitempty"`
	FileExtension string `json:"file_extension,omitempty"`
	FilenameUploaded string `json:"filename_uploaded,omitempty"`
	Filesize string `json:"filesize,omitempty"`
	S3 int32 `json:"s3,omitempty"`
	ServerDocumentId int32 `json:"server_document_id,omitempty"`
	TeamUserId int32 `json:"team_user_id,omitempty"`
	TotalPages int32 `json:"total_pages,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}