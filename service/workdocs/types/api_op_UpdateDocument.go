// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/enums"
)

type UpdateDocumentInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// The name of the document.
	Name *string `min:"1" type:"string"`

	// The ID of the parent folder.
	ParentFolderId *string `min:"1" type:"string"`

	// The resource state of the document. Only ACTIVE and RECYCLED are supported.
	ResourceState enums.ResourceStateType `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDocumentInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ParentFolderId != nil && len(*s.ParentFolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParentFolderId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDocumentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDocumentOutput) String() string {
	return awsutil.Prettify(s)
}