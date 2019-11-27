// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDocumentVersionInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// A comma-separated list of values. Specify "SOURCE" to include a URL for the
	// source document.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// Set this to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool `location:"querystring" locationName:"includeCustomMetadata" type:"boolean"`

	// The version ID of the document.
	//
	// VersionId is a required field
	VersionId *string `location:"uri" locationName:"VersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDocumentVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentVersionInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDocumentVersionOutput struct {
	_ struct{} `type:"structure"`

	// The custom metadata on the document version.
	CustomMetadata map[string]string `min:"1" type:"map"`

	// The version metadata.
	Metadata *DocumentVersionMetadata `type:"structure"`
}

// String returns the string representation
func (s GetDocumentVersionOutput) String() string {
	return awsutil.Prettify(s)
}
