// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartOnDemandAuditTaskInput struct {
	_ struct{} `type:"structure"`

	// Which checks are performed during the audit. The checks you specify must
	// be enabled for your account or an exception occurs. Use DescribeAccountAuditConfiguration
	// to see the list of all checks, including those that are enabled or UpdateAccountAuditConfiguration
	// to select which checks are enabled.
	//
	// TargetCheckNames is a required field
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list" required:"true"`
}

// String returns the string representation
func (s StartOnDemandAuditTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartOnDemandAuditTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartOnDemandAuditTaskInput"}

	if s.TargetCheckNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetCheckNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartOnDemandAuditTaskOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the on-demand audit you started.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartOnDemandAuditTaskOutput) String() string {
	return awsutil.Prettify(s)
}
