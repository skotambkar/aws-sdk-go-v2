// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteOTAUpdateInput struct {
	_ struct{} `type:"structure"`

	// Specifies if the stream associated with an OTA update should be deleted when
	// the OTA update is deleted.
	DeleteStream *bool `location:"querystring" locationName:"deleteStream" type:"boolean"`

	// Specifies if the AWS Job associated with the OTA update should be deleted
	// with the OTA update is deleted.
	ForceDeleteAWSJob *bool `location:"querystring" locationName:"forceDeleteAWSJob" type:"boolean"`

	// The OTA update ID to delete.
	//
	// OtaUpdateId is a required field
	OtaUpdateId *string `location:"uri" locationName:"otaUpdateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOTAUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOTAUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOTAUpdateInput"}

	if s.OtaUpdateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OtaUpdateId"))
	}
	if s.OtaUpdateId != nil && len(*s.OtaUpdateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OtaUpdateId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOTAUpdateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOTAUpdateOutput) String() string {
	return awsutil.Prettify(s)
}