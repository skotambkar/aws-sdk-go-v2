// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/ram/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
)

const opAssociateResourceShare = "AssociateResourceShare"

// AssociateResourceShareRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Associates the specified resource share with the specified principals and
// resources.
//
//    // Example sending a request using AssociateResourceShareRequest.
//    req := client.AssociateResourceShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AssociateResourceShare
func (c *Client) AssociateResourceShareRequest(input *types.AssociateResourceShareInput) AssociateResourceShareRequest {
	op := &aws.Operation{
		Name:       opAssociateResourceShare,
		HTTPMethod: "POST",
		HTTPPath:   "/associateresourceshare",
	}

	if input == nil {
		input = &types.AssociateResourceShareInput{}
	}

	req := c.newRequest(op, input, &types.AssociateResourceShareOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.AssociateResourceShareMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateResourceShareRequest{Request: req, Input: input, Copy: c.AssociateResourceShareRequest}
}

// AssociateResourceShareRequest is the request type for the
// AssociateResourceShare API operation.
type AssociateResourceShareRequest struct {
	*aws.Request
	Input *types.AssociateResourceShareInput
	Copy  func(*types.AssociateResourceShareInput) AssociateResourceShareRequest
}

// Send marshals and sends the AssociateResourceShare API request.
func (r AssociateResourceShareRequest) Send(ctx context.Context) (*AssociateResourceShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateResourceShareResponse{
		AssociateResourceShareOutput: r.Request.Data.(*types.AssociateResourceShareOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateResourceShareResponse is the response type for the
// AssociateResourceShare API operation.
type AssociateResourceShareResponse struct {
	*types.AssociateResourceShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateResourceShare request.
func (r *AssociateResourceShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
