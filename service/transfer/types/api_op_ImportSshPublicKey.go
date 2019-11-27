// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ImportSshPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// The public key portion of an SSH key pair.
	//
	// SshPublicKeyBody is a required field
	SshPublicKeyBody *string `type:"string" required:"true"`

	// The name of the user account that is assigned to one or more servers.
	//
	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ImportSshPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportSshPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportSshPublicKeyInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}

	if s.SshPublicKeyBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("SshPublicKeyBody"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This response identifies the user, the server they belong to, and the identifier
// of the SSH public key associated with that user. A user can have more than
// one key on each server that they are associated with.
type ImportSshPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// This identifier is the name given to a public key by the system that was
	// imported.
	//
	// SshPublicKeyId is a required field
	SshPublicKeyId *string `type:"string" required:"true"`

	// A user name assigned to the ServerID value that you specified.
	//
	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ImportSshPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}
