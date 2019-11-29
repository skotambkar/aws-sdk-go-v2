// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteTransitGatewayVpcAttachment = "DeleteTransitGatewayVpcAttachment"

// DeleteTransitGatewayVpcAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified VPC attachment.
//
//    // Example sending a request using DeleteTransitGatewayVpcAttachmentRequest.
//    req := client.DeleteTransitGatewayVpcAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayVpcAttachment
func (c *Client) DeleteTransitGatewayVpcAttachmentRequest(input *types.DeleteTransitGatewayVpcAttachmentInput) DeleteTransitGatewayVpcAttachmentRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGatewayVpcAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTransitGatewayVpcAttachmentInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTransitGatewayVpcAttachmentOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteTransitGatewayVpcAttachmentMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteTransitGatewayVpcAttachmentRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayVpcAttachmentRequest}
}

// DeleteTransitGatewayVpcAttachmentRequest is the request type for the
// DeleteTransitGatewayVpcAttachment API operation.
type DeleteTransitGatewayVpcAttachmentRequest struct {
	*aws.Request
	Input *types.DeleteTransitGatewayVpcAttachmentInput
	Copy  func(*types.DeleteTransitGatewayVpcAttachmentInput) DeleteTransitGatewayVpcAttachmentRequest
}

// Send marshals and sends the DeleteTransitGatewayVpcAttachment API request.
func (r DeleteTransitGatewayVpcAttachmentRequest) Send(ctx context.Context) (*DeleteTransitGatewayVpcAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayVpcAttachmentResponse{
		DeleteTransitGatewayVpcAttachmentOutput: r.Request.Data.(*types.DeleteTransitGatewayVpcAttachmentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayVpcAttachmentResponse is the response type for the
// DeleteTransitGatewayVpcAttachment API operation.
type DeleteTransitGatewayVpcAttachmentResponse struct {
	*types.DeleteTransitGatewayVpcAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGatewayVpcAttachment request.
func (r *DeleteTransitGatewayVpcAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
