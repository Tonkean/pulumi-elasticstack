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
    'IngestProcessorFingerprintResult',
    'AwaitableIngestProcessorFingerprintResult',
    'ingest_processor_fingerprint',
    'ingest_processor_fingerprint_output',
]

@pulumi.output_type
class IngestProcessorFingerprintResult:
    """
    A collection of values returned by IngestProcessorFingerprint.
    """
    def __init__(__self__, description=None, fields=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, method=None, on_failures=None, salt=None, tag=None, target_field=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fields and not isinstance(fields, list):
            raise TypeError("Expected argument 'fields' to be a list")
        pulumi.set(__self__, "fields", fields)
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
        if method and not isinstance(method, str):
            raise TypeError("Expected argument 'method' to be a str")
        pulumi.set(__self__, "method", method)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if salt and not isinstance(salt, str):
            raise TypeError("Expected argument 'salt' to be a str")
        pulumi.set(__self__, "salt", salt)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if target_field and not isinstance(target_field, str):
            raise TypeError("Expected argument 'target_field' to be a str")
        pulumi.set(__self__, "target_field", target_field)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fields(self) -> Sequence[str]:
        """
        Array of fields to include in the fingerprint.
        """
        return pulumi.get(self, "fields")

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
        If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
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
    @pulumi.getter
    def method(self) -> Optional[str]:
        """
        The hash method used to compute the fingerprint.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter(name="onFailures")
    def on_failures(self) -> Optional[Sequence[str]]:
        """
        Handle failures for the processor.
        """
        return pulumi.get(self, "on_failures")

    @property
    @pulumi.getter
    def salt(self) -> Optional[str]:
        """
        Salt value for the hash function.
        """
        return pulumi.get(self, "salt")

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
        Output field for the fingerprint.
        """
        return pulumi.get(self, "target_field")


class AwaitableIngestProcessorFingerprintResult(IngestProcessorFingerprintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorFingerprintResult(
            description=self.description,
            fields=self.fields,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            method=self.method,
            on_failures=self.on_failures,
            salt=self.salt,
            tag=self.tag,
            target_field=self.target_field)


def ingest_processor_fingerprint(description: Optional[str] = None,
                                 fields: Optional[Sequence[str]] = None,
                                 if_: Optional[str] = None,
                                 ignore_failure: Optional[bool] = None,
                                 ignore_missing: Optional[bool] = None,
                                 method: Optional[str] = None,
                                 on_failures: Optional[Sequence[str]] = None,
                                 salt: Optional[str] = None,
                                 tag: Optional[str] = None,
                                 target_field: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorFingerprintResult:
    """
    Computes a hash of the document’s content. You can use this hash for content fingerprinting.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fingerprint-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    fingerprint = elasticstack.ingest_processor_fingerprint(fields=["user"])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[fingerprint.json])
    ```


    :param str description: Description of the processor.
    :param Sequence[str] fields: Array of fields to include in the fingerprint.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
    :param str method: The hash method used to compute the fingerprint.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str salt: Salt value for the hash function.
    :param str tag: Identifier for the processor.
    :param str target_field: Output field for the fingerprint.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['fields'] = fields
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['method'] = method
    __args__['onFailures'] = on_failures
    __args__['salt'] = salt
    __args__['tag'] = tag
    __args__['targetField'] = target_field
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorFingerprint:IngestProcessorFingerprint', __args__, opts=opts, typ=IngestProcessorFingerprintResult).value

    return AwaitableIngestProcessorFingerprintResult(
        description=__ret__.description,
        fields=__ret__.fields,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        method=__ret__.method,
        on_failures=__ret__.on_failures,
        salt=__ret__.salt,
        tag=__ret__.tag,
        target_field=__ret__.target_field)


@_utilities.lift_output_func(ingest_processor_fingerprint)
def ingest_processor_fingerprint_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                        fields: Optional[pulumi.Input[Sequence[str]]] = None,
                                        if_: Optional[pulumi.Input[Optional[str]]] = None,
                                        ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                        ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                        method: Optional[pulumi.Input[Optional[str]]] = None,
                                        on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                        salt: Optional[pulumi.Input[Optional[str]]] = None,
                                        tag: Optional[pulumi.Input[Optional[str]]] = None,
                                        target_field: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorFingerprintResult]:
    """
    Computes a hash of the document’s content. You can use this hash for content fingerprinting.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fingerprint-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    fingerprint = elasticstack.ingest_processor_fingerprint(fields=["user"])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[fingerprint.json])
    ```


    :param str description: Description of the processor.
    :param Sequence[str] fields: Array of fields to include in the fingerprint.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true`, the processor ignores any missing `fields`. If all fields are missing, the processor silently exits without modifying the document.
    :param str method: The hash method used to compute the fingerprint.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str salt: Salt value for the hash function.
    :param str tag: Identifier for the processor.
    :param str target_field: Output field for the fingerprint.
    """
    ...
