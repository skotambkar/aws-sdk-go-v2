// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opRegisterResource = "RegisterResource"

// RegisterResourceRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Registers the resource as managed by the Data Catalog.
//
// To add or update data, Lake Formation needs read/write access to the chosen
// Amazon S3 path. Choose a role that you know has permission to do this, or
// choose the AWSServiceRoleForLakeFormationDataAccess service-linked role.
// When you register the first Amazon S3 path, the service-linked role and a
// new inline policy are created on your behalf. Lake Formation adds the first
// path to the inline policy and attaches it to the service-linked role. When
// you register subsequent paths, Lake Formation adds the path to the existing
// policy.
//
//    // Example sending a request using RegisterResourceRequest.
//    req := client.RegisterResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/RegisterResource
func (c *Client) RegisterResourceRequest(input *types.RegisterResourceInput) RegisterResourceRequest {
	op := &aws.Operation{
		Name:       opRegisterResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterResourceInput{}
	}

	req := c.newRequest(op, input, &types.RegisterResourceOutput{})
	return RegisterResourceRequest{Request: req, Input: input, Copy: c.RegisterResourceRequest}
}

// RegisterResourceRequest is the request type for the
// RegisterResource API operation.
type RegisterResourceRequest struct {
	*aws.Request
	Input *types.RegisterResourceInput
	Copy  func(*types.RegisterResourceInput) RegisterResourceRequest
}

// Send marshals and sends the RegisterResource API request.
func (r RegisterResourceRequest) Send(ctx context.Context) (*RegisterResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterResourceResponse{
		RegisterResourceOutput: r.Request.Data.(*types.RegisterResourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterResourceResponse is the response type for the
// RegisterResource API operation.
type RegisterResourceResponse struct {
	*types.RegisterResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterResource request.
func (r *RegisterResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
