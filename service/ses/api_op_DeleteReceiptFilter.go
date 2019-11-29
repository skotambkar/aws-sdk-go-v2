// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opDeleteReceiptFilter = "DeleteReceiptFilter"

// DeleteReceiptFilterRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes the specified IP address filter.
//
// For information about managing IP address filters, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-ip-filters.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteReceiptFilterRequest.
//    req := client.DeleteReceiptFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteReceiptFilter
func (c *Client) DeleteReceiptFilterRequest(input *types.DeleteReceiptFilterInput) DeleteReceiptFilterRequest {
	op := &aws.Operation{
		Name:       opDeleteReceiptFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteReceiptFilterInput{}
	}

	req := c.newRequest(op, input, &types.DeleteReceiptFilterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteReceiptFilterMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteReceiptFilterRequest{Request: req, Input: input, Copy: c.DeleteReceiptFilterRequest}
}

// DeleteReceiptFilterRequest is the request type for the
// DeleteReceiptFilter API operation.
type DeleteReceiptFilterRequest struct {
	*aws.Request
	Input *types.DeleteReceiptFilterInput
	Copy  func(*types.DeleteReceiptFilterInput) DeleteReceiptFilterRequest
}

// Send marshals and sends the DeleteReceiptFilter API request.
func (r DeleteReceiptFilterRequest) Send(ctx context.Context) (*DeleteReceiptFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReceiptFilterResponse{
		DeleteReceiptFilterOutput: r.Request.Data.(*types.DeleteReceiptFilterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReceiptFilterResponse is the response type for the
// DeleteReceiptFilter API operation.
type DeleteReceiptFilterResponse struct {
	*types.DeleteReceiptFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReceiptFilter request.
func (r *DeleteReceiptFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
