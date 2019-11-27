// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opListGeoLocations = "ListGeoLocations"

// ListGeoLocationsRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Retrieves a list of supported geographic locations.
//
// Countries are listed first, and continents are listed last. If Amazon Route
// 53 supports subdivisions for a country (for example, states or provinces),
// the subdivisions for that country are listed in alphabetical order immediately
// after the corresponding country.
//
//    // Example sending a request using ListGeoLocationsRequest.
//    req := client.ListGeoLocationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListGeoLocations
func (c *Client) ListGeoLocationsRequest(input *types.ListGeoLocationsInput) ListGeoLocationsRequest {
	op := &aws.Operation{
		Name:       opListGeoLocations,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/geolocations",
	}

	if input == nil {
		input = &types.ListGeoLocationsInput{}
	}

	req := c.newRequest(op, input, &types.ListGeoLocationsOutput{})
	return ListGeoLocationsRequest{Request: req, Input: input, Copy: c.ListGeoLocationsRequest}
}

// ListGeoLocationsRequest is the request type for the
// ListGeoLocations API operation.
type ListGeoLocationsRequest struct {
	*aws.Request
	Input *types.ListGeoLocationsInput
	Copy  func(*types.ListGeoLocationsInput) ListGeoLocationsRequest
}

// Send marshals and sends the ListGeoLocations API request.
func (r ListGeoLocationsRequest) Send(ctx context.Context) (*ListGeoLocationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGeoLocationsResponse{
		ListGeoLocationsOutput: r.Request.Data.(*types.ListGeoLocationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGeoLocationsResponse is the response type for the
// ListGeoLocations API operation.
type ListGeoLocationsResponse struct {
	*types.ListGeoLocationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGeoLocations request.
func (r *ListGeoLocationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
