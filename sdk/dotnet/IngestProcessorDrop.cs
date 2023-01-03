// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorDrop
    {
        /// <summary>
        /// Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
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
        ///     var drop = Elasticstack.IngestProcessorDrop.Invoke(new()
        ///     {
        ///         If = "ctx.network_name == 'Guest'",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             drop.Apply(ingestProcessorDropResult =&gt; ingestProcessorDropResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorDropResult> InvokeAsync(IngestProcessorDropArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorDropResult>("elasticstack:index/ingestProcessorDrop:IngestProcessorDrop", args ?? new IngestProcessorDropArgs(), options.WithDefaults());

        /// <summary>
        /// Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
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
        ///     var drop = Elasticstack.IngestProcessorDrop.Invoke(new()
        ///     {
        ///         If = "ctx.network_name == 'Guest'",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             drop.Apply(ingestProcessorDropResult =&gt; ingestProcessorDropResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorDropResult> Invoke(IngestProcessorDropInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorDropResult>("elasticstack:index/ingestProcessorDrop:IngestProcessorDrop", args ?? new IngestProcessorDropInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorDropArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

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

        public IngestProcessorDropArgs()
        {
        }
        public static new IngestProcessorDropArgs Empty => new IngestProcessorDropArgs();
    }

    public sealed class IngestProcessorDropInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        public IngestProcessorDropInvokeArgs()
        {
        }
        public static new IngestProcessorDropInvokeArgs Empty => new IngestProcessorDropInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorDropResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Internal identifier of the resource
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

        [OutputConstructor]
        private IngestProcessorDropResult(
            string? description,

            string id,

            string? @if,

            bool? ignoreFailure,

            string json,

            ImmutableArray<string> onFailures,

            string? tag)
        {
            Description = description;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            Json = json;
            OnFailures = onFailures;
            Tag = tag;
        }
    }
}