// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.
//
// Dissect matches a single text field against a defined pattern.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/Tonkean/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dissect, err := elasticstack.IngestProcessorDissect(ctx, &elasticstack.IngestProcessorDissectArgs{
//				Field:   "message",
//				Pattern: fmt.Sprintf("%vclientip} %vident} %vauth} [%v@timestamp}] \"%vverb} %vrequest} HTTP/%vhttpversion}\" %vstatus} %vsize}", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(dissect.Json),
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
func IngestProcessorDissect(ctx *pulumi.Context, args *IngestProcessorDissectArgs, opts ...pulumi.InvokeOption) (*IngestProcessorDissectResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv IngestProcessorDissectResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorDissect:IngestProcessorDissect", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorDissect.
type IngestProcessorDissectArgs struct {
	// The character(s) that separate the appended fields.
	AppendSeparator *string `pulumi:"appendSeparator"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to dissect.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The pattern to apply to the field.
	Pattern string `pulumi:"pattern"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by IngestProcessorDissect.
type IngestProcessorDissectResult struct {
	// The character(s) that separate the appended fields.
	AppendSeparator *string `pulumi:"appendSeparator"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to dissect.
	Field string `pulumi:"field"`
	// Internal identifier of the resource
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
	// The pattern to apply to the field.
	Pattern string `pulumi:"pattern"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func IngestProcessorDissectOutput(ctx *pulumi.Context, args IngestProcessorDissectOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorDissectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorDissectResult, error) {
			args := v.(IngestProcessorDissectArgs)
			r, err := IngestProcessorDissect(ctx, &args, opts...)
			var s IngestProcessorDissectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorDissectResultOutput)
}

// A collection of arguments for invoking IngestProcessorDissect.
type IngestProcessorDissectOutputArgs struct {
	// The character(s) that separate the appended fields.
	AppendSeparator pulumi.StringPtrInput `pulumi:"appendSeparator"`
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to dissect.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The pattern to apply to the field.
	Pattern pulumi.StringInput `pulumi:"pattern"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (IngestProcessorDissectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDissectArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorDissect.
type IngestProcessorDissectResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorDissectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDissectResult)(nil)).Elem()
}

func (o IngestProcessorDissectResultOutput) ToIngestProcessorDissectResultOutput() IngestProcessorDissectResultOutput {
	return o
}

func (o IngestProcessorDissectResultOutput) ToIngestProcessorDissectResultOutputWithContext(ctx context.Context) IngestProcessorDissectResultOutput {
	return o
}

// The character(s) that separate the appended fields.
func (o IngestProcessorDissectResultOutput) AppendSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *string { return v.AppendSeparator }).(pulumi.StringPtrOutput)
}

// Description of the processor.
func (o IngestProcessorDissectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to dissect.
func (o IngestProcessorDissectResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o IngestProcessorDissectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorDissectResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorDissectResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o IngestProcessorDissectResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorDissectResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorDissectResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The pattern to apply to the field.
func (o IngestProcessorDissectResultOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) string { return v.Pattern }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o IngestProcessorDissectResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDissectResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorDissectResultOutput{})
}
