// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetCampaigns = "GetCampaigns"

// GetCampaignsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for all the campaigns that are associated with an application.
//
//    // Example sending a request using GetCampaignsRequest.
//    req := client.GetCampaignsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaigns
func (c *Client) GetCampaignsRequest(input *types.GetCampaignsInput) GetCampaignsRequest {
	op := &aws.Operation{
		Name:       opGetCampaigns,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns",
	}

	if input == nil {
		input = &types.GetCampaignsInput{}
	}

	req := c.newRequest(op, input, &types.GetCampaignsOutput{})
	return GetCampaignsRequest{Request: req, Input: input, Copy: c.GetCampaignsRequest}
}

// GetCampaignsRequest is the request type for the
// GetCampaigns API operation.
type GetCampaignsRequest struct {
	*aws.Request
	Input *types.GetCampaignsInput
	Copy  func(*types.GetCampaignsInput) GetCampaignsRequest
}

// Send marshals and sends the GetCampaigns API request.
func (r GetCampaignsRequest) Send(ctx context.Context) (*GetCampaignsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignsResponse{
		GetCampaignsOutput: r.Request.Data.(*types.GetCampaignsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignsResponse is the response type for the
// GetCampaigns API operation.
type GetCampaignsResponse struct {
	*types.GetCampaignsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaigns request.
func (r *GetCampaignsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
