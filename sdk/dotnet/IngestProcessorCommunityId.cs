// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorCommunityId
    {
        /// <summary>
        /// Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec). 
        /// You can use a community ID to correlate network events related to a single flow.
        /// 
        /// The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html
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
        ///     var community = Elasticstack.IngestProcessorCommunityId.Invoke();
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             community.Apply(ingestProcessorCommunityIdResult =&gt; ingestProcessorCommunityIdResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorCommunityIdResult> InvokeAsync(IngestProcessorCommunityIdArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorCommunityIdResult>("elasticstack:index/ingestProcessorCommunityId:IngestProcessorCommunityId", args ?? new IngestProcessorCommunityIdArgs(), options.WithDefaults());

        /// <summary>
        /// Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec). 
        /// You can use a community ID to correlate network events related to a single flow.
        /// 
        /// The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html
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
        ///     var community = Elasticstack.IngestProcessorCommunityId.Invoke();
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             community.Apply(ingestProcessorCommunityIdResult =&gt; ingestProcessorCommunityIdResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorCommunityIdResult> Invoke(IngestProcessorCommunityIdInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorCommunityIdResult>("elasticstack:index/ingestProcessorCommunityId:IngestProcessorCommunityId", args ?? new IngestProcessorCommunityIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorCommunityIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Field containing the destination IP address.
        /// </summary>
        [Input("destinationIp")]
        public string? DestinationIp { get; set; }

        /// <summary>
        /// Field containing the destination port.
        /// </summary>
        [Input("destinationPort")]
        public int? DestinationPort { get; set; }

        /// <summary>
        /// Field containing the IANA number.
        /// </summary>
        [Input("ianaNumber")]
        public int? IanaNumber { get; set; }

        /// <summary>
        /// Field containing the ICMP code.
        /// </summary>
        [Input("icmpCode")]
        public int? IcmpCode { get; set; }

        /// <summary>
        /// Field containing the ICMP type.
        /// </summary>
        [Input("icmpType")]
        public int? IcmpType { get; set; }

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
        /// Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
        /// </summary>
        [Input("seed")]
        public int? Seed { get; set; }

        /// <summary>
        /// Field containing the source IP address.
        /// </summary>
        [Input("sourceIp")]
        public string? SourceIp { get; set; }

        /// <summary>
        /// Field containing the source port.
        /// </summary>
        [Input("sourcePort")]
        public int? SourcePort { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// Output field for the community ID.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        /// <summary>
        /// Field containing the transport protocol. Used only when the `iana_number` field is not present.
        /// </summary>
        [Input("transport")]
        public string? Transport { get; set; }

        public IngestProcessorCommunityIdArgs()
        {
        }
        public static new IngestProcessorCommunityIdArgs Empty => new IngestProcessorCommunityIdArgs();
    }

    public sealed class IngestProcessorCommunityIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field containing the destination IP address.
        /// </summary>
        [Input("destinationIp")]
        public Input<string>? DestinationIp { get; set; }

        /// <summary>
        /// Field containing the destination port.
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        /// <summary>
        /// Field containing the IANA number.
        /// </summary>
        [Input("ianaNumber")]
        public Input<int>? IanaNumber { get; set; }

        /// <summary>
        /// Field containing the ICMP code.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// Field containing the ICMP type.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

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
        /// Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
        /// </summary>
        [Input("seed")]
        public Input<int>? Seed { get; set; }

        /// <summary>
        /// Field containing the source IP address.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Field containing the source port.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Output field for the community ID.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        /// <summary>
        /// Field containing the transport protocol. Used only when the `iana_number` field is not present.
        /// </summary>
        [Input("transport")]
        public Input<string>? Transport { get; set; }

        public IngestProcessorCommunityIdInvokeArgs()
        {
        }
        public static new IngestProcessorCommunityIdInvokeArgs Empty => new IngestProcessorCommunityIdInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorCommunityIdResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Field containing the destination IP address.
        /// </summary>
        public readonly string? DestinationIp;
        /// <summary>
        /// Field containing the destination port.
        /// </summary>
        public readonly int? DestinationPort;
        /// <summary>
        /// Field containing the IANA number.
        /// </summary>
        public readonly int? IanaNumber;
        /// <summary>
        /// Field containing the ICMP code.
        /// </summary>
        public readonly int? IcmpCode;
        /// <summary>
        /// Field containing the ICMP type.
        /// </summary>
        public readonly int? IcmpType;
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
        /// Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
        /// </summary>
        public readonly int? Seed;
        /// <summary>
        /// Field containing the source IP address.
        /// </summary>
        public readonly string? SourceIp;
        /// <summary>
        /// Field containing the source port.
        /// </summary>
        public readonly int? SourcePort;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Output field for the community ID.
        /// </summary>
        public readonly string? TargetField;
        /// <summary>
        /// Field containing the transport protocol. Used only when the `iana_number` field is not present.
        /// </summary>
        public readonly string? Transport;

        [OutputConstructor]
        private IngestProcessorCommunityIdResult(
            string? description,

            string? destinationIp,

            int? destinationPort,

            int? ianaNumber,

            int? icmpCode,

            int? icmpType,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            int? seed,

            string? sourceIp,

            int? sourcePort,

            string? tag,

            string? targetField,

            string? transport)
        {
            Description = description;
            DestinationIp = destinationIp;
            DestinationPort = destinationPort;
            IanaNumber = ianaNumber;
            IcmpCode = icmpCode;
            IcmpType = icmpType;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IgnoreMissing = ignoreMissing;
            Json = json;
            OnFailures = onFailures;
            Seed = seed;
            SourceIp = sourceIp;
            SourcePort = sourcePort;
            Tag = tag;
            TargetField = targetField;
            Transport = transport;
        }
    }
}