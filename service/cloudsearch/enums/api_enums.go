// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AlgorithmicStemming string

// Enum values for AlgorithmicStemming
const (
	AlgorithmicStemmingNone    AlgorithmicStemming = "none"
	AlgorithmicStemmingMinimal AlgorithmicStemming = "minimal"
	AlgorithmicStemmingLight   AlgorithmicStemming = "light"
	AlgorithmicStemmingFull    AlgorithmicStemming = "full"
)

func (enum AlgorithmicStemming) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AlgorithmicStemming) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// An IETF RFC 4646 (http://tools.ietf.org/html/rfc4646) language code or mul
// for multiple languages.
type AnalysisSchemeLanguage string

// Enum values for AnalysisSchemeLanguage
const (
	AnalysisSchemeLanguageAr     AnalysisSchemeLanguage = "ar"
	AnalysisSchemeLanguageBg     AnalysisSchemeLanguage = "bg"
	AnalysisSchemeLanguageCa     AnalysisSchemeLanguage = "ca"
	AnalysisSchemeLanguageCs     AnalysisSchemeLanguage = "cs"
	AnalysisSchemeLanguageDa     AnalysisSchemeLanguage = "da"
	AnalysisSchemeLanguageDe     AnalysisSchemeLanguage = "de"
	AnalysisSchemeLanguageEl     AnalysisSchemeLanguage = "el"
	AnalysisSchemeLanguageEn     AnalysisSchemeLanguage = "en"
	AnalysisSchemeLanguageEs     AnalysisSchemeLanguage = "es"
	AnalysisSchemeLanguageEu     AnalysisSchemeLanguage = "eu"
	AnalysisSchemeLanguageFa     AnalysisSchemeLanguage = "fa"
	AnalysisSchemeLanguageFi     AnalysisSchemeLanguage = "fi"
	AnalysisSchemeLanguageFr     AnalysisSchemeLanguage = "fr"
	AnalysisSchemeLanguageGa     AnalysisSchemeLanguage = "ga"
	AnalysisSchemeLanguageGl     AnalysisSchemeLanguage = "gl"
	AnalysisSchemeLanguageHe     AnalysisSchemeLanguage = "he"
	AnalysisSchemeLanguageHi     AnalysisSchemeLanguage = "hi"
	AnalysisSchemeLanguageHu     AnalysisSchemeLanguage = "hu"
	AnalysisSchemeLanguageHy     AnalysisSchemeLanguage = "hy"
	AnalysisSchemeLanguageId     AnalysisSchemeLanguage = "id"
	AnalysisSchemeLanguageIt     AnalysisSchemeLanguage = "it"
	AnalysisSchemeLanguageJa     AnalysisSchemeLanguage = "ja"
	AnalysisSchemeLanguageKo     AnalysisSchemeLanguage = "ko"
	AnalysisSchemeLanguageLv     AnalysisSchemeLanguage = "lv"
	AnalysisSchemeLanguageMul    AnalysisSchemeLanguage = "mul"
	AnalysisSchemeLanguageNl     AnalysisSchemeLanguage = "nl"
	AnalysisSchemeLanguageNo     AnalysisSchemeLanguage = "no"
	AnalysisSchemeLanguagePt     AnalysisSchemeLanguage = "pt"
	AnalysisSchemeLanguageRo     AnalysisSchemeLanguage = "ro"
	AnalysisSchemeLanguageRu     AnalysisSchemeLanguage = "ru"
	AnalysisSchemeLanguageSv     AnalysisSchemeLanguage = "sv"
	AnalysisSchemeLanguageTh     AnalysisSchemeLanguage = "th"
	AnalysisSchemeLanguageTr     AnalysisSchemeLanguage = "tr"
	AnalysisSchemeLanguageZhHans AnalysisSchemeLanguage = "zh-Hans"
	AnalysisSchemeLanguageZhHant AnalysisSchemeLanguage = "zh-Hant"
)

func (enum AnalysisSchemeLanguage) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AnalysisSchemeLanguage) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of field. The valid options for a field depend on the field type.
// For more information about the supported field types, see Configuring Index
// Fields (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html)
// in the Amazon CloudSearch Developer Guide.
type IndexFieldType string

// Enum values for IndexFieldType
const (
	IndexFieldTypeInt          IndexFieldType = "int"
	IndexFieldTypeDouble       IndexFieldType = "double"
	IndexFieldTypeLiteral      IndexFieldType = "literal"
	IndexFieldTypeText         IndexFieldType = "text"
	IndexFieldTypeDate         IndexFieldType = "date"
	IndexFieldTypeLatlon       IndexFieldType = "latlon"
	IndexFieldTypeIntArray     IndexFieldType = "int-array"
	IndexFieldTypeDoubleArray  IndexFieldType = "double-array"
	IndexFieldTypeLiteralArray IndexFieldType = "literal-array"
	IndexFieldTypeTextArray    IndexFieldType = "text-array"
	IndexFieldTypeDateArray    IndexFieldType = "date-array"
)

func (enum IndexFieldType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IndexFieldType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The state of processing a change to an option. One of:
//
//    * RequiresIndexDocuments: The option's latest value will not be deployed
//    until IndexDocuments has been called and indexing is complete.
//
//    * Processing: The option's latest value is in the process of being activated.
//
//    * Active: The option's latest value is fully deployed.
//
//    * FailedToValidate: The option value is not compatible with the domain's
//    data and cannot be used to index the data. You must either modify the
//    option value or update or remove the incompatible documents.
type OptionState string

// Enum values for OptionState
const (
	OptionStateRequiresIndexDocuments OptionState = "RequiresIndexDocuments"
	OptionStateProcessing             OptionState = "Processing"
	OptionStateActive                 OptionState = "Active"
	OptionStateFailedToValidate       OptionState = "FailedToValidate"
)

func (enum OptionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OptionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The instance type (such as search.m1.small) on which an index partition is
// hosted.
type PartitionInstanceType string

// Enum values for PartitionInstanceType
const (
	PartitionInstanceTypeSearchM1Small   PartitionInstanceType = "search.m1.small"
	PartitionInstanceTypeSearchM1Large   PartitionInstanceType = "search.m1.large"
	PartitionInstanceTypeSearchM2Xlarge  PartitionInstanceType = "search.m2.xlarge"
	PartitionInstanceTypeSearchM22xlarge PartitionInstanceType = "search.m2.2xlarge"
	PartitionInstanceTypeSearchM3Medium  PartitionInstanceType = "search.m3.medium"
	PartitionInstanceTypeSearchM3Large   PartitionInstanceType = "search.m3.large"
	PartitionInstanceTypeSearchM3Xlarge  PartitionInstanceType = "search.m3.xlarge"
	PartitionInstanceTypeSearchM32xlarge PartitionInstanceType = "search.m3.2xlarge"
)

func (enum PartitionInstanceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PartitionInstanceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SuggesterFuzzyMatching string

// Enum values for SuggesterFuzzyMatching
const (
	SuggesterFuzzyMatchingNone SuggesterFuzzyMatching = "none"
	SuggesterFuzzyMatchingLow  SuggesterFuzzyMatching = "low"
	SuggesterFuzzyMatchingHigh SuggesterFuzzyMatching = "high"
)

func (enum SuggesterFuzzyMatching) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SuggesterFuzzyMatching) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
