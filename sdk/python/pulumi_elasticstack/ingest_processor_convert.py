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
    'IngestProcessorConvertResult',
    'AwaitableIngestProcessorConvertResult',
    'ingest_processor_convert',
    'ingest_processor_convert_output',
]

@pulumi.output_type
class IngestProcessorConvertResult:
    """
    A collection of values returned by IngestProcessorConvert.
    """
    def __init__(__self__, description=None, field=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, on_failures=None, tag=None, target_field=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if field and not isinstance(field, str):
            raise TypeError("Expected argument 'field' to be a str")
        pulumi.set(__self__, "field", field)
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
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if target_field and not isinstance(target_field, str):
            raise TypeError("Expected argument 'target_field' to be a str")
        pulumi.set(__self__, "target_field", target_field)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field whose value is to be converted.
        """
        return pulumi.get(self, "field")

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
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="targetField")
    def target_field(self) -> Optional[str]:
        """
        The field to assign the converted value to.
        """
        return pulumi.get(self, "target_field")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type to convert the existing value to
        """
        return pulumi.get(self, "type")


class AwaitableIngestProcessorConvertResult(IngestProcessorConvertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorConvertResult(
            description=self.description,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            on_failures=self.on_failures,
            tag=self.tag,
            target_field=self.target_field,
            type=self.type)


def ingest_processor_convert(description: Optional[str] = None,
                             field: Optional[str] = None,
                             if_: Optional[str] = None,
                             ignore_failure: Optional[bool] = None,
                             ignore_missing: Optional[bool] = None,
                             on_failures: Optional[Sequence[str]] = None,
                             tag: Optional[str] = None,
                             target_field: Optional[str] = None,
                             type: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorConvertResult:
    """
    Helper data source to which can be used to convert a field in the currently ingested document to a different type, such as converting a string to an integer. If the field value is an array, all members will be converted.

    The supported types include: `integer`, `long`, `float`, `double`, `string`, `boolean`, `ip`, and `auto`.

    Specifying `boolean` will set the field to true if its string value is equal to true (ignore case), to false if its string value is equal to false (ignore case), or it will throw an exception otherwise.

    Specifying `ip` will set the target field to the value of `field` if it contains a valid IPv4 or IPv6 address that can be indexed into an IP field type.

    Specifying `auto` will attempt to convert the string-valued `field` into the closest non-string, non-IP type. For example, a field whose value is "true" will be converted to its respective boolean type: true. Do note that float takes precedence of double in auto. A value of "242.15" will "automatically" be converted to 242.15 of type `float`. If a provided field cannot be appropriately converted, the processor will still process successfully and leave the field value as-is. In such a case, `target_field` will be updated with the unconverted field value.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/convert-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    convert = elasticstack.ingest_processor_convert(description="converts the content of the id field to an integer",
        field="id",
        type="integer")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[convert.json])
    ```


    :param str description: Description of the processor.
    :param str field: The field whose value is to be converted.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param str target_field: The field to assign the converted value to.
    :param str type: The type to convert the existing value to
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['onFailures'] = on_failures
    __args__['tag'] = tag
    __args__['targetField'] = target_field
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorConvert:IngestProcessorConvert', __args__, opts=opts, typ=IngestProcessorConvertResult).value

    return AwaitableIngestProcessorConvertResult(
        description=__ret__.description,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        on_failures=__ret__.on_failures,
        tag=__ret__.tag,
        target_field=__ret__.target_field,
        type=__ret__.type)


@_utilities.lift_output_func(ingest_processor_convert)
def ingest_processor_convert_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                    field: Optional[pulumi.Input[str]] = None,
                                    if_: Optional[pulumi.Input[Optional[str]]] = None,
                                    ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                    ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                    on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    tag: Optional[pulumi.Input[Optional[str]]] = None,
                                    target_field: Optional[pulumi.Input[Optional[str]]] = None,
                                    type: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorConvertResult]:
    """
    Helper data source to which can be used to convert a field in the currently ingested document to a different type, such as converting a string to an integer. If the field value is an array, all members will be converted.

    The supported types include: `integer`, `long`, `float`, `double`, `string`, `boolean`, `ip`, and `auto`.

    Specifying `boolean` will set the field to true if its string value is equal to true (ignore case), to false if its string value is equal to false (ignore case), or it will throw an exception otherwise.

    Specifying `ip` will set the target field to the value of `field` if it contains a valid IPv4 or IPv6 address that can be indexed into an IP field type.

    Specifying `auto` will attempt to convert the string-valued `field` into the closest non-string, non-IP type. For example, a field whose value is "true" will be converted to its respective boolean type: true. Do note that float takes precedence of double in auto. A value of "242.15" will "automatically" be converted to 242.15 of type `float`. If a provided field cannot be appropriately converted, the processor will still process successfully and leave the field value as-is. In such a case, `target_field` will be updated with the unconverted field value.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/convert-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    convert = elasticstack.ingest_processor_convert(description="converts the content of the id field to an integer",
        field="id",
        type="integer")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[convert.json])
    ```


    :param str description: Description of the processor.
    :param str field: The field whose value is to be converted.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param str target_field: The field to assign the converted value to.
    :param str type: The type to convert the existing value to
    """
    ...
