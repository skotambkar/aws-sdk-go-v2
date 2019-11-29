// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opCreateReplicationSubnetGroup = "CreateReplicationSubnetGroup"

// CreateReplicationSubnetGroupRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Creates a replication subnet group given a list of the subnet IDs in a VPC.
//
//    // Example sending a request using CreateReplicationSubnetGroupRequest.
//    req := client.CreateReplicationSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateReplicationSubnetGroup
func (c *Client) CreateReplicationSubnetGroupRequest(input *types.CreateReplicationSubnetGroupInput) CreateReplicationSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateReplicationSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateReplicationSubnetGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateReplicationSubnetGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateReplicationSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateReplicationSubnetGroupRequest}
}

// CreateReplicationSubnetGroupRequest is the request type for the
// CreateReplicationSubnetGroup API operation.
type CreateReplicationSubnetGroupRequest struct {
	*aws.Request
	Input *types.CreateReplicationSubnetGroupInput
	Copy  func(*types.CreateReplicationSubnetGroupInput) CreateReplicationSubnetGroupRequest
}

// Send marshals and sends the CreateReplicationSubnetGroup API request.
func (r CreateReplicationSubnetGroupRequest) Send(ctx context.Context) (*CreateReplicationSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationSubnetGroupResponse{
		CreateReplicationSubnetGroupOutput: r.Request.Data.(*types.CreateReplicationSubnetGroupOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationSubnetGroupResponse is the response type for the
// CreateReplicationSubnetGroup API operation.
type CreateReplicationSubnetGroupResponse struct {
	*types.CreateReplicationSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationSubnetGroup request.
func (r *CreateReplicationSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
