// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorTrim
    {
        /// <summary>
        /// Trims whitespace from field. If the field is an array of strings, all members of the array will be trimmed.
        /// 
        /// **NOTE:** This only works on leading and trailing whitespace.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/trim-processor.html
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var trim = Elasticstack.IngestProcessorTrim.Invoke(new()
        ///     {
        ///         Field = "foo",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             trim.Apply(ingestProcessorTrimResult =&gt; ingestProcessorTrimResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorTrimResult> InvokeAsync(IngestProcessorTrimArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorTrimResult>("elasticstack:index/ingestProcessorTrim:IngestProcessorTrim", args ?? new IngestProcessorTrimArgs(), options.WithDefaults());

        /// <summary>
        /// Trims whitespace from field. If the field is an array of strings, all members of the array will be trimmed.
        /// 
        /// **NOTE:** This only works on leading and trailing whitespace.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/trim-processor.html
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var trim = Elasticstack.IngestProcessorTrim.Invoke(new()
        ///     {
        ///         Field = "foo",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             trim.Apply(ingestProcessorTrimResult =&gt; ingestProcessorTrimResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorTrimResult> Invoke(IngestProcessorTrimInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorTrimResult>("elasticstack:index/ingestProcessorTrim:IngestProcessorTrim", args ?? new IngestProcessorTrimInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorTrimArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The string-valued field to trim whitespace from.
        /// </summary>
        [Input("field", required: true)]
        public string Field { get; set; } = null!;

        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        [Input("if")]
        public string? If { get; set; }

        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        [Input("ignoreFailure")]
        public bool? IgnoreFailure { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public bool? IgnoreMissing { get; set; }

        [Input("onFailures")]
        private List<string>? _onFailures;

        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public List<string> OnFailures
        {
            get => _onFailures ?? (_onFailures = new List<string>());
            set => _onFailures = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// The field to assign the trimmed value to, by default `field` is updated in-place.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        public IngestProcessorTrimArgs()
        {
        }
        public static new IngestProcessorTrimArgs Empty => new IngestProcessorTrimArgs();
    }

    public sealed class IngestProcessorTrimInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The string-valued field to trim whitespace from.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        [Input("if")]
        public Input<string>? If { get; set; }

        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        [Input("ignoreFailure")]
        public Input<bool>? IgnoreFailure { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public Input<bool>? IgnoreMissing { get; set; }

        [Input("onFailures")]
        private InputList<string>? _onFailures;

        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public InputList<string> OnFailures
        {
            get => _onFailures ?? (_onFailures = new InputList<string>());
            set => _onFailures = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The field to assign the trimmed value to, by default `field` is updated in-place.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        public IngestProcessorTrimInvokeArgs()
        {
        }
        public static new IngestProcessorTrimInvokeArgs Empty => new IngestProcessorTrimInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorTrimResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The string-valued field to trim whitespace from.
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// Internal identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        public readonly string? If;
        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        public readonly bool? IgnoreFailure;
        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        public readonly bool? IgnoreMissing;
        /// <summary>
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// The field to assign the trimmed value to, by default `field` is updated in-place.
        /// </summary>
        public readonly string? TargetField;

        [OutputConstructor]
        private IngestProcessorTrimResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            string? tag,

            string? targetField)
        {
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IgnoreMissing = ignoreMissing;
            Json = json;
            OnFailures = onFailures;
            Tag = tag;
            TargetField = targetField;
        }
    }
}