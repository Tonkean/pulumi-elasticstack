// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class IngestProcessorSetSecurityUser
    {
        /// <summary>
        /// Sets user-related details (such as `username`, `roles`, `email`, `full_name`, `metadata`, `api_key`, `realm` and `authentication_typ`e) from the current authenticated user to the current document by pre-processing the ingest. The `api_key` property exists only if the user authenticates with an API key. It is an object containing the id, name and metadata (if it exists and is non-empty) fields of the API key. The realm property is also an object with two fields, name and type. When using API key authentication, the realm property refers to the realm from which the API key is created. The `authentication_type property` is a string that can take value from `REALM`, `API_KEY`, `TOKEN` and `ANONYMOUS`.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-node-set-security-user-processor.html
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
        ///     var user = Elasticstack.IngestProcessorSetSecurityUser.Invoke(new()
        ///     {
        ///         Field = "user",
        ///         Properties = new[]
        ///         {
        ///             "username",
        ///             "realm",
        ///         },
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             user.Apply(ingestProcessorSetSecurityUserResult =&gt; ingestProcessorSetSecurityUserResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IngestProcessorSetSecurityUserResult> InvokeAsync(IngestProcessorSetSecurityUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IngestProcessorSetSecurityUserResult>("elasticstack:index/ingestProcessorSetSecurityUser:IngestProcessorSetSecurityUser", args ?? new IngestProcessorSetSecurityUserArgs(), options.WithDefaults());

        /// <summary>
        /// Sets user-related details (such as `username`, `roles`, `email`, `full_name`, `metadata`, `api_key`, `realm` and `authentication_typ`e) from the current authenticated user to the current document by pre-processing the ingest. The `api_key` property exists only if the user authenticates with an API key. It is an object containing the id, name and metadata (if it exists and is non-empty) fields of the API key. The realm property is also an object with two fields, name and type. When using API key authentication, the realm property refers to the realm from which the API key is created. The `authentication_type property` is a string that can take value from `REALM`, `API_KEY`, `TOKEN` and `ANONYMOUS`.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-node-set-security-user-processor.html
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
        ///     var user = Elasticstack.IngestProcessorSetSecurityUser.Invoke(new()
        ///     {
        ///         Field = "user",
        ///         Properties = new[]
        ///         {
        ///             "username",
        ///             "realm",
        ///         },
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             user.Apply(ingestProcessorSetSecurityUserResult =&gt; ingestProcessorSetSecurityUserResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IngestProcessorSetSecurityUserResult> Invoke(IngestProcessorSetSecurityUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IngestProcessorSetSecurityUserResult>("elasticstack:index/ingestProcessorSetSecurityUser:IngestProcessorSetSecurityUser", args ?? new IngestProcessorSetSecurityUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class IngestProcessorSetSecurityUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to store the user information into.
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

        [Input("properties")]
        private List<string>? _properties;

        /// <summary>
        /// Controls what user related properties are added to the `field`.
        /// </summary>
        public List<string> Properties
        {
            get => _properties ?? (_properties = new List<string>());
            set => _properties = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public IngestProcessorSetSecurityUserArgs()
        {
        }
        public static new IngestProcessorSetSecurityUserArgs Empty => new IngestProcessorSetSecurityUserArgs();
    }

    public sealed class IngestProcessorSetSecurityUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to store the user information into.
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

        [Input("properties")]
        private InputList<string>? _properties;

        /// <summary>
        /// Controls what user related properties are added to the `field`.
        /// </summary>
        public InputList<string> Properties
        {
            get => _properties ?? (_properties = new InputList<string>());
            set => _properties = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public IngestProcessorSetSecurityUserInvokeArgs()
        {
        }
        public static new IngestProcessorSetSecurityUserInvokeArgs Empty => new IngestProcessorSetSecurityUserInvokeArgs();
    }


    [OutputType]
    public sealed class IngestProcessorSetSecurityUserResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to store the user information into.
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
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Controls what user related properties are added to the `field`.
        /// </summary>
        public readonly ImmutableArray<string> Properties;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private IngestProcessorSetSecurityUserResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            string json,

            ImmutableArray<string> onFailures,

            ImmutableArray<string> properties,

            string? tag)
        {
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            Json = json;
            OnFailures = onFailures;
            Properties = properties;
            Tag = tag;
        }
    }
}