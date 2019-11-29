// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRepositoryInput struct {
	_ struct{} `type:"structure"`

	// If a repository contains images, forces the deletion.
	Force *bool `locationName:"force" type:"boolean"`

	// The AWS account ID associated with the registry that contains the repository
	// to delete. If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository to delete.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
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
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// The repository that was deleted.
	Repository *Repository `locationName:"repository" type:"structure"`
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}