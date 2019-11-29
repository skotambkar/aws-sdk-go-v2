// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCommentsForComparedCommitInput struct {
	_ struct{} `type:"structure"`

	// To establish the directionality of the comparison, the full commit ID of
	// the after commit.
	//
	// AfterCommitId is a required field
	AfterCommitId *string `locationName:"afterCommitId" type:"string" required:"true"`

	// To establish the directionality of the comparison, the full commit ID of
	// the before commit.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments, but you can configure up to 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository where you want to compare commits.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCommentsForComparedCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommentsForComparedCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommentsForComparedCommitInput"}

	if s.AfterCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AfterCommitId"))
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

type GetCommentsForComparedCommitOutput struct {
	_ struct{} `type:"structure"`

	// A list of comment objects on the compared commit.
	CommentsForComparedCommitData []CommentsForComparedCommit `locationName:"commentsForComparedCommitData" type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetCommentsForComparedCommitOutput) String() string {
	return awsutil.Prettify(s)
}