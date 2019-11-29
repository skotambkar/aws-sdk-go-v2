// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opGetDashboardEmbedUrl = "GetDashboardEmbedUrl"

// GetDashboardEmbedUrlRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Generates a server-side embeddable URL and authorization code. Before this
// can work properly, first you need to configure the dashboards and user permissions.
// For more information, see Embedding Amazon QuickSight Dashboards (https://docs.aws.example.com/en_us/quicksight/latest/user/embedding.html).
//
// Currently, you can use GetDashboardEmbedURL only from the server, not from
// the user’s browser.
//
// CLI Sample:
//
// Assume the role with permissions enabled for actions: quickSight:RegisterUser
// and quicksight:GetDashboardEmbedURL. You can use assume-role, assume-role-with-web-identity,
// or assume-role-with-saml.
//
// aws sts assume-role --role-arn "arn:aws:iam::111122223333:role/embedding_quicksight_dashboard_role"
// --role-session-name embeddingsession
//
// If the user does not exist in QuickSight, register the user:
//
// aws quicksight register-user --aws-account-id 111122223333 --namespace default
// --identity-type IAM --iam-arn "arn:aws:iam::111122223333:role/embedding_quicksight_dashboard_role"
// --user-role READER --session-name "embeddingsession" --email user123@example.com
// --region us-east-1
//
// Get the URL for the embedded dashboard (IAM identity authentication):
//
// aws quicksight get-dashboard-embed-url --aws-account-id 111122223333 --dashboard-id
// 1a1ac2b2-3fc3-4b44-5e5d-c6db6778df89 --identity-type IAM
//
// Get the URL for the embedded dashboard (QUICKSIGHT identity authentication):
//
// aws quicksight get-dashboard-embed-url --aws-account-id 111122223333 --dashboard-id
// 1a1ac2b2-3fc3-4b44-5e5d-c6db6778df89 --identity-type QUICKSIGHT --user-arn
// arn:aws:quicksight:us-east-1:111122223333:user/default/embedding_quicksight_dashboard_role/embeddingsession
//
//    // Example sending a request using GetDashboardEmbedUrlRequest.
//    req := client.GetDashboardEmbedUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/GetDashboardEmbedUrl
func (c *Client) GetDashboardEmbedUrlRequest(input *types.GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest {
	op := &aws.Operation{
		Name:       opGetDashboardEmbedUrl,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/embed-url",
	}

	if input == nil {
		input = &types.GetDashboardEmbedUrlInput{}
	}

	req := c.newRequest(op, input, &types.GetDashboardEmbedUrlOutput{})
	return GetDashboardEmbedUrlRequest{Request: req, Input: input, Copy: c.GetDashboardEmbedUrlRequest}
}

// GetDashboardEmbedUrlRequest is the request type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlRequest struct {
	*aws.Request
	Input *types.GetDashboardEmbedUrlInput
	Copy  func(*types.GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest
}

// Send marshals and sends the GetDashboardEmbedUrl API request.
func (r GetDashboardEmbedUrlRequest) Send(ctx context.Context) (*GetDashboardEmbedUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDashboardEmbedUrlResponse{
		GetDashboardEmbedUrlOutput: r.Request.Data.(*types.GetDashboardEmbedUrlOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDashboardEmbedUrlResponse is the response type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlResponse struct {
	*types.GetDashboardEmbedUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDashboardEmbedUrl request.
func (r *GetDashboardEmbedUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
