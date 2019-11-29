// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePullRequestInput struct {
	_ struct{} `type:"structure"`

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// is received with the same parameters and a token is included, the request
	// returns information about the initial request that used that token.
	//
	// The AWS SDKs prepopulate client request tokens. If you are using an AWS SDK,
	// an idempotency token is created for you.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// A description of the pull request.
	Description *string `locationName:"description" type:"string"`

	// The targets for the pull request, including the source of the code to be
	// reviewed (the source branch) and the destination where the creator of the
	// pull request intends the code to be merged after the pull request is closed
	// (the destination branch).
	//
	// Targets is a required field
	Targets []Target `locationName:"targets" type:"list" required:"true"`

	// The title of the pull request. This title is used to identify the pull request
	// to other users in the repository.
	//
	// Title is a required field
	Title *string `locationName:"title" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePullRequestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePullRequestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePullRequestInput"}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}

	if s.Title == nil {
		invalidParams.Add(aws.NewErrParamRequired("Title"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePullRequestOutput struct {
	_ struct{} `type:"structure"`

	// Information about the newly created pull request.
	//
	// PullRequest is a required field
	PullRequest *PullRequest `locationName:"pullRequest" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreatePullRequestOutput) String() string {
	return awsutil.Prettify(s)
}
