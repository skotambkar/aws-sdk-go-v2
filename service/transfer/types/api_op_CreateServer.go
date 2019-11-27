// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/transfer/enums"
)

type CreateServerInput struct {
	_ struct{} `type:"structure"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure
	// for your SFTP server. This parameter is required when you specify a value
	// for the EndpointType parameter.
	EndpointDetails *EndpointDetails `type:"structure"`

	// The type of VPC endpoint that you want your SFTP server to connect to. If
	// you connect to a VPC endpoint, your SFTP server isn't accessible over the
	// public internet.
	EndpointType enums.EndpointType `type:"string" enum:"true"`

	// The RSA private key as generated by the ssh-keygen -N "" -f my-new-server-key
	// command.
	//
	// If you aren't planning to migrate existing users from an existing SFTP server
	// to a new AWS SFTP server, don't update the host key. Accidentally changing
	// a server's host key can be disruptive.
	//
	// For more information, see "https://docs.aws.amazon.com/transfer/latest/userguide/change-host-key"
	// in the AWS SFTP User Guide.
	HostKey *string `type:"string" sensitive:"true"`

	// This parameter is required when the IdentityProviderType is set to API_GATEWAY.
	// Accepts an array containing all of the information required to call a customer-supplied
	// authentication API, including the API Gateway URL. This property is not required
	// when the IdentityProviderType is set to SERVICE_MANAGED.
	IdentityProviderDetails *IdentityProviderDetails `type:"structure"`

	// Specifies the mode of authentication for the SFTP server. The default value
	// is SERVICE_MANAGED, which allows you to store and access SFTP user credentials
	// within the AWS Transfer for SFTP service. Use the API_GATEWAY value to integrate
	// with an identity provider of your choosing. The API_GATEWAY setting requires
	// you to provide an API Gateway endpoint URL to call for authentication using
	// the IdentityProviderDetails parameter.
	IdentityProviderType enums.IdentityProviderType `type:"string" enum:"true"`

	// A value that allows the service to write your SFTP users' activity to your
	// Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole *string `type:"string"`

	// Key-value pairs that can be used to group and search for servers.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServerInput"}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateServerOutput struct {
	_ struct{} `type:"structure"`

	// The service-assigned ID of the SFTP server that is created.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServerOutput) String() string {
	return awsutil.Prettify(s)
}
