// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/enums"
)

type DescribeNodeAssociationStatusInput struct {
	_ struct{} `type:"structure"`

	// The token returned in either the AssociateNodeResponse or the DisassociateNodeResponse.
	//
	// NodeAssociationStatusToken is a required field
	NodeAssociationStatusToken *string `type:"string" required:"true"`

	// The name of the server from which to disassociate the node.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNodeAssociationStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNodeAssociationStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNodeAssociationStatusInput"}

	if s.NodeAssociationStatusToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeAssociationStatusToken"))
	}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeNodeAssociationStatusOutput struct {
	_ struct{} `type:"structure"`

	// Attributes specific to the node association. In Puppet, the attibute PUPPET_NODE_CERT
	// contains the signed certificate (the result of the CSR).
	EngineAttributes []EngineAttribute `type:"list"`

	// The status of the association or disassociation request.
	//
	// Possible values:
	//
	//    * SUCCESS: The association or disassociation succeeded.
	//
	//    * FAILED: The association or disassociation failed.
	//
	//    * IN_PROGRESS: The association or disassociation is still in progress.
	NodeAssociationStatus enums.NodeAssociationStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeNodeAssociationStatusOutput) String() string {
	return awsutil.Prettify(s)
}
