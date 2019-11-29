// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opListVPCAssociationAuthorizations = "ListVPCAssociationAuthorizations"

// ListVPCAssociationAuthorizationsRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets a list of the VPCs that were created by other accounts and that can
// be associated with a specified hosted zone because you've submitted one or
// more CreateVPCAssociationAuthorization requests.
//
// The response includes a VPCs element with a VPC child element for each VPC
// that can be associated with the hosted zone.
//
//    // Example sending a request using ListVPCAssociationAuthorizationsRequest.
//    req := client.ListVPCAssociationAuthorizationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListVPCAssociationAuthorizations
func (c *Client) ListVPCAssociationAuthorizationsRequest(input *types.ListVPCAssociationAuthorizationsInput) ListVPCAssociationAuthorizationsRequest {
	op := &aws.Operation{
		Name:       opListVPCAssociationAuthorizations,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}/authorizevpcassociation",
	}

	if input == nil {
		input = &types.ListVPCAssociationAuthorizationsInput{}
	}

	req := c.newRequest(op, input, &types.ListVPCAssociationAuthorizationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.ListVPCAssociationAuthorizationsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListVPCAssociationAuthorizationsRequest{Request: req, Input: input, Copy: c.ListVPCAssociationAuthorizationsRequest}
}

// ListVPCAssociationAuthorizationsRequest is the request type for the
// ListVPCAssociationAuthorizations API operation.
type ListVPCAssociationAuthorizationsRequest struct {
	*aws.Request
	Input *types.ListVPCAssociationAuthorizationsInput
	Copy  func(*types.ListVPCAssociationAuthorizationsInput) ListVPCAssociationAuthorizationsRequest
}

// Send marshals and sends the ListVPCAssociationAuthorizations API request.
func (r ListVPCAssociationAuthorizationsRequest) Send(ctx context.Context) (*ListVPCAssociationAuthorizationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVPCAssociationAuthorizationsResponse{
		ListVPCAssociationAuthorizationsOutput: r.Request.Data.(*types.ListVPCAssociationAuthorizationsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVPCAssociationAuthorizationsResponse is the response type for the
// ListVPCAssociationAuthorizations API operation.
type ListVPCAssociationAuthorizationsResponse struct {
	*types.ListVPCAssociationAuthorizationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVPCAssociationAuthorizations request.
func (r *ListVPCAssociationAuthorizationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
