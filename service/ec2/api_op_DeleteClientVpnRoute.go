// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteClientVpnRoute = "DeleteClientVpnRoute"

// DeleteClientVpnRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes a route from a Client VPN endpoint. You can only delete routes that
// you manually added using the CreateClientVpnRoute action. You cannot delete
// routes that were automatically added when associating a subnet. To remove
// routes that have been automatically added, disassociate the target subnet
// from the Client VPN endpoint.
//
//    // Example sending a request using DeleteClientVpnRouteRequest.
//    req := client.DeleteClientVpnRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteClientVpnRoute
func (c *Client) DeleteClientVpnRouteRequest(input *types.DeleteClientVpnRouteInput) DeleteClientVpnRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteClientVpnRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteClientVpnRouteInput{}
	}

	req := c.newRequest(op, input, &types.DeleteClientVpnRouteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteClientVpnRouteMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteClientVpnRouteRequest{Request: req, Input: input, Copy: c.DeleteClientVpnRouteRequest}
}

// DeleteClientVpnRouteRequest is the request type for the
// DeleteClientVpnRoute API operation.
type DeleteClientVpnRouteRequest struct {
	*aws.Request
	Input *types.DeleteClientVpnRouteInput
	Copy  func(*types.DeleteClientVpnRouteInput) DeleteClientVpnRouteRequest
}

// Send marshals and sends the DeleteClientVpnRoute API request.
func (r DeleteClientVpnRouteRequest) Send(ctx context.Context) (*DeleteClientVpnRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClientVpnRouteResponse{
		DeleteClientVpnRouteOutput: r.Request.Data.(*types.DeleteClientVpnRouteOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClientVpnRouteResponse is the response type for the
// DeleteClientVpnRoute API operation.
type DeleteClientVpnRouteResponse struct {
	*types.DeleteClientVpnRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteClientVpnRoute request.
func (r *DeleteClientVpnRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
