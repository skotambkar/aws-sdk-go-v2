// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opUpdateDataSource = "UpdateDataSource"

// UpdateDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates a data source.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:datasource/data-source-id
//
// CLI syntax:
//
// aws quicksight update-data-source \
//
// --aws-account-id=111122223333 \
//
// --data-source-id=unique-data-source-id \
//
// --name='My Data Source' \
//
// --data-source-parameters='{"PostgreSqlParameters":{"Host":"my-db-host.example.com","Port":1234,"Database":"my-db"}}'
// \
//
// --credentials='{"CredentialPair":{"Username":"username","Password":"password"}}
//
//    // Example sending a request using UpdateDataSourceRequest.
//    req := client.UpdateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDataSource
func (c *Client) UpdateDataSourceRequest(input *types.UpdateDataSourceInput) UpdateDataSourceRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSource,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources/{DataSourceId}",
	}

	if input == nil {
		input = &types.UpdateDataSourceInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDataSourceOutput{})
	return UpdateDataSourceRequest{Request: req, Input: input, Copy: c.UpdateDataSourceRequest}
}

// UpdateDataSourceRequest is the request type for the
// UpdateDataSource API operation.
type UpdateDataSourceRequest struct {
	*aws.Request
	Input *types.UpdateDataSourceInput
	Copy  func(*types.UpdateDataSourceInput) UpdateDataSourceRequest
}

// Send marshals and sends the UpdateDataSource API request.
func (r UpdateDataSourceRequest) Send(ctx context.Context) (*UpdateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSourceResponse{
		UpdateDataSourceOutput: r.Request.Data.(*types.UpdateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSourceResponse is the response type for the
// UpdateDataSource API operation.
type UpdateDataSourceResponse struct {
	*types.UpdateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSource request.
func (r *UpdateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
