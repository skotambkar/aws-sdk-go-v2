// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

const opGetExclusionsPreview = "GetExclusionsPreview"

// GetExclusionsPreviewRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Retrieves the exclusions preview (a list of ExclusionPreview objects) specified
// by the preview token. You can obtain the preview token by running the CreateExclusionsPreview
// API.
//
//    // Example sending a request using GetExclusionsPreviewRequest.
//    req := client.GetExclusionsPreviewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/GetExclusionsPreview
func (c *Client) GetExclusionsPreviewRequest(input *types.GetExclusionsPreviewInput) GetExclusionsPreviewRequest {
	op := &aws.Operation{
		Name:       opGetExclusionsPreview,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetExclusionsPreviewInput{}
	}

	req := c.newRequest(op, input, &types.GetExclusionsPreviewOutput{})
	return GetExclusionsPreviewRequest{Request: req, Input: input, Copy: c.GetExclusionsPreviewRequest}
}

// GetExclusionsPreviewRequest is the request type for the
// GetExclusionsPreview API operation.
type GetExclusionsPreviewRequest struct {
	*aws.Request
	Input *types.GetExclusionsPreviewInput
	Copy  func(*types.GetExclusionsPreviewInput) GetExclusionsPreviewRequest
}

// Send marshals and sends the GetExclusionsPreview API request.
func (r GetExclusionsPreviewRequest) Send(ctx context.Context) (*GetExclusionsPreviewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExclusionsPreviewResponse{
		GetExclusionsPreviewOutput: r.Request.Data.(*types.GetExclusionsPreviewOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetExclusionsPreviewRequestPaginator returns a paginator for GetExclusionsPreview.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetExclusionsPreviewRequest(input)
//   p := inspector.NewGetExclusionsPreviewRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetExclusionsPreviewPaginator(req GetExclusionsPreviewRequest) GetExclusionsPreviewPaginator {
	return GetExclusionsPreviewPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetExclusionsPreviewInput
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

// GetExclusionsPreviewPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetExclusionsPreviewPaginator struct {
	aws.Pager
}

func (p *GetExclusionsPreviewPaginator) CurrentPage() *types.GetExclusionsPreviewOutput {
	return p.Pager.CurrentPage().(*types.GetExclusionsPreviewOutput)
}

// GetExclusionsPreviewResponse is the response type for the
// GetExclusionsPreview API operation.
type GetExclusionsPreviewResponse struct {
	*types.GetExclusionsPreviewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExclusionsPreview request.
func (r *GetExclusionsPreviewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
