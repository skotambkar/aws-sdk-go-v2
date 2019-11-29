// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opAddTagsToOnPremisesInstances = "AddTagsToOnPremisesInstances"

// AddTagsToOnPremisesInstancesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Adds tags to on-premises instances.
//
//    // Example sending a request using AddTagsToOnPremisesInstancesRequest.
//    req := client.AddTagsToOnPremisesInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/AddTagsToOnPremisesInstances
func (c *Client) AddTagsToOnPremisesInstancesRequest(input *types.AddTagsToOnPremisesInstancesInput) AddTagsToOnPremisesInstancesRequest {
	op := &aws.Operation{
		Name:       opAddTagsToOnPremisesInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddTagsToOnPremisesInstancesInput{}
	}

	req := c.newRequest(op, input, &types.AddTagsToOnPremisesInstancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AddTagsToOnPremisesInstancesMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddTagsToOnPremisesInstancesRequest{Request: req, Input: input, Copy: c.AddTagsToOnPremisesInstancesRequest}
}

// AddTagsToOnPremisesInstancesRequest is the request type for the
// AddTagsToOnPremisesInstances API operation.
type AddTagsToOnPremisesInstancesRequest struct {
	*aws.Request
	Input *types.AddTagsToOnPremisesInstancesInput
	Copy  func(*types.AddTagsToOnPremisesInstancesInput) AddTagsToOnPremisesInstancesRequest
}

// Send marshals and sends the AddTagsToOnPremisesInstances API request.
func (r AddTagsToOnPremisesInstancesRequest) Send(ctx context.Context) (*AddTagsToOnPremisesInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToOnPremisesInstancesResponse{
		AddTagsToOnPremisesInstancesOutput: r.Request.Data.(*types.AddTagsToOnPremisesInstancesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToOnPremisesInstancesResponse is the response type for the
// AddTagsToOnPremisesInstances API operation.
type AddTagsToOnPremisesInstancesResponse struct {
	*types.AddTagsToOnPremisesInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToOnPremisesInstances request.
func (r *AddTagsToOnPremisesInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
