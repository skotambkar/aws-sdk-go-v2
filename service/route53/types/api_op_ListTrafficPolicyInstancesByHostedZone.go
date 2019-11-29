// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/route53/enums"
)

// A request for the traffic policy instances that you created in a specified
// hosted zone.
type ListTrafficPolicyInstancesByHostedZoneInput struct {
	_ struct{} `type:"structure"`

	// The ID of the hosted zone that you want to list traffic policy instances
	// for.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"querystring" locationName:"id" type:"string" required:"true"`

	// The maximum number of traffic policy instances to be included in the response
	// body for this request. If you have more than MaxItems traffic policy instances,
	// the value of the IsTruncated element in the response is true, and the values
	// of HostedZoneIdMarker, TrafficPolicyInstanceNameMarker, and TrafficPolicyInstanceTypeMarker
	// represent the first traffic policy instance that Amazon Route 53 will return
	// if you submit another request.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`

	// If the value of IsTruncated in the previous response is true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancename,
	// specify the value of TrafficPolicyInstanceNameMarker from the previous response,
	// which is the name of the first traffic policy instance in the next group
	// of traffic policy instances.
	//
	// If the value of IsTruncated in the previous response was false, there are
	// no more traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string `location:"querystring" locationName:"trafficpolicyinstancename" type:"string"`

	// If the value of IsTruncated in the previous response is true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancetype,
	// specify the value of TrafficPolicyInstanceTypeMarker from the previous response,
	// which is the type of the first traffic policy instance in the next group
	// of traffic policy instances.
	//
	// If the value of IsTruncated in the previous response was false, there are
	// no more traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker enums.RRType `location:"querystring" locationName:"trafficpolicyinstancetype" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTrafficPolicyInstancesByHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTrafficPolicyInstancesByHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTrafficPolicyInstancesByHostedZoneInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesByHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more traffic policy instances to
	// be listed. If the response was truncated, you can get the next group of traffic
	// policy instances by submitting another ListTrafficPolicyInstancesByHostedZone
	// request and specifying the values of HostedZoneIdMarker, TrafficPolicyInstanceNameMarker,
	// and TrafficPolicyInstanceTypeMarker in the corresponding request parameters.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// The value that you specified for the MaxItems parameter in the ListTrafficPolicyInstancesByHostedZone
	// request that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If IsTruncated is true, TrafficPolicyInstanceNameMarker is the name of the
	// first traffic policy instance in the next group of traffic policy instances.
	TrafficPolicyInstanceNameMarker *string `type:"string"`

	// If IsTruncated is true, TrafficPolicyInstanceTypeMarker is the DNS type of
	// the resource record sets that are associated with the first traffic policy
	// instance in the next group of traffic policy instances.
	TrafficPolicyInstanceTypeMarker enums.RRType `type:"string" enum:"true"`

	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	//
	// TrafficPolicyInstances is a required field
	TrafficPolicyInstances []TrafficPolicyInstance `locationNameList:"TrafficPolicyInstance" type:"list" required:"true"`
}

// String returns the string representation
func (s ListTrafficPolicyInstancesByHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}