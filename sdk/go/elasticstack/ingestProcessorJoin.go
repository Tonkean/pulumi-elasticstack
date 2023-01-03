// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Joins each element of an array into a single string using a separator character between each element. Throws an error when the field is not an array.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/join-processor.html
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
//			join, err := elasticstack.IngestProcessorJoin(ctx, &elasticstack.IngestProcessorJoinArgs{
//				Field:     "joined_array_field",
//				Separator: "-",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(join.Json),
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
func IngestProcessorJoin(ctx *pulumi.Context, args *IngestProcessorJoinArgs, opts ...pulumi.InvokeOption) (*IngestProcessorJoinResult, error) {
	var rv IngestProcessorJoinResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorJoin:IngestProcessorJoin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorJoin.
type IngestProcessorJoinArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Field containing array values to join.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The separator character.
	Separator string `pulumi:"separator"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by IngestProcessorJoin.
type IngestProcessorJoinResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Field containing array values to join.
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
	// The separator character.
	Separator string `pulumi:"separator"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

func IngestProcessorJoinOutput(ctx *pulumi.Context, args IngestProcessorJoinOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorJoinResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorJoinResult, error) {
			args := v.(IngestProcessorJoinArgs)
			r, err := IngestProcessorJoin(ctx, &args, opts...)
			var s IngestProcessorJoinResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorJoinResultOutput)
}

// A collection of arguments for invoking IngestProcessorJoin.
type IngestProcessorJoinOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Field containing array values to join.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The separator character.
	Separator pulumi.StringInput `pulumi:"separator"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (IngestProcessorJoinOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorJoinArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorJoin.
type IngestProcessorJoinResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorJoinResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorJoinResult)(nil)).Elem()
}

func (o IngestProcessorJoinResultOutput) ToIngestProcessorJoinResultOutput() IngestProcessorJoinResultOutput {
	return o
}

func (o IngestProcessorJoinResultOutput) ToIngestProcessorJoinResultOutputWithContext(ctx context.Context) IngestProcessorJoinResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorJoinResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Field containing array values to join.
func (o IngestProcessorJoinResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o IngestProcessorJoinResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorJoinResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorJoinResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorJoinResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorJoinResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The separator character.
func (o IngestProcessorJoinResultOutput) Separator() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) string { return v.Separator }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o IngestProcessorJoinResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the converted value to, by default `field` is updated in-place.
func (o IngestProcessorJoinResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorJoinResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorJoinResultOutput{})
}
