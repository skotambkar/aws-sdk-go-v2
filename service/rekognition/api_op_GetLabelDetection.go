// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opGetLabelDetection = "GetLabelDetection"

// GetLabelDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection.
//
// The label detection operation is started by a call to StartLabelDetection
// which returns a job identifier (JobId). When the label detection operation
// finishes, Amazon Rekognition publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to StartlabelDetection.
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLabelDetection and pass the job identifier (JobId) from the initial call
// to StartLabelDetection.
//
// GetLabelDetection returns an array of detected labels (Labels) sorted by
// the time the labels were detected. You can also sort by the label name by
// specifying NAME for the SortBy input parameter.
//
// The labels returned include the label name, the percentage confidence in
// the accuracy of the detected label, and the time the label was detected in
// the video.
//
// The returned labels also include bounding box information for common objects,
// a hierarchical taxonomy of detected labels, and the version of the label
// model used for detection.
//
// Use MaxResults parameter to limit the number of labels returned. If there
// are more results than specified in MaxResults, the value of NextToken in
// the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetlabelDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetLabelDetection.
//
//    // Example sending a request using GetLabelDetectionRequest.
//    req := client.GetLabelDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetLabelDetectionRequest(input *types.GetLabelDetectionInput) GetLabelDetectionRequest {
	op := &aws.Operation{
		Name:       opGetLabelDetection,
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
		input = &types.GetLabelDetectionInput{}
	}

	req := c.newRequest(op, input, &types.GetLabelDetectionOutput{})
	return GetLabelDetectionRequest{Request: req, Input: input, Copy: c.GetLabelDetectionRequest}
}

// GetLabelDetectionRequest is the request type for the
// GetLabelDetection API operation.
type GetLabelDetectionRequest struct {
	*aws.Request
	Input *types.GetLabelDetectionInput
	Copy  func(*types.GetLabelDetectionInput) GetLabelDetectionRequest
}

// Send marshals and sends the GetLabelDetection API request.
func (r GetLabelDetectionRequest) Send(ctx context.Context) (*GetLabelDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLabelDetectionResponse{
		GetLabelDetectionOutput: r.Request.Data.(*types.GetLabelDetectionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetLabelDetectionRequestPaginator returns a paginator for GetLabelDetection.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetLabelDetectionRequest(input)
//   p := rekognition.NewGetLabelDetectionRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetLabelDetectionPaginator(req GetLabelDetectionRequest) GetLabelDetectionPaginator {
	return GetLabelDetectionPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetLabelDetectionInput
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

// GetLabelDetectionPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetLabelDetectionPaginator struct {
	aws.Pager
}

func (p *GetLabelDetectionPaginator) CurrentPage() *types.GetLabelDetectionOutput {
	return p.Pager.CurrentPage().(*types.GetLabelDetectionOutput)
}

// GetLabelDetectionResponse is the response type for the
// GetLabelDetection API operation.
type GetLabelDetectionResponse struct {
	*types.GetLabelDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLabelDetection request.
func (r *GetLabelDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
