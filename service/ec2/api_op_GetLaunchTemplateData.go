// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opGetLaunchTemplateData = "GetLaunchTemplateData"

// GetLaunchTemplateDataRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Retrieves the configuration data of the specified instance. You can use this
// data to create a launch template.
//
//    // Example sending a request using GetLaunchTemplateDataRequest.
//    req := client.GetLaunchTemplateDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetLaunchTemplateData
func (c *Client) GetLaunchTemplateDataRequest(input *types.GetLaunchTemplateDataInput) GetLaunchTemplateDataRequest {
	op := &aws.Operation{
		Name:       opGetLaunchTemplateData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetLaunchTemplateDataInput{}
	}

	req := c.newRequest(op, input, &types.GetLaunchTemplateDataOutput{})
	return GetLaunchTemplateDataRequest{Request: req, Input: input, Copy: c.GetLaunchTemplateDataRequest}
}

// GetLaunchTemplateDataRequest is the request type for the
// GetLaunchTemplateData API operation.
type GetLaunchTemplateDataRequest struct {
	*aws.Request
	Input *types.GetLaunchTemplateDataInput
	Copy  func(*types.GetLaunchTemplateDataInput) GetLaunchTemplateDataRequest
}

// Send marshals and sends the GetLaunchTemplateData API request.
func (r GetLaunchTemplateDataRequest) Send(ctx context.Context) (*GetLaunchTemplateDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLaunchTemplateDataResponse{
		GetLaunchTemplateDataOutput: r.Request.Data.(*types.GetLaunchTemplateDataOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLaunchTemplateDataResponse is the response type for the
// GetLaunchTemplateData API operation.
type GetLaunchTemplateDataResponse struct {
	*types.GetLaunchTemplateDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLaunchTemplateData request.
func (r *GetLaunchTemplateDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
