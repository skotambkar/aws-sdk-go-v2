// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a create repository operation.
type CreateRepositoryInput struct {
	_ struct{} `type:"structure"`

	// A comment or description about the new repository.
	//
	// The description field for a repository accepts all HTML characters and all
	// valid Unicode characters. Applications that do not HTML-encode the description
	// and display it in a webpage can expose users to potentially malicious code.
	// Make sure that you HTML-encode the description field in any application that
	// uses this API to display the repository description on a webpage.
	RepositoryDescription *string `locationName:"repositoryDescription" type:"string"`

	// The name of the new repository to be created.
	//
	// The repository name must be unique across the calling AWS account. Repository
	// names are limited to 100 alphanumeric, dash, and underscore characters, and
	// cannot include certain characters. For more information about the limits
	// on repository names, see Limits (https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html)
	// in the AWS CodeCommit User Guide. The suffix .git is prohibited.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// One or more tag key-value pairs to use when tagging this repository.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRepositoryInput"}

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

// Represents the output of a create repository operation.
type CreateRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the newly created repository.
	RepositoryMetadata *RepositoryMetadata `locationName:"repositoryMetadata" type:"structure"`
}

// String returns the string representation
func (s CreateRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}
