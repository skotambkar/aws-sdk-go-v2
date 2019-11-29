// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const opDeleteTable = "DeleteTable"

// DeleteTableRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// The DeleteTable operation deletes a table and all of its items. After a DeleteTable
// request, the specified table is in the DELETING state until DynamoDB completes
// the deletion. If the table is in the ACTIVE state, you can delete it. If
// a table is in CREATING or UPDATING states, then DynamoDB returns a ResourceInUseException.
// If the specified table does not exist, DynamoDB returns a ResourceNotFoundException.
// If table is already in the DELETING state, no error is returned.
//
// DynamoDB might continue to accept data read and write operations, such as
// GetItem and PutItem, on a table in the DELETING state until the table deletion
// is complete.
//
// When you delete a table, any indexes on that table are also deleted.
//
// If you have DynamoDB Streams enabled on the table, then the corresponding
// stream on that table goes into the DISABLED state, and the stream is automatically
// deleted after 24 hours.
//
// Use the DescribeTable action to check the status of the table.
//
//    // Example sending a request using DeleteTableRequest.
//    req := client.DeleteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteTable
func (c *Client) DeleteTableRequest(input *types.DeleteTableInput) DeleteTableRequest {
	op := &aws.Operation{
		Name:       opDeleteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTableInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTableOutput{})
	return DeleteTableRequest{Request: req, Input: input, Copy: c.DeleteTableRequest}
}

// DeleteTableRequest is the request type for the
// DeleteTable API operation.
type DeleteTableRequest struct {
	*aws.Request
	Input *types.DeleteTableInput
	Copy  func(*types.DeleteTableInput) DeleteTableRequest
}

// Send marshals and sends the DeleteTable API request.
func (r DeleteTableRequest) Send(ctx context.Context) (*DeleteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTableResponse{
		DeleteTableOutput: r.Request.Data.(*types.DeleteTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTableResponse is the response type for the
// DeleteTable API operation.
type DeleteTableResponse struct {
	*types.DeleteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTable request.
func (r *DeleteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
