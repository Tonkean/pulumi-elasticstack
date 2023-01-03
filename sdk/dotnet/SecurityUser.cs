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
    /// Adds and updates users in the native realm. These users are commonly referred to as native users. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
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
    ///     var user = new Elasticstack.SecurityUser("user", new()
    ///     {
    ///         Username = "testuser",
    ///         PasswordHash = "$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO",
    ///         Roles = new[]
    ///         {
    ///             "kibana_user",
    ///         },
    ///         Metadata = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["env"] = "testing",
    ///             ["open"] = false,
    ///             ["number"] = 49,
    ///         }),
    ///         ElasticsearchConnection = new Elasticstack.Inputs.SecurityUserElasticsearchConnectionArgs
    ///         {
    ///             Endpoints = new[]
    ///             {
    ///                 "http://localhost:9200",
    ///             },
    ///             Username = "elastic",
    ///             Password = "changeme",
    ///         },
    ///     });
    /// 
    ///     var dev = new Elasticstack.SecurityUser("dev", new()
    ///     {
    ///         Username = "devuser",
    ///         Password = "1234567890",
    ///         Roles = new[]
    ///         {
    ///             "kibana_user",
    ///         },
    ///         Metadata = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["env"] = "testing",
    ///             ["open"] = false,
    ///             ["number"] = 49,
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import elasticstack:index/securityUser:SecurityUser user &lt;cluster_uuid&gt;/elastic
    /// ```
    /// </summary>
    [ElasticstackResourceType("elasticstack:index/securityUser:SecurityUser")]
    public partial class SecurityUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Output("elasticsearchConnection")]
        public Output<Outputs.SecurityUserElasticsearchConnection?> ElasticsearchConnection { get; private set; } = null!;

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The full name of the user.
        /// </summary>
        [Output("fullName")]
        public Output<string?> FullName { get; private set; } = null!;

        /// <summary>
        /// Arbitrary metadata that you want to associate with the user.
        /// </summary>
        [Output("metadata")]
        public Output<string> Metadata { get; private set; } = null!;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        [Output("passwordHash")]
        public Output<string?> PasswordHash { get; private set; } = null!;

        /// <summary>
        /// A set of roles the user has. The roles determine the user’s access permissions. Default is [].
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityUser(string name, SecurityUserArgs args, CustomResourceOptions? options = null)
            : base("elasticstack:index/securityUser:SecurityUser", name, args ?? new SecurityUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityUser(string name, Input<string> id, SecurityUserState? state = null, CustomResourceOptions? options = null)
            : base("elasticstack:index/securityUser:SecurityUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/Tonkean/pulumi-elasticstack",
                AdditionalSecretOutputs =
                {
                    "password",
                    "passwordHash",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityUser Get(string name, Input<string> id, SecurityUserState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityUser(name, id, state, options);
        }
    }

    public sealed class SecurityUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.SecurityUserElasticsearchConnectionArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The full name of the user.
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// Arbitrary metadata that you want to associate with the user.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// A set of roles the user has. The roles determine the user’s access permissions. Default is [].
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecurityUserArgs()
        {
        }
        public static new SecurityUserArgs Empty => new SecurityUserArgs();
    }

    public sealed class SecurityUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.SecurityUserElasticsearchConnectionGetArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The full name of the user.
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// Arbitrary metadata that you want to associate with the user.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A set of roles the user has. The roles determine the user’s access permissions. Default is [].
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecurityUserState()
        {
        }
        public static new SecurityUserState Empty => new SecurityUserState();
    }
}
