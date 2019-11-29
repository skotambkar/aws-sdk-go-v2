// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opSearchTransitGatewayRoutes = "SearchTransitGatewayRoutes"

// SearchTransitGatewayRoutesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Searches for routes in the specified transit gateway route table.
//
//    // Example sending a request using SearchTransitGatewayRoutesRequest.
//    req := client.SearchTransitGatewayRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/SearchTransitGatewayRoutes
func (c *Client) SearchTransitGatewayRoutesRequest(input *types.SearchTransitGatewayRoutesInput) SearchTransitGatewayRoutesRequest {
	op := &aws.Operation{
		Name:       opSearchTransitGatewayRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SearchTransitGatewayRoutesInput{}
	}

	req := c.newRequest(op, input, &types.SearchTransitGatewayRoutesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.SearchTransitGatewayRoutesMarshaler{Input: input}.GetNamedBuildHandler())

	return SearchTransitGatewayRoutesRequest{Request: req, Input: input, Copy: c.SearchTransitGatewayRoutesRequest}
}

// SearchTransitGatewayRoutesRequest is the request type for the
// SearchTransitGatewayRoutes API operation.
type SearchTransitGatewayRoutesRequest struct {
	*aws.Request
	Input *types.SearchTransitGatewayRoutesInput
	Copy  func(*types.SearchTransitGatewayRoutesInput) SearchTransitGatewayRoutesRequest
}

// Send marshals and sends the SearchTransitGatewayRoutes API request.
func (r SearchTransitGatewayRoutesRequest) Send(ctx context.Context) (*SearchTransitGatewayRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchTransitGatewayRoutesResponse{
		SearchTransitGatewayRoutesOutput: r.Request.Data.(*types.SearchTransitGatewayRoutesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchTransitGatewayRoutesResponse is the response type for the
// SearchTransitGatewayRoutes API operation.
type SearchTransitGatewayRoutesResponse struct {
	*types.SearchTransitGatewayRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchTransitGatewayRoutes request.
func (r *SearchTransitGatewayRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
