// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.
//
// The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.
//
// First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html
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
//			dateIndexName, err := elasticstack.IngestProcessorDateIndexName(ctx, &elasticstack.IngestProcessorDateIndexNameArgs{
//				Description:     pulumi.StringRef("monthly date-time index naming"),
//				Field:           "date1",
//				IndexNamePrefix: pulumi.StringRef("my-index-"),
//				DateRounding:    "M",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(dateIndexName.Json),
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
func IngestProcessorDateIndexName(ctx *pulumi.Context, args *IngestProcessorDateIndexNameArgs, opts ...pulumi.InvokeOption) (*IngestProcessorDateIndexNameResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv IngestProcessorDateIndexNameResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorDateIndexName:IngestProcessorDateIndexName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorDateIndexName.
type IngestProcessorDateIndexNameArgs struct {
	// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
	DateFormats []string `pulumi:"dateFormats"`
	// How to round the date when formatting the date into the index name.
	DateRounding string `pulumi:"dateRounding"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date or timestamp from.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// The format to be used when printing the parsed date into the index name.
	IndexNameFormat *string `pulumi:"indexNameFormat"`
	// A prefix of the index name to be prepended before the printed date.
	IndexNamePrefix *string `pulumi:"indexNamePrefix"`
	// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
	Timezone *string `pulumi:"timezone"`
}

// A collection of values returned by IngestProcessorDateIndexName.
type IngestProcessorDateIndexNameResult struct {
	// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
	DateFormats []string `pulumi:"dateFormats"`
	// How to round the date when formatting the date into the index name.
	DateRounding string `pulumi:"dateRounding"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date or timestamp from.
	Field string `pulumi:"field"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// The format to be used when printing the parsed date into the index name.
	IndexNameFormat *string `pulumi:"indexNameFormat"`
	// A prefix of the index name to be prepended before the printed date.
	IndexNamePrefix *string `pulumi:"indexNamePrefix"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
	Timezone *string `pulumi:"timezone"`
}

func IngestProcessorDateIndexNameOutput(ctx *pulumi.Context, args IngestProcessorDateIndexNameOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorDateIndexNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorDateIndexNameResult, error) {
			args := v.(IngestProcessorDateIndexNameArgs)
			r, err := IngestProcessorDateIndexName(ctx, &args, opts...)
			var s IngestProcessorDateIndexNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorDateIndexNameResultOutput)
}

// A collection of arguments for invoking IngestProcessorDateIndexName.
type IngestProcessorDateIndexNameOutputArgs struct {
	// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
	DateFormats pulumi.StringArrayInput `pulumi:"dateFormats"`
	// How to round the date when formatting the date into the index name.
	DateRounding pulumi.StringInput `pulumi:"dateRounding"`
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to get the date or timestamp from.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// The format to be used when printing the parsed date into the index name.
	IndexNameFormat pulumi.StringPtrInput `pulumi:"indexNameFormat"`
	// A prefix of the index name to be prepended before the printed date.
	IndexNamePrefix pulumi.StringPtrInput `pulumi:"indexNamePrefix"`
	// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
	Locale pulumi.StringPtrInput `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
	Timezone pulumi.StringPtrInput `pulumi:"timezone"`
}

func (IngestProcessorDateIndexNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDateIndexNameArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorDateIndexName.
type IngestProcessorDateIndexNameResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorDateIndexNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorDateIndexNameResult)(nil)).Elem()
}

func (o IngestProcessorDateIndexNameResultOutput) ToIngestProcessorDateIndexNameResultOutput() IngestProcessorDateIndexNameResultOutput {
	return o
}

func (o IngestProcessorDateIndexNameResultOutput) ToIngestProcessorDateIndexNameResultOutputWithContext(ctx context.Context) IngestProcessorDateIndexNameResultOutput {
	return o
}

// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
func (o IngestProcessorDateIndexNameResultOutput) DateFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) []string { return v.DateFormats }).(pulumi.StringArrayOutput)
}

// How to round the date when formatting the date into the index name.
func (o IngestProcessorDateIndexNameResultOutput) DateRounding() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) string { return v.DateRounding }).(pulumi.StringOutput)
}

// Description of the processor.
func (o IngestProcessorDateIndexNameResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to get the date or timestamp from.
func (o IngestProcessorDateIndexNameResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o IngestProcessorDateIndexNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorDateIndexNameResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorDateIndexNameResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// The format to be used when printing the parsed date into the index name.
func (o IngestProcessorDateIndexNameResultOutput) IndexNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.IndexNameFormat }).(pulumi.StringPtrOutput)
}

// A prefix of the index name to be prepended before the printed date.
func (o IngestProcessorDateIndexNameResultOutput) IndexNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.IndexNamePrefix }).(pulumi.StringPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorDateIndexNameResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) string { return v.Json }).(pulumi.StringOutput)
}

// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
func (o IngestProcessorDateIndexNameResultOutput) Locale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.Locale }).(pulumi.StringPtrOutput)
}

// Handle failures for the processor.
func (o IngestProcessorDateIndexNameResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o IngestProcessorDateIndexNameResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
func (o IngestProcessorDateIndexNameResultOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorDateIndexNameResult) *string { return v.Timezone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorDateIndexNameResultOutput{})
}
