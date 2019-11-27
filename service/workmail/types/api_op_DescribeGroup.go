// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/workmail/enums"
)

type DescribeGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the group to be described.
	//
	// GroupId is a required field
	GroupId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the group exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}
	if s.GroupId != nil && len(*s.GroupId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeGroupOutput struct {
	_ struct{} `type:"structure"`

	// The date and time when a user was deregistered from WorkMail, in UNIX epoch
	// time format.
	DisabledDate *time.Time `type:"timestamp"`

	// The email of the described group.
	Email *string `min:"1" type:"string"`

	// The date and time when a user was registered to WorkMail, in UNIX epoch time
	// format.
	EnabledDate *time.Time `type:"timestamp"`

	// The identifier of the described group.
	GroupId *string `min:"12" type:"string"`

	// The name of the described group.
	Name *string `min:"1" type:"string"`

	// The state of the user: enabled (registered to Amazon WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State enums.EntityState `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeGroupOutput) String() string {
	return awsutil.Prettify(s)
}
