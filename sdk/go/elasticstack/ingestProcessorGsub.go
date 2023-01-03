// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts a string field by applying a regular expression and a replacement. If the field is an array of string, all members of the array will be converted. If any non-string values are encountered, the processor will throw an exception.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/gsub-processor.html
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
//			gsub, err := elasticstack.IngestProcessorGsub(ctx, &elasticstack.IngestProcessorGsubArgs{
//				Field:       "field1",
//				Pattern:     "\\.",
//				Replacement: "-",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(gsub.Json),
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
func IngestProcessorGsub(ctx *pulumi.Context, args *IngestProcessorGsubArgs, opts ...pulumi.InvokeOption) (*IngestProcessorGsubResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv IngestProcessorGsubResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorGsub:IngestProcessorGsub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorGsub.
type IngestProcessorGsubArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to apply the replacement to.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The pattern to be replaced.
	Pattern string `pulumi:"pattern"`
	// The string to replace the matching patterns with.
	Replacement string `pulumi:"replacement"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by IngestProcessorGsub.
type IngestProcessorGsubResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to apply the replacement to.
	Field string `pulumi:"field"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The pattern to be replaced.
	Pattern string `pulumi:"pattern"`
	// The string to replace the matching patterns with.
	Replacement string `pulumi:"replacement"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

func IngestProcessorGsubOutput(ctx *pulumi.Context, args IngestProcessorGsubOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorGsubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorGsubResult, error) {
			args := v.(IngestProcessorGsubArgs)
			r, err := IngestProcessorGsub(ctx, &args, opts...)
			var s IngestProcessorGsubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorGsubResultOutput)
}

// A collection of arguments for invoking IngestProcessorGsub.
type IngestProcessorGsubOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to apply the replacement to.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The pattern to be replaced.
	Pattern pulumi.StringInput `pulumi:"pattern"`
	// The string to replace the matching patterns with.
	Replacement pulumi.StringInput `pulumi:"replacement"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (IngestProcessorGsubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorGsubArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorGsub.
type IngestProcessorGsubResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorGsubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorGsubResult)(nil)).Elem()
}

func (o IngestProcessorGsubResultOutput) ToIngestProcessorGsubResultOutput() IngestProcessorGsubResultOutput {
	return o
}

func (o IngestProcessorGsubResultOutput) ToIngestProcessorGsubResultOutputWithContext(ctx context.Context) IngestProcessorGsubResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorGsubResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to apply the replacement to.
func (o IngestProcessorGsubResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o IngestProcessorGsubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorGsubResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorGsubResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o IngestProcessorGsubResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorGsubResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorGsubResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The pattern to be replaced.
func (o IngestProcessorGsubResultOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) string { return v.Pattern }).(pulumi.StringOutput)
}

// The string to replace the matching patterns with.
func (o IngestProcessorGsubResultOutput) Replacement() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) string { return v.Replacement }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o IngestProcessorGsubResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the converted value to, by default `field` is updated in-place.
func (o IngestProcessorGsubResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorGsubResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorGsubResultOutput{})
}
