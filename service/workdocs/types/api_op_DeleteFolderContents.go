// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFolderContentsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the folder.
	//
	// FolderId is a required field
	FolderId *string `location:"uri" locationName:"FolderId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFolderContentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFolderContentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFolderContentsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.FolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderId"))
	}
	if s.FolderId != nil && len(*s.FolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFolderContentsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFolderContentsOutput) String() string {
	return awsutil.Prettify(s)
}