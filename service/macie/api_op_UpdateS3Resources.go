// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/macie/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
)

const opUpdateS3Resources = "UpdateS3Resources"

// UpdateS3ResourcesRequest returns a request value for making API operation for
// Amazon Macie.
//
// Updates the classification types for the specified S3 resources. If memberAccountId
// isn't specified, the action updates the classification types of the S3 resources
// associated with Amazon Macie for the current master account. If memberAccountId
// is specified, the action updates the classification types of the S3 resources
// associated with Amazon Macie for the specified member account.
//
//    // Example sending a request using UpdateS3ResourcesRequest.
//    req := client.UpdateS3ResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/UpdateS3Resources
func (c *Client) UpdateS3ResourcesRequest(input *types.UpdateS3ResourcesInput) UpdateS3ResourcesRequest {
	op := &aws.Operation{
		Name:       opUpdateS3Resources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateS3ResourcesInput{}
	}

	req := c.newRequest(op, input, &types.UpdateS3ResourcesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateS3ResourcesMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateS3ResourcesRequest{Request: req, Input: input, Copy: c.UpdateS3ResourcesRequest}
}

// UpdateS3ResourcesRequest is the request type for the
// UpdateS3Resources API operation.
type UpdateS3ResourcesRequest struct {
	*aws.Request
	Input *types.UpdateS3ResourcesInput
	Copy  func(*types.UpdateS3ResourcesInput) UpdateS3ResourcesRequest
}

// Send marshals and sends the UpdateS3Resources API request.
func (r UpdateS3ResourcesRequest) Send(ctx context.Context) (*UpdateS3ResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateS3ResourcesResponse{
		UpdateS3ResourcesOutput: r.Request.Data.(*types.UpdateS3ResourcesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateS3ResourcesResponse is the response type for the
// UpdateS3Resources API operation.
type UpdateS3ResourcesResponse struct {
	*types.UpdateS3ResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateS3Resources request.
func (r *UpdateS3ResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
