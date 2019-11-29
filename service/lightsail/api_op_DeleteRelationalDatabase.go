// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opDeleteRelationalDatabase = "DeleteRelationalDatabase"

// DeleteRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a database in Amazon Lightsail.
//
// The delete relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteRelationalDatabaseRequest.
//    req := client.DeleteRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteRelationalDatabase
func (c *Client) DeleteRelationalDatabaseRequest(input *types.DeleteRelationalDatabaseInput) DeleteRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opDeleteRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRelationalDatabaseOutput{})
	return DeleteRelationalDatabaseRequest{Request: req, Input: input, Copy: c.DeleteRelationalDatabaseRequest}
}

// DeleteRelationalDatabaseRequest is the request type for the
// DeleteRelationalDatabase API operation.
type DeleteRelationalDatabaseRequest struct {
	*aws.Request
	Input *types.DeleteRelationalDatabaseInput
	Copy  func(*types.DeleteRelationalDatabaseInput) DeleteRelationalDatabaseRequest
}

// Send marshals and sends the DeleteRelationalDatabase API request.
func (r DeleteRelationalDatabaseRequest) Send(ctx context.Context) (*DeleteRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRelationalDatabaseResponse{
		DeleteRelationalDatabaseOutput: r.Request.Data.(*types.DeleteRelationalDatabaseOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRelationalDatabaseResponse is the response type for the
// DeleteRelationalDatabase API operation.
type DeleteRelationalDatabaseResponse struct {
	*types.DeleteRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRelationalDatabase request.
func (r *DeleteRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
