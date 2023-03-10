// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Calculates the network direction given a source IP address, destination IP address, and a list of internal networks.
//
// The network direction processor reads IP addresses from Elastic Common Schema (ECS) fields by default. If you use the ECS, only the `internalNetworks` option must be specified.
//
// One of either `internalNetworks` or `internalNetworksField` must be specified. If `internalNetworksField` is specified, it follows the behavior specified by `ignoreMissing`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/Tonkean/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			networkDirection, err := elasticstack.IngestProcessorNetworkDirection(ctx, &elasticstack.IngestProcessorNetworkDirectionArgs{
//				InternalNetworks: []string{
//					"private",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(networkDirection.Json),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func IngestProcessorNetworkDirection(ctx *pulumi.Context, args *IngestProcessorNetworkDirectionArgs, opts ...pulumi.InvokeOption) (*IngestProcessorNetworkDirectionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv IngestProcessorNetworkDirectionResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorNetworkDirection:IngestProcessorNetworkDirection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorNetworkDirection.
type IngestProcessorNetworkDirectionArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Field containing the destination IP address.
	DestinationIp *string `pulumi:"destinationIp"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of internal networks.
	InternalNetworks []string `pulumi:"internalNetworks"`
	// A field on the given document to read the internalNetworks configuration from.
	InternalNetworksField *string `pulumi:"internalNetworksField"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Field containing the source IP address.
	SourceIp *string `pulumi:"sourceIp"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Output field for the network direction.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by IngestProcessorNetworkDirection.
type IngestProcessorNetworkDirectionResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Field containing the destination IP address.
	DestinationIp *string `pulumi:"destinationIp"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of internal networks.
	InternalNetworks []string `pulumi:"internalNetworks"`
	// A field on the given document to read the internalNetworks configuration from.
	InternalNetworksField *string `pulumi:"internalNetworksField"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Field containing the source IP address.
	SourceIp *string `pulumi:"sourceIp"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Output field for the network direction.
	TargetField *string `pulumi:"targetField"`
}

func IngestProcessorNetworkDirectionOutput(ctx *pulumi.Context, args IngestProcessorNetworkDirectionOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorNetworkDirectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorNetworkDirectionResult, error) {
			args := v.(IngestProcessorNetworkDirectionArgs)
			r, err := IngestProcessorNetworkDirection(ctx, &args, opts...)
			var s IngestProcessorNetworkDirectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorNetworkDirectionResultOutput)
}

// A collection of arguments for invoking IngestProcessorNetworkDirection.
type IngestProcessorNetworkDirectionOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Field containing the destination IP address.
	DestinationIp pulumi.StringPtrInput `pulumi:"destinationIp"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// List of internal networks.
	InternalNetworks pulumi.StringArrayInput `pulumi:"internalNetworks"`
	// A field on the given document to read the internalNetworks configuration from.
	InternalNetworksField pulumi.StringPtrInput `pulumi:"internalNetworksField"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Field containing the source IP address.
	SourceIp pulumi.StringPtrInput `pulumi:"sourceIp"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// Output field for the network direction.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (IngestProcessorNetworkDirectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorNetworkDirectionArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorNetworkDirection.
type IngestProcessorNetworkDirectionResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorNetworkDirectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorNetworkDirectionResult)(nil)).Elem()
}

func (o IngestProcessorNetworkDirectionResultOutput) ToIngestProcessorNetworkDirectionResultOutput() IngestProcessorNetworkDirectionResultOutput {
	return o
}

func (o IngestProcessorNetworkDirectionResultOutput) ToIngestProcessorNetworkDirectionResultOutputWithContext(ctx context.Context) IngestProcessorNetworkDirectionResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorNetworkDirectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Field containing the destination IP address.
func (o IngestProcessorNetworkDirectionResultOutput) DestinationIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.DestinationIp }).(pulumi.StringPtrOutput)
}

// Internal identifier of the resource.
func (o IngestProcessorNetworkDirectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorNetworkDirectionResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorNetworkDirectionResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o IngestProcessorNetworkDirectionResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// List of internal networks.
func (o IngestProcessorNetworkDirectionResultOutput) InternalNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) []string { return v.InternalNetworks }).(pulumi.StringArrayOutput)
}

// A field on the given document to read the internalNetworks configuration from.
func (o IngestProcessorNetworkDirectionResultOutput) InternalNetworksField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.InternalNetworksField }).(pulumi.StringPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorNetworkDirectionResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorNetworkDirectionResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Field containing the source IP address.
func (o IngestProcessorNetworkDirectionResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o IngestProcessorNetworkDirectionResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// Output field for the network direction.
func (o IngestProcessorNetworkDirectionResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorNetworkDirectionResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorNetworkDirectionResultOutput{})
}
