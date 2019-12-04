// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeIAMPolicyAssignmentInput struct {
	_ struct{} `type:"structure"`

	// The name of the assignment.
	//
	// AssignmentName is a required field
	AssignmentName *string `location:"uri" locationName:"AssignmentName" min:"1" type:"string" required:"true"`

	// The AWS account ID that contains the assignment you want to describe.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The namespace that contains the assignment.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIAMPolicyAssignmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIAMPolicyAssignmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIAMPolicyAssignmentInput"}

	if s.AssignmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentName"))
	}
	if s.AssignmentName != nil && len(*s.AssignmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentName", 1))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIAMPolicyAssignmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssignmentName != nil {
		v := *s.AssignmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AssignmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeIAMPolicyAssignmentOutput struct {
	_ struct{} `type:"structure"`

	// Information describing the IAM policy assignment.
	IAMPolicyAssignment *IAMPolicyAssignment `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DescribeIAMPolicyAssignmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIAMPolicyAssignmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IAMPolicyAssignment != nil {
		v := s.IAMPolicyAssignment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IAMPolicyAssignment", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeIAMPolicyAssignment = "DescribeIAMPolicyAssignment"

// DescribeIAMPolicyAssignmentRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes an existing IAMPolicy Assignment by specified assignment name.
//
// CLI syntax:
//
// aws quicksight describe-iam-policy-assignment --aws-account-id=111122223333
// --assignment-name=testtest --namespace=default --region=us-east-1
//
//    // Example sending a request using DescribeIAMPolicyAssignmentRequest.
//    req := client.DescribeIAMPolicyAssignmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeIAMPolicyAssignment
func (c *Client) DescribeIAMPolicyAssignmentRequest(input *DescribeIAMPolicyAssignmentInput) DescribeIAMPolicyAssignmentRequest {
	op := &aws.Operation{
		Name:       opDescribeIAMPolicyAssignment,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/iam-policy-assignments/{AssignmentName}",
	}

	if input == nil {
		input = &DescribeIAMPolicyAssignmentInput{}
	}

	req := c.newRequest(op, input, &DescribeIAMPolicyAssignmentOutput{})
	return DescribeIAMPolicyAssignmentRequest{Request: req, Input: input, Copy: c.DescribeIAMPolicyAssignmentRequest}
}

// DescribeIAMPolicyAssignmentRequest is the request type for the
// DescribeIAMPolicyAssignment API operation.
type DescribeIAMPolicyAssignmentRequest struct {
	*aws.Request
	Input *DescribeIAMPolicyAssignmentInput
	Copy  func(*DescribeIAMPolicyAssignmentInput) DescribeIAMPolicyAssignmentRequest
}

// Send marshals and sends the DescribeIAMPolicyAssignment API request.
func (r DescribeIAMPolicyAssignmentRequest) Send(ctx context.Context) (*DescribeIAMPolicyAssignmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIAMPolicyAssignmentResponse{
		DescribeIAMPolicyAssignmentOutput: r.Request.Data.(*DescribeIAMPolicyAssignmentOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIAMPolicyAssignmentResponse is the response type for the
// DescribeIAMPolicyAssignment API operation.
type DescribeIAMPolicyAssignmentResponse struct {
	*DescribeIAMPolicyAssignmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIAMPolicyAssignment request.
func (r *DescribeIAMPolicyAssignmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
