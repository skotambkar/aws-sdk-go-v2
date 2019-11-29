// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

const opCreatePrivateDnsNamespace = "CreatePrivateDnsNamespace"

// CreatePrivateDnsNamespaceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Creates a private namespace based on DNS, which will be visible only inside
// a specified Amazon VPC. The namespace defines your service naming scheme.
// For example, if you name your namespace example.com and name your service
// backend, the resulting DNS name for the service will be backend.example.com.
// For the current limit on the number of namespaces that you can create using
// the same AWS account, see AWS Cloud Map Limits (http://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the AWS Cloud Map Developer Guide.
//
//    // Example sending a request using CreatePrivateDnsNamespaceRequest.
//    req := client.CreatePrivateDnsNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/CreatePrivateDnsNamespace
func (c *Client) CreatePrivateDnsNamespaceRequest(input *types.CreatePrivateDnsNamespaceInput) CreatePrivateDnsNamespaceRequest {
	op := &aws.Operation{
		Name:       opCreatePrivateDnsNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePrivateDnsNamespaceInput{}
	}

	req := c.newRequest(op, input, &types.CreatePrivateDnsNamespaceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreatePrivateDnsNamespaceMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePrivateDnsNamespaceRequest{Request: req, Input: input, Copy: c.CreatePrivateDnsNamespaceRequest}
}

// CreatePrivateDnsNamespaceRequest is the request type for the
// CreatePrivateDnsNamespace API operation.
type CreatePrivateDnsNamespaceRequest struct {
	*aws.Request
	Input *types.CreatePrivateDnsNamespaceInput
	Copy  func(*types.CreatePrivateDnsNamespaceInput) CreatePrivateDnsNamespaceRequest
}

// Send marshals and sends the CreatePrivateDnsNamespace API request.
func (r CreatePrivateDnsNamespaceRequest) Send(ctx context.Context) (*CreatePrivateDnsNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePrivateDnsNamespaceResponse{
		CreatePrivateDnsNamespaceOutput: r.Request.Data.(*types.CreatePrivateDnsNamespaceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePrivateDnsNamespaceResponse is the response type for the
// CreatePrivateDnsNamespace API operation.
type CreatePrivateDnsNamespaceResponse struct {
	*types.CreatePrivateDnsNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePrivateDnsNamespace request.
func (r *CreatePrivateDnsNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
