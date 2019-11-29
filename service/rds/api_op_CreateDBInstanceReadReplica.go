// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opCreateDBInstanceReadReplica = "CreateDBInstanceReadReplica"

// CreateDBInstanceReadReplicaRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB instance that acts as a Read Replica for an existing source
// DB instance. You can create a Read Replica for a DB instance running MySQL,
// MariaDB, Oracle, or PostgreSQL. For more information, see Working with Read
// Replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
// in the Amazon RDS User Guide.
//
// Amazon Aurora doesn't support this action. You must call the CreateDBInstance
// action to create a DB instance for an Aurora DB cluster.
//
// All Read Replica DB instances are created with backups disabled. All other
// DB instance attributes (including DB security groups and DB parameter groups)
// are inherited from the source DB instance, except as specified following.
//
// Your source DB instance must have backup retention enabled.
//
//    // Example sending a request using CreateDBInstanceReadReplicaRequest.
//    req := client.CreateDBInstanceReadReplicaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBInstanceReadReplica
func (c *Client) CreateDBInstanceReadReplicaRequest(input *types.CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest {
	op := &aws.Operation{
		Name:       opCreateDBInstanceReadReplica,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDBInstanceReadReplicaInput{}
	}

	req := c.newRequest(op, input, &types.CreateDBInstanceReadReplicaOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateDBInstanceReadReplicaMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDBInstanceReadReplicaRequest{Request: req, Input: input, Copy: c.CreateDBInstanceReadReplicaRequest}
}

// CreateDBInstanceReadReplicaRequest is the request type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaRequest struct {
	*aws.Request
	Input *types.CreateDBInstanceReadReplicaInput
	Copy  func(*types.CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest
}

// Send marshals and sends the CreateDBInstanceReadReplica API request.
func (r CreateDBInstanceReadReplicaRequest) Send(ctx context.Context) (*CreateDBInstanceReadReplicaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBInstanceReadReplicaResponse{
		CreateDBInstanceReadReplicaOutput: r.Request.Data.(*types.CreateDBInstanceReadReplicaOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBInstanceReadReplicaResponse is the response type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaResponse struct {
	*types.CreateDBInstanceReadReplicaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBInstanceReadReplica request.
func (r *CreateDBInstanceReadReplicaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
