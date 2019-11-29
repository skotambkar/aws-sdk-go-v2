// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHITTypeInput struct {
	_ struct{} `type:"structure"`

	// The amount of time, in seconds, that a Worker has to complete the HIT after
	// accepting it. If a Worker does not complete the assignment within the specified
	// duration, the assignment is considered abandoned. If the HIT is still active
	// (that is, its lifetime has not elapsed), the assignment becomes available
	// for other users to find and accept.
	//
	// AssignmentDurationInSeconds is a required field
	AssignmentDurationInSeconds *int64 `type:"long" required:"true"`

	// The number of seconds after an assignment for the HIT has been submitted,
	// after which the assignment is considered Approved automatically unless the
	// Requester explicitly rejects it.
	AutoApprovalDelayInSeconds *int64 `type:"long"`

	// A general description of the HIT. A description includes detailed information
	// about the kind of task the HIT contains. On the Amazon Mechanical Turk web
	// site, the HIT description appears in the expanded view of search results,
	// and in the HIT and assignment screens. A good description gives the user
	// enough information to evaluate the HIT before accepting it.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// One or more words or phrases that describe the HIT, separated by commas.
	// These words are used in searches to find HITs.
	Keywords *string `type:"string"`

	// Conditions that a Worker's Qualifications must meet in order to accept the
	// HIT. A HIT can have between zero and ten Qualification requirements. All
	// requirements must be met in order for a Worker to accept the HIT. Additionally,
	// other actions can be restricted using the ActionsGuarded field on each QualificationRequirement
	// structure.
	QualificationRequirements []QualificationRequirement `type:"list"`

	// The amount of money the Requester will pay a Worker for successfully completing
	// the HIT.
	//
	// Reward is a required field
	Reward *string `type:"string" required:"true"`

	// The title of the HIT. A title should be short and descriptive about the kind
	// of task the HIT contains. On the Amazon Mechanical Turk web site, the HIT
	// title appears in search results, and everywhere the HIT is mentioned.
	//
	// Title is a required field
	Title *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHITTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHITTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHITTypeInput"}

	if s.AssignmentDurationInSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentDurationInSeconds"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Reward == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reward"))
	}

	if s.Title == nil {
		invalidParams.Add(aws.NewErrParamRequired("Title"))
	}
	if s.QualificationRequirements != nil {
		for i, v := range s.QualificationRequirements {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "QualificationRequirements", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHITTypeOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the newly registered HIT type.
	HITTypeId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateHITTypeOutput) String() string {
	return awsutil.Prettify(s)
}
