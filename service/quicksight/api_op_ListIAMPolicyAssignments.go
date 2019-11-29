// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListIAMPolicyAssignments = "ListIAMPolicyAssignments"

// ListIAMPolicyAssignmentsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists assignments in current QuickSight account.
//
// CLI syntax:
//
// aws quicksight list-iam-policy-assignments --aws-account-id=111122223333
// --max-result=5 --assignment-status=ENABLED --namespace=default --region=us-east-1
// --next-token=3
//
//    // Example sending a request using ListIAMPolicyAssignmentsRequest.
//    req := client.ListIAMPolicyAssignmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListIAMPolicyAssignments
func (c *Client) ListIAMPolicyAssignmentsRequest(input *types.ListIAMPolicyAssignmentsInput) ListIAMPolicyAssignmentsRequest {
	op := &aws.Operation{
		Name:       opListIAMPolicyAssignments,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/iam-policy-assignments",
	}

	if input == nil {
		input = &types.ListIAMPolicyAssignmentsInput{}
	}

	req := c.newRequest(op, input, &types.ListIAMPolicyAssignmentsOutput{})
	return ListIAMPolicyAssignmentsRequest{Request: req, Input: input, Copy: c.ListIAMPolicyAssignmentsRequest}
}

// ListIAMPolicyAssignmentsRequest is the request type for the
// ListIAMPolicyAssignments API operation.
type ListIAMPolicyAssignmentsRequest struct {
	*aws.Request
	Input *types.ListIAMPolicyAssignmentsInput
	Copy  func(*types.ListIAMPolicyAssignmentsInput) ListIAMPolicyAssignmentsRequest
}

// Send marshals and sends the ListIAMPolicyAssignments API request.
func (r ListIAMPolicyAssignmentsRequest) Send(ctx context.Context) (*ListIAMPolicyAssignmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIAMPolicyAssignmentsResponse{
		ListIAMPolicyAssignmentsOutput: r.Request.Data.(*types.ListIAMPolicyAssignmentsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIAMPolicyAssignmentsResponse is the response type for the
// ListIAMPolicyAssignments API operation.
type ListIAMPolicyAssignmentsResponse struct {
	*types.ListIAMPolicyAssignmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIAMPolicyAssignments request.
func (r *ListIAMPolicyAssignmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
