// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type EncryptionStatus string

// Enum values for EncryptionStatus
const (
	EncryptionStatusUpdating EncryptionStatus = "UPDATING"
	EncryptionStatusActive   EncryptionStatus = "ACTIVE"
)

func (enum EncryptionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncryptionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeNone EncryptionType = "NONE"
	EncryptionTypeKms  EncryptionType = "KMS"
)

func (enum EncryptionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncryptionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SamplingStrategyName string

// Enum values for SamplingStrategyName
const (
	SamplingStrategyNamePartialScan SamplingStrategyName = "PartialScan"
	SamplingStrategyNameFixedRate   SamplingStrategyName = "FixedRate"
)

func (enum SamplingStrategyName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SamplingStrategyName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TimeRangeType string

// Enum values for TimeRangeType
const (
	TimeRangeTypeTraceId TimeRangeType = "TraceId"
	TimeRangeTypeEvent   TimeRangeType = "Event"
)

func (enum TimeRangeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TimeRangeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
