# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'IngestProcessorCommunityIdResult',
    'AwaitableIngestProcessorCommunityIdResult',
    'ingest_processor_community_id',
    'ingest_processor_community_id_output',
]

@pulumi.output_type
class IngestProcessorCommunityIdResult:
    """
    A collection of values returned by IngestProcessorCommunityId.
    """
    def __init__(__self__, description=None, destination_ip=None, destination_port=None, iana_number=None, icmp_code=None, icmp_type=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, on_failures=None, seed=None, source_ip=None, source_port=None, tag=None, target_field=None, transport=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination_ip and not isinstance(destination_ip, str):
            raise TypeError("Expected argument 'destination_ip' to be a str")
        pulumi.set(__self__, "destination_ip", destination_ip)
        if destination_port and not isinstance(destination_port, int):
            raise TypeError("Expected argument 'destination_port' to be a int")
        pulumi.set(__self__, "destination_port", destination_port)
        if iana_number and not isinstance(iana_number, int):
            raise TypeError("Expected argument 'iana_number' to be a int")
        pulumi.set(__self__, "iana_number", iana_number)
        if icmp_code and not isinstance(icmp_code, int):
            raise TypeError("Expected argument 'icmp_code' to be a int")
        pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type and not isinstance(icmp_type, int):
            raise TypeError("Expected argument 'icmp_type' to be a int")
        pulumi.set(__self__, "icmp_type", icmp_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if if_ and not isinstance(if_, str):
            raise TypeError("Expected argument 'if_' to be a str")
        pulumi.set(__self__, "if_", if_)
        if ignore_failure and not isinstance(ignore_failure, bool):
            raise TypeError("Expected argument 'ignore_failure' to be a bool")
        pulumi.set(__self__, "ignore_failure", ignore_failure)
        if ignore_missing and not isinstance(ignore_missing, bool):
            raise TypeError("Expected argument 'ignore_missing' to be a bool")
        pulumi.set(__self__, "ignore_missing", ignore_missing)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if seed and not isinstance(seed, int):
            raise TypeError("Expected argument 'seed' to be a int")
        pulumi.set(__self__, "seed", seed)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if source_port and not isinstance(source_port, int):
            raise TypeError("Expected argument 'source_port' to be a int")
        pulumi.set(__self__, "source_port", source_port)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if target_field and not isinstance(target_field, str):
            raise TypeError("Expected argument 'target_field' to be a str")
        pulumi.set(__self__, "target_field", target_field)
        if transport and not isinstance(transport, str):
            raise TypeError("Expected argument 'transport' to be a str")
        pulumi.set(__self__, "transport", transport)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationIp")
    def destination_ip(self) -> Optional[str]:
        """
        Field containing the destination IP address.
        """
        return pulumi.get(self, "destination_ip")

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[int]:
        """
        Field containing the destination port.
        """
        return pulumi.get(self, "destination_port")

    @property
    @pulumi.getter(name="ianaNumber")
    def iana_number(self) -> Optional[int]:
        """
        Field containing the IANA number.
        """
        return pulumi.get(self, "iana_number")

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[int]:
        """
        Field containing the ICMP code.
        """
        return pulumi.get(self, "icmp_code")

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[int]:
        """
        Field containing the ICMP type.
        """
        return pulumi.get(self, "icmp_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Internal identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="if")
    def if_(self) -> Optional[str]:
        """
        Conditionally execute the processor
        """
        return pulumi.get(self, "if_")

    @property
    @pulumi.getter(name="ignoreFailure")
    def ignore_failure(self) -> Optional[bool]:
        """
        Ignore failures for the processor.
        """
        return pulumi.get(self, "ignore_failure")

    @property
    @pulumi.getter(name="ignoreMissing")
    def ignore_missing(self) -> Optional[bool]:
        """
        If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        """
        return pulumi.get(self, "ignore_missing")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        JSON representation of this data source.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="onFailures")
    def on_failures(self) -> Optional[Sequence[str]]:
        """
        Handle failures for the processor.
        """
        return pulumi.get(self, "on_failures")

    @property
    @pulumi.getter
    def seed(self) -> Optional[int]:
        """
        Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
        """
        return pulumi.get(self, "seed")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[str]:
        """
        Field containing the source IP address.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[int]:
        """
        Field containing the source port.
        """
        return pulumi.get(self, "source_port")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="targetField")
    def target_field(self) -> Optional[str]:
        """
        Output field for the community ID.
        """
        return pulumi.get(self, "target_field")

    @property
    @pulumi.getter
    def transport(self) -> Optional[str]:
        """
        Field containing the transport protocol. Used only when the `iana_number` field is not present.
        """
        return pulumi.get(self, "transport")


class AwaitableIngestProcessorCommunityIdResult(IngestProcessorCommunityIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorCommunityIdResult(
            description=self.description,
            destination_ip=self.destination_ip,
            destination_port=self.destination_port,
            iana_number=self.iana_number,
            icmp_code=self.icmp_code,
            icmp_type=self.icmp_type,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            on_failures=self.on_failures,
            seed=self.seed,
            source_ip=self.source_ip,
            source_port=self.source_port,
            tag=self.tag,
            target_field=self.target_field,
            transport=self.transport)


def ingest_processor_community_id(description: Optional[str] = None,
                                  destination_ip: Optional[str] = None,
                                  destination_port: Optional[int] = None,
                                  iana_number: Optional[int] = None,
                                  icmp_code: Optional[int] = None,
                                  icmp_type: Optional[int] = None,
                                  if_: Optional[str] = None,
                                  ignore_failure: Optional[bool] = None,
                                  ignore_missing: Optional[bool] = None,
                                  on_failures: Optional[Sequence[str]] = None,
                                  seed: Optional[int] = None,
                                  source_ip: Optional[str] = None,
                                  source_port: Optional[int] = None,
                                  tag: Optional[str] = None,
                                  target_field: Optional[str] = None,
                                  transport: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorCommunityIdResult:
    """
    Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec).
    You can use a community ID to correlate network events related to a single flow.

    The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    community = elasticstack.ingest_processor_community_id()
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[community.json])
    ```


    :param str description: Description of the processor.
    :param str destination_ip: Field containing the destination IP address.
    :param int destination_port: Field containing the destination port.
    :param int iana_number: Field containing the IANA number.
    :param int icmp_code: Field containing the ICMP code.
    :param int icmp_type: Field containing the ICMP type.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param int seed: Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
    :param str source_ip: Field containing the source IP address.
    :param int source_port: Field containing the source port.
    :param str tag: Identifier for the processor.
    :param str target_field: Output field for the community ID.
    :param str transport: Field containing the transport protocol. Used only when the `iana_number` field is not present.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['destinationIp'] = destination_ip
    __args__['destinationPort'] = destination_port
    __args__['ianaNumber'] = iana_number
    __args__['icmpCode'] = icmp_code
    __args__['icmpType'] = icmp_type
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['onFailures'] = on_failures
    __args__['seed'] = seed
    __args__['sourceIp'] = source_ip
    __args__['sourcePort'] = source_port
    __args__['tag'] = tag
    __args__['targetField'] = target_field
    __args__['transport'] = transport
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorCommunityId:IngestProcessorCommunityId', __args__, opts=opts, typ=IngestProcessorCommunityIdResult).value

    return AwaitableIngestProcessorCommunityIdResult(
        description=__ret__.description,
        destination_ip=__ret__.destination_ip,
        destination_port=__ret__.destination_port,
        iana_number=__ret__.iana_number,
        icmp_code=__ret__.icmp_code,
        icmp_type=__ret__.icmp_type,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        on_failures=__ret__.on_failures,
        seed=__ret__.seed,
        source_ip=__ret__.source_ip,
        source_port=__ret__.source_port,
        tag=__ret__.tag,
        target_field=__ret__.target_field,
        transport=__ret__.transport)


@_utilities.lift_output_func(ingest_processor_community_id)
def ingest_processor_community_id_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                         destination_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                         destination_port: Optional[pulumi.Input[Optional[int]]] = None,
                                         iana_number: Optional[pulumi.Input[Optional[int]]] = None,
                                         icmp_code: Optional[pulumi.Input[Optional[int]]] = None,
                                         icmp_type: Optional[pulumi.Input[Optional[int]]] = None,
                                         if_: Optional[pulumi.Input[Optional[str]]] = None,
                                         ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                         ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                         on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                         seed: Optional[pulumi.Input[Optional[int]]] = None,
                                         source_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                         source_port: Optional[pulumi.Input[Optional[int]]] = None,
                                         tag: Optional[pulumi.Input[Optional[str]]] = None,
                                         target_field: Optional[pulumi.Input[Optional[str]]] = None,
                                         transport: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorCommunityIdResult]:
    """
    Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec).
    You can use a community ID to correlate network events related to a single flow.

    The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    community = elasticstack.ingest_processor_community_id()
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[community.json])
    ```


    :param str description: Description of the processor.
    :param str destination_ip: Field containing the destination IP address.
    :param int destination_port: Field containing the destination port.
    :param int iana_number: Field containing the IANA number.
    :param int icmp_code: Field containing the ICMP code.
    :param int icmp_type: Field containing the ICMP type.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param int seed: Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
    :param str source_ip: Field containing the source IP address.
    :param int source_port: Field containing the source port.
    :param str tag: Identifier for the processor.
    :param str target_field: Output field for the community ID.
    :param str transport: Field containing the transport protocol. Used only when the `iana_number` field is not present.
    """
    ...
