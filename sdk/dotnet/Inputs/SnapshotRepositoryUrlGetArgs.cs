// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class SnapshotRepositoryUrlGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum size of files in snapshots.
        /// </summary>
        [Input("chunkSize")]
        public Input<string>? ChunkSize { get; set; }

        /// <summary>
        /// If true, metadata files, such as index mappings and settings, are compressed in snapshots.
        /// </summary>
        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        /// <summary>
        /// Maximum number of retries for http and https URLs.
        /// </summary>
        [Input("httpMaxRetries")]
        public Input<int>? HttpMaxRetries { get; set; }

        /// <summary>
        /// Maximum wait time for data transfers over a connection.
        /// </summary>
        [Input("httpSocketTimeout")]
        public Input<string>? HttpSocketTimeout { get; set; }

        /// <summary>
        /// Maximum number of snapshots the repository can contain.
        /// </summary>
        [Input("maxNumberOfSnapshots")]
        public Input<int>? MaxNumberOfSnapshots { get; set; }

        /// <summary>
        /// Maximum snapshot restore rate per node.
        /// </summary>
        [Input("maxRestoreBytesPerSec")]
        public Input<string>? MaxRestoreBytesPerSec { get; set; }

        /// <summary>
        /// Maximum snapshot creation rate per node.
        /// </summary>
        [Input("maxSnapshotBytesPerSec")]
        public Input<string>? MaxSnapshotBytesPerSec { get; set; }

        /// <summary>
        /// If true, the repository is read-only.
        /// </summary>
        [Input("readonly")]
        public Input<bool>? Readonly { get; set; }

        /// <summary>
        /// URL location of the root of the shared filesystem repository.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public SnapshotRepositoryUrlGetArgs()
        {
        }
        public static new SnapshotRepositoryUrlGetArgs Empty => new SnapshotRepositoryUrlGetArgs();
    }
}
