// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDeleteDatastore = "DeleteDatastore"

// DeleteDatastoreRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Deletes the specified data store.
//
//    // Example sending a request using DeleteDatastoreRequest.
//    req := client.DeleteDatastoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DeleteDatastore
func (c *Client) DeleteDatastoreRequest(input *types.DeleteDatastoreInput) DeleteDatastoreRequest {
	op := &aws.Operation{
		Name:       opDeleteDatastore,
		HTTPMethod: "DELETE",
		HTTPPath:   "/datastores/{datastoreName}",
	}

	if input == nil {
		input = &types.DeleteDatastoreInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDatastoreOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDatastoreRequest{Request: req, Input: input, Copy: c.DeleteDatastoreRequest}
}

// DeleteDatastoreRequest is the request type for the
// DeleteDatastore API operation.
type DeleteDatastoreRequest struct {
	*aws.Request
	Input *types.DeleteDatastoreInput
	Copy  func(*types.DeleteDatastoreInput) DeleteDatastoreRequest
}

// Send marshals and sends the DeleteDatastore API request.
func (r DeleteDatastoreRequest) Send(ctx context.Context) (*DeleteDatastoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatastoreResponse{
		DeleteDatastoreOutput: r.Request.Data.(*types.DeleteDatastoreOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatastoreResponse is the response type for the
// DeleteDatastore API operation.
type DeleteDatastoreResponse struct {
	*types.DeleteDatastoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDatastore request.
func (r *DeleteDatastoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
