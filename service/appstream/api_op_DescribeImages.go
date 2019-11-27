// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDescribeImages = "DescribeImages"

// DescribeImagesRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes one or more specified images, if the image
// names or image ARNs are provided. Otherwise, all images in the account are
// described.
//
//    // Example sending a request using DescribeImagesRequest.
//    req := client.DescribeImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeImages
func (c *Client) DescribeImagesRequest(input *types.DescribeImagesInput) DescribeImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeImages,
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
		input = &types.DescribeImagesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeImagesOutput{})
	return DescribeImagesRequest{Request: req, Input: input, Copy: c.DescribeImagesRequest}
}

// DescribeImagesRequest is the request type for the
// DescribeImages API operation.
type DescribeImagesRequest struct {
	*aws.Request
	Input *types.DescribeImagesInput
	Copy  func(*types.DescribeImagesInput) DescribeImagesRequest
}

// Send marshals and sends the DescribeImages API request.
func (r DescribeImagesRequest) Send(ctx context.Context) (*DescribeImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImagesResponse{
		DescribeImagesOutput: r.Request.Data.(*types.DescribeImagesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeImagesRequestPaginator returns a paginator for DescribeImages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeImagesRequest(input)
//   p := appstream.NewDescribeImagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeImagesPaginator(req DescribeImagesRequest) DescribeImagesPaginator {
	return DescribeImagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeImagesInput
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

// DescribeImagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeImagesPaginator struct {
	aws.Pager
}

func (p *DescribeImagesPaginator) CurrentPage() *types.DescribeImagesOutput {
	return p.Pager.CurrentPage().(*types.DescribeImagesOutput)
}

// DescribeImagesResponse is the response type for the
// DescribeImages API operation.
type DescribeImagesResponse struct {
	*types.DescribeImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImages request.
func (r *DescribeImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
