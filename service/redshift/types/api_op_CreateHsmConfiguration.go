// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHsmConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A text description of the HSM configuration to be created.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The identifier to be assigned to the new Amazon Redshift HSM configuration.
	//
	// HsmConfigurationIdentifier is a required field
	HsmConfigurationIdentifier *string `type:"string" required:"true"`

	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	//
	// HsmIpAddress is a required field
	HsmIpAddress *string `type:"string" required:"true"`

	// The name of the partition in the HSM where the Amazon Redshift clusters will
	// store their database encryption keys.
	//
	// HsmPartitionName is a required field
	HsmPartitionName *string `type:"string" required:"true"`

	// The password required to access the HSM partition.
	//
	// HsmPartitionPassword is a required field
	HsmPartitionPassword *string `type:"string" required:"true"`

	// The HSMs public certificate file. When using Cloud HSM, the file name is
	// server.pem.
	//
	// HsmServerPublicCertificate is a required field
	HsmServerPublicCertificate *string `type:"string" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateHsmConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHsmConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHsmConfigurationInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.HsmConfigurationIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmConfigurationIdentifier"))
	}

	if s.HsmIpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmIpAddress"))
	}

	if s.HsmPartitionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmPartitionName"))
	}

	if s.HsmPartitionPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmPartitionPassword"))
	}

	if s.HsmServerPublicCertificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmServerPublicCertificate"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHsmConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about an HSM configuration, which is an object that describes
	// to Amazon Redshift clusters the information they require to connect to an
	// HSM where they can store database encryption keys.
	HsmConfiguration *HsmConfiguration `type:"structure"`
}

// String returns the string representation
func (s CreateHsmConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
