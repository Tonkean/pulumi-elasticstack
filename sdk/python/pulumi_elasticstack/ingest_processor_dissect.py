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
    'IngestProcessorDissectResult',
    'AwaitableIngestProcessorDissectResult',
    'ingest_processor_dissect',
    'ingest_processor_dissect_output',
]

@pulumi.output_type
class IngestProcessorDissectResult:
    """
    A collection of values returned by IngestProcessorDissect.
    """
    def __init__(__self__, append_separator=None, description=None, field=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, on_failures=None, pattern=None, tag=None):
        if append_separator and not isinstance(append_separator, str):
            raise TypeError("Expected argument 'append_separator' to be a str")
        pulumi.set(__self__, "append_separator", append_separator)
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
        if pattern and not isinstance(pattern, str):
            raise TypeError("Expected argument 'pattern' to be a str")
        pulumi.set(__self__, "pattern", pattern)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="appendSeparator")
    def append_separator(self) -> Optional[str]:
        """
        The character(s) that separate the appended fields.
        """
        return pulumi.get(self, "append_separator")

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
        The field to dissect.
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
    def pattern(self) -> str:
        """
        The pattern to apply to the field.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")


class AwaitableIngestProcessorDissectResult(IngestProcessorDissectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorDissectResult(
            append_separator=self.append_separator,
            description=self.description,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            on_failures=self.on_failures,
            pattern=self.pattern,
            tag=self.tag)


def ingest_processor_dissect(append_separator: Optional[str] = None,
                             description: Optional[str] = None,
                             field: Optional[str] = None,
                             if_: Optional[str] = None,
                             ignore_failure: Optional[bool] = None,
                             ignore_missing: Optional[bool] = None,
                             on_failures: Optional[Sequence[str]] = None,
                             pattern: Optional[str] = None,
                             tag: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorDissectResult:
    """
    Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.

    Dissect matches a single text field against a defined pattern.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    dissect = elasticstack.ingest_processor_dissect(field="message",
        pattern="%{clientip} %{ident} %{auth} [%{@timestamp}] \\"%{verb} %{request} HTTP/%{httpversion}\\" %{status} %{size}")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[dissect.json])
    ```


    :param str append_separator: The character(s) that separate the appended fields.
    :param str description: Description of the processor.
    :param str field: The field to dissect.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str pattern: The pattern to apply to the field.
    :param str tag: Identifier for the processor.
    """
    __args__ = dict()
    __args__['appendSeparator'] = append_separator
    __args__['description'] = description
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['onFailures'] = on_failures
    __args__['pattern'] = pattern
    __args__['tag'] = tag
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorDissect:IngestProcessorDissect', __args__, opts=opts, typ=IngestProcessorDissectResult).value

    return AwaitableIngestProcessorDissectResult(
        append_separator=__ret__.append_separator,
        description=__ret__.description,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        on_failures=__ret__.on_failures,
        pattern=__ret__.pattern,
        tag=__ret__.tag)


@_utilities.lift_output_func(ingest_processor_dissect)
def ingest_processor_dissect_output(append_separator: Optional[pulumi.Input[Optional[str]]] = None,
                                    description: Optional[pulumi.Input[Optional[str]]] = None,
                                    field: Optional[pulumi.Input[str]] = None,
                                    if_: Optional[pulumi.Input[Optional[str]]] = None,
                                    ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                    ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                    on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    pattern: Optional[pulumi.Input[str]] = None,
                                    tag: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorDissectResult]:
    """
    Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.

    Dissect matches a single text field against a defined pattern.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    dissect = elasticstack.ingest_processor_dissect(field="message",
        pattern="%{clientip} %{ident} %{auth} [%{@timestamp}] \\"%{verb} %{request} HTTP/%{httpversion}\\" %{status} %{size}")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[dissect.json])
    ```


    :param str append_separator: The character(s) that separate the appended fields.
    :param str description: Description of the processor.
    :param str field: The field to dissect.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str pattern: The pattern to apply to the field.
    :param str tag: Identifier for the processor.
    """
    ...