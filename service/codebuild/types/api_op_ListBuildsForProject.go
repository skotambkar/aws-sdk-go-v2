// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/enums"
)

type ListBuildsForProjectInput struct {
	_ struct{} `type:"structure"`

	// During a previous call, if there are more than 100 items in the list, only
	// the first 100 items are returned, along with a unique string called a next
	// token. To get the next batch of items in the list, call this operation again,
	// adding the next token to the call. To get all of the items in the list, keep
	// calling this operation with each subsequent next token that is returned,
	// until no more next tokens are returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the AWS CodeBuild project.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The order to list build IDs. Valid values include:
	//
	//    * ASCENDING: List the build IDs in ascending order by build ID.
	//
	//    * DESCENDING: List the build IDs in descending order by build ID.
	SortOrder enums.SortOrderType `locationName:"sortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListBuildsForProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBuildsForProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBuildsForProjectInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListBuildsForProjectOutput struct {
	_ struct{} `type:"structure"`

	// A list of build IDs for the specified build project, with each build ID representing
	// a single build.
	Ids []string `locationName:"ids" min:"1" type:"list"`

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a next token. To get the next
	// batch of items in the list, call this operation again, adding the next token
	// to the call.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBuildsForProjectOutput) String() string {
	return awsutil.Prettify(s)
}
