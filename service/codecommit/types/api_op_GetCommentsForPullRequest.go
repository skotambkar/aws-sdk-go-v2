// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCommentsForPullRequestInput struct {
	_ struct{} `type:"structure"`

	// The full commit ID of the commit in the source branch that was the tip of
	// the branch at the time the comment was made.
	AfterCommitId *string `locationName:"afterCommitId" type:"string"`

	// The full commit ID of the commit in the destination branch that was the tip
	// of the branch at the time the pull request was created.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The name of the repository that contains the pull request.
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string"`
}

// String returns the string representation
func (s GetCommentsForPullRequestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommentsForPullRequestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommentsForPullRequestInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCommentsForPullRequestOutput struct {
	_ struct{} `type:"structure"`

	// An array of comment objects on the pull request.
	CommentsForPullRequestData []CommentsForPullRequest `locationName:"commentsForPullRequestData" type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetCommentsForPullRequestOutput) String() string {
	return awsutil.Prettify(s)
}
