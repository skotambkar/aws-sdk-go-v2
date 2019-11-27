// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeFpgaImages = "DescribeFpgaImages"

// DescribeFpgaImagesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Amazon FPGA Images (AFIs) available to you. These include public
// AFIs, private AFIs that you own, and AFIs owned by other AWS accounts for
// which you have load permissions.
//
//    // Example sending a request using DescribeFpgaImagesRequest.
//    req := client.DescribeFpgaImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImages
func (c *Client) DescribeFpgaImagesRequest(input *types.DescribeFpgaImagesInput) DescribeFpgaImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeFpgaImages,
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
		input = &types.DescribeFpgaImagesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeFpgaImagesOutput{})
	return DescribeFpgaImagesRequest{Request: req, Input: input, Copy: c.DescribeFpgaImagesRequest}
}

// DescribeFpgaImagesRequest is the request type for the
// DescribeFpgaImages API operation.
type DescribeFpgaImagesRequest struct {
	*aws.Request
	Input *types.DescribeFpgaImagesInput
	Copy  func(*types.DescribeFpgaImagesInput) DescribeFpgaImagesRequest
}

// Send marshals and sends the DescribeFpgaImages API request.
func (r DescribeFpgaImagesRequest) Send(ctx context.Context) (*DescribeFpgaImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFpgaImagesResponse{
		DescribeFpgaImagesOutput: r.Request.Data.(*types.DescribeFpgaImagesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFpgaImagesRequestPaginator returns a paginator for DescribeFpgaImages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFpgaImagesRequest(input)
//   p := ec2.NewDescribeFpgaImagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFpgaImagesPaginator(req DescribeFpgaImagesRequest) DescribeFpgaImagesPaginator {
	return DescribeFpgaImagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeFpgaImagesInput
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

// DescribeFpgaImagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFpgaImagesPaginator struct {
	aws.Pager
}

func (p *DescribeFpgaImagesPaginator) CurrentPage() *types.DescribeFpgaImagesOutput {
	return p.Pager.CurrentPage().(*types.DescribeFpgaImagesOutput)
}

// DescribeFpgaImagesResponse is the response type for the
// DescribeFpgaImages API operation.
type DescribeFpgaImagesResponse struct {
	*types.DescribeFpgaImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFpgaImages request.
func (r *DescribeFpgaImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
