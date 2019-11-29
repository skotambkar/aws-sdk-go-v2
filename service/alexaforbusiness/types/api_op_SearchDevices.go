// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchDevicesInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to list a specified set of devices. Supported filter keys
	// are DeviceName, DeviceStatus, DeviceStatusDetailCode, RoomName, DeviceType,
	// DeviceSerialNumber, UnassociatedOnly, ConnectionStatus (ONLINE and OFFLINE),
	// NetworkProfileName, NetworkProfileArn, Feature, and FailureCode.
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use in listing the specified set of devices. Supported
	// sort keys are DeviceName, DeviceStatus, RoomName, DeviceType, DeviceSerialNumber,
	// ConnectionStatus, NetworkProfileName, NetworkProfileArn, Feature, and FailureCode.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchDevicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SortCriteria != nil {
		for i, v := range s.SortCriteria {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortCriteria", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchDevicesOutput struct {
	_ struct{} `type:"structure"`

	// The devices that meet the specified set of filter criteria, in sort order.
	Devices []DeviceData `type:"list"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The total number of devices returned.
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s SearchDevicesOutput) String() string {
	return awsutil.Prettify(s)
}