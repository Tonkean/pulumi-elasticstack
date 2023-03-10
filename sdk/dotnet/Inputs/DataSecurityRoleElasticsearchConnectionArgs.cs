// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class DataSecurityRoleElasticsearchConnectionInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// API Key to use for authentication to Elasticsearch
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// PEM-encoded custom Certificate Authority certificate
        /// </summary>
        [Input("caData")]
        public Input<string>? CaData { get; set; }

        /// <summary>
        /// Path to a custom Certificate Authority certificate
        /// </summary>
        [Input("caFile")]
        public Input<string>? CaFile { get; set; }

        /// <summary>
        /// PEM encoded certificate for client auth
        /// </summary>
        [Input("certData")]
        public Input<string>? CertData { get; set; }

        /// <summary>
        /// Path to a file containing the PEM encoded certificate for client auth
        /// </summary>
        [Input("certFile")]
        public Input<string>? CertFile { get; set; }

        [Input("endpoints")]
        private InputList<string>? _endpoints;
        public InputList<string> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<string>());
                _endpoints = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Disable TLS certificate validation
        /// </summary>
        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        [Input("keyData")]
        private Input<string>? _keyData;

        /// <summary>
        /// PEM encoded private key for client auth
        /// </summary>
        public Input<string>? KeyData
        {
            get => _keyData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Path to a file containing the PEM encoded private key for client auth
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password to use for API authentication to Elasticsearch.
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

        /// <summary>
        /// Username to use for API authentication to Elasticsearch.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public DataSecurityRoleElasticsearchConnectionInputArgs()
        {
        }
        public static new DataSecurityRoleElasticsearchConnectionInputArgs Empty => new DataSecurityRoleElasticsearchConnectionInputArgs();
    }
}
