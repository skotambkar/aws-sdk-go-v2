// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified set of tags from the specified set of resources.
//
// To list the current tags, use DescribeTags. For more information about tags,
// see Tagging Your Resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTags
func (c *Client) DeleteTagsRequest(input *types.DeleteTagsInput) DeleteTagsRequest {
	op := &aws.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTagsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTagsOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteTagsRequest{Request: req, Input: input, Copy: c.DeleteTagsRequest}
}

// DeleteTagsRequest is the request type for the
// DeleteTags API operation.
type DeleteTagsRequest struct {
	*aws.Request
	Input *types.DeleteTagsInput
	Copy  func(*types.DeleteTagsInput) DeleteTagsRequest
}

// Send marshals and sends the DeleteTags API request.
func (r DeleteTagsRequest) Send(ctx context.Context) (*DeleteTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsResponse{
		DeleteTagsOutput: r.Request.Data.(*types.DeleteTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsResponse is the response type for the
// DeleteTags API operation.
type DeleteTagsResponse struct {
	*types.DeleteTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTags request.
func (r *DeleteTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
