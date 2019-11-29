// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the DescribeDirectories operation.
type DescribeDirectoriesInput struct {
	_ struct{} `type:"structure"`

	// A list of identifiers of the directories for which to obtain the information.
	// If this member is null, all directories that belong to the current account
	// are returned.
	//
	// An empty list results in an InvalidParameterException being thrown.
	DirectoryIds []string `type:"list"`

	// The maximum number of items to return. If this value is zero, the maximum
	// number of items is specified by the limitations of the operation.
	Limit *int64 `type:"integer"`

	// The DescribeDirectoriesResult.NextToken value from a previous call to DescribeDirectories.
	// Pass null if this is the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeDirectoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the results of the DescribeDirectories operation.
type DescribeDirectoriesOutput struct {
	_ struct{} `type:"structure"`

	// The list of DirectoryDescription objects that were retrieved.
	//
	// It is possible that this list contains less than the number of items specified
	// in the Limit member of the request. This occurs if there are less than the
	// requested number of items left to retrieve, or if the limitations of the
	// operation have been exceeded.
	DirectoryDescriptions []DirectoryDescription `type:"list"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeDirectories to retrieve the next
	// set of items.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeDirectoriesOutput) String() string {
	return awsutil.Prettify(s)
}