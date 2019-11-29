// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opDeleteVpcPeeringConnection = "DeleteVpcPeeringConnection"

// DeleteVpcPeeringConnectionRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Removes a VPC peering connection. To delete the connection, you must have
// a valid authorization for the VPC peering connection that you want to delete.
// You can check for an authorization by calling DescribeVpcPeeringAuthorizations
// or request a new one using CreateVpcPeeringAuthorization.
//
// Once a valid authorization exists, call this operation from the AWS account
// that is used to manage the Amazon GameLift fleets. Identify the connection
// to delete by the connection ID and fleet ID. If successful, the connection
// is removed.
//
//    * CreateVpcPeeringAuthorization
//
//    * DescribeVpcPeeringAuthorizations
//
//    * DeleteVpcPeeringAuthorization
//
//    * CreateVpcPeeringConnection
//
//    * DescribeVpcPeeringConnections
//
//    * DeleteVpcPeeringConnection
//
//    // Example sending a request using DeleteVpcPeeringConnectionRequest.
//    req := client.DeleteVpcPeeringConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteVpcPeeringConnection
func (c *Client) DeleteVpcPeeringConnectionRequest(input *types.DeleteVpcPeeringConnectionInput) DeleteVpcPeeringConnectionRequest {
	op := &aws.Operation{
		Name:       opDeleteVpcPeeringConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVpcPeeringConnectionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVpcPeeringConnectionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteVpcPeeringConnectionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteVpcPeeringConnectionRequest{Request: req, Input: input, Copy: c.DeleteVpcPeeringConnectionRequest}
}

// DeleteVpcPeeringConnectionRequest is the request type for the
// DeleteVpcPeeringConnection API operation.
type DeleteVpcPeeringConnectionRequest struct {
	*aws.Request
	Input *types.DeleteVpcPeeringConnectionInput
	Copy  func(*types.DeleteVpcPeeringConnectionInput) DeleteVpcPeeringConnectionRequest
}

// Send marshals and sends the DeleteVpcPeeringConnection API request.
func (r DeleteVpcPeeringConnectionRequest) Send(ctx context.Context) (*DeleteVpcPeeringConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcPeeringConnectionResponse{
		DeleteVpcPeeringConnectionOutput: r.Request.Data.(*types.DeleteVpcPeeringConnectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcPeeringConnectionResponse is the response type for the
// DeleteVpcPeeringConnection API operation.
type DeleteVpcPeeringConnectionResponse struct {
	*types.DeleteVpcPeeringConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpcPeeringConnection request.
func (r *DeleteVpcPeeringConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
