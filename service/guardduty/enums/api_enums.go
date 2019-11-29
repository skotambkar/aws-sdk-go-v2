// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type DestinationType string

// Enum values for DestinationType
const (
	DestinationTypeS3 DestinationType = "S3"
)

func (enum DestinationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DestinationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DetectorStatus string

// Enum values for DetectorStatus
const (
	DetectorStatusEnabled  DetectorStatus = "ENABLED"
	DetectorStatusDisabled DetectorStatus = "DISABLED"
)

func (enum DetectorStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DetectorStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Feedback string

// Enum values for Feedback
const (
	FeedbackUseful    Feedback = "USEFUL"
	FeedbackNotUseful Feedback = "NOT_USEFUL"
)

func (enum Feedback) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Feedback) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FilterAction string

// Enum values for FilterAction
const (
	FilterActionNoop    FilterAction = "NOOP"
	FilterActionArchive FilterAction = "ARCHIVE"
)

func (enum FilterAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FilterAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FindingPublishingFrequency string

// Enum values for FindingPublishingFrequency
const (
	FindingPublishingFrequencyFifteenMinutes FindingPublishingFrequency = "FIFTEEN_MINUTES"
	FindingPublishingFrequencyOneHour        FindingPublishingFrequency = "ONE_HOUR"
	FindingPublishingFrequencySixHours       FindingPublishingFrequency = "SIX_HOURS"
)

func (enum FindingPublishingFrequency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FindingPublishingFrequency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FindingStatisticType string

// Enum values for FindingStatisticType
const (
	FindingStatisticTypeCountBySeverity FindingStatisticType = "COUNT_BY_SEVERITY"
)

func (enum FindingStatisticType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FindingStatisticType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IpSetFormat string

// Enum values for IpSetFormat
const (
	IpSetFormatTxt        IpSetFormat = "TXT"
	IpSetFormatStix       IpSetFormat = "STIX"
	IpSetFormatOtxCsv     IpSetFormat = "OTX_CSV"
	IpSetFormatAlienVault IpSetFormat = "ALIEN_VAULT"
	IpSetFormatProofPoint IpSetFormat = "PROOF_POINT"
	IpSetFormatFireEye    IpSetFormat = "FIRE_EYE"
)

func (enum IpSetFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IpSetFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IpSetStatus string

// Enum values for IpSetStatus
const (
	IpSetStatusInactive      IpSetStatus = "INACTIVE"
	IpSetStatusActivating    IpSetStatus = "ACTIVATING"
	IpSetStatusActive        IpSetStatus = "ACTIVE"
	IpSetStatusDeactivating  IpSetStatus = "DEACTIVATING"
	IpSetStatusError         IpSetStatus = "ERROR"
	IpSetStatusDeletePending IpSetStatus = "DELETE_PENDING"
	IpSetStatusDeleted       IpSetStatus = "DELETED"
)

func (enum IpSetStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IpSetStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

func (enum OrderBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrderBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PublishingStatus string

// Enum values for PublishingStatus
const (
	PublishingStatusPendingVerification                   PublishingStatus = "PENDING_VERIFICATION"
	PublishingStatusPublishing                            PublishingStatus = "PUBLISHING"
	PublishingStatusUnableToPublishFixDestinationProperty PublishingStatus = "UNABLE_TO_PUBLISH_FIX_DESTINATION_PROPERTY"
	PublishingStatusStopped                               PublishingStatus = "STOPPED"
)

func (enum PublishingStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PublishingStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThreatIntelSetFormat string

// Enum values for ThreatIntelSetFormat
const (
	ThreatIntelSetFormatTxt        ThreatIntelSetFormat = "TXT"
	ThreatIntelSetFormatStix       ThreatIntelSetFormat = "STIX"
	ThreatIntelSetFormatOtxCsv     ThreatIntelSetFormat = "OTX_CSV"
	ThreatIntelSetFormatAlienVault ThreatIntelSetFormat = "ALIEN_VAULT"
	ThreatIntelSetFormatProofPoint ThreatIntelSetFormat = "PROOF_POINT"
	ThreatIntelSetFormatFireEye    ThreatIntelSetFormat = "FIRE_EYE"
)

func (enum ThreatIntelSetFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThreatIntelSetFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThreatIntelSetStatus string

// Enum values for ThreatIntelSetStatus
const (
	ThreatIntelSetStatusInactive      ThreatIntelSetStatus = "INACTIVE"
	ThreatIntelSetStatusActivating    ThreatIntelSetStatus = "ACTIVATING"
	ThreatIntelSetStatusActive        ThreatIntelSetStatus = "ACTIVE"
	ThreatIntelSetStatusDeactivating  ThreatIntelSetStatus = "DEACTIVATING"
	ThreatIntelSetStatusError         ThreatIntelSetStatus = "ERROR"
	ThreatIntelSetStatusDeletePending ThreatIntelSetStatus = "DELETE_PENDING"
	ThreatIntelSetStatusDeleted       ThreatIntelSetStatus = "DELETED"
)

func (enum ThreatIntelSetStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThreatIntelSetStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
