// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_jsonrpc service, marshal.go contains the generated operation marshalers and methods on it.

package aws_jsonrpc

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AddTagsToStreamMarshaler defines marshaler for AddTagsToStream operation
type AddTagsToStreamMarshaler struct {
	Input *types.AddTagsToStreamInput
}

// CreateStreamMarshaler defines marshaler for CreateStream operation
type CreateStreamMarshaler struct {
	Input *types.CreateStreamInput
}

// DecreaseStreamRetentionPeriodMarshaler defines marshaler for DecreaseStreamRetentionPeriod operation
type DecreaseStreamRetentionPeriodMarshaler struct {
	Input *types.DecreaseStreamRetentionPeriodInput
}

// DeleteStreamMarshaler defines marshaler for DeleteStream operation
type DeleteStreamMarshaler struct {
	Input *types.DeleteStreamInput
}

// DeregisterStreamConsumerMarshaler defines marshaler for DeregisterStreamConsumer operation
type DeregisterStreamConsumerMarshaler struct {
	Input *types.DeregisterStreamConsumerInput
}

// DescribeLimitsMarshaler defines marshaler for DescribeLimits operation
type DescribeLimitsMarshaler struct {
	Input *types.DescribeLimitsInput
}

// DescribeStreamMarshaler defines marshaler for DescribeStream operation
type DescribeStreamMarshaler struct {
	Input *types.DescribeStreamInput
}

// DescribeStreamConsumerMarshaler defines marshaler for DescribeStreamConsumer operation
type DescribeStreamConsumerMarshaler struct {
	Input *types.DescribeStreamConsumerInput
}

// DescribeStreamSummaryMarshaler defines marshaler for DescribeStreamSummary operation
type DescribeStreamSummaryMarshaler struct {
	Input *types.DescribeStreamSummaryInput
}

// DisableEnhancedMonitoringMarshaler defines marshaler for DisableEnhancedMonitoring operation
type DisableEnhancedMonitoringMarshaler struct {
	Input *types.DisableEnhancedMonitoringInput
}

// EnableEnhancedMonitoringMarshaler defines marshaler for EnableEnhancedMonitoring operation
type EnableEnhancedMonitoringMarshaler struct {
	Input *types.EnableEnhancedMonitoringInput
}

// GetRecordsMarshaler defines marshaler for GetRecords operation
type GetRecordsMarshaler struct {
	Input *types.GetRecordsInput
}

// GetShardIteratorMarshaler defines marshaler for GetShardIterator operation
type GetShardIteratorMarshaler struct {
	Input *types.GetShardIteratorInput
}

// IncreaseStreamRetentionPeriodMarshaler defines marshaler for IncreaseStreamRetentionPeriod operation
type IncreaseStreamRetentionPeriodMarshaler struct {
	Input *types.IncreaseStreamRetentionPeriodInput
}

// ListShardsMarshaler defines marshaler for ListShards operation
type ListShardsMarshaler struct {
	Input *types.ListShardsInput
}

// ListStreamConsumersMarshaler defines marshaler for ListStreamConsumers operation
type ListStreamConsumersMarshaler struct {
	Input *types.ListStreamConsumersInput
}

// ListStreamsMarshaler defines marshaler for ListStreams operation
type ListStreamsMarshaler struct {
	Input *types.ListStreamsInput
}

// ListTagsForStreamMarshaler defines marshaler for ListTagsForStream operation
type ListTagsForStreamMarshaler struct {
	Input *types.ListTagsForStreamInput
}

// MergeShardsMarshaler defines marshaler for MergeShards operation
type MergeShardsMarshaler struct {
	Input *types.MergeShardsInput
}

// PutRecordMarshaler defines marshaler for PutRecord operation
type PutRecordMarshaler struct {
	Input *types.PutRecordInput
}

// PutRecordsMarshaler defines marshaler for PutRecords operation
type PutRecordsMarshaler struct {
	Input *types.PutRecordsInput
}

// RegisterStreamConsumerMarshaler defines marshaler for RegisterStreamConsumer operation
type RegisterStreamConsumerMarshaler struct {
	Input *types.RegisterStreamConsumerInput
}

// RemoveTagsFromStreamMarshaler defines marshaler for RemoveTagsFromStream operation
type RemoveTagsFromStreamMarshaler struct {
	Input *types.RemoveTagsFromStreamInput
}

// SplitShardMarshaler defines marshaler for SplitShard operation
type SplitShardMarshaler struct {
	Input *types.SplitShardInput
}

// StartStreamEncryptionMarshaler defines marshaler for StartStreamEncryption operation
type StartStreamEncryptionMarshaler struct {
	Input *types.StartStreamEncryptionInput
}

// StopStreamEncryptionMarshaler defines marshaler for StopStreamEncryption operation
type StopStreamEncryptionMarshaler struct {
	Input *types.StopStreamEncryptionInput
}

