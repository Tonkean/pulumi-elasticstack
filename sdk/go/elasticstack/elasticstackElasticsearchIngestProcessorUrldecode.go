// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// URL-decodes a string. If the field is an array of strings, all members of the array will be decoded.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/urldecode-processor.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			urldecode, err := elasticstack.ElasticstackElasticsearchIngestProcessorUrldecode(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorUrldecodeArgs{
//				Field: "my_url_to_decode",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(urldecode.Json),
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
func ElasticstackElasticsearchIngestProcessorUrldecode(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorUrldecodeArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorUrldecodeResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorUrldecodeResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorUrldecode:ElasticstackElasticsearchIngestProcessorUrldecode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorUrldecode.
type ElasticstackElasticsearchIngestProcessorUrldecodeArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to decode
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorUrldecode.
type ElasticstackElasticsearchIngestProcessorUrldecodeResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to decode
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
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

func ElasticstackElasticsearchIngestProcessorUrldecodeOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorUrldecodeOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorUrldecodeResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorUrldecodeArgs)
			r, err := ElasticstackElasticsearchIngestProcessorUrldecode(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorUrldecodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorUrldecode.
type ElasticstackElasticsearchIngestProcessorUrldecodeOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to decode
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (ElasticstackElasticsearchIngestProcessorUrldecodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorUrldecodeArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorUrldecode.
type ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorUrldecodeResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) ToElasticstackElasticsearchIngestProcessorUrldecodeResultOutput() ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) ToElasticstackElasticsearchIngestProcessorUrldecodeResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to decode
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the converted value to, by default `field` is updated in-place.
func (o ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUrldecodeResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorUrldecodeResultOutput{})
}