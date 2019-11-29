// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to list origin access identities.
type ListCloudFrontOriginAccessIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of origin access identities. The results include identities in the list that
	// occur after the marker. To get the next page of results, set the Marker to
	// the value of the NextMarker from the current page's response (which is also
	// the ID of the last identity on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of origin access identities you want in the response body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListCloudFrontOriginAccessIdentitiesInput) String() string {
	return awsutil.Prettify(s)
}

// The returned result of the corresponding request.
type ListCloudFrontOriginAccessIdentitiesOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentityList"`

	// The CloudFrontOriginAccessIdentityList type.
	CloudFrontOriginAccessIdentityList *CloudFrontOriginAccessIdentityList `type:"structure"`
}

// String returns the string representation
func (s ListCloudFrontOriginAccessIdentitiesOutput) String() string {
	return awsutil.Prettify(s)
}
