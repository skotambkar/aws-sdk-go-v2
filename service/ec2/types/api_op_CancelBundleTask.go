// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CancelBundleTask.
type CancelBundleTaskInput struct {
	_ struct{} `type:"structure"`

	// The ID of the bundle task.
	//
	// BundleId is a required field
	BundleId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s CancelBundleTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelBundleTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelBundleTaskInput"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CancelBundleTask.
type CancelBundleTaskOutput struct {
	_ struct{} `type:"structure"`

	// Information about the bundle task.
	BundleTask *BundleTask `locationName:"bundleInstanceTask" type:"structure"`
}

// String returns the string representation
func (s CancelBundleTaskOutput) String() string {
	return awsutil.Prettify(s)
}
