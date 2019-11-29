// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

const opCreateVirtualRouter = "CreateVirtualRouter"

// CreateVirtualRouterRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual router within a service mesh.
//
// Any inbound traffic that your virtual router expects should be specified
// as a listener.
//
// Virtual routers handle traffic for one or more virtual services within your
// mesh. After you create your virtual router, create and associate routes for
// your virtual router that direct incoming requests to different virtual nodes.
//
//    // Example sending a request using CreateVirtualRouterRequest.
//    req := client.CreateVirtualRouterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualRouter
func (c *Client) CreateVirtualRouterRequest(input *types.CreateVirtualRouterInput) CreateVirtualRouterRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualRouter,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouters",
	}

	if input == nil {
		input = &types.CreateVirtualRouterInput{}
	}

	req := c.newRequest(op, input, &types.CreateVirtualRouterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateVirtualRouterMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateVirtualRouterRequest{Request: req, Input: input, Copy: c.CreateVirtualRouterRequest}
}

// CreateVirtualRouterRequest is the request type for the
// CreateVirtualRouter API operation.
type CreateVirtualRouterRequest struct {
	*aws.Request
	Input *types.CreateVirtualRouterInput
	Copy  func(*types.CreateVirtualRouterInput) CreateVirtualRouterRequest
}

// Send marshals and sends the CreateVirtualRouter API request.
func (r CreateVirtualRouterRequest) Send(ctx context.Context) (*CreateVirtualRouterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualRouterResponse{
		CreateVirtualRouterOutput: r.Request.Data.(*types.CreateVirtualRouterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualRouterResponse is the response type for the
// CreateVirtualRouter API operation.
type CreateVirtualRouterResponse struct {
	*types.CreateVirtualRouterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualRouter request.
func (r *CreateVirtualRouterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
