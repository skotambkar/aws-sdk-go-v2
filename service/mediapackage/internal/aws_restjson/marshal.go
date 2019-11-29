// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
)

var _ bytes.Buffer
var _ awserr.Error

// CreateChannelMarshaler defines marshaler for CreateChannel operation
type CreateChannelMarshaler struct {
	Input *types.CreateChannelInput
}

// CreateHarvestJobMarshaler defines marshaler for CreateHarvestJob operation
type CreateHarvestJobMarshaler struct {
	Input *types.CreateHarvestJobInput
}

// CreateOriginEndpointMarshaler defines marshaler for CreateOriginEndpoint operation
type CreateOriginEndpointMarshaler struct {
	Input *types.CreateOriginEndpointInput
}

// DeleteChannelMarshaler defines marshaler for DeleteChannel operation
type DeleteChannelMarshaler struct {
	Input *types.DeleteChannelInput
}

// DeleteOriginEndpointMarshaler defines marshaler for DeleteOriginEndpoint operation
type DeleteOriginEndpointMarshaler struct {
	Input *types.DeleteOriginEndpointInput
}

// DescribeChannelMarshaler defines marshaler for DescribeChannel operation
type DescribeChannelMarshaler struct {
	Input *types.DescribeChannelInput
}

// DescribeHarvestJobMarshaler defines marshaler for DescribeHarvestJob operation
type DescribeHarvestJobMarshaler struct {
	Input *types.DescribeHarvestJobInput
}

// DescribeOriginEndpointMarshaler defines marshaler for DescribeOriginEndpoint operation
type DescribeOriginEndpointMarshaler struct {
	Input *types.DescribeOriginEndpointInput
}

// ListChannelsMarshaler defines marshaler for ListChannels operation
type ListChannelsMarshaler struct {
	Input *types.ListChannelsInput
}

// ListHarvestJobsMarshaler defines marshaler for ListHarvestJobs operation
type ListHarvestJobsMarshaler struct {
	Input *types.ListHarvestJobsInput
}

// ListOriginEndpointsMarshaler defines marshaler for ListOriginEndpoints operation
type ListOriginEndpointsMarshaler struct {
	Input *types.ListOriginEndpointsInput
}

// ListTagsForResourceMarshaler defines marshaler for ListTagsForResource operation
type ListTagsForResourceMarshaler struct {
	Input *types.ListTagsForResourceInput
}

// RotateChannelCredentialsMarshaler defines marshaler for RotateChannelCredentials operation
type RotateChannelCredentialsMarshaler struct {
	Input *types.RotateChannelCredentialsInput
}

// RotateIngestEndpointCredentialsMarshaler defines marshaler for RotateIngestEndpointCredentials operation
type RotateIngestEndpointCredentialsMarshaler struct {
	Input *types.RotateIngestEndpointCredentialsInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

// UpdateChannelMarshaler defines marshaler for UpdateChannel operation
type UpdateChannelMarshaler struct {
	Input *types.UpdateChannelInput
}

// UpdateOriginEndpointMarshaler defines marshaler for UpdateOriginEndpoint operation
type UpdateOriginEndpointMarshaler struct {
	Input *types.UpdateOriginEndpointInput
}

func MarshalCreateChannelInputShapeAWSREST(v *types.CreateChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateChannelInputShapeAWSJSON(v *types.CreateChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalCreateHarvestJobInputShapeAWSREST(v *types.CreateHarvestJobInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateHarvestJobInputShapeAWSJSON(v *types.CreateHarvestJobInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalCreateOriginEndpointInputShapeAWSREST(v *types.CreateOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateOriginEndpointInputShapeAWSJSON(v *types.CreateOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteChannelInputShapeAWSREST(v *types.DeleteChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteChannelInputShapeAWSJSON(v *types.DeleteChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteOriginEndpointInputShapeAWSREST(v *types.DeleteOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteOriginEndpointInputShapeAWSJSON(v *types.DeleteOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeChannelInputShapeAWSREST(v *types.DescribeChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeChannelInputShapeAWSJSON(v *types.DescribeChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeHarvestJobInputShapeAWSREST(v *types.DescribeHarvestJobInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeHarvestJobInputShapeAWSJSON(v *types.DescribeHarvestJobInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeOriginEndpointInputShapeAWSREST(v *types.DescribeOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeOriginEndpointInputShapeAWSJSON(v *types.DescribeOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListChannelsInputShapeAWSREST(v *types.ListChannelsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListChannelsInputShapeAWSJSON(v *types.ListChannelsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListHarvestJobsInputShapeAWSREST(v *types.ListHarvestJobsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListHarvestJobsInputShapeAWSJSON(v *types.ListHarvestJobsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListOriginEndpointsInputShapeAWSREST(v *types.ListOriginEndpointsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListOriginEndpointsInputShapeAWSJSON(v *types.ListOriginEndpointsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListTagsForResourceInputShapeAWSREST(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListTagsForResourceInputShapeAWSJSON(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalRotateChannelCredentialsInputShapeAWSREST(v *types.RotateChannelCredentialsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalRotateChannelCredentialsInputShapeAWSJSON(v *types.RotateChannelCredentialsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalRotateIngestEndpointCredentialsInputShapeAWSREST(v *types.RotateIngestEndpointCredentialsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalRotateIngestEndpointCredentialsInputShapeAWSJSON(v *types.RotateIngestEndpointCredentialsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalTagResourceInputShapeAWSREST(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalTagResourceInputShapeAWSJSON(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUntagResourceInputShapeAWSREST(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUntagResourceInputShapeAWSJSON(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateChannelInputShapeAWSREST(v *types.UpdateChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateChannelInputShapeAWSJSON(v *types.UpdateChannelInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateOriginEndpointInputShapeAWSREST(v *types.UpdateOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateOriginEndpointInputShapeAWSJSON(v *types.UpdateOriginEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m CreateChannelMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateChannelInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateChannelInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateChannelMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateChannel.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateHarvestJobMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateHarvestJobInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateHarvestJobInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateHarvestJobMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateHarvestJob.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateOriginEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateOriginEndpointInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateOriginEndpointInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateOriginEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateOriginEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteChannelMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteChannelInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteChannelInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteChannelMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteChannel.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteOriginEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteOriginEndpointInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteOriginEndpointInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteOriginEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteOriginEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeChannelMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeChannelInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeChannelInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeChannelMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeChannel.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeHarvestJobMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeHarvestJobInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeHarvestJobInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeHarvestJobMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeHarvestJob.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeOriginEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeOriginEndpointInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeOriginEndpointInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeOriginEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeOriginEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListChannelsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListChannelsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListChannelsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListChannelsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListChannels.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListHarvestJobsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListHarvestJobsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListHarvestJobsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListHarvestJobsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListHarvestJobs.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListOriginEndpointsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListOriginEndpointsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListOriginEndpointsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListOriginEndpointsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListOriginEndpoints.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListTagsForResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListTagsForResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListTagsForResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RotateChannelCredentialsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRotateChannelCredentialsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalRotateChannelCredentialsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RotateChannelCredentialsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RotateChannelCredentials.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RotateIngestEndpointCredentialsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRotateIngestEndpointCredentialsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalRotateIngestEndpointCredentialsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RotateIngestEndpointCredentialsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RotateIngestEndpointCredentials.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalTagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m TagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "TagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UntagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUntagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUntagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UntagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UntagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateChannelMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateChannelInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateChannelInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateChannelMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateChannel.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateOriginEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateOriginEndpointInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateOriginEndpointInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateOriginEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateOriginEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
