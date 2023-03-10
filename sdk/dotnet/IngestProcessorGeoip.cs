// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorGeoip
    {
        /// <summary>
        /// The geoip processor adds information about the geographical location of an IPv4 or IPv6 address.
        /// 
        /// By default, the processor uses the GeoLite2 City, GeoLite2 Country, and GeoLite2 ASN GeoIP2 databases from MaxMind, shared under the CC BY-SA 4.0 license. Elasticsearch automatically downloads updates for these databases from the Elastic GeoIP endpoint: https://geoip.elastic.co/v1/database. To get download statistics for these updates, use the GeoIP stats API.
        /// 
        /// If your cluster can’t connect to the Elastic GeoIP endpoint or you want to manage your own updates, [see Manage your own GeoIP2 database updates](https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html#manage-geoip-database-updates).
        /// 
        /// If Elasticsearch can’t connect to the endpoint for 30 days all updated databases will become invalid. Elasticsearch will stop enriching documents with geoip data and will add tags: ["_geoip_expired_database"] field instead.
        /// 
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
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
        ///     var geoip = Elasticstack.IngestProcessorGeoip.Invoke(new()
        ///     {
        ///         Field = "ip",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             geoip.Apply(ingestProcessorGeoipResult =&gt; ingestProcessorGeoipResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorGeoipResult> InvokeAsync(IngestProcessorGeoipArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorGeoipResult>("elasticstack:index/ingestProcessorGeoip:IngestProcessorGeoip", args ?? new IngestProcessorGeoipArgs(), options.WithDefaults());

        /// <summary>
        /// The geoip processor adds information about the geographical location of an IPv4 or IPv6 address.
        /// 
        /// By default, the processor uses the GeoLite2 City, GeoLite2 Country, and GeoLite2 ASN GeoIP2 databases from MaxMind, shared under the CC BY-SA 4.0 license. Elasticsearch automatically downloads updates for these databases from the Elastic GeoIP endpoint: https://geoip.elastic.co/v1/database. To get download statistics for these updates, use the GeoIP stats API.
        /// 
        /// If your cluster can’t connect to the Elastic GeoIP endpoint or you want to manage your own updates, [see Manage your own GeoIP2 database updates](https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html#manage-geoip-database-updates).
        /// 
        /// If Elasticsearch can’t connect to the endpoint for 30 days all updated databases will become invalid. Elasticsearch will stop enriching documents with geoip data and will add tags: ["_geoip_expired_database"] field instead.
        /// 
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
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
        ///     var geoip = Elasticstack.IngestProcessorGeoip.Invoke(new()
        ///     {
        ///         Field = "ip",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             geoip.Apply(ingestProcessorGeoipResult =&gt; ingestProcessorGeoipResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorGeoipResult> Invoke(IngestProcessorGeoipInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorGeoipResult>("elasticstack:index/ingestProcessorGeoip:IngestProcessorGeoip", args ?? new IngestProcessorGeoipInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorGeoipArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
        /// </summary>
        [Input("databaseFile")]
        public string? DatabaseFile { get; set; }

        /// <summary>
        /// The field to get the ip address from for the geographical lookup.
        /// </summary>
        [Input("field", required: true)]
        public string Field { get; set; } = null!;

        /// <summary>
        /// If `true` only first found geoip data will be returned, even if field contains array.
        /// </summary>
        [Input("firstOnly")]
        public bool? FirstOnly { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public bool? IgnoreMissing { get; set; }

        [Input("properties")]
        private List<string>? _properties;

        /// <summary>
        /// Controls what properties are added to the `target_field` based on the geoip lookup.
        /// </summary>
        public List<string> Properties
        {
            get => _properties ?? (_properties = new List<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The field that will hold the geographical information looked up from the MaxMind database.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        public IngestProcessorGeoipArgs()
        {
        }
        public static new IngestProcessorGeoipArgs Empty => new IngestProcessorGeoipArgs();
    }

    public sealed class IngestProcessorGeoipInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
        /// </summary>
        [Input("databaseFile")]
        public Input<string>? DatabaseFile { get; set; }

        /// <summary>
        /// The field to get the ip address from for the geographical lookup.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// If `true` only first found geoip data will be returned, even if field contains array.
        /// </summary>
        [Input("firstOnly")]
        public Input<bool>? FirstOnly { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public Input<bool>? IgnoreMissing { get; set; }

        [Input("properties")]
        private InputList<string>? _properties;

        /// <summary>
        /// Controls what properties are added to the `target_field` based on the geoip lookup.
        /// </summary>
        public InputList<string> Properties
        {
            get => _properties ?? (_properties = new InputList<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The field that will hold the geographical information looked up from the MaxMind database.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        public IngestProcessorGeoipInvokeArgs()
        {
        }
        public static new IngestProcessorGeoipInvokeArgs Empty => new IngestProcessorGeoipInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorGeoipResult
    {
        /// <summary>
        /// The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
        /// </summary>
        public readonly string? DatabaseFile;
        /// <summary>
        /// The field to get the ip address from for the geographical lookup.
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// If `true` only first found geoip data will be returned, even if field contains array.
        /// </summary>
        public readonly bool? FirstOnly;
        /// <summary>
        /// Internal identifier of the resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// If `true` and `field` does not exist, the processor quietly exits without modifying the document.
        /// </summary>
        public readonly bool? IgnoreMissing;
        /// <summary>
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// Controls what properties are added to the `target_field` based on the geoip lookup.
        /// </summary>
        public readonly ImmutableArray<string> Properties;
        /// <summary>
        /// The field that will hold the geographical information looked up from the MaxMind database.
        /// </summary>
        public readonly string? TargetField;

        [OutputConstructor]
        private IngestProcessorGeoipResult(
            string? databaseFile,

            string field,

            bool? firstOnly,

            string id,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> properties,

            string? targetField)
        {
            DatabaseFile = databaseFile;
            Field = field;
            FirstOnly = firstOnly;
            Id = id;
            IgnoreMissing = ignoreMissing;
            Json = json;
            Properties = properties;
            TargetField = targetField;
        }
    }
}
