// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDeleteUserByPrincipalId = "DeleteUserByPrincipalId"

// DeleteUserByPrincipalIdRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes a user identified by its principal ID.
//
// CLI Sample:
//
// aws quicksight delete-user-by-principal-id --aws-account-id=111122223333
// --namespace=default --principal-id=ABCDEFJA26JLI7EUUOEHS
//
//    // Example sending a request using DeleteUserByPrincipalIdRequest.
//    req := client.DeleteUserByPrincipalIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteUserByPrincipalId
func (c *Client) DeleteUserByPrincipalIdRequest(input *types.DeleteUserByPrincipalIdInput) DeleteUserByPrincipalIdRequest {
	op := &aws.Operation{
		Name:       opDeleteUserByPrincipalId,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/user-principals/{PrincipalId}",
	}

	if input == nil {
		input = &types.DeleteUserByPrincipalIdInput{}
	}

	req := c.newRequest(op, input, &types.DeleteUserByPrincipalIdOutput{})
	return DeleteUserByPrincipalIdRequest{Request: req, Input: input, Copy: c.DeleteUserByPrincipalIdRequest}
}

// DeleteUserByPrincipalIdRequest is the request type for the
// DeleteUserByPrincipalId API operation.
type DeleteUserByPrincipalIdRequest struct {
	*aws.Request
	Input *types.DeleteUserByPrincipalIdInput
	Copy  func(*types.DeleteUserByPrincipalIdInput) DeleteUserByPrincipalIdRequest
}

// Send marshals and sends the DeleteUserByPrincipalId API request.
func (r DeleteUserByPrincipalIdRequest) Send(ctx context.Context) (*DeleteUserByPrincipalIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserByPrincipalIdResponse{
		DeleteUserByPrincipalIdOutput: r.Request.Data.(*types.DeleteUserByPrincipalIdOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserByPrincipalIdResponse is the response type for the
// DeleteUserByPrincipalId API operation.
type DeleteUserByPrincipalIdResponse struct {
	*types.DeleteUserByPrincipalIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserByPrincipalId request.
func (r *DeleteUserByPrincipalIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
