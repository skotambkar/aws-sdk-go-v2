// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAccountModificationsInput struct {
	_ struct{} `type:"structure"`

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAccountModificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAccountModificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAccountModificationsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAccountModificationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of modifications to the configuration of BYOL.
	AccountModifications []AccountModification `type:"list"`

	// The token to use to retrieve the next set of results, or null if no more
	// results are available.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAccountModificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountModifications = "DescribeAccountModifications"

// DescribeAccountModificationsRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Retrieves a list that describes modifications to the configuration of Bring
// Your Own License (BYOL) for the specified account.
//
//    // Example sending a request using DescribeAccountModificationsRequest.
//    req := client.DescribeAccountModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeAccountModifications
func (c *Client) DescribeAccountModificationsRequest(input *DescribeAccountModificationsInput) DescribeAccountModificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountModifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountModificationsInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountModificationsOutput{})
	return DescribeAccountModificationsRequest{Request: req, Input: input, Copy: c.DescribeAccountModificationsRequest}
}

// DescribeAccountModificationsRequest is the request type for the
// DescribeAccountModifications API operation.
type DescribeAccountModificationsRequest struct {
	*aws.Request
	Input *DescribeAccountModificationsInput
	Copy  func(*DescribeAccountModificationsInput) DescribeAccountModificationsRequest
}

// Send marshals and sends the DescribeAccountModifications API request.
func (r DescribeAccountModificationsRequest) Send(ctx context.Context) (*DescribeAccountModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountModificationsResponse{
		DescribeAccountModificationsOutput: r.Request.Data.(*DescribeAccountModificationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountModificationsResponse is the response type for the
// DescribeAccountModifications API operation.
type DescribeAccountModificationsResponse struct {
	*DescribeAccountModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountModifications request.
func (r *DescribeAccountModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
