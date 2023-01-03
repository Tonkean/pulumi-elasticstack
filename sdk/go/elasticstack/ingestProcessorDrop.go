// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
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
//			drop, err := elasticstack.IngestProcessorDrop(ctx, &elasticstack.IngestProcessorDropArgs{
//				If: pulumi.StringRef("ctx.network_name == 'Guest'"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(drop.Json),
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
func IngestProcessorDrop(ctx *pulumi.Context, args *IngestProcessorDropArgs, opts ...pulumi.InvokeOption) (*IngestProcessorDropResult, error) {
	var rv IngestProcessorDropResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorDrop:IngestProcessorDrop", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorDrop.
type IngestProcessorDropArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by IngestProcessorDrop.
type IngestProcessorDropResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func IngestProcessorDropOutput(ctx *pulumi.Context, args IngestProcessorDropOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorDropResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorDropResult, error) {
			args := v.(IngestProcessorDropArgs)
			r, err := IngestProcessorDrop(ctx, &args, opts...)
			var s IngestProcessorDropResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorDropResultOutput)
}

// A collection of arguments for invoking IngestProcessorDrop.
type IngestProcessorDropOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (IngestProcessorDropOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDropArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorDrop.
type IngestProcessorDropResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorDropResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDropResult)(nil)).Elem()
}

func (o IngestProcessorDropResultOutput) ToIngestProcessorDropResultOutput() IngestProcessorDropResultOutput {
	return o
}

func (o IngestProcessorDropResultOutput) ToIngestProcessorDropResultOutputWithContext(ctx context.Context) IngestProcessorDropResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorDropResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Internal identifier of the resource
func (o IngestProcessorDropResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorDropResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorDropResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorDropResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorDropResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o IngestProcessorDropResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDropResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorDropResultOutput{})
}