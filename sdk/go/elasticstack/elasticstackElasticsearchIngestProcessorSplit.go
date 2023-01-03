// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Splits a field into an array using a separator character. Only works on string fields.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/split-processor.html
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
//			split, err := elasticstack.ElasticstackElasticsearchIngestProcessorSplit(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorSplitArgs{
//				Field:     "my_field",
//				Separator: "\\s+",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(split.Json),
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
func ElasticstackElasticsearchIngestProcessorSplit(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorSplitArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorSplitResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorSplitResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorSplit:ElasticstackElasticsearchIngestProcessorSplit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorSplit.
type ElasticstackElasticsearchIngestProcessorSplitArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to split
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Preserves empty trailing fields, if any.
	PreserveTrailing *bool `pulumi:"preserveTrailing"`
	// A regex which matches the separator, eg `,` or `\s+`
	Separator string `pulumi:"separator"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorSplit.
type ElasticstackElasticsearchIngestProcessorSplitResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to split
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
	// Preserves empty trailing fields, if any.
	PreserveTrailing *bool `pulumi:"preserveTrailing"`
	// A regex which matches the separator, eg `,` or `\s+`
	Separator string `pulumi:"separator"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

func ElasticstackElasticsearchIngestProcessorSplitOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorSplitOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorSplitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorSplitResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorSplitArgs)
			r, err := ElasticstackElasticsearchIngestProcessorSplit(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorSplitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorSplitResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorSplit.
type ElasticstackElasticsearchIngestProcessorSplitOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to split
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Preserves empty trailing fields, if any.
	PreserveTrailing pulumi.BoolPtrInput `pulumi:"preserveTrailing"`
	// A regex which matches the separator, eg `,` or `\s+`
	Separator pulumi.StringInput `pulumi:"separator"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (ElasticstackElasticsearchIngestProcessorSplitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorSplitArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorSplit.
type ElasticstackElasticsearchIngestProcessorSplitResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorSplitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorSplitResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) ToElasticstackElasticsearchIngestProcessorSplitResultOutput() ElasticstackElasticsearchIngestProcessorSplitResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) ToElasticstackElasticsearchIngestProcessorSplitResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorSplitResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to split
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Preserves empty trailing fields, if any.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) PreserveTrailing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *bool { return v.PreserveTrailing }).(pulumi.BoolPtrOutput)
}

// A regex which matches the separator, eg `,` or `\s+`
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Separator() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) string { return v.Separator }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the converted value to, by default `field` is updated in-place.
func (o ElasticstackElasticsearchIngestProcessorSplitResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSplitResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorSplitResultOutput{})
}