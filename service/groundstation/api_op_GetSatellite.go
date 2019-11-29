// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
)

const opGetSatellite = "GetSatellite"

// GetSatelliteRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a satellite.
//
//    // Example sending a request using GetSatelliteRequest.
//    req := client.GetSatelliteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetSatellite
func (c *Client) GetSatelliteRequest(input *types.GetSatelliteInput) GetSatelliteRequest {
	op := &aws.Operation{
		Name:       opGetSatellite,
		HTTPMethod: "GET",
		HTTPPath:   "/satellite/{satelliteId}",
	}

	if input == nil {
		input = &types.GetSatelliteInput{}
	}

	req := c.newRequest(op, input, &types.GetSatelliteOutput{})
	return GetSatelliteRequest{Request: req, Input: input, Copy: c.GetSatelliteRequest}
}

// GetSatelliteRequest is the request type for the
// GetSatellite API operation.
type GetSatelliteRequest struct {
	*aws.Request
	Input *types.GetSatelliteInput
	Copy  func(*types.GetSatelliteInput) GetSatelliteRequest
}

// Send marshals and sends the GetSatellite API request.
func (r GetSatelliteRequest) Send(ctx context.Context) (*GetSatelliteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSatelliteResponse{
		GetSatelliteOutput: r.Request.Data.(*types.GetSatelliteOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSatelliteResponse is the response type for the
// GetSatellite API operation.
type GetSatelliteResponse struct {
	*types.GetSatelliteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSatellite request.
func (r *GetSatelliteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
