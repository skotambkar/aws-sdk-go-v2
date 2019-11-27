// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

const opCreateHttpNamespace = "CreateHttpNamespace"

// CreateHttpNamespaceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Creates an HTTP namespace. Service instances that you register using an HTTP
// namespace can be discovered using a DiscoverInstances request but can't be
// discovered using DNS.
//
// For the current limit on the number of namespaces that you can create using
// the same AWS account, see AWS Cloud Map Limits (http://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the AWS Cloud Map Developer Guide.
//
//    // Example sending a request using CreateHttpNamespaceRequest.
//    req := client.CreateHttpNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/CreateHttpNamespace
func (c *Client) CreateHttpNamespaceRequest(input *types.CreateHttpNamespaceInput) CreateHttpNamespaceRequest {
	op := &aws.Operation{
		Name:       opCreateHttpNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateHttpNamespaceInput{}
	}

	req := c.newRequest(op, input, &types.CreateHttpNamespaceOutput{})
	return CreateHttpNamespaceRequest{Request: req, Input: input, Copy: c.CreateHttpNamespaceRequest}
}

// CreateHttpNamespaceRequest is the request type for the
// CreateHttpNamespace API operation.
type CreateHttpNamespaceRequest struct {
	*aws.Request
	Input *types.CreateHttpNamespaceInput
	Copy  func(*types.CreateHttpNamespaceInput) CreateHttpNamespaceRequest
}

// Send marshals and sends the CreateHttpNamespace API request.
func (r CreateHttpNamespaceRequest) Send(ctx context.Context) (*CreateHttpNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHttpNamespaceResponse{
		CreateHttpNamespaceOutput: r.Request.Data.(*types.CreateHttpNamespaceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHttpNamespaceResponse is the response type for the
// CreateHttpNamespace API operation.
type CreateHttpNamespaceResponse struct {
	*types.CreateHttpNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHttpNamespace request.
func (r *CreateHttpNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