// UpdateShardCountMarshaler defines marshaler for UpdateShardCount operation
type UpdateShardCountMarshaler struct {
	Input *types.UpdateShardCountInput
}

func MarshalAddTagsToStreamInputShapeAWSJSON(v *types.AddTagsToStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateStreamInputShapeAWSJSON(v *types.CreateStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDecreaseStreamRetentionPeriodInputShapeAWSJSON(v *types.DecreaseStreamRetentionPeriodInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteStreamInputShapeAWSJSON(v *types.DeleteStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeregisterStreamConsumerInputShapeAWSJSON(v *types.DeregisterStreamConsumerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeLimitsInputShapeAWSJSON(v *types.DescribeLimitsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeStreamInputShapeAWSJSON(v *types.DescribeStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeStreamConsumerInputShapeAWSJSON(v *types.DescribeStreamConsumerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeStreamSummaryInputShapeAWSJSON(v *types.DescribeStreamSummaryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDisableEnhancedMonitoringInputShapeAWSJSON(v *types.DisableEnhancedMonitoringInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalEnableEnhancedMonitoringInputShapeAWSJSON(v *types.EnableEnhancedMonitoringInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalGetRecordsInputShapeAWSJSON(v *types.GetRecordsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalGetShardIteratorInputShapeAWSJSON(v *types.GetShardIteratorInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalIncreaseStreamRetentionPeriodInputShapeAWSJSON(v *types.IncreaseStreamRetentionPeriodInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListShardsInputShapeAWSJSON(v *types.ListShardsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListStreamConsumersInputShapeAWSJSON(v *types.ListStreamConsumersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListStreamsInputShapeAWSJSON(v *types.ListStreamsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListTagsForStreamInputShapeAWSJSON(v *types.ListTagsForStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalMergeShardsInputShapeAWSJSON(v *types.MergeShardsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalPutRecordInputShapeAWSJSON(v *types.PutRecordInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalPutRecordsInputShapeAWSJSON(v *types.PutRecordsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRegisterStreamConsumerInputShapeAWSJSON(v *types.RegisterStreamConsumerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRemoveTagsFromStreamInputShapeAWSJSON(v *types.RemoveTagsFromStreamInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalSplitShardInputShapeAWSJSON(v *types.SplitShardInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalStartStreamEncryptionInputShapeAWSJSON(v *types.StartStreamEncryptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalStopStreamEncryptionInputShapeAWSJSON(v *types.StopStreamEncryptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateShardCountInputShapeAWSJSON(v *types.UpdateShardCountInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}
func (m AddTagsToStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAddTagsToStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AddTagsToStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AddTagsToStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DecreaseStreamRetentionPeriodMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDecreaseStreamRetentionPeriodInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DecreaseStreamRetentionPeriodMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DecreaseStreamRetentionPeriod.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeregisterStreamConsumerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeregisterStreamConsumerInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeregisterStreamConsumerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeregisterStreamConsumer.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeLimitsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeLimitsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeLimitsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeLimits.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeStreamConsumerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeStreamConsumerInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeStreamConsumerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeStreamConsumer.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeStreamSummaryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeStreamSummaryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeStreamSummaryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeStreamSummary.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DisableEnhancedMonitoringMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDisableEnhancedMonitoringInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DisableEnhancedMonitoringMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DisableEnhancedMonitoring.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m EnableEnhancedMonitoringMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalEnableEnhancedMonitoringInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m EnableEnhancedMonitoringMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "EnableEnhancedMonitoring.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetRecordsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetRecordsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetRecordsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetRecords.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetShardIteratorMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetShardIteratorInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetShardIteratorMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetShardIterator.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m IncreaseStreamRetentionPeriodMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalIncreaseStreamRetentionPeriodInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m IncreaseStreamRetentionPeriodMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "IncreaseStreamRetentionPeriod.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListShardsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListShardsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListShardsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListShards.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListStreamConsumersMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListStreamConsumersInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListStreamConsumersMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListStreamConsumers.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListStreamsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListStreamsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListStreamsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListStreams.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListTagsForStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListTagsForStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m MergeShardsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalMergeShardsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m MergeShardsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "MergeShards.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m PutRecordMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalPutRecordInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m PutRecordMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "PutRecord.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m PutRecordsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalPutRecordsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m PutRecordsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "PutRecords.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RegisterStreamConsumerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRegisterStreamConsumerInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RegisterStreamConsumerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RegisterStreamConsumer.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RemoveTagsFromStreamMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRemoveTagsFromStreamInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RemoveTagsFromStreamMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RemoveTagsFromStream.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SplitShardMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSplitShardInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SplitShardMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SplitShard.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StartStreamEncryptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStartStreamEncryptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StartStreamEncryptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StartStreamEncryption.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StopStreamEncryptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStopStreamEncryptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StopStreamEncryptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StopStreamEncryption.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateShardCountMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateShardCountInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateShardCountMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateShardCount.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
