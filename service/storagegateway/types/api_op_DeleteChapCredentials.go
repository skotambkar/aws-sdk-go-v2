// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * DeleteChapCredentialsInput$InitiatorName
//
//    * DeleteChapCredentialsInput$TargetARN
type DeleteChapCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The iSCSI initiator that connects to the target.
	//
	// InitiatorName is a required field
	InitiatorName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the DescribeStorediSCSIVolumes
	// operation to return to retrieve the TargetARN for specified VolumeARN.
	//
	// TargetARN is a required field
	TargetARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteChapCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteChapCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteChapCredentialsInput"}

	if s.InitiatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InitiatorName"))
	}
	if s.InitiatorName != nil && len(*s.InitiatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InitiatorName", 1))
	}

	if s.TargetARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetARN"))
	}
	if s.TargetARN != nil && len(*s.TargetARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
type DeleteChapCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// The iSCSI initiator that connects to the target.
	InitiatorName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the target.
	TargetARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteChapCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}