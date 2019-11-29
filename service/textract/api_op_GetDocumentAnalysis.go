// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package textract

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/textract/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
)

const opGetDocumentAnalysis = "GetDocumentAnalysis"

// GetDocumentAnalysisRequest returns a request value for making API operation for
// Amazon Textract.
//
// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a document.
//
// You start asynchronous text analysis by calling StartDocumentAnalysis, which
// returns a job identifier (JobId). When the text analysis operation finishes,
// Amazon Textract publishes a completion status to the Amazon Simple Notification
// Service (Amazon SNS) topic that's registered in the initial call to StartDocumentAnalysis.
// To get the results of the text-detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetDocumentAnalysis, and pass the job identifier (JobId) from the initial
// call to StartDocumentAnalysis.
//
// GetDocumentAnalysis returns an array of Block objects. The following types
// of information are returned:
//
//    * Words and lines that are related to nearby lines and words. The related
//    information is returned in two Block objects each of type KEY_VALUE_SET:
//    a KEY Block object and a VALUE Block object. For example, Name: Ana Silva
//    Carolina contains a key and value. Name: is the key. Ana Silva Carolina
//    is the value.
//
//    * Table and table cell data. A TABLE Block object contains information
//    about a detected table. A CELL Block object is returned for each cell
//    in a table.
//
//    * Selectable elements such as checkboxes and radio buttons. A SELECTION_ELEMENT
//    Block object contains information about a selectable element.
//
//    * Lines and words of text. A LINE Block object contains one or more WORD
//    Block objects.
//
// Use the MaxResults parameter to limit the number of blocks returned. If there
// are more results than specified in MaxResults, the value of NextToken in
// the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetDocumentAnalysis, and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetDocumentAnalysis.
//
// For more information, see Document Text Analysis (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html).
//
//    // Example sending a request using GetDocumentAnalysisRequest.
//    req := client.GetDocumentAnalysisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/textract-2018-06-27/GetDocumentAnalysis
func (c *Client) GetDocumentAnalysisRequest(input *types.GetDocumentAnalysisInput) GetDocumentAnalysisRequest {
	op := &aws.Operation{
		Name:       opGetDocumentAnalysis,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetDocumentAnalysisInput{}
	}

	req := c.newRequest(op, input, &types.GetDocumentAnalysisOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetDocumentAnalysisMarshaler{Input: input}.GetNamedBuildHandler())

	return GetDocumentAnalysisRequest{Request: req, Input: input, Copy: c.GetDocumentAnalysisRequest}
}

// GetDocumentAnalysisRequest is the request type for the
// GetDocumentAnalysis API operation.
type GetDocumentAnalysisRequest struct {
	*aws.Request
	Input *types.GetDocumentAnalysisInput
	Copy  func(*types.GetDocumentAnalysisInput) GetDocumentAnalysisRequest
}

// Send marshals and sends the GetDocumentAnalysis API request.
func (r GetDocumentAnalysisRequest) Send(ctx context.Context) (*GetDocumentAnalysisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentAnalysisResponse{
		GetDocumentAnalysisOutput: r.Request.Data.(*types.GetDocumentAnalysisOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentAnalysisResponse is the response type for the
// GetDocumentAnalysis API operation.
type GetDocumentAnalysisResponse struct {
	*types.GetDocumentAnalysisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocumentAnalysis request.
func (r *GetDocumentAnalysisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
