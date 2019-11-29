// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type GetCapacityReservationUsageInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Capacity Reservation.
	//
	// CapacityReservationId is a required field
	CapacityReservationId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the returned
	// nextToken value.
	//
	// Valid range: Minimum value of 1. Maximum value of 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCapacityReservationUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCapacityReservationUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCapacityReservationUsageInput"}

	if s.CapacityReservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CapacityReservationId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCapacityReservationUsageOutput struct {
	_ struct{} `type:"structure"`

	// The remaining capacity. Indicates the number of instances that can be launched
	// in the Capacity Reservation.
	AvailableInstanceCount *int64 `locationName:"availableInstanceCount" type:"integer"`

	// The ID of the Capacity Reservation.
	CapacityReservationId *string `locationName:"capacityReservationId" type:"string"`

	// The type of instance for which the Capacity Reservation reserves capacity.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// Information about the Capacity Reservation usage.
	InstanceUsages []InstanceUsage `locationName:"instanceUsageSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The current state of the Capacity Reservation. A Capacity Reservation can
	// be in one of the following states:
	//
	//    * active - The Capacity Reservation is active and the capacity is available
	//    for your use.
	//
	//    * expired - The Capacity Reservation expired automatically at the date
	//    and time specified in your request. The reserved capacity is no longer
	//    available for your use.
	//
	//    * cancelled - The Capacity Reservation was manually cancelled. The reserved
	//    capacity is no longer available for your use.
	//
	//    * pending - The Capacity Reservation request was successful but the capacity
	//    provisioning is still pending.
	//
	//    * failed - The Capacity Reservation request has failed. A request might
	//    fail due to invalid request parameters, capacity constraints, or instance
	//    limit constraints. Failed requests are retained for 60 minutes.
	State enums.CapacityReservationState `locationName:"state" type:"string" enum:"true"`

	// The number of instances for which the Capacity Reservation reserves capacity.
	TotalInstanceCount *int64 `locationName:"totalInstanceCount" type:"integer"`
}

// String returns the string representation
func (s GetCapacityReservationUsageOutput) String() string {
	return awsutil.Prettify(s)
}