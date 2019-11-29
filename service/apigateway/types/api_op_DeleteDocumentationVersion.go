// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Deletes an existing documentation version of an API.
type DeleteDocumentationVersionInput struct {
	_ struct{} `type:"structure"`

	// [Required] The version identifier of a to-be-deleted documentation snapshot.
	//
	// DocumentationVersion is a required field
	DocumentationVersion *string `location:"uri" locationName:"doc_version" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDocumentationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDocumentationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDocumentationVersionInput"}

	if s.DocumentationVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentationVersion"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDocumentationVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDocumentationVersionOutput) String() string {
	return awsutil.Prettify(s)
}