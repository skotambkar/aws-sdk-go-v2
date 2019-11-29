// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opDeleteDBInstance = "DeleteDBInstance"

// DeleteDBInstanceRequest returns a request value for making API operation for
// Amazon Neptune.
//
// The DeleteDBInstance action deletes a previously provisioned DB instance.
// When you delete a DB instance, all automated backups for that instance are
// deleted and can't be recovered. Manual DB snapshots of the DB instance to
// be deleted by DeleteDBInstance are not deleted.
//
// If you request a final DB snapshot the status of the Amazon Neptune DB instance
// is deleting until the DB snapshot is created. The API action DescribeDBInstance
// is used to monitor the status of this operation. The action can't be canceled
// or reverted once submitted.
//
// Note that when a DB instance is in a failure state and has a status of failed,
// incompatible-restore, or incompatible-network, you can only delete it when
// the SkipFinalSnapshot parameter is set to true.
//
// You can't delete a DB instance if it is the only instance in the DB cluster.
//
//    // Example sending a request using DeleteDBInstanceRequest.
//    req := client.DeleteDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DeleteDBInstance
func (c *Client) DeleteDBInstanceRequest(input *types.DeleteDBInstanceInput) DeleteDBInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDBInstanceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDBInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteDBInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteDBInstanceRequest{Request: req, Input: input, Copy: c.DeleteDBInstanceRequest}
}

// DeleteDBInstanceRequest is the request type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceRequest struct {
	*aws.Request
	Input *types.DeleteDBInstanceInput
	Copy  func(*types.DeleteDBInstanceInput) DeleteDBInstanceRequest
}

// Send marshals and sends the DeleteDBInstance API request.
func (r DeleteDBInstanceRequest) Send(ctx context.Context) (*DeleteDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBInstanceResponse{
		DeleteDBInstanceOutput: r.Request.Data.(*types.DeleteDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBInstanceResponse is the response type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceResponse struct {
	*types.DeleteDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBInstance request.
func (r *DeleteDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
