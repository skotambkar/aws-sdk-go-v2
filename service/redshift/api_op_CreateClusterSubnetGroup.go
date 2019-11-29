// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opCreateClusterSubnetGroup = "CreateClusterSubnetGroup"

// CreateClusterSubnetGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new Amazon Redshift subnet group. You must provide a list of one
// or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC)
// when creating Amazon Redshift subnet group.
//
// For information about subnet groups, go to Amazon Redshift Cluster Subnet
// Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateClusterSubnetGroupRequest.
//    req := client.CreateClusterSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSubnetGroup
func (c *Client) CreateClusterSubnetGroupRequest(input *types.CreateClusterSubnetGroupInput) CreateClusterSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateClusterSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateClusterSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateClusterSubnetGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateClusterSubnetGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateClusterSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateClusterSubnetGroupRequest}
}

// CreateClusterSubnetGroupRequest is the request type for the
// CreateClusterSubnetGroup API operation.
type CreateClusterSubnetGroupRequest struct {
	*aws.Request
	Input *types.CreateClusterSubnetGroupInput
	Copy  func(*types.CreateClusterSubnetGroupInput) CreateClusterSubnetGroupRequest
}

// Send marshals and sends the CreateClusterSubnetGroup API request.
func (r CreateClusterSubnetGroupRequest) Send(ctx context.Context) (*CreateClusterSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterSubnetGroupResponse{
		CreateClusterSubnetGroupOutput: r.Request.Data.(*types.CreateClusterSubnetGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterSubnetGroupResponse is the response type for the
// CreateClusterSubnetGroup API operation.
type CreateClusterSubnetGroupResponse struct {
	*types.CreateClusterSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClusterSubnetGroup request.
func (r *CreateClusterSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
