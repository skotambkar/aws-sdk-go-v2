// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The ViewBilling request includes the following elements.
type ViewBillingInput struct {
	_ struct{} `type:"structure"`

	// The end date and time for the time period for which you want a list of billing
	// records. Specify the date and time in Coordinated Universal time (UTC).
	End *time.Time `type:"timestamp"`

	// For an initial request for a list of billing records, omit this element.
	// If the number of billing records that are associated with the current AWS
	// account during the specified period is greater than the value that you specified
	// for MaxItems, you can use Marker to return additional billing records. Get
	// the value of NextPageMarker from the previous response, and submit another
	// request that includes the value of NextPageMarker in the Marker element.
	//
	// Constraints: The marker must match the value of NextPageMarker that was returned
	// in the previous response.
	Marker *string `type:"string"`

	// The number of billing records to be returned.
	//
	// Default: 20
	MaxItems *int64 `type:"integer"`

	// The beginning date and time for the time period for which you want a list
	// of billing records. Specify the date and time in Coordinated Universal time
	// (UTC).
	Start *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ViewBillingInput) String() string {
	return awsutil.Prettify(s)
}

// The ViewBilling response includes the following elements.
type ViewBillingOutput struct {
	_ struct{} `type:"structure"`

	// A summary of billing records.
	BillingRecords []BillingRecord `type:"list"`

	// If there are more billing records than you specified for MaxItems in the
	// request, submit another request and include the value of NextPageMarker in
	// the value of Marker.
	NextPageMarker *string `type:"string"`
}

// String returns the string representation
func (s ViewBillingOutput) String() string {
	return awsutil.Prettify(s)
}