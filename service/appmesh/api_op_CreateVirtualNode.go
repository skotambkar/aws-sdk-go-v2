// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

const opCreateVirtualNode = "CreateVirtualNode"

// CreateVirtualNodeRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such
// as an Amazon ECS service or a Kubernetes deployment. When you create a virtual
// node, you can specify the service discovery information for your task group.
//
// Any inbound traffic that your virtual node expects should be specified as
// a listener. Any outbound traffic that your virtual node expects to reach
// should be specified as a backend.
//
// The response metadata for your new virtual node contains the arn that is
// associated with the virtual node. Set this value (either the full ARN or
// the truncated resource name: for example, mesh/default/virtualNode/simpleapp)
// as the APPMESH_VIRTUAL_NODE_NAME environment variable for your task group's
// Envoy proxy container in your task definition or pod spec. This is then mapped
// to the node.id and node.cluster Envoy parameters.
//
// If you require your Envoy stats or tracing to use a different name, you can
// override the node.cluster value that is set by APPMESH_VIRTUAL_NODE_NAME
// with the APPMESH_VIRTUAL_NODE_CLUSTER environment variable.
//
//    // Example sending a request using CreateVirtualNodeRequest.
//    req := client.CreateVirtualNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualNode
func (c *Client) CreateVirtualNodeRequest(input *types.CreateVirtualNodeInput) CreateVirtualNodeRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualNode,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualNodes",
	}

	if input == nil {
		input = &types.CreateVirtualNodeInput{}
	}

	req := c.newRequest(op, input, &types.CreateVirtualNodeOutput{})
	return CreateVirtualNodeRequest{Request: req, Input: input, Copy: c.CreateVirtualNodeRequest}
}

// CreateVirtualNodeRequest is the request type for the
// CreateVirtualNode API operation.
type CreateVirtualNodeRequest struct {
	*aws.Request
	Input *types.CreateVirtualNodeInput
	Copy  func(*types.CreateVirtualNodeInput) CreateVirtualNodeRequest
}

// Send marshals and sends the CreateVirtualNode API request.
func (r CreateVirtualNodeRequest) Send(ctx context.Context) (*CreateVirtualNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualNodeResponse{
		CreateVirtualNodeOutput: r.Request.Data.(*types.CreateVirtualNodeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualNodeResponse is the response type for the
// CreateVirtualNode API operation.
type CreateVirtualNodeResponse struct {
	*types.CreateVirtualNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualNode request.
func (r *CreateVirtualNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
