// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kafka/enums"
)

// Creates a cluster.
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// Information about the broker nodes in the cluster.
	//
	// BrokerNodeGroupInfo is a required field
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `locationName:"brokerNodeGroupInfo" type:"structure" required:"true"`

	// Includes all client authentication related information.
	ClientAuthentication *Authentication `locationName:"clientAuthentication" type:"structure"`

	// The name of the cluster.
	//
	// ClusterName is a required field
	ClusterName *string `locationName:"clusterName" min:"1" type:"string" required:"true"`

	// Represents the configuration that you want MSK to use for the cluster.
	ConfigurationInfo *ConfigurationInfo `locationName:"configurationInfo" type:"structure"`

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo `locationName:"encryptionInfo" type:"structure"`

	// Specifies the level of monitoring for the MSK cluster. The possible values
	// are DEFAULT, PER_BROKER, and PER_TOPIC_PER_BROKER.
	EnhancedMonitoring enums.EnhancedMonitoring `locationName:"enhancedMonitoring" type:"string" enum:"true"`

	// The version of Apache Kafka.
	//
	// KafkaVersion is a required field
	KafkaVersion *string `locationName:"kafkaVersion" min:"1" type:"string" required:"true"`

	// The number of Kafka broker nodes in the Amazon MSK cluster.
	//
	// NumberOfBrokerNodes is a required field
	NumberOfBrokerNodes *int64 `locationName:"numberOfBrokerNodes" min:"1" type:"integer" required:"true"`

	// Create tags when creating the cluster.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.BrokerNodeGroupInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerNodeGroupInfo"))
	}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.ClusterName != nil && len(*s.ClusterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterName", 1))
	}

	if s.KafkaVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("KafkaVersion"))
	}
	if s.KafkaVersion != nil && len(*s.KafkaVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KafkaVersion", 1))
	}

	if s.NumberOfBrokerNodes == nil {
		invalidParams.Add(aws.NewErrParamRequired("NumberOfBrokerNodes"))
	}
	if s.NumberOfBrokerNodes != nil && *s.NumberOfBrokerNodes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBrokerNodes", 1))
	}
	if s.BrokerNodeGroupInfo != nil {
		if err := s.BrokerNodeGroupInfo.Validate(); err != nil {
			invalidParams.AddNested("BrokerNodeGroupInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.ConfigurationInfo != nil {
		if err := s.ConfigurationInfo.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.EncryptionInfo != nil {
		if err := s.EncryptionInfo.Validate(); err != nil {
			invalidParams.AddNested("EncryptionInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns information about the created cluster.
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The name of the MSK cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The state of the cluster. The possible states are CREATING, ACTIVE, and FAILED.
	State enums.ClusterState `locationName:"state" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}