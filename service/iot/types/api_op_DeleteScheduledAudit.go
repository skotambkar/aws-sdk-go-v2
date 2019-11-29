// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The name of the scheduled audit you want to delete.
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScheduledAuditInput"}

	if s.ScheduledAuditName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledAuditName"))
	}
	if s.ScheduledAuditName != nil && len(*s.ScheduledAuditName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledAuditName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteScheduledAuditOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}