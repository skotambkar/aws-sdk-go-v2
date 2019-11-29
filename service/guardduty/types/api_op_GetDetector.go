// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/enums"
)

type GetDetectorInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector that you want to get.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDetectorInput"}

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

type GetDetectorOutput struct {
	_ struct{} `type:"structure"`

	// Detector creation timestamp.
	CreatedAt *string `locationName:"createdAt" type:"string"`

	// Finding publishing frequency.
	FindingPublishingFrequency enums.FindingPublishingFrequency `locationName:"findingPublishingFrequency" type:"string" enum:"true"`

	// The GuardDuty service role.
	//
	// ServiceRole is a required field
	ServiceRole *string `locationName:"serviceRole" type:"string" required:"true"`

	// The detector status.
	//
	// Status is a required field
	Status enums.DetectorStatus `locationName:"status" min:"1" type:"string" required:"true" enum:"true"`

	// The tags of the detector resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`

	// Detector last update timestamp.
	UpdatedAt *string `locationName:"updatedAt" type:"string"`
}

// String returns the string representation
func (s GetDetectorOutput) String() string {
	return awsutil.Prettify(s)
}
