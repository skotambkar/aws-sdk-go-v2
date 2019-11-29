// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/enums"
)

type DescribePullRequestEventsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user whose actions resulted in the
	// event. Examples include updating the pull request with more commits or changing
	// the status of a pull request.
	ActorArn *string `locationName:"actorArn" type:"string"`

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 events, which is also the maximum number of events that
	// can be returned in a result.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Optional. The pull request event type about which you want to return information.
	PullRequestEventType enums.PullRequestEventType `locationName:"pullRequestEventType" type:"string" enum:"true"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePullRequestEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePullRequestEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePullRequestEventsInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePullRequestEventsOutput struct {
	_ struct{} `type:"structure"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the pull request events.
	//
	// PullRequestEvents is a required field
	PullRequestEvents []PullRequestEvent `locationName:"pullRequestEvents" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribePullRequestEventsOutput) String() string {
	return awsutil.Prettify(s)
}
