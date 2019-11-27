// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecr/enums"
)

type CreateRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The tag mutability setting for the repository. If this parameter is omitted,
	// the default setting of MUTABLE will be used which will allow image tags to
	// be overwritten. If IMMUTABLE is specified, all image tags within the repository
	// will be immutable which will prevent them from being overwritten.
	ImageTagMutability enums.ImageTagMutability `locationName:"imageTagMutability" type:"string" enum:"true"`

	// The name to use for the repository. The repository name may be specified
	// on its own (such as nginx-web-app) or it can be prepended with a namespace
	// to group the repository into a category (such as project-a/nginx-web-app).
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`

	// The metadata that you apply to the repository to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of
	// which you define. Tag keys can have a maximum character length of 128 characters,
	// and tag values can have a maximum length of 256 characters.
	Tags []Tag `locationName:"tags" type:"list"`
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
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// The repository that was created.
	Repository *Repository `locationName:"repository" type:"structure"`
}

// String returns the string representation
func (s CreateRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}
