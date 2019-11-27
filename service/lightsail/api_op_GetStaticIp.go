// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetStaticIp = "GetStaticIp"

// GetStaticIpRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about a specific static IP.
//
//    // Example sending a request using GetStaticIpRequest.
//    req := client.GetStaticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetStaticIp
func (c *Client) GetStaticIpRequest(input *types.GetStaticIpInput) GetStaticIpRequest {
	op := &aws.Operation{
		Name:       opGetStaticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetStaticIpInput{}
	}

	req := c.newRequest(op, input, &types.GetStaticIpOutput{})
	return GetStaticIpRequest{Request: req, Input: input, Copy: c.GetStaticIpRequest}
}

// GetStaticIpRequest is the request type for the
// GetStaticIp API operation.
type GetStaticIpRequest struct {
	*aws.Request
	Input *types.GetStaticIpInput
	Copy  func(*types.GetStaticIpInput) GetStaticIpRequest
}

// Send marshals and sends the GetStaticIp API request.
func (r GetStaticIpRequest) Send(ctx context.Context) (*GetStaticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStaticIpResponse{
		GetStaticIpOutput: r.Request.Data.(*types.GetStaticIpOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStaticIpResponse is the response type for the
// GetStaticIp API operation.
type GetStaticIpResponse struct {
	*types.GetStaticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStaticIp request.
func (r *GetStaticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
