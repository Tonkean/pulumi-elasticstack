// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorRegisteredDomain
    {
        /// <summary>
        /// Extracts the registered domain (also known as the effective top-level domain or eTLD), sub-domain, and top-level domain from a fully qualified domain name (FQDN). Uses the registered domains defined in the Mozilla Public Suffix List.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/registered-domain-processor.html
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
        ///     var domain = Elasticstack.IngestProcessorRegisteredDomain.Invoke(new()
        ///     {
        ///         Field = "fqdn",
        ///         TargetField = "url",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             domain.Apply(ingestProcessorRegisteredDomainResult =&gt; ingestProcessorRegisteredDomainResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorRegisteredDomainResult> InvokeAsync(IngestProcessorRegisteredDomainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorRegisteredDomainResult>("elasticstack:index/ingestProcessorRegisteredDomain:IngestProcessorRegisteredDomain", args ?? new IngestProcessorRegisteredDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Extracts the registered domain (also known as the effective top-level domain or eTLD), sub-domain, and top-level domain from a fully qualified domain name (FQDN). Uses the registered domains defined in the Mozilla Public Suffix List.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/registered-domain-processor.html
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
        ///     var domain = Elasticstack.IngestProcessorRegisteredDomain.Invoke(new()
        ///     {
        ///         Field = "fqdn",
        ///         TargetField = "url",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             domain.Apply(ingestProcessorRegisteredDomainResult =&gt; ingestProcessorRegisteredDomainResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorRegisteredDomainResult> Invoke(IngestProcessorRegisteredDomainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorRegisteredDomainResult>("elasticstack:index/ingestProcessorRegisteredDomain:IngestProcessorRegisteredDomain", args ?? new IngestProcessorRegisteredDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorRegisteredDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Field containing the source FQDN.
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
        /// Object field containing extracted domain components. If an `&lt;empty string&gt;`, the processor adds components to the document’s root.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        public IngestProcessorRegisteredDomainArgs()
        {
        }
        public static new IngestProcessorRegisteredDomainArgs Empty => new IngestProcessorRegisteredDomainArgs();
    }

    public sealed class IngestProcessorRegisteredDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field containing the source FQDN.
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
        /// Object field containing extracted domain components. If an `&lt;empty string&gt;`, the processor adds components to the document’s root.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        public IngestProcessorRegisteredDomainInvokeArgs()
        {
        }
        public static new IngestProcessorRegisteredDomainInvokeArgs Empty => new IngestProcessorRegisteredDomainInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorRegisteredDomainResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Field containing the source FQDN.
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
        /// Object field containing extracted domain components. If an `&lt;empty string&gt;`, the processor adds components to the document’s root.
        /// </summary>
        public readonly string? TargetField;

        [OutputConstructor]
        private IngestProcessorRegisteredDomainResult(
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