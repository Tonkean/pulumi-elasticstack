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
    /// Creates or updates an index template. Index templates define settings, mappings, and aliases that can be applied automatically to new indices. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
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
    ///     var myTemplate = new Elasticstack.IndexTemplate("myTemplate", new()
    ///     {
    ///         Priority = 42,
    ///         IndexPatterns = new[]
    ///         {
    ///             "logstash*",
    ///             "filebeat*",
    ///         },
    ///         Template = new Elasticstack.Inputs.IndexTemplateTemplateArgs
    ///         {
    ///             Aliases = new[]
    ///             {
    ///                 new Elasticstack.Inputs.IndexTemplateTemplateAliasArgs
    ///                 {
    ///                     Name = "my_template_test",
    ///                 },
    ///                 new Elasticstack.Inputs.IndexTemplateTemplateAliasArgs
    ///                 {
    ///                     Name = "another_test",
    ///                 },
    ///             },
    ///             Settings = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["number_of_shards"] = "3",
    ///             }),
    ///         },
    ///     });
    /// 
    ///     var myDataStream = new Elasticstack.IndexTemplate("myDataStream", new()
    ///     {
    ///         IndexPatterns = new[]
    ///         {
    ///             "stream*",
    ///         },
    ///         DataStream = null,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import elasticstack:index/indexTemplate:IndexTemplate my_template &lt;cluster_uuid&gt;/&lt;template_name&gt;
    /// ```
    /// </summary>
    [ElasticstackResourceType("elasticstack:index/indexTemplate:IndexTemplate")]
    public partial class IndexTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An ordered list of component template names.
        /// </summary>
        [Output("composedOfs")]
        public Output<ImmutableArray<string>> ComposedOfs { get; private set; } = null!;

        /// <summary>
        /// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
        /// </summary>
        [Output("dataStream")]
        public Output<Outputs.IndexTemplateDataStream?> DataStream { get; private set; } = null!;

        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Output("elasticsearchConnection")]
        public Output<Outputs.IndexTemplateElasticsearchConnection?> ElasticsearchConnection { get; private set; } = null!;

        /// <summary>
        /// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
        /// </summary>
        [Output("indexPatterns")]
        public Output<ImmutableArray<string>> IndexPatterns { get; private set; } = null!;

        /// <summary>
        /// Optional user metadata about the index template.
        /// </summary>
        [Output("metadata")]
        public Output<string?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Name of the index template to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Priority to determine index template precedence when a new data stream or index is created.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
        /// </summary>
        [Output("template")]
        public Output<Outputs.IndexTemplateTemplate?> Template { get; private set; } = null!;

        /// <summary>
        /// Version number used to manage index templates externally.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a IndexTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IndexTemplate(string name, IndexTemplateArgs args, CustomResourceOptions? options = null)
            : base("elasticstack:index/indexTemplate:IndexTemplate", name, args ?? new IndexTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IndexTemplate(string name, Input<string> id, IndexTemplateState? state = null, CustomResourceOptions? options = null)
            : base("elasticstack:index/indexTemplate:IndexTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IndexTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IndexTemplate Get(string name, Input<string> id, IndexTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new IndexTemplate(name, id, state, options);
        }
    }

    public sealed class IndexTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("composedOfs")]
        private InputList<string>? _composedOfs;

        /// <summary>
        /// An ordered list of component template names.
        /// </summary>
        public InputList<string> ComposedOfs
        {
            get => _composedOfs ?? (_composedOfs = new InputList<string>());
            set => _composedOfs = value;
        }

        /// <summary>
        /// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
        /// </summary>
        [Input("dataStream")]
        public Input<Inputs.IndexTemplateDataStreamArgs>? DataStream { get; set; }

        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.IndexTemplateElasticsearchConnectionArgs>? ElasticsearchConnection { get; set; }

        [Input("indexPatterns", required: true)]
        private InputList<string>? _indexPatterns;

        /// <summary>
        /// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
        /// </summary>
        public InputList<string> IndexPatterns
        {
            get => _indexPatterns ?? (_indexPatterns = new InputList<string>());
            set => _indexPatterns = value;
        }

        /// <summary>
        /// Optional user metadata about the index template.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        /// <summary>
        /// Name of the index template to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority to determine index template precedence when a new data stream or index is created.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
        /// </summary>
        [Input("template")]
        public Input<Inputs.IndexTemplateTemplateArgs>? Template { get; set; }

        /// <summary>
        /// Version number used to manage index templates externally.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public IndexTemplateArgs()
        {
        }
        public static new IndexTemplateArgs Empty => new IndexTemplateArgs();
    }

    public sealed class IndexTemplateState : global::Pulumi.ResourceArgs
    {
        [Input("composedOfs")]
        private InputList<string>? _composedOfs;

        /// <summary>
        /// An ordered list of component template names.
        /// </summary>
        public InputList<string> ComposedOfs
        {
            get => _composedOfs ?? (_composedOfs = new InputList<string>());
            set => _composedOfs = value;
        }

        /// <summary>
        /// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
        /// </summary>
        [Input("dataStream")]
        public Input<Inputs.IndexTemplateDataStreamGetArgs>? DataStream { get; set; }

        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.IndexTemplateElasticsearchConnectionGetArgs>? ElasticsearchConnection { get; set; }

        [Input("indexPatterns")]
        private InputList<string>? _indexPatterns;

        /// <summary>
        /// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
        /// </summary>
        public InputList<string> IndexPatterns
        {
            get => _indexPatterns ?? (_indexPatterns = new InputList<string>());
            set => _indexPatterns = value;
        }

        /// <summary>
        /// Optional user metadata about the index template.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        /// <summary>
        /// Name of the index template to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority to determine index template precedence when a new data stream or index is created.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
        /// </summary>
        [Input("template")]
        public Input<Inputs.IndexTemplateTemplateGetArgs>? Template { get; set; }

        /// <summary>
        /// Version number used to manage index templates externally.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public IndexTemplateState()
        {
        }
        public static new IndexTemplateState Empty => new IndexTemplateState();
    }
}