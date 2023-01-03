// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This processor helps automatically parse messages (or specific event fields) which are of the `foo=bar` variety.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/kv-processor.html
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
//			kv, err := elasticstack.ElasticstackElasticsearchIngestProcessorKv(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorKvArgs{
//				Field:      "message",
//				FieldSplit: " ",
//				ValueSplit: "=",
//				ExcludeKeys: []string{
//					"tags",
//				},
//				Prefix: pulumi.StringRef("setting_"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(kv.Json),
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
func ElasticstackElasticsearchIngestProcessorKv(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorKvArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorKvResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorKvResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorKv:ElasticstackElasticsearchIngestProcessorKv", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorKv.
type ElasticstackElasticsearchIngestProcessorKvArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys []string `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field string `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit string `pulumi:"fieldSplit"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys []string `pulumi:"includeKeys"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix *string `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets *bool `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField *string `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey *string `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue *string `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit string `pulumi:"valueSplit"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorKv.
type ElasticstackElasticsearchIngestProcessorKvResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys []string `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field string `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit string `pulumi:"fieldSplit"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys []string `pulumi:"includeKeys"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix *string `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets *bool `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField *string `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey *string `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue *string `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit string `pulumi:"valueSplit"`
}

func ElasticstackElasticsearchIngestProcessorKvOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorKvOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorKvResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorKvResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorKvArgs)
			r, err := ElasticstackElasticsearchIngestProcessorKv(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorKvResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorKvResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorKv.
type ElasticstackElasticsearchIngestProcessorKvOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys pulumi.StringArrayInput `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field pulumi.StringInput `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit pulumi.StringInput `pulumi:"fieldSplit"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys pulumi.StringArrayInput `pulumi:"includeKeys"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets pulumi.BoolPtrInput `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey pulumi.StringPtrInput `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue pulumi.StringPtrInput `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit pulumi.StringInput `pulumi:"valueSplit"`
}

func (ElasticstackElasticsearchIngestProcessorKvOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorKvArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorKv.
type ElasticstackElasticsearchIngestProcessorKvResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorKvResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorKvResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) ToElasticstackElasticsearchIngestProcessorKvResultOutput() ElasticstackElasticsearchIngestProcessorKvResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) ToElasticstackElasticsearchIngestProcessorKvResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorKvResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// List of keys to exclude from document
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) ExcludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) []string { return v.ExcludeKeys }).(pulumi.StringArrayOutput)
}

// The field to be parsed. Supports template snippets.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) string { return v.Field }).(pulumi.StringOutput)
}

// Regex pattern to use for splitting key-value pairs.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) FieldSplit() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) string { return v.FieldSplit }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// List of keys to filter and insert into document. Defaults to including all keys
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) IncludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) []string { return v.IncludeKeys }).(pulumi.StringArrayOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Prefix to be added to extracted keys.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) StripBrackets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *bool { return v.StripBrackets }).(pulumi.BoolPtrOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to insert the extracted keys into. Defaults to the root of the document.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted keys.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) TrimKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.TrimKey }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted values.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) TrimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) *string { return v.TrimValue }).(pulumi.StringPtrOutput)
}

// Regex pattern to use for splitting the key from the value within a key-value pair.
func (o ElasticstackElasticsearchIngestProcessorKvResultOutput) ValueSplit() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorKvResult) string { return v.ValueSplit }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorKvResultOutput{})
}