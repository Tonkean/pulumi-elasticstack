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
    public sealed class DataSnapshotRepositoryFResult
    {
        public readonly string ChunkSize;
        public readonly bool Compress;
        public readonly string Location;
        public readonly int MaxNumberOfSnapshots;
        public readonly string MaxRestoreBytesPerSec;
        public readonly string MaxSnapshotBytesPerSec;
        public readonly bool Readonly;

        [OutputConstructor]
        private DataSnapshotRepositoryFResult(
            string chunkSize,

            bool compress,

            string location,

            int maxNumberOfSnapshots,

            string maxRestoreBytesPerSec,

            string maxSnapshotBytesPerSec,

            bool @readonly)
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