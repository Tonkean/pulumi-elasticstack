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
//	"github.com/Tonkean/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kv, err := elasticstack.IngestProcessorKv(ctx, &elasticstack.IngestProcessorKvArgs{
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
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
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
func IngestProcessorKv(ctx *pulumi.Context, args *IngestProcessorKvArgs, opts ...pulumi.InvokeOption) (*IngestProcessorKvResult, error) {
	var rv IngestProcessorKvResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorKv:IngestProcessorKv", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorKv.
type IngestProcessorKvArgs struct {
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

// A collection of values returned by IngestProcessorKv.
type IngestProcessorKvResult struct {
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

func IngestProcessorKvOutput(ctx *pulumi.Context, args IngestProcessorKvOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorKvResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorKvResult, error) {
			args := v.(IngestProcessorKvArgs)
			r, err := IngestProcessorKv(ctx, &args, opts...)
			var s IngestProcessorKvResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorKvResultOutput)
}

// A collection of arguments for invoking IngestProcessorKv.
type IngestProcessorKvOutputArgs struct {
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

func (IngestProcessorKvOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorKvArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorKv.
type IngestProcessorKvResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorKvResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorKvResult)(nil)).Elem()
}

func (o IngestProcessorKvResultOutput) ToIngestProcessorKvResultOutput() IngestProcessorKvResultOutput {
	return o
}

func (o IngestProcessorKvResultOutput) ToIngestProcessorKvResultOutputWithContext(ctx context.Context) IngestProcessorKvResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorKvResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// List of keys to exclude from document
func (o IngestProcessorKvResultOutput) ExcludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) []string { return v.ExcludeKeys }).(pulumi.StringArrayOutput)
}

// The field to be parsed. Supports template snippets.
func (o IngestProcessorKvResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) string { return v.Field }).(pulumi.StringOutput)
}

// Regex pattern to use for splitting key-value pairs.
func (o IngestProcessorKvResultOutput) FieldSplit() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) string { return v.FieldSplit }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o IngestProcessorKvResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorKvResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorKvResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o IngestProcessorKvResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// List of keys to filter and insert into document. Defaults to including all keys
func (o IngestProcessorKvResultOutput) IncludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) []string { return v.IncludeKeys }).(pulumi.StringArrayOutput)
}

// JSON representation of this data source.
func (o IngestProcessorKvResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorKvResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Prefix to be added to extracted keys.
func (o IngestProcessorKvResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
func (o IngestProcessorKvResultOutput) StripBrackets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *bool { return v.StripBrackets }).(pulumi.BoolPtrOutput)
}

// Identifier for the processor.
func (o IngestProcessorKvResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to insert the extracted keys into. Defaults to the root of the document.
func (o IngestProcessorKvResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted keys.
func (o IngestProcessorKvResultOutput) TrimKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.TrimKey }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted values.
func (o IngestProcessorKvResultOutput) TrimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) *string { return v.TrimValue }).(pulumi.StringPtrOutput)
}

// Regex pattern to use for splitting the key from the value within a key-value pair.
func (o IngestProcessorKvResultOutput) ValueSplit() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorKvResult) string { return v.ValueSplit }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorKvResultOutput{})
}
