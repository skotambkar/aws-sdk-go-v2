// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/fsx/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
)

const opDescribeBackups = "DescribeBackups"

// DescribeBackupsRequest returns a request value for making API operation for
// Amazon FSx.
//
// Returns the description of specific Amazon FSx for Windows File Server backups,
// if a BackupIds value is provided for that backup. Otherwise, it returns all
// backups owned by your AWS account in the AWS Region of the endpoint that
// you're calling.
//
// When retrieving all backups, you can optionally specify the MaxResults parameter
// to limit the number of backups in a response. If more backups remain, Amazon
// FSx returns a NextToken value in the response. In this case, send a later
// request with the NextToken request parameter set to the value of NextToken
// from the last response.
//
// This action is used in an iterative process to retrieve a list of your backups.
// DescribeBackups is called first without a NextTokenvalue. Then the action
// continues to be called with the NextToken parameter set to the value of the
// last NextToken value until a response has no NextToken.
//
// When using this action, keep the following in mind:
//
//    * The implementation might return fewer than MaxResults file system descriptions
//    while still including a NextToken value.
//
//    * The order of backups returned in the response of one DescribeBackups
//    call and the order of backups returned across the responses of a multi-call
//    iteration is unspecified.
//
//    // Example sending a request using DescribeBackupsRequest.
//    req := client.DescribeBackupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/DescribeBackups
func (c *Client) DescribeBackupsRequest(input *types.DescribeBackupsInput) DescribeBackupsRequest {
	op := &aws.Operation{
		Name:       opDescribeBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeBackupsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeBackupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeBackupsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeBackupsRequest{Request: req, Input: input, Copy: c.DescribeBackupsRequest}
}

// DescribeBackupsRequest is the request type for the
// DescribeBackups API operation.
type DescribeBackupsRequest struct {
	*aws.Request
	Input *types.DescribeBackupsInput
	Copy  func(*types.DescribeBackupsInput) DescribeBackupsRequest
}

// Send marshals and sends the DescribeBackups API request.
func (r DescribeBackupsRequest) Send(ctx context.Context) (*DescribeBackupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBackupsResponse{
		DescribeBackupsOutput: r.Request.Data.(*types.DescribeBackupsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeBackupsRequestPaginator returns a paginator for DescribeBackups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeBackupsRequest(input)
//   p := fsx.NewDescribeBackupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeBackupsPaginator(req DescribeBackupsRequest) DescribeBackupsPaginator {
	return DescribeBackupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeBackupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeBackupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeBackupsPaginator struct {
	aws.Pager
}

func (p *DescribeBackupsPaginator) CurrentPage() *types.DescribeBackupsOutput {
	return p.Pager.CurrentPage().(*types.DescribeBackupsOutput)
}

// DescribeBackupsResponse is the response type for the
// DescribeBackups API operation.
type DescribeBackupsResponse struct {
	*types.DescribeBackupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBackups request.
func (r *DescribeBackupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
