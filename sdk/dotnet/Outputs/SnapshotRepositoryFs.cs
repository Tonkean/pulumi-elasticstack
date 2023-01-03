// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Outputs
{

    [OutputType]
    public sealed class SnapshotRepositoryFs
    {
        /// <summary>
        /// Maximum size of files in snapshots.
        /// </summary>
        public readonly string? ChunkSize;
        /// <summary>
        /// If true, metadata files, such as index mappings and settings, are compressed in snapshots.
        /// </summary>
        public readonly bool? Compress;
        /// <summary>
        /// Location of the shared filesystem used to store and retrieve snapshots.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maximum number of snapshots the repository can contain.
        /// </summary>
        public readonly int? MaxNumberOfSnapshots;
        /// <summary>
        /// Maximum snapshot restore rate per node.
        /// </summary>
        public readonly string? MaxRestoreBytesPerSec;
        /// <summary>
        /// Maximum snapshot creation rate per node.
        /// </summary>
        public readonly string? MaxSnapshotBytesPerSec;
        /// <summary>
        /// If true, the repository is read-only.
        /// </summary>
        public readonly bool? Readonly;

        [OutputConstructor]
        private SnapshotRepositoryFs(
            string? chunkSize,

            bool? compress,

            string location,

            int? maxNumberOfSnapshots,

            string? maxRestoreBytesPerSec,

            string? maxSnapshotBytesPerSec,

            bool? @readonly)
        {
            ChunkSize = chunkSize;
            Compress = compress;
            Location = location;
            MaxNumberOfSnapshots = maxNumberOfSnapshots;
            MaxRestoreBytesPerSec = maxRestoreBytesPerSec;
            MaxSnapshotBytesPerSec = maxSnapshotBytesPerSec;
            Readonly = @readonly;
        }
    }
}
