// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMissionProfileInput struct {
	_ struct{} `type:"structure"`

	// MissionProfileId is a required field
	MissionProfileId *string `location:"uri" locationName:"missionProfileId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMissionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMissionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMissionProfileInput"}

	if s.MissionProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MissionProfileId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMissionProfileOutput struct {
	_ struct{} `type:"structure"`

	ContactPostPassDurationSeconds *int64 `locationName:"contactPostPassDurationSeconds" min:"1" type:"integer"`

	ContactPrePassDurationSeconds *int64 `locationName:"contactPrePassDurationSeconds" min:"1" type:"integer"`

	DataflowEdges [][]string `locationName:"dataflowEdges" type:"list"`

	MinimumViableContactDurationSeconds *int64 `locationName:"minimumViableContactDurationSeconds" min:"1" type:"integer"`

	MissionProfileArn *string `locationName:"missionProfileArn" type:"string"`

	MissionProfileId *string `locationName:"missionProfileId" type:"string"`

	Name *string `locationName:"name" type:"string"`

	Region *string `locationName:"region" type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`

	TrackingConfigArn *string `locationName:"trackingConfigArn" type:"string"`
}

// String returns the string representation
func (s GetMissionProfileOutput) String() string {
	return awsutil.Prettify(s)
}
