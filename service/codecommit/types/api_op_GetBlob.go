// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get blob operation.
type GetBlobInput struct {
	_ struct{} `type:"structure"`

	// The ID of the blob, which is its SHA-1 pointer.
	//
	// BlobId is a required field
	BlobId *string `locationName:"blobId" type:"string" required:"true"`

	// The name of the repository that contains the blob.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBlobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBlobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBlobInput"}

	if s.BlobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BlobId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a get blob operation.
type GetBlobOutput struct {
	_ struct{} `type:"structure"`

	// The content of the blob, usually a file.
	//
	// Content is automatically base64 encoded/decoded by the SDK.
	//
	// Content is a required field
	Content []byte `locationName:"content" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetBlobOutput) String() string {
	return awsutil.Prettify(s)
}