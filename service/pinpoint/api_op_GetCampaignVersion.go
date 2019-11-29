// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetCampaignVersion = "GetCampaignVersion"

// GetCampaignVersionRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for a specific version of a campaign.
//
//    // Example sending a request using GetCampaignVersionRequest.
//    req := client.GetCampaignVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersion
func (c *Client) GetCampaignVersionRequest(input *types.GetCampaignVersionInput) GetCampaignVersionRequest {
	op := &aws.Operation{
		Name:       opGetCampaignVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns/{campaign-id}/versions/{version}",
	}

	if input == nil {
		input = &types.GetCampaignVersionInput{}
	}

	req := c.newRequest(op, input, &types.GetCampaignVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetCampaignVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCampaignVersionRequest{Request: req, Input: input, Copy: c.GetCampaignVersionRequest}
}

// GetCampaignVersionRequest is the request type for the
// GetCampaignVersion API operation.
type GetCampaignVersionRequest struct {
	*aws.Request
	Input *types.GetCampaignVersionInput
	Copy  func(*types.GetCampaignVersionInput) GetCampaignVersionRequest
}

// Send marshals and sends the GetCampaignVersion API request.
func (r GetCampaignVersionRequest) Send(ctx context.Context) (*GetCampaignVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignVersionResponse{
		GetCampaignVersionOutput: r.Request.Data.(*types.GetCampaignVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignVersionResponse is the response type for the
// GetCampaignVersion API operation.
type GetCampaignVersionResponse struct {
	*types.GetCampaignVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaignVersion request.
func (r *GetCampaignVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
