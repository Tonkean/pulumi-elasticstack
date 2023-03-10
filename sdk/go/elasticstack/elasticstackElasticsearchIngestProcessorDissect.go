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
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dissect, err := elasticstack.ElasticstackElasticsearchIngestProcessorDissect(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorDissectArgs{
//				Field:   "message",
//				Pattern: fmt.Sprintf("%vclientip} %vident} %vauth} [%v@timestamp}] \"%vverb} %vrequest} HTTP/%vhttpversion}\" %vstatus} %vsize}", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
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
func ElasticstackElasticsearchIngestProcessorDissect(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorDissectArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorDissectResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorDissectResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorDissect:ElasticstackElasticsearchIngestProcessorDissect", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDissect.
type ElasticstackElasticsearchIngestProcessorDissectArgs struct {
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

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDissect.
type ElasticstackElasticsearchIngestProcessorDissectResult struct {
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

func ElasticstackElasticsearchIngestProcessorDissectOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorDissectOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorDissectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorDissectResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorDissectArgs)
			r, err := ElasticstackElasticsearchIngestProcessorDissect(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorDissectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorDissectResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDissect.
type ElasticstackElasticsearchIngestProcessorDissectOutputArgs struct {
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

func (ElasticstackElasticsearchIngestProcessorDissectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDissectArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDissect.
type ElasticstackElasticsearchIngestProcessorDissectResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorDissectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDissectResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) ToElasticstackElasticsearchIngestProcessorDissectResultOutput() ElasticstackElasticsearchIngestProcessorDissectResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) ToElasticstackElasticsearchIngestProcessorDissectResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorDissectResultOutput {
	return o
}

// The character(s) that separate the appended fields.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) AppendSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *string { return v.AppendSeparator }).(pulumi.StringPtrOutput)
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to dissect.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The pattern to apply to the field.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) string { return v.Pattern }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorDissectResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDissectResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorDissectResultOutput{})
}
