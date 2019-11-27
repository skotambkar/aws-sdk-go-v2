// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableAlarmActionsInput struct {
	_ struct{} `type:"structure"`

	// The names of the alarms.
	//
	// AlarmNames is a required field
	AlarmNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s EnableAlarmActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableAlarmActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableAlarmActionsInput"}

	if s.AlarmNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("AlarmNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableAlarmActionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableAlarmActionsOutput) String() string {
	return awsutil.Prettify(s)
}
