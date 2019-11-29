// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type CompressionFormat string

// Enum values for CompressionFormat
const (
	CompressionFormatUncompressed CompressionFormat = "UNCOMPRESSED"
	CompressionFormatGzip         CompressionFormat = "GZIP"
	CompressionFormatZip          CompressionFormat = "ZIP"
	CompressionFormatSnappy       CompressionFormat = "Snappy"
)

func (enum CompressionFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CompressionFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStreamEncryptionStatus string

// Enum values for DeliveryStreamEncryptionStatus
const (
	DeliveryStreamEncryptionStatusEnabled         DeliveryStreamEncryptionStatus = "ENABLED"
	DeliveryStreamEncryptionStatusEnabling        DeliveryStreamEncryptionStatus = "ENABLING"
	DeliveryStreamEncryptionStatusEnablingFailed  DeliveryStreamEncryptionStatus = "ENABLING_FAILED"
	DeliveryStreamEncryptionStatusDisabled        DeliveryStreamEncryptionStatus = "DISABLED"
	DeliveryStreamEncryptionStatusDisabling       DeliveryStreamEncryptionStatus = "DISABLING"
	DeliveryStreamEncryptionStatusDisablingFailed DeliveryStreamEncryptionStatus = "DISABLING_FAILED"
)

func (enum DeliveryStreamEncryptionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStreamEncryptionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStreamFailureType string

// Enum values for DeliveryStreamFailureType
const (
	DeliveryStreamFailureTypeRetireKmsGrantFailed DeliveryStreamFailureType = "RETIRE_KMS_GRANT_FAILED"
	DeliveryStreamFailureTypeCreateKmsGrantFailed DeliveryStreamFailureType = "CREATE_KMS_GRANT_FAILED"
	DeliveryStreamFailureTypeKmsAccessDenied      DeliveryStreamFailureType = "KMS_ACCESS_DENIED"
	DeliveryStreamFailureTypeDisabledKmsKey       DeliveryStreamFailureType = "DISABLED_KMS_KEY"
	DeliveryStreamFailureTypeInvalidKmsKey        DeliveryStreamFailureType = "INVALID_KMS_KEY"
	DeliveryStreamFailureTypeKmsKeyNotFound       DeliveryStreamFailureType = "KMS_KEY_NOT_FOUND"
	DeliveryStreamFailureTypeKmsOptInRequired     DeliveryStreamFailureType = "KMS_OPT_IN_REQUIRED"
	DeliveryStreamFailureTypeUnknownError         DeliveryStreamFailureType = "UNKNOWN_ERROR"
)

func (enum DeliveryStreamFailureType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStreamFailureType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStreamStatus string

// Enum values for DeliveryStreamStatus
const (
	DeliveryStreamStatusCreating       DeliveryStreamStatus = "CREATING"
	DeliveryStreamStatusCreatingFailed DeliveryStreamStatus = "CREATING_FAILED"
	DeliveryStreamStatusDeleting       DeliveryStreamStatus = "DELETING"
	DeliveryStreamStatusDeletingFailed DeliveryStreamStatus = "DELETING_FAILED"
	DeliveryStreamStatusActive         DeliveryStreamStatus = "ACTIVE"
)

func (enum DeliveryStreamStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStreamStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStreamType string

// Enum values for DeliveryStreamType
const (
	DeliveryStreamTypeDirectPut             DeliveryStreamType = "DirectPut"
	DeliveryStreamTypeKinesisStreamAsSource DeliveryStreamType = "KinesisStreamAsSource"
)

func (enum DeliveryStreamType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStreamType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ElasticsearchIndexRotationPeriod string

// Enum values for ElasticsearchIndexRotationPeriod
const (
	ElasticsearchIndexRotationPeriodNoRotation ElasticsearchIndexRotationPeriod = "NoRotation"
	ElasticsearchIndexRotationPeriodOneHour    ElasticsearchIndexRotationPeriod = "OneHour"
	ElasticsearchIndexRotationPeriodOneDay     ElasticsearchIndexRotationPeriod = "OneDay"
	ElasticsearchIndexRotationPeriodOneWeek    ElasticsearchIndexRotationPeriod = "OneWeek"
	ElasticsearchIndexRotationPeriodOneMonth   ElasticsearchIndexRotationPeriod = "OneMonth"
)

func (enum ElasticsearchIndexRotationPeriod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ElasticsearchIndexRotationPeriod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ElasticsearchS3BackupMode string

// Enum values for ElasticsearchS3BackupMode
const (
	ElasticsearchS3BackupModeFailedDocumentsOnly ElasticsearchS3BackupMode = "FailedDocumentsOnly"
	ElasticsearchS3BackupModeAllDocuments        ElasticsearchS3BackupMode = "AllDocuments"
)

func (enum ElasticsearchS3BackupMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ElasticsearchS3BackupMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HECEndpointType string

// Enum values for HECEndpointType
const (
	HECEndpointTypeRaw   HECEndpointType = "Raw"
	HECEndpointTypeEvent HECEndpointType = "Event"
)

func (enum HECEndpointType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HECEndpointType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeAwsOwnedCmk        KeyType = "AWS_OWNED_CMK"
	KeyTypeCustomerManagedCmk KeyType = "CUSTOMER_MANAGED_CMK"
)

func (enum KeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NoEncryptionConfig string

// Enum values for NoEncryptionConfig
const (
	NoEncryptionConfigNoEncryption NoEncryptionConfig = "NoEncryption"
)

func (enum NoEncryptionConfig) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NoEncryptionConfig) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrcCompression string

// Enum values for OrcCompression
const (
	OrcCompressionNone   OrcCompression = "NONE"
	OrcCompressionZlib   OrcCompression = "ZLIB"
	OrcCompressionSnappy OrcCompression = "SNAPPY"
)

func (enum OrcCompression) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrcCompression) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrcFormatVersion string

// Enum values for OrcFormatVersion
const (
	OrcFormatVersionV011 OrcFormatVersion = "V0_11"
	OrcFormatVersionV012 OrcFormatVersion = "V0_12"
)

func (enum OrcFormatVersion) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrcFormatVersion) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ParquetCompression string

// Enum values for ParquetCompression
const (
	ParquetCompressionUncompressed ParquetCompression = "UNCOMPRESSED"
	ParquetCompressionGzip         ParquetCompression = "GZIP"
	ParquetCompressionSnappy       ParquetCompression = "SNAPPY"
)

func (enum ParquetCompression) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ParquetCompression) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ParquetWriterVersion string

// Enum values for ParquetWriterVersion
const (
	ParquetWriterVersionV1 ParquetWriterVersion = "V1"
	ParquetWriterVersionV2 ParquetWriterVersion = "V2"
)

func (enum ParquetWriterVersion) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ParquetWriterVersion) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProcessorParameterName string

// Enum values for ProcessorParameterName
const (
	ProcessorParameterNameLambdaArn               ProcessorParameterName = "LambdaArn"
	ProcessorParameterNameNumberOfRetries         ProcessorParameterName = "NumberOfRetries"
	ProcessorParameterNameRoleArn                 ProcessorParameterName = "RoleArn"
	ProcessorParameterNameBufferSizeInMbs         ProcessorParameterName = "BufferSizeInMBs"
	ProcessorParameterNameBufferIntervalInSeconds ProcessorParameterName = "BufferIntervalInSeconds"
)

func (enum ProcessorParameterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProcessorParameterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProcessorType string

// Enum values for ProcessorType
const (
	ProcessorTypeLambda ProcessorType = "Lambda"
)

func (enum ProcessorType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProcessorType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RedshiftS3BackupMode string

// Enum values for RedshiftS3BackupMode
const (
	RedshiftS3BackupModeDisabled RedshiftS3BackupMode = "Disabled"
	RedshiftS3BackupModeEnabled  RedshiftS3BackupMode = "Enabled"
)

func (enum RedshiftS3BackupMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RedshiftS3BackupMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type S3BackupMode string

// Enum values for S3BackupMode
const (
	S3BackupModeDisabled S3BackupMode = "Disabled"
	S3BackupModeEnabled  S3BackupMode = "Enabled"
)

func (enum S3BackupMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum S3BackupMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SplunkS3BackupMode string

// Enum values for SplunkS3BackupMode
const (
	SplunkS3BackupModeFailedEventsOnly SplunkS3BackupMode = "FailedEventsOnly"
	SplunkS3BackupModeAllEvents        SplunkS3BackupMode = "AllEvents"
)

func (enum SplunkS3BackupMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SplunkS3BackupMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
