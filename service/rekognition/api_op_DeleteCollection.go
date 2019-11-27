// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opDeleteCollection = "DeleteCollection"

// DeleteCollectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Deletes the specified collection. Note that this operation removes all faces
// in the collection. For an example, see delete-collection-procedure.
//
// This operation requires permissions to perform the rekognition:DeleteCollection
// action.
//
//    // Example sending a request using DeleteCollectionRequest.
//    req := client.DeleteCollectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteCollectionRequest(input *types.DeleteCollectionInput) DeleteCollectionRequest {
	op := &aws.Operation{
		Name:       opDeleteCollection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCollectionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCollectionOutput{})
	return DeleteCollectionRequest{Request: req, Input: input, Copy: c.DeleteCollectionRequest}
}

// DeleteCollectionRequest is the request type for the
// DeleteCollection API operation.
type DeleteCollectionRequest struct {
	*aws.Request
	Input *types.DeleteCollectionInput
	Copy  func(*types.DeleteCollectionInput) DeleteCollectionRequest
}

// Send marshals and sends the DeleteCollection API request.
func (r DeleteCollectionRequest) Send(ctx context.Context) (*DeleteCollectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCollectionResponse{
		DeleteCollectionOutput: r.Request.Data.(*types.DeleteCollectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCollectionResponse is the response type for the
// DeleteCollection API operation.
type DeleteCollectionResponse struct {
	*types.DeleteCollectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCollection request.
func (r *DeleteCollectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
