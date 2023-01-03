// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Executes another pipeline.
//
// The name of the current pipeline can be accessed from the `_ingest.pipeline` ingest metadata key.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/pipeline-processor.html
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
//			appendTags, err := elasticstack.IngestProcessorAppend(ctx, &elasticstack.IngestProcessorAppendArgs{
//				Field: "tags",
//				Values: []string{
//					"production",
//					"{{{app}}}",
//					"{{{owner}}}",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			pipelineA, err := elasticstack.NewIngestPipeline(ctx, "pipelineA", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(appendTags.Json),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fingerprint, err := elasticstack.IngestProcessorFingerprint(ctx, &elasticstack.IngestProcessorFingerprintArgs{
//				Fields: []string{
//					"owner",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			pipeline := elasticstack.IngestProcessorPipelineOutput(ctx, elasticstack.IngestProcessorPipelineOutputArgs{
//				Name: pipelineA.Name,
//			}, nil)
//			_, err = elasticstack.NewIngestPipeline(ctx, "pipelineB", &elasticstack.IngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					pipeline.ApplyT(func(pipeline elasticstack.IngestProcessorPipelineResult) (*string, error) {
//						return &pipeline.Json, nil
//					}).(pulumi.StringPtrOutput),
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
func IngestProcessorPipeline(ctx *pulumi.Context, args *IngestProcessorPipelineArgs, opts ...pulumi.InvokeOption) (*IngestProcessorPipelineResult, error) {
	var rv IngestProcessorPipelineResult
	err := ctx.Invoke("elasticstack:index/ingestProcessorPipeline:IngestProcessorPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IngestProcessorPipeline.
type IngestProcessorPipelineArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// The name of the pipeline to execute.
	Name string `pulumi:"name"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by IngestProcessorPipeline.
type IngestProcessorPipelineResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The name of the pipeline to execute.
	Name string `pulumi:"name"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func IngestProcessorPipelineOutput(ctx *pulumi.Context, args IngestProcessorPipelineOutputArgs, opts ...pulumi.InvokeOption) IngestProcessorPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IngestProcessorPipelineResult, error) {
			args := v.(IngestProcessorPipelineArgs)
			r, err := IngestProcessorPipeline(ctx, &args, opts...)
			var s IngestProcessorPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IngestProcessorPipelineResultOutput)
}

// A collection of arguments for invoking IngestProcessorPipeline.
type IngestProcessorPipelineOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// The name of the pipeline to execute.
	Name pulumi.StringInput `pulumi:"name"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (IngestProcessorPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorPipelineArgs)(nil)).Elem()
}

// A collection of values returned by IngestProcessorPipeline.
type IngestProcessorPipelineResultOutput struct{ *pulumi.OutputState }

func (IngestProcessorPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestProcessorPipelineResult)(nil)).Elem()
}

func (o IngestProcessorPipelineResultOutput) ToIngestProcessorPipelineResultOutput() IngestProcessorPipelineResultOutput {
	return o
}

func (o IngestProcessorPipelineResultOutput) ToIngestProcessorPipelineResultOutputWithContext(ctx context.Context) IngestProcessorPipelineResultOutput {
	return o
}

// Description of the processor.
func (o IngestProcessorPipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Internal identifier of the resource.
func (o IngestProcessorPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o IngestProcessorPipelineResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o IngestProcessorPipelineResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o IngestProcessorPipelineResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) string { return v.Json }).(pulumi.StringOutput)
}

// The name of the pipeline to execute.
func (o IngestProcessorPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o IngestProcessorPipelineResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o IngestProcessorPipelineResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestProcessorPipelineResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestProcessorPipelineResultOutput{})
}
