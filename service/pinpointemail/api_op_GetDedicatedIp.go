// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opGetDedicatedIp = "GetDedicatedIp"

// GetDedicatedIpRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Get information about a dedicated IP address, including the name of the dedicated
// IP pool that it's associated with, as well information about the automatic
// warm-up process for the address.
//
//    // Example sending a request using GetDedicatedIpRequest.
//    req := client.GetDedicatedIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDedicatedIp
func (c *Client) GetDedicatedIpRequest(input *types.GetDedicatedIpInput) GetDedicatedIpRequest {
	op := &aws.Operation{
		Name:       opGetDedicatedIp,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/dedicated-ips/{IP}",
	}

	if input == nil {
		input = &types.GetDedicatedIpInput{}
	}

	req := c.newRequest(op, input, &types.GetDedicatedIpOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetDedicatedIpMarshaler{Input: input}.GetNamedBuildHandler())

	return GetDedicatedIpRequest{Request: req, Input: input, Copy: c.GetDedicatedIpRequest}
}

// GetDedicatedIpRequest is the request type for the
// GetDedicatedIp API operation.
type GetDedicatedIpRequest struct {
	*aws.Request
	Input *types.GetDedicatedIpInput
	Copy  func(*types.GetDedicatedIpInput) GetDedicatedIpRequest
}

// Send marshals and sends the GetDedicatedIp API request.
func (r GetDedicatedIpRequest) Send(ctx context.Context) (*GetDedicatedIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDedicatedIpResponse{
		GetDedicatedIpOutput: r.Request.Data.(*types.GetDedicatedIpOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDedicatedIpResponse is the response type for the
// GetDedicatedIp API operation.
type GetDedicatedIpResponse struct {
	*types.GetDedicatedIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDedicatedIp request.
func (r *GetDedicatedIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
