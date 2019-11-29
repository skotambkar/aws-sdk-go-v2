// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAuditStreamConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Amazon Kinesis data stream that receives the audit events.
	AuditStreamArn *string `type:"string"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAuditStreamConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAuditStreamConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAuditStreamConfigurationInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAuditStreamConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAuditStreamConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
