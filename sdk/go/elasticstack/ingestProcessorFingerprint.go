// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Computes a hash of the document’s content. You can use this hash for content fingerprinting.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fingerprint-processor.html
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
//			fingerprint, err := elasticstack.IngestProcessorFingerprint(ctx, &elasticstack.IngestProcessorFingerprintArgs{
//				Fields: []string{
//					"user",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(fingerprint.Json),
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
func IngestProcessorFingerprint(ctx *pulumi.Context, args *IngestProcessorFingerprintArgs, opts ...pulumi.InvokeOption) (*IngestProcessorFingerprintResult, error) {
	var rv IngestProcessorFingerprintResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorFingerprint:IngestProcessorFingerprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorFingerprint.
type IngestProcessorFingerprintArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Array of fields to include in the fingerprint.
	Fields []string `pulumi:"fields"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// The hash method used to compute the fingerprint.
	Method *string `pulumi:"method"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Salt value for the hash function.
	Salt *string `pulumi:"salt"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Output field for the fingerprint.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by IngestProcessorFingerprint.
type IngestProcessorFingerprintResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Array of fields to include in the fingerprint.
	Fields []string `pulumi:"fields"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The hash method used to compute the fingerprint.
	Method *string `pulumi:"method"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Salt value for the hash function.
	Salt *string `pulumi:"salt"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Output field for the fingerprint.
	TargetField *string `pulumi:"targetField"`
}

func IngestProcessorFingerprintOutput(ctx *pulumi.Context, args IngestProcessorFingerprintOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorFingerprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorFingerprintResult, error) {
			args := v.(IngestProcessorFingerprintArgs)
			r, err := IngestProcessorFingerprint(ctx, &args, opts...)
			var s IngestProcessorFingerprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorFingerprintResultOutput)
}

// A collection of arguments for invoking IngestProcessorFingerprint.
type IngestProcessorFingerprintOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Array of fields to include in the fingerprint.
	Fields pulumi.StringArrayInput `pulumi:"fields"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// The hash method used to compute the fingerprint.
	Method pulumi.StringPtrInput `pulumi:"method"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Salt value for the hash function.
	Salt pulumi.StringPtrInput `pulumi:"salt"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// Output field for the fingerprint.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (IngestProcessorFingerprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorFingerprintArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorFingerprint.
type IngestProcessorFingerprintResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorFingerprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorFingerprintResult)(nil)).Elem()
}

func (o IngestProcessorFingerprintResultOutput) ToIngestProcessorFingerprintResultOutput() IngestProcessorFingerprintResultOutput {
	return o
}

func (o IngestProcessorFingerprintResultOutput) ToIngestProcessorFingerprintResultOutputWithContext(ctx context.Context) IngestProcessorFingerprintResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorFingerprintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Array of fields to include in the fingerprint.
func (o IngestProcessorFingerprintResultOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) []string { return v.Fields }).(pulumi.StringArrayOutput)
}

// Internal identifier of the resource
func (o IngestProcessorFingerprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorFingerprintResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorFingerprintResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
func (o IngestProcessorFingerprintResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorFingerprintResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) string { return v.Json }).(pulumi.StringOutput)
}

// The hash method used to compute the fingerprint.
func (o IngestProcessorFingerprintResultOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.Method }).(pulumi.StringPtrOutput)
}

// Handle failures for the processor.
func (o IngestProcessorFingerprintResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Salt value for the hash function.
func (o IngestProcessorFingerprintResultOutput) Salt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.Salt }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o IngestProcessorFingerprintResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// Output field for the fingerprint.
func (o IngestProcessorFingerprintResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorFingerprintResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorFingerprintResultOutput{})
}
