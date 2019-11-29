// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opDeleteChangeSet = "DeleteChangeSet"

// DeleteChangeSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set.
//
// If the call successfully completes, AWS CloudFormation successfully deleted
// the change set.
//
//    // Example sending a request using DeleteChangeSetRequest.
//    req := client.DeleteChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteChangeSet
func (c *Client) DeleteChangeSetRequest(input *types.DeleteChangeSetInput) DeleteChangeSetRequest {
	op := &aws.Operation{
		Name:       opDeleteChangeSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteChangeSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteChangeSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteChangeSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteChangeSetRequest{Request: req, Input: input, Copy: c.DeleteChangeSetRequest}
}

// DeleteChangeSetRequest is the request type for the
// DeleteChangeSet API operation.
type DeleteChangeSetRequest struct {
	*aws.Request
	Input *types.DeleteChangeSetInput
	Copy  func(*types.DeleteChangeSetInput) DeleteChangeSetRequest
}

// Send marshals and sends the DeleteChangeSet API request.
func (r DeleteChangeSetRequest) Send(ctx context.Context) (*DeleteChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteChangeSetResponse{
		DeleteChangeSetOutput: r.Request.Data.(*types.DeleteChangeSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteChangeSetResponse is the response type for the
// DeleteChangeSet API operation.
type DeleteChangeSetResponse struct {
	*types.DeleteChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteChangeSet request.
func (r *DeleteChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
