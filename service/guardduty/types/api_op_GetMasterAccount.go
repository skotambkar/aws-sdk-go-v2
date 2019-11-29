// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMasterAccountInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector of the GuardDuty member account.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMasterAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMasterAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMasterAccountInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMasterAccountOutput struct {
	_ struct{} `type:"structure"`

	// Master account details.
	//
	// Master is a required field
	Master *Master `locationName:"master" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetMasterAccountOutput) String() string {
	return awsutil.Prettify(s)
}