// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    /// <summary>
    /// Manages data streams. This resource can create, delete and show the information about the created data stream. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/data-stream-apis.html
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Elasticstack = Pulumi.Elasticstack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an ILM policy for our data stream
    ///     var myIlm = new Elasticstack.IndexLifecycle("myIlm", new()
    ///     {
    ///         Hot = new Elasticstack.Inputs.IndexLifecycleHotArgs
    ///         {
    ///             MinAge = "1h",
    ///             SetPriority = new Elasticstack.Inputs.IndexLifecycleHotSetPriorityArgs
    ///             {
    ///                 Priority = 10,
    ///             },
    ///             Rollover = new Elasticstack.Inputs.IndexLifecycleHotRolloverArgs
    ///             {
    ///                 MaxAge = "1d",
    ///             },
    ///             Readonly = null,
    ///         },
    ///         Delete = new Elasticstack.Inputs.IndexLifecycleDeleteArgs
    ///         {
    ///             MinAge = "2d",
    ///             Delete = null,
    ///         },
    ///     });
    /// 
    ///     // First we must have a index template created
    ///     var myDataStreamTemplate = new Elasticstack.IndexTemplate("myDataStreamTemplate", new()
    ///     {
    ///         IndexPatterns = new[]
    ///         {
    ///             "my-stream*",
    ///         },
    ///         Template = new Elasticstack.Inputs.IndexTemplateTemplateArgs
    ///         {
    ///             Settings = myIlm.Name.Apply(name =&gt; JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["lifecycle.name"] = name,
    ///             })),
    ///         },
    ///         DataStream = null,
    ///     });
    /// 
    ///     // and now we can create data stream based on the index template
    ///     var myDataStream = new Elasticstack.DataStream("myDataStream", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             myDataStreamTemplate,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import elasticstack:index/dataStream:DataStream my_data_stream &lt;cluster_uuid&gt;/&lt;data_stream_name&gt;
    /// ```
    /// </summary>
    [ElasticstackResourceType("elasticstack:index/dataStream:DataStream")]
    public partial class DataStream : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Output("elasticsearchConnection")]
        public Output<Outputs.DataStreamElasticsearchConnection?> ElasticsearchConnection { get; private set; } = null!;

        /// <summary>
        /// Current generation for the data stream.
        /// </summary>
        [Output("generation")]
        public Output<int> Generation { get; private set; } = null!;

        /// <summary>
        /// If `true`, the data stream is hidden.
        /// </summary>
        [Output("hidden")]
        public Output<bool> Hidden { get; private set; } = null!;

        /// <summary>
        /// Name of the current ILM lifecycle policy in the stream’s matching index template.
        /// </summary>
        [Output("ilmPolicy")]
        public Output<string> IlmPolicy { get; private set; } = null!;

        /// <summary>
        /// Array of objects containing information about the data stream’s backing indices. The last item in this array contains information about the stream’s current write index.
        /// </summary>
        [Output("indices")]
        public Output<ImmutableArray<Outputs.DataStreamIndex>> Indices { get; private set; } = null!;

        /// <summary>
        /// Custom metadata for the stream, copied from the _meta object of the stream’s matching index template.
        /// </summary>
        [Output("metadata")]
        public Output<string> Metadata { get; private set; } = null!;

        /// <summary>
        /// Name of the data stream to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If `true`, the data stream is created and managed by cross-cluster replication and the local cluster can not write into this data stream or change its mappings.
        /// </summary>
        [Output("replicated")]
        public Output<bool> Replicated { get; private set; } = null!;

        /// <summary>
        /// Health status of the data stream.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// If `true`, the data stream is created and managed by an Elastic stack component and cannot be modified through normal user interaction.
        /// </summary>
        [Output("system")]
        public Output<bool> System { get; private set; } = null!;

        /// <summary>
        /// Name of the index template used to create the data stream’s backing indices.
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;

        /// <summary>
        /// Contains information about the data stream’s @timestamp field.
        /// </summary>
        [Output("timestampField")]
        public Output<string> TimestampField { get; private set; } = null!;


        /// <summary>
        /// Create a DataStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataStream(string name, DataStreamArgs? args = null, CustomResourceOptions? options = null)
            : base("elasticstack:index/dataStream:DataStream", name, args ?? new DataStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataStream(string name, Input<string> id, DataStreamState? state = null, CustomResourceOptions? options = null)
            : base("elasticstack:index/dataStream:DataStream", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/Tonkean/pulumi-elasticstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataStream Get(string name, Input<string> id, DataStreamState? state = null, CustomResourceOptions? options = null)
        {
            return new DataStream(name, id, state, options);
        }
    }

    public sealed class DataStreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.DataStreamElasticsearchConnectionArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// Name of the data stream to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DataStreamArgs()
        {
        }
        public static new DataStreamArgs Empty => new DataStreamArgs();
    }

    public sealed class DataStreamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.DataStreamElasticsearchConnectionGetArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// Current generation for the data stream.
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        /// <summary>
        /// If `true`, the data stream is hidden.
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        /// <summary>
        /// Name of the current ILM lifecycle policy in the stream’s matching index template.
        /// </summary>
        [Input("ilmPolicy")]
        public Input<string>? IlmPolicy { get; set; }

        [Input("indices")]
        private InputList<Inputs.DataStreamIndexGetArgs>? _indices;

        /// <summary>
        /// Array of objects containing information about the data stream’s backing indices. The last item in this array contains information about the stream’s current write index.
        /// </summary>
        public InputList<Inputs.DataStreamIndexGetArgs> Indices
        {
            get => _indices ?? (_indices = new InputList<Inputs.DataStreamIndexGetArgs>());
            set => _indices = value;
        }

        /// <summary>
        /// Custom metadata for the stream, copied from the _meta object of the stream’s matching index template.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        /// <summary>
        /// Name of the data stream to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If `true`, the data stream is created and managed by cross-cluster replication and the local cluster can not write into this data stream or change its mappings.
        /// </summary>
        [Input("replicated")]
        public Input<bool>? Replicated { get; set; }

        /// <summary>
        /// Health status of the data stream.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// If `true`, the data stream is created and managed by an Elastic stack component and cannot be modified through normal user interaction.
        /// </summary>
        [Input("system")]
        public Input<bool>? System { get; set; }

        /// <summary>
        /// Name of the index template used to create the data stream’s backing indices.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Contains information about the data stream’s @timestamp field.
        /// </summary>
        [Input("timestampField")]
        public Input<string>? TimestampField { get; set; }

        public DataStreamState()
        {
        }
        public static new DataStreamState Empty => new DataStreamState();
    }
}
