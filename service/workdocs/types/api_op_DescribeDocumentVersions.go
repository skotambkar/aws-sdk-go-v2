// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDocumentVersionsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// Specify "SOURCE" to include initialized versions and a URL for the source
	// document.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// A comma-separated list of values. Specify "INITIALIZED" to include incomplete
	// versions.
	Include *string `location:"querystring" locationName:"include" min:"1" type:"string"`

	// The maximum number of versions to return with this call.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDocumentVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDocumentVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDocumentVersionsInput"}
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
	if s.Include != nil && len(*s.Include) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Include", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDocumentVersionsOutput struct {
	_ struct{} `type:"structure"`

	// The document versions.
	DocumentVersions []DocumentVersionMetadata `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDocumentVersionsOutput) String() string {
	return awsutil.Prettify(s)
}
