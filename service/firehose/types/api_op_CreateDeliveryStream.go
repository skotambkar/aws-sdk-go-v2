// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/firehose/enums"
)

type CreateDeliveryStreamInput struct {
	_ struct{} `type:"structure"`

	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *DeliveryStreamEncryptionConfigurationInput `type:"structure"`

	// The name of the delivery stream. This name must be unique per AWS account
	// in the same AWS Region. If the delivery streams are in different accounts
	// or different Regions, you can have multiple delivery streams with the same
	// name.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// The delivery stream type. This parameter can be one of the following values:
	//
	//    * DirectPut: Provider applications access the delivery stream directly.
	//
	//    * KinesisStreamAsSource: The delivery stream uses a Kinesis data stream
	//    as a source.
	DeliveryStreamType enums.DeliveryStreamType `type:"string" enum:"true"`

	// The destination in Amazon ES. You can specify only one destination.
	ElasticsearchDestinationConfiguration *ElasticsearchDestinationConfiguration `type:"structure"`

	// The destination in Amazon S3. You can specify only one destination.
	ExtendedS3DestinationConfiguration *ExtendedS3DestinationConfiguration `type:"structure"`

	// When a Kinesis data stream is used as the source for the delivery stream,
	// a KinesisStreamSourceConfiguration containing the Kinesis data stream Amazon
	// Resource Name (ARN) and the role ARN for the source stream.
	KinesisStreamSourceConfiguration *KinesisStreamSourceConfiguration `type:"structure"`

	// The destination in Amazon Redshift. You can specify only one destination.
	RedshiftDestinationConfiguration *RedshiftDestinationConfiguration `type:"structure"`

	// [Deprecated] The destination in Amazon S3. You can specify only one destination.
	S3DestinationConfiguration *S3DestinationConfiguration `deprecated:"true" type:"structure"`

	// The destination in Splunk. You can specify only one destination.
	SplunkDestinationConfiguration *SplunkDestinationConfiguration `type:"structure"`

	// A set of tags to assign to the delivery stream. A tag is a key-value pair
	// that you can define and assign to AWS resources. Tags are metadata. For example,
	// you can add friendly names and descriptions or other types of information
	// that can help you distinguish the delivery stream. For more information about
	// tags, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateDeliveryStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeliveryStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeliveryStreamInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.DeliveryStreamEncryptionConfigurationInput != nil {
		if err := s.DeliveryStreamEncryptionConfigurationInput.Validate(); err != nil {
			invalidParams.AddNested("DeliveryStreamEncryptionConfigurationInput", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchDestinationConfiguration != nil {
		if err := s.ElasticsearchDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExtendedS3DestinationConfiguration != nil {
		if err := s.ExtendedS3DestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ExtendedS3DestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.KinesisStreamSourceConfiguration != nil {
		if err := s.KinesisStreamSourceConfiguration.Validate(); err != nil {
			invalidParams.AddNested("KinesisStreamSourceConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.RedshiftDestinationConfiguration != nil {
		if err := s.RedshiftDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RedshiftDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.S3DestinationConfiguration != nil {
		if err := s.S3DestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("S3DestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.SplunkDestinationConfiguration != nil {
		if err := s.SplunkDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SplunkDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
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

type CreateDeliveryStreamOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the delivery stream.
	DeliveryStreamARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDeliveryStreamOutput) String() string {
	return awsutil.Prettify(s)
}