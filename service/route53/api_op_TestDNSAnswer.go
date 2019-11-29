// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opTestDNSAnswer = "TestDNSAnswer"

// TestDNSAnswerRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the value that Amazon Route 53 returns in response to a DNS request
// for a specified record name and type. You can optionally specify the IP address
// of a DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
//
//    // Example sending a request using TestDNSAnswerRequest.
//    req := client.TestDNSAnswerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/TestDNSAnswer
func (c *Client) TestDNSAnswerRequest(input *types.TestDNSAnswerInput) TestDNSAnswerRequest {
	op := &aws.Operation{
		Name:       opTestDNSAnswer,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/testdnsanswer",
	}

	if input == nil {
		input = &types.TestDNSAnswerInput{}
	}

	req := c.newRequest(op, input, &types.TestDNSAnswerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.TestDNSAnswerMarshaler{Input: input}.GetNamedBuildHandler())

	return TestDNSAnswerRequest{Request: req, Input: input, Copy: c.TestDNSAnswerRequest}
}

// TestDNSAnswerRequest is the request type for the
// TestDNSAnswer API operation.
type TestDNSAnswerRequest struct {
	*aws.Request
	Input *types.TestDNSAnswerInput
	Copy  func(*types.TestDNSAnswerInput) TestDNSAnswerRequest
}

// Send marshals and sends the TestDNSAnswer API request.
func (r TestDNSAnswerRequest) Send(ctx context.Context) (*TestDNSAnswerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestDNSAnswerResponse{
		TestDNSAnswerOutput: r.Request.Data.(*types.TestDNSAnswerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestDNSAnswerResponse is the response type for the
// TestDNSAnswer API operation.
type TestDNSAnswerResponse struct {
	*types.TestDNSAnswerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestDNSAnswer request.
func (r *TestDNSAnswerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
