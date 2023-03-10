// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
// By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `targetField` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
//
// ## Example Usage
//
// Here is an example that adds the parsed date to the `timestamp` field based on the `initialDate` field:
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
//			date, err := elasticstack.ElasticstackElasticsearchIngestProcessorDate(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorDateArgs{
//				Field:       "initial_date",
//				TargetField: pulumi.StringRef("timestamp"),
//				Formats: []string{
//					"dd/MM/yyyy HH:mm:ss",
//				},
//				Timezone: pulumi.StringRef("Europe/Amsterdam"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(date.Json),
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
func ElasticstackElasticsearchIngestProcessorDate(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorDateArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorDateResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorDateResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorDate:ElasticstackElasticsearchIngestProcessorDate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDate.
type ElasticstackElasticsearchIngestProcessorDateArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date from.
	Field string `pulumi:"field"`
	// An array of the expected date formats.
	Formats []string `pulumi:"formats"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat *string `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField *string `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone *string `pulumi:"timezone"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDate.
type ElasticstackElasticsearchIngestProcessorDateResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date from.
	Field string `pulumi:"field"`
	// An array of the expected date formats.
	Formats []string `pulumi:"formats"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat *string `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField *string `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone *string `pulumi:"timezone"`
}

func ElasticstackElasticsearchIngestProcessorDateOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorDateOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorDateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorDateResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorDateArgs)
			r, err := ElasticstackElasticsearchIngestProcessorDate(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorDateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorDateResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDate.
type ElasticstackElasticsearchIngestProcessorDateOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to get the date from.
	Field pulumi.StringInput `pulumi:"field"`
	// An array of the expected date formats.
	Formats pulumi.StringArrayInput `pulumi:"formats"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale pulumi.StringPtrInput `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat pulumi.StringPtrInput `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone pulumi.StringPtrInput `pulumi:"timezone"`
}

func (ElasticstackElasticsearchIngestProcessorDateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDateArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDate.
type ElasticstackElasticsearchIngestProcessorDateResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorDateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDateResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) ToElasticstackElasticsearchIngestProcessorDateResultOutput() ElasticstackElasticsearchIngestProcessorDateResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) ToElasticstackElasticsearchIngestProcessorDateResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorDateResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to get the date from.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) string { return v.Field }).(pulumi.StringOutput)
}

// An array of the expected date formats.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Formats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) []string { return v.Formats }).(pulumi.StringArrayOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) string { return v.Json }).(pulumi.StringOutput)
}

// The locale to use when parsing the date, relevant when parsing month names or week days.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Locale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.Locale }).(pulumi.StringPtrOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The format to use when writing the date to `targetField`.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) OutputFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.OutputFormat }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field that will hold the parsed date.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

// The timezone to use when parsing the date.
func (o ElasticstackElasticsearchIngestProcessorDateResultOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDateResult) *string { return v.Timezone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorDateResultOutput{})
}
