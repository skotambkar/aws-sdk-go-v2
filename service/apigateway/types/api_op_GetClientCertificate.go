// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to get information about the current ClientCertificate resource.
type GetClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the ClientCertificate resource to be described.
	//
	// ClientCertificateId is a required field
	ClientCertificateId *string `location:"uri" locationName:"clientcertificate_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetClientCertificateInput"}

	if s.ClientCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint.
//
// Client certificates are used to authenticate an API by the backend server.
// To authenticate an API client (or user), use IAM roles and policies, a custom
// Authorizer or an Amazon Cognito user pool.
//
// Use Client-Side Certificate (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GetClientCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the client certificate.
	ClientCertificateId *string `locationName:"clientCertificateId" type:"string"`

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// The description of the client certificate.
	Description *string `locationName:"description" type:"string"`

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time `locationName:"expirationDate" type:"timestamp"`

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string `locationName:"pemEncodedCertificate" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}