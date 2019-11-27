// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opCreateVPCEConfiguration = "CreateVPCEConfiguration"

// CreateVPCEConfigurationRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Creates a configuration record in Device Farm for your Amazon Virtual Private
// Cloud (VPC) endpoint.
//
//    // Example sending a request using CreateVPCEConfigurationRequest.
//    req := client.CreateVPCEConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateVPCEConfiguration
func (c *Client) CreateVPCEConfigurationRequest(input *types.CreateVPCEConfigurationInput) CreateVPCEConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateVPCEConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateVPCEConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.CreateVPCEConfigurationOutput{})
	return CreateVPCEConfigurationRequest{Request: req, Input: input, Copy: c.CreateVPCEConfigurationRequest}
}

// CreateVPCEConfigurationRequest is the request type for the
// CreateVPCEConfiguration API operation.
type CreateVPCEConfigurationRequest struct {
	*aws.Request
	Input *types.CreateVPCEConfigurationInput
	Copy  func(*types.CreateVPCEConfigurationInput) CreateVPCEConfigurationRequest
}

// Send marshals and sends the CreateVPCEConfiguration API request.
func (r CreateVPCEConfigurationRequest) Send(ctx context.Context) (*CreateVPCEConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVPCEConfigurationResponse{
		CreateVPCEConfigurationOutput: r.Request.Data.(*types.CreateVPCEConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVPCEConfigurationResponse is the response type for the
// CreateVPCEConfiguration API operation.
type CreateVPCEConfigurationResponse struct {
	*types.CreateVPCEConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVPCEConfiguration request.
func (r *CreateVPCEConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
