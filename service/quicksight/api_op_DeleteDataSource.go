// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDeleteDataSource = "DeleteDataSource"

// DeleteDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes the data source permanently. This action breaks all the datasets
// that reference the deleted data source.
//
// CLI syntax:
//
// aws quicksight delete-data-source \
//
// --aws-account-id=111122223333 \
//
// --data-source-id=unique-data-source-id
//
//    // Example sending a request using DeleteDataSourceRequest.
//    req := client.DeleteDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteDataSource
func (c *Client) DeleteDataSourceRequest(input *types.DeleteDataSourceInput) DeleteDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources/{DataSourceId}",
	}

	if input == nil {
		input = &types.DeleteDataSourceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDataSourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteDataSourceMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteDataSourceRequest{Request: req, Input: input, Copy: c.DeleteDataSourceRequest}
}

// DeleteDataSourceRequest is the request type for the
// DeleteDataSource API operation.
type DeleteDataSourceRequest struct {
	*aws.Request
	Input *types.DeleteDataSourceInput
	Copy  func(*types.DeleteDataSourceInput) DeleteDataSourceRequest
}

// Send marshals and sends the DeleteDataSource API request.
func (r DeleteDataSourceRequest) Send(ctx context.Context) (*DeleteDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDataSourceResponse{
		DeleteDataSourceOutput: r.Request.Data.(*types.DeleteDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDataSourceResponse is the response type for the
// DeleteDataSource API operation.
type DeleteDataSourceResponse struct {
	*types.DeleteDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataSource request.
func (r *DeleteDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
