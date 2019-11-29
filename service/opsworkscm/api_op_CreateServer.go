// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
)

const opCreateServer = "CreateServer"

// CreateServerRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Creates and immedately starts a new server. The server is ready to use when
// it is in the HEALTHY state. By default, you can create a maximum of 10 servers.
//
// This operation is asynchronous.
//
// A LimitExceededException is thrown when you have created the maximum number
// of servers (10). A ResourceAlreadyExistsException is thrown when a server
// with the same name already exists in the account. A ResourceNotFoundException
// is thrown when you specify a backup ID that is not valid or is for a backup
// that does not exist. A ValidationException is thrown when parameters of the
// request are not valid.
//
// If you do not specify a security group by adding the SecurityGroupIds parameter,
// AWS OpsWorks creates a new security group.
//
// Chef Automate: The default security group opens the Chef server to the world
// on TCP port 443. If a KeyName is present, AWS OpsWorks enables SSH access.
// SSH is also open to the world on TCP port 22.
//
// Puppet Enterprise: The default security group opens TCP ports 22, 443, 4433,
// 8140, 8142, 8143, and 8170. If a KeyName is present, AWS OpsWorks enables
// SSH access. SSH is also open to the world on TCP port 22.
//
// By default, your server is accessible from any IP address. We recommend that
// you update your security group rules to allow access from known IP addresses
// and address ranges only. To edit security group rules, open Security Groups
// in the navigation pane of the EC2 management console.
//
// To specify your own domain for a server, and provide your own self-signed
// or CA-signed certificate and private key, specify values for CustomDomain,
// CustomCertificate, and CustomPrivateKey.
//
//    // Example sending a request using CreateServerRequest.
//    req := client.CreateServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/CreateServer
func (c *Client) CreateServerRequest(input *types.CreateServerInput) CreateServerRequest {
	op := &aws.Operation{
		Name:       opCreateServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateServerInput{}
	}

	req := c.newRequest(op, input, &types.CreateServerOutput{})
	return CreateServerRequest{Request: req, Input: input, Copy: c.CreateServerRequest}
}

// CreateServerRequest is the request type for the
// CreateServer API operation.
type CreateServerRequest struct {
	*aws.Request
	Input *types.CreateServerInput
	Copy  func(*types.CreateServerInput) CreateServerRequest
}

// Send marshals and sends the CreateServer API request.
func (r CreateServerRequest) Send(ctx context.Context) (*CreateServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServerResponse{
		CreateServerOutput: r.Request.Data.(*types.CreateServerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServerResponse is the response type for the
// CreateServer API operation.
type CreateServerResponse struct {
	*types.CreateServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServer request.
func (r *CreateServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
