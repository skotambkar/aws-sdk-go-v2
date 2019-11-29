// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a list branches operation.
type ListBranchesInput struct {
	_ struct{} `type:"structure"`

	// An enumeration token that allows the operation to batch the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository that contains the branches.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListBranchesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBranchesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBranchesInput"}

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

// Represents the output of a list branches operation.
type ListBranchesOutput struct {
	_ struct{} `type:"structure"`

	// The list of branch names.
	Branches []string `locationName:"branches" type:"list"`

	// An enumeration token that returns the batch of the results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBranchesOutput) String() string {
	return awsutil.Prettify(s)
}