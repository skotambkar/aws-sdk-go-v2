// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opCreateDBCluster = "CreateDBCluster"

// CreateDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new Amazon Aurora DB cluster.
//
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon RDS MySQL DB instance.
// For cross-region replication where the DB cluster identified by ReplicationSourceIdentifier
// is encrypted, you must also specify the PreSignedUrl parameter.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using CreateDBClusterRequest.
//    req := client.CreateDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBCluster
func (c *Client) CreateDBClusterRequest(input *types.CreateDBClusterInput) CreateDBClusterRequest {
	op := &aws.Operation{
		Name:       opCreateDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDBClusterInput{}
	}

	req := c.newRequest(op, input, &types.CreateDBClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateDBClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDBClusterRequest{Request: req, Input: input, Copy: c.CreateDBClusterRequest}
}

// CreateDBClusterRequest is the request type for the
// CreateDBCluster API operation.
type CreateDBClusterRequest struct {
	*aws.Request
	Input *types.CreateDBClusterInput
	Copy  func(*types.CreateDBClusterInput) CreateDBClusterRequest
}

// Send marshals and sends the CreateDBCluster API request.
func (r CreateDBClusterRequest) Send(ctx context.Context) (*CreateDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterResponse{
		CreateDBClusterOutput: r.Request.Data.(*types.CreateDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterResponse is the response type for the
// CreateDBCluster API operation.
type CreateDBClusterResponse struct {
	*types.CreateDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBCluster request.
func (r *CreateDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
