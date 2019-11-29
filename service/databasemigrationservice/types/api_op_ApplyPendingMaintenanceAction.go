// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ApplyPendingMaintenanceActionInput struct {
	_ struct{} `type:"structure"`

	// The pending maintenance action to apply to this resource.
	//
	// ApplyAction is a required field
	ApplyAction *string `type:"string" required:"true"`

	// A value that specifies the type of opt-in request, or undoes an opt-in request.
	// You can't undo an opt-in request of type immediate.
	//
	// Valid values:
	//
	//    * immediate - Apply the maintenance action immediately.
	//
	//    * next-maintenance - Apply the maintenance action during the next maintenance
	//    window for the resource.
	//
	//    * undo-opt-in - Cancel any existing next-maintenance opt-in requests.
	//
	// OptInType is a required field
	OptInType *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the AWS DMS resource that the pending maintenance
	// action applies to.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplyPendingMaintenanceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplyPendingMaintenanceActionInput"}

	if s.ApplyAction == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyAction"))
	}

	if s.OptInType == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptInType"))
	}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ApplyPendingMaintenanceActionOutput struct {
	_ struct{} `type:"structure"`

	// The AWS DMS resource that the pending maintenance action will be applied
	// to.
	ResourcePendingMaintenanceActions *ResourcePendingMaintenanceActions `type:"structure"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionOutput) String() string {
	return awsutil.Prettify(s)
}