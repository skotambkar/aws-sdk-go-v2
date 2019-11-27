// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a delete repository operation.
type DeleteRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the repository to delete.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRepositoryInput"}

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

// Represents the output of a delete repository operation.
type DeleteRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the repository that was deleted.
	RepositoryId *string `locationName:"repositoryId" type:"string"`
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}
