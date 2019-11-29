// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a delete branch operation.
type DeleteBranchInput struct {
	_ struct{} `type:"structure"`

	// The name of the branch to delete.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// The name of the repository that contains the branch to be deleted.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBranchInput"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
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

// Represents the output of a delete branch operation.
type DeleteBranchOutput struct {
	_ struct{} `type:"structure"`

	// Information about the branch deleted by the operation, including the branch
	// name and the commit ID that was the tip of the branch.
	DeletedBranch *BranchInfo `locationName:"deletedBranch" type:"structure"`
}

// String returns the string representation
func (s DeleteBranchOutput) String() string {
	return awsutil.Prettify(s)
}