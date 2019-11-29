// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

const opCreateVirtualService = "CreateVirtualService"

// CreateVirtualServiceRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual service within a service mesh.
//
// A virtual service is an abstraction of a real service that is provided by
// a virtual node directly or indirectly by means of a virtual router. Dependent
// services call your virtual service by its virtualServiceName, and those requests
// are routed to the virtual node or virtual router that is specified as the
// provider for the virtual service.
//
//    // Example sending a request using CreateVirtualServiceRequest.
//    req := client.CreateVirtualServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualService
func (c *Client) CreateVirtualServiceRequest(input *types.CreateVirtualServiceInput) CreateVirtualServiceRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualService,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualServices",
	}

	if input == nil {
		input = &types.CreateVirtualServiceInput{}
	}

	req := c.newRequest(op, input, &types.CreateVirtualServiceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateVirtualServiceMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateVirtualServiceRequest{Request: req, Input: input, Copy: c.CreateVirtualServiceRequest}
}

// CreateVirtualServiceRequest is the request type for the
// CreateVirtualService API operation.
type CreateVirtualServiceRequest struct {
	*aws.Request
	Input *types.CreateVirtualServiceInput
	Copy  func(*types.CreateVirtualServiceInput) CreateVirtualServiceRequest
}

// Send marshals and sends the CreateVirtualService API request.
func (r CreateVirtualServiceRequest) Send(ctx context.Context) (*CreateVirtualServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualServiceResponse{
		CreateVirtualServiceOutput: r.Request.Data.(*types.CreateVirtualServiceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualServiceResponse is the response type for the
// CreateVirtualService API operation.
type CreateVirtualServiceResponse struct {
	*types.CreateVirtualServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualService request.
func (r *CreateVirtualServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
