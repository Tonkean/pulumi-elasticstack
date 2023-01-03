// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sorts the elements of an array ascending or descending. Homogeneous arrays of numbers will be sorted numerically, while arrays of strings or heterogeneous arrays of strings + numbers will be sorted lexicographically. Throws an error when the field is not an array.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-processor.html
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
//			sort, err := elasticstack.ElasticstackElasticsearchIngestProcessorSort(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorSortArgs{
//				Field: "array_field_to_sort",
//				Order: pulumi.StringRef("desc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(sort.Json),
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
func ElasticstackElasticsearchIngestProcessorSort(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorSortArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorSortResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorSortResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorSort:ElasticstackElasticsearchIngestProcessorSort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorSort.
type ElasticstackElasticsearchIngestProcessorSortArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to be sorted
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The sort order to use. Accepts `asc` or `desc`.
	Order *string `pulumi:"order"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the sorted value to, by default `field` is updated in-place
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorSort.
type ElasticstackElasticsearchIngestProcessorSortResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to be sorted
	Field string `pulumi:"field"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The sort order to use. Accepts `asc` or `desc`.
	Order *string `pulumi:"order"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the sorted value to, by default `field` is updated in-place
	TargetField *string `pulumi:"targetField"`
}

func ElasticstackElasticsearchIngestProcessorSortOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorSortOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorSortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorSortResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorSortArgs)
			r, err := ElasticstackElasticsearchIngestProcessorSort(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorSortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorSortResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorSort.
type ElasticstackElasticsearchIngestProcessorSortOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to be sorted
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The sort order to use. Accepts `asc` or `desc`.
	Order pulumi.StringPtrInput `pulumi:"order"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to assign the sorted value to, by default `field` is updated in-place
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (ElasticstackElasticsearchIngestProcessorSortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorSortArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorSort.
type ElasticstackElasticsearchIngestProcessorSortResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorSortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorSortResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) ToElasticstackElasticsearchIngestProcessorSortResultOutput() ElasticstackElasticsearchIngestProcessorSortResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) ToElasticstackElasticsearchIngestProcessorSortResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorSortResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to be sorted
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The sort order to use. Accepts `asc` or `desc`.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the sorted value to, by default `field` is updated in-place
func (o ElasticstackElasticsearchIngestProcessorSortResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorSortResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorSortResultOutput{})
}