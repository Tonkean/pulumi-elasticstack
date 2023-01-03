// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Expands a field with dots into an object field. This processor allows fields with dots in the name to be accessible by other processors in the pipeline. Otherwise these fields can’t be accessed by any processor.
//
// See: elastic.co/guide/en/elasticsearch/reference/current/dot-expand-processor.html
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
//			dotExpander, err := elasticstack.IngestProcessorDotExpander(ctx, &elasticstack.IngestProcessorDotExpanderArgs{
//				Field: "foo.bar",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(dotExpander.Json),
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
func IngestProcessorDotExpander(ctx *pulumi.Context, args *IngestProcessorDotExpanderArgs, opts ...pulumi.InvokeOption) (*IngestProcessorDotExpanderResult, error) {
	var rv IngestProcessorDotExpanderResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorDotExpander:IngestProcessorDotExpander", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorDotExpander.
type IngestProcessorDotExpanderArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to expand into an object field. If set to *, all top-level fields will be expanded.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
	Override *bool `pulumi:"override"`
	// The field that contains the field to expand.
	Path *string `pulumi:"path"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by IngestProcessorDotExpander.
type IngestProcessorDotExpanderResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to expand into an object field. If set to *, all top-level fields will be expanded.
	Field string `pulumi:"field"`
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
	// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
	Override *bool `pulumi:"override"`
	// The field that contains the field to expand.
	Path *string `pulumi:"path"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func IngestProcessorDotExpanderOutput(ctx *pulumi.Context, args IngestProcessorDotExpanderOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorDotExpanderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorDotExpanderResult, error) {
			args := v.(IngestProcessorDotExpanderArgs)
			r, err := IngestProcessorDotExpander(ctx, &args, opts...)
			var s IngestProcessorDotExpanderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorDotExpanderResultOutput)
}

// A collection of arguments for invoking IngestProcessorDotExpander.
type IngestProcessorDotExpanderOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to expand into an object field. If set to *, all top-level fields will be expanded.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
	Override pulumi.BoolPtrInput `pulumi:"override"`
	// The field that contains the field to expand.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (IngestProcessorDotExpanderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDotExpanderArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorDotExpander.
type IngestProcessorDotExpanderResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorDotExpanderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDotExpanderResult)(nil)).Elem()
}

func (o IngestProcessorDotExpanderResultOutput) ToIngestProcessorDotExpanderResultOutput() IngestProcessorDotExpanderResultOutput {
	return o
}

func (o IngestProcessorDotExpanderResultOutput) ToIngestProcessorDotExpanderResultOutputWithContext(ctx context.Context) IngestProcessorDotExpanderResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorDotExpanderResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to expand into an object field. If set to *, all top-level fields will be expanded.
func (o IngestProcessorDotExpanderResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o IngestProcessorDotExpanderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorDotExpanderResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorDotExpanderResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorDotExpanderResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorDotExpanderResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
func (o IngestProcessorDotExpanderResultOutput) Override() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *bool { return v.Override }).(pulumi.BoolPtrOutput)
}

// The field that contains the field to expand.
func (o IngestProcessorDotExpanderResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o IngestProcessorDotExpanderResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDotExpanderResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorDotExpanderResultOutput{})
}