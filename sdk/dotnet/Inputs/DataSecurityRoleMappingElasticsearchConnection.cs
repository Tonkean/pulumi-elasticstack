// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class DataSecurityRoleMappingElasticsearchConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiKey")]
        private string? _apiKey;

        /// <summary>
        /// API Key to use for authentication to Elasticsearch
        /// </summary>
        public string? ApiKey
        {
            get => _apiKey;
            set => _apiKey = value;
        }

        /// <summary>
        /// PEM-encoded custom Certificate Authority certificate
        /// </summary>
        [Input("caData")]
        public string? CaData { get; set; }

        /// <summary>
        /// Path to a custom Certificate Authority certificate
        /// </summary>
        [Input("caFile")]
        public string? CaFile { get; set; }

        /// <summary>
        /// PEM encoded certificate for client auth
        /// </summary>
        [Input("certData")]
        public string? CertData { get; set; }

        /// <summary>
        /// Path to a file containing the PEM encoded certificate for client auth
        /// </summary>
        [Input("certFile")]
        public string? CertFile { get; set; }

        [Input("endpoints")]
        private List<string>? _endpoints;
        public List<string> Endpoints
        {
            get => _endpoints ?? (_endpoints = new List<string>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Disable TLS certificate validation
        /// </summary>
        [Input("insecure")]
        public bool? Insecure { get; set; }

        [Input("keyData")]
        private string? _keyData;

        /// <summary>
        /// PEM encoded private key for client auth
        /// </summary>
        public string? KeyData
        {
            get => _keyData;
            set => _keyData = value;
        }

        /// <summary>
        /// Path to a file containing the PEM encoded private key for client auth
        /// </summary>
        [Input("keyFile")]
        public string? KeyFile { get; set; }

        [Input("password")]
        private string? _password;

        /// <summary>
        /// Password to use for API authentication to Elasticsearch.
        /// </summary>
        public string? Password
        {
            get => _password;
            set => _password = value;
        }

        /// <summary>
        /// Username to use for API authentication to Elasticsearch.
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        public DataSecurityRoleMappingElasticsearchConnectionArgs()
        {
        }
        public static new DataSecurityRoleMappingElasticsearchConnectionArgs Empty => new DataSecurityRoleMappingElasticsearchConnectionArgs();
    }
}
