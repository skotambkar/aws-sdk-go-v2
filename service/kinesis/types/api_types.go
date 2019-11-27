// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/enums"
)

// An object that represents the details of the consumer you registered.
type Consumer struct {
	_ struct{} `type:"structure"`

	// When you register a consumer, Kinesis Data Streams generates an ARN for it.
	// You need this ARN to be able to call SubscribeToShard.
	//
	// If you delete a consumer and then create a new one with the same name, it
	// won't have the same ARN. That's because consumer ARNs contain the creation
	// timestamp. This is important to keep in mind if you have IAM policies that
	// reference consumer ARNs.
	//
	// ConsumerARN is a required field
	ConsumerARN *string `min:"1" type:"string" required:"true"`

	// ConsumerCreationTimestamp is a required field
	ConsumerCreationTimestamp *time.Time `type:"timestamp" required:"true"`

	// The name of the consumer is something you choose when you register the consumer.
	//
	// ConsumerName is a required field
	ConsumerName *string `min:"1" type:"string" required:"true"`

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// ConsumerStatus is a required field
	ConsumerStatus enums.ConsumerStatus `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Consumer) String() string {
	return awsutil.Prettify(s)
}

// An object that represents the details of a registered consumer.
type ConsumerDescription struct {
	_ struct{} `type:"structure"`

	// When you register a consumer, Kinesis Data Streams generates an ARN for it.
	// You need this ARN to be able to call SubscribeToShard.
	//
	// If you delete a consumer and then create a new one with the same name, it
	// won't have the same ARN. That's because consumer ARNs contain the creation
	// timestamp. This is important to keep in mind if you have IAM policies that
	// reference consumer ARNs.
	//
	// ConsumerARN is a required field
	ConsumerARN *string `min:"1" type:"string" required:"true"`

	// ConsumerCreationTimestamp is a required field
	ConsumerCreationTimestamp *time.Time `type:"timestamp" required:"true"`

	// The name of the consumer is something you choose when you register the consumer.
	//
	// ConsumerName is a required field
	ConsumerName *string `min:"1" type:"string" required:"true"`

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// ConsumerStatus is a required field
	ConsumerStatus enums.ConsumerStatus `type:"string" required:"true" enum:"true"`

	// The ARN of the stream with which you registered the consumer.
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ConsumerDescription) String() string {
	return awsutil.Prettify(s)
}

// Represents enhanced metrics types.
type EnhancedMetrics struct {
	_ struct{} `type:"structure"`

	// List of shard-level metrics.
	//
	// The following are the valid shard-level metrics. The value "ALL" enhances
	// every metric.
	//
	//    * IncomingBytes
	//
	//    * IncomingRecords
	//
	//    * OutgoingBytes
	//
	//    * OutgoingRecords
	//
	//    * WriteProvisionedThroughputExceeded
	//
	//    * ReadProvisionedThroughputExceeded
	//
	//    * IteratorAgeMilliseconds
	//
	//    * ALL
	//
	// For more information, see Monitoring the Amazon Kinesis Data Streams Service
	// with Amazon CloudWatch (http://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	ShardLevelMetrics []enums.MetricsName `min:"1" type:"list"`
}

// String returns the string representation
func (s EnhancedMetrics) String() string {
	return awsutil.Prettify(s)
}

// The range of possible hash key values for the shard, which is a set of ordered
// contiguous positive integers.
type HashKeyRange struct {
	_ struct{} `type:"structure"`

	// The ending hash key of the hash key range.
	//
	// EndingHashKey is a required field
	EndingHashKey *string `type:"string" required:"true"`

	// The starting hash key of the hash key range.
	//
	// StartingHashKey is a required field
	StartingHashKey *string `type:"string" required:"true"`
}

// String returns the string representation
func (s HashKeyRange) String() string {
	return awsutil.Prettify(s)
}

// Represents the output for PutRecords.
type PutRecordsRequestEntry struct {
	_ struct{} `type:"structure"`

	// The data blob to put into the record, which is base64-encoded when the blob
	// is serialized. When the data blob (the payload before base64-encoding) is
	// added to the partition key size, the total size must not exceed the maximum
	// record size (1 MB).
	//
	// Data is automatically base64 encoded/decoded by the SDK.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`

	// The hash value used to determine explicitly the shard that the data record
	// is assigned to by overriding the partition key hash.
	ExplicitHashKey *string `type:"string"`

	// Determines which shard in the stream the data record is assigned to. Partition
	// keys are Unicode strings with a maximum length limit of 256 characters for
	// each key. Amazon Kinesis Data Streams uses the partition key as input to
	// a hash function that maps the partition key and associated data to a specific
	// shard. Specifically, an MD5 hash function is used to map partition keys to
	// 128-bit integer values and to map associated data records to shards. As a
	// result of this hashing mechanism, all data records with the same partition
	// key map to the same shard within the stream.
	//
	// PartitionKey is a required field
	PartitionKey *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecordsRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecordsRequestEntry) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecordsRequestEntry"}

	if s.Data == nil {
		invalidParams.Add(aws.NewErrParamRequired("Data"))
	}

	if s.PartitionKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionKey"))
	}
	if s.PartitionKey != nil && len(*s.PartitionKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PartitionKey", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of an individual record from a PutRecords request.
// A record that is successfully added to a stream includes SequenceNumber and
// ShardId in the result. A record that fails to be added to the stream includes
// ErrorCode and ErrorMessage in the result.
type PutRecordsResultEntry struct {
	_ struct{} `type:"structure"`

	// The error code for an individual record result. ErrorCodes can be either
	// ProvisionedThroughputExceededException or InternalFailure.
	ErrorCode *string `type:"string"`

	// The error message for an individual record result. An ErrorCode value of
	// ProvisionedThroughputExceededException has an error message that includes
	// the account ID, stream name, and shard ID. An ErrorCode value of InternalFailure
	// has the error message "Internal Service Failure".
	ErrorMessage *string `type:"string"`

	// The sequence number for an individual record result.
	SequenceNumber *string `type:"string"`

	// The shard ID for an individual record result.
	ShardId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutRecordsResultEntry) String() string {
	return awsutil.Prettify(s)
}

// The unit of data of the Kinesis data stream, which is composed of a sequence
// number, a partition key, and a data blob.
type Record struct {
	_ struct{} `type:"structure"`

	// The approximate time that the record was inserted into the stream.
	ApproximateArrivalTimestamp *time.Time `type:"timestamp"`

	// The data blob. The data in the blob is both opaque and immutable to Kinesis
	// Data Streams, which does not inspect, interpret, or change the data in the
	// blob in any way. When the data blob (the payload before base64-encoding)
	// is added to the partition key size, the total size must not exceed the maximum
	// record size (1 MB).
	//
	// Data is automatically base64 encoded/decoded by the SDK.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`

	// The encryption type used on the record. This parameter can be one of the
	// following values:
	//
	//    * NONE: Do not encrypt the records in the stream.
	//
	//    * KMS: Use server-side encryption on the records in the stream using a
	//    customer-managed AWS KMS key.
	EncryptionType enums.EncryptionType `type:"string" enum:"true"`

	// Identifies which shard in the stream the data record is assigned to.
	//
	// PartitionKey is a required field
	PartitionKey *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the record within its shard.
	//
	// SequenceNumber is a required field
	SequenceNumber *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Record) String() string {
	return awsutil.Prettify(s)
}

// The range of possible sequence numbers for the shard.
type SequenceNumberRange struct {
	_ struct{} `type:"structure"`

	// The ending sequence number for the range. Shards that are in the OPEN state
	// have an ending sequence number of null.
	EndingSequenceNumber *string `type:"string"`

	// The starting sequence number for the range.
	//
	// StartingSequenceNumber is a required field
	StartingSequenceNumber *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SequenceNumberRange) String() string {
	return awsutil.Prettify(s)
}

// A uniquely identified group of data records in a Kinesis data stream.
type Shard struct {
	_ struct{} `type:"structure"`

	// The shard ID of the shard adjacent to the shard's parent.
	AdjacentParentShardId *string `min:"1" type:"string"`

	// The range of possible hash key values for the shard, which is a set of ordered
	// contiguous positive integers.
	//
	// HashKeyRange is a required field
	HashKeyRange *HashKeyRange `type:"structure" required:"true"`

	// The shard ID of the shard's parent.
	ParentShardId *string `min:"1" type:"string"`

	// The range of possible sequence numbers for the shard.
	//
	// SequenceNumberRange is a required field
	SequenceNumberRange *SequenceNumberRange `type:"structure" required:"true"`

	// The unique identifier of the shard within the stream.
	//
	// ShardId is a required field
	ShardId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Shard) String() string {
	return awsutil.Prettify(s)
}

// Represents the output for DescribeStream.
type StreamDescription struct {
	_ struct{} `type:"structure"`

	// The server-side encryption type used on the stream. This parameter can be
	// one of the following values:
	//
	//    * NONE: Do not encrypt the records in the stream.
	//
	//    * KMS: Use server-side encryption on the records in the stream using a
	//    customer-managed AWS KMS key.
	EncryptionType enums.EncryptionType `type:"string" enum:"true"`

	// Represents the current enhanced monitoring settings of the stream.
	//
	// EnhancedMonitoring is a required field
	EnhancedMonitoring []EnhancedMetrics `type:"list" required:"true"`

	// If set to true, more shards in the stream are available to describe.
	//
	// HasMoreShards is a required field
	HasMoreShards *bool `type:"boolean" required:"true"`

	// The GUID for the customer-managed AWS KMS key to use for encryption. This
	// value can be a globally unique identifier, a fully specified ARN to either
	// an alias or a key, or an alias name prefixed by "alias/".You can also use
	// a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
	//
	//    * Key ARN example: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//    * Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//    * Alias name example: alias/MyAliasName
	//
	//    * Master key owned by Kinesis Data Streams: alias/aws/kinesis
	KeyId *string `min:"1" type:"string"`

	// The current retention period, in hours.
	//
	// RetentionPeriodHours is a required field
	RetentionPeriodHours *int64 `min:"1" type:"integer" required:"true"`

	// The shards that comprise the stream.
	//
	// Shards is a required field
	Shards []Shard `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`

	// The approximate time that the stream was created.
	//
	// StreamCreationTimestamp is a required field
	StreamCreationTimestamp *time.Time `type:"timestamp" required:"true"`

	// The name of the stream being described.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// The current status of the stream being described. The stream status is one
	// of the following states:
	//
	//    * CREATING - The stream is being created. Kinesis Data Streams immediately
	//    returns and sets StreamStatus to CREATING.
	//
	//    * DELETING - The stream is being deleted. The specified stream is in the
	//    DELETING state until Kinesis Data Streams completes the deletion.
	//
	//    * ACTIVE - The stream exists and is ready for read and write operations
	//    or deletion. You should perform read and write operations only on an ACTIVE
	//    stream.
	//
	//    * UPDATING - Shards in the stream are being merged or split. Read and
	//    write operations continue to work while the stream is in the UPDATING
	//    state.
	//
	// StreamStatus is a required field
	StreamStatus enums.StreamStatus `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StreamDescription) String() string {
	return awsutil.Prettify(s)
}

// Represents the output for DescribeStreamSummary
type StreamDescriptionSummary struct {
	_ struct{} `type:"structure"`

	// The number of enhanced fan-out consumers registered with the stream.
	ConsumerCount *int64 `type:"integer"`

	// The encryption type used. This value is one of the following:
	//
	//    * KMS
	//
	//    * NONE
	EncryptionType enums.EncryptionType `type:"string" enum:"true"`

	// Represents the current enhanced monitoring settings of the stream.
	//
	// EnhancedMonitoring is a required field
	EnhancedMonitoring []EnhancedMetrics `type:"list" required:"true"`

	// The GUID for the customer-managed AWS KMS key to use for encryption. This
	// value can be a globally unique identifier, a fully specified ARN to either
	// an alias or a key, or an alias name prefixed by "alias/".You can also use
	// a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
	//
	//    * Key ARN example: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//    * Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//    * Alias name example: alias/MyAliasName
	//
	//    * Master key owned by Kinesis Data Streams: alias/aws/kinesis
	KeyId *string `min:"1" type:"string"`

	// The number of open shards in the stream.
	//
	// OpenShardCount is a required field
	OpenShardCount *int64 `type:"integer" required:"true"`

	// The current retention period, in hours.
	//
	// RetentionPeriodHours is a required field
	RetentionPeriodHours *int64 `min:"1" type:"integer" required:"true"`

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`

	// The approximate time that the stream was created.
	//
	// StreamCreationTimestamp is a required field
	StreamCreationTimestamp *time.Time `type:"timestamp" required:"true"`

	// The name of the stream being described.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// The current status of the stream being described. The stream status is one
	// of the following states:
	//
	//    * CREATING - The stream is being created. Kinesis Data Streams immediately
	//    returns and sets StreamStatus to CREATING.
	//
	//    * DELETING - The stream is being deleted. The specified stream is in the
	//    DELETING state until Kinesis Data Streams completes the deletion.
	//
	//    * ACTIVE - The stream exists and is ready for read and write operations
	//    or deletion. You should perform read and write operations only on an ACTIVE
	//    stream.
	//
	//    * UPDATING - Shards in the stream are being merged or split. Read and
	//    write operations continue to work while the stream is in the UPDATING
	//    state.
	//
	// StreamStatus is a required field
	StreamStatus enums.StreamStatus `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StreamDescriptionSummary) String() string {
	return awsutil.Prettify(s)
}

// Metadata assigned to the stream, consisting of a key-value pair.
type Tag struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the tag. Maximum length: 128 characters. Valid characters:
	// Unicode letters, digits, white space, _ . / = + - % @
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// An optional string, typically used to describe or define the tag. Maximum
	// length: 256 characters. Valid characters: Unicode letters, digits, white
	// space, _ . / = + - % @
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}
