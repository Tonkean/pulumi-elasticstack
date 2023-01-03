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
    'IngestProcessorGrokResult',
    'AwaitableIngestProcessorGrokResult',
    'ingest_processor_grok',
    'ingest_processor_grok_output',
]

@pulumi.output_type
class IngestProcessorGrokResult:
    """
    A collection of values returned by IngestProcessorGrok.
    """
    def __init__(__self__, description=None, ecs_compatibility=None, field=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, on_failures=None, pattern_definitions=None, patterns=None, tag=None, trace_match=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if ecs_compatibility and not isinstance(ecs_compatibility, str):
            raise TypeError("Expected argument 'ecs_compatibility' to be a str")
        pulumi.set(__self__, "ecs_compatibility", ecs_compatibility)
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
        if pattern_definitions and not isinstance(pattern_definitions, dict):
            raise TypeError("Expected argument 'pattern_definitions' to be a dict")
        pulumi.set(__self__, "pattern_definitions", pattern_definitions)
        if patterns and not isinstance(patterns, list):
            raise TypeError("Expected argument 'patterns' to be a list")
        pulumi.set(__self__, "patterns", patterns)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if trace_match and not isinstance(trace_match, bool):
            raise TypeError("Expected argument 'trace_match' to be a bool")
        pulumi.set(__self__, "trace_match", trace_match)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ecsCompatibility")
    def ecs_compatibility(self) -> Optional[str]:
        """
        Must be disabled or v1. If v1, the processor uses patterns with Elastic Common Schema (ECS) field names. **NOTE:** Supported only starting from version of Elasticsearch **7.16.x**.
        """
        return pulumi.get(self, "ecs_compatibility")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field to use for grok expression parsing
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
        If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document
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
    @pulumi.getter(name="patternDefinitions")
    def pattern_definitions(self) -> Optional[Mapping[str, str]]:
        """
        A map of pattern-name and pattern tuples defining custom patterns to be used by the current processor. Patterns matching existing names will override the pre-existing definition.
        """
        return pulumi.get(self, "pattern_definitions")

    @property
    @pulumi.getter
    def patterns(self) -> Sequence[str]:
        """
        An ordered list of grok expression to match and extract named captures with. Returns on the first expression in the list that matches.
        """
        return pulumi.get(self, "patterns")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="traceMatch")
    def trace_match(self) -> Optional[bool]:
        """
        when true, `_ingest._grok_match_index` will be inserted into your matched document’s metadata with the index into the pattern found in `patterns` that matched.
        """
        return pulumi.get(self, "trace_match")


class AwaitableIngestProcessorGrokResult(IngestProcessorGrokResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorGrokResult(
            description=self.description,
            ecs_compatibility=self.ecs_compatibility,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            on_failures=self.on_failures,
            pattern_definitions=self.pattern_definitions,
            patterns=self.patterns,
            tag=self.tag,
            trace_match=self.trace_match)


def ingest_processor_grok(description: Optional[str] = None,
                          ecs_compatibility: Optional[str] = None,
                          field: Optional[str] = None,
                          if_: Optional[str] = None,
                          ignore_failure: Optional[bool] = None,
                          ignore_missing: Optional[bool] = None,
                          on_failures: Optional[Sequence[str]] = None,
                          pattern_definitions: Optional[Mapping[str, str]] = None,
                          patterns: Optional[Sequence[str]] = None,
                          tag: Optional[str] = None,
                          trace_match: Optional[bool] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorGrokResult:
    """
    Extracts structured fields out of a single text field within a document. You choose which field to extract matched fields from, as well as the grok pattern you expect will match. A grok pattern is like a regular expression that supports aliased expressions that can be reused.

    This processor comes packaged with many [reusable patterns](https://github.com/elastic/elasticsearch/blob/master/libs/grok/src/main/resources/patterns).

    If you need help building patterns to match your logs, you will find the [Grok Debugger](https://www.elastic.co/guide/en/kibana/master/xpack-grokdebugger.html) tool quite useful! [The Grok Constructor](https://grokconstructor.appspot.com/) is also a useful tool.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    grok = elasticstack.ingest_processor_grok(field="message",
        patterns=[
            "%{FAVORITE_DOG:pet}",
            "%{FAVORITE_CAT:pet}",
        ],
        pattern_definitions={
            "FAVORITE_DOG": "beagle",
            "FAVORITE_CAT": "burmese",
        })
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[grok.json])
    ```


    :param str description: Description of the processor.
    :param str ecs_compatibility: Must be disabled or v1. If v1, the processor uses patterns with Elastic Common Schema (ECS) field names. **NOTE:** Supported only starting from version of Elasticsearch **7.16.x**.
    :param str field: The field to use for grok expression parsing
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param Mapping[str, str] pattern_definitions: A map of pattern-name and pattern tuples defining custom patterns to be used by the current processor. Patterns matching existing names will override the pre-existing definition.
    :param Sequence[str] patterns: An ordered list of grok expression to match and extract named captures with. Returns on the first expression in the list that matches.
    :param str tag: Identifier for the processor.
    :param bool trace_match: when true, `_ingest._grok_match_index` will be inserted into your matched document’s metadata with the index into the pattern found in `patterns` that matched.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['ecsCompatibility'] = ecs_compatibility
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['onFailures'] = on_failures
    __args__['patternDefinitions'] = pattern_definitions
    __args__['patterns'] = patterns
    __args__['tag'] = tag
    __args__['traceMatch'] = trace_match
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorGrok:IngestProcessorGrok', __args__, opts=opts, typ=IngestProcessorGrokResult).value

    return AwaitableIngestProcessorGrokResult(
        description=__ret__.description,
        ecs_compatibility=__ret__.ecs_compatibility,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        on_failures=__ret__.on_failures,
        pattern_definitions=__ret__.pattern_definitions,
        patterns=__ret__.patterns,
        tag=__ret__.tag,
        trace_match=__ret__.trace_match)


@_utilities.lift_output_func(ingest_processor_grok)
def ingest_processor_grok_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                 ecs_compatibility: Optional[pulumi.Input[Optional[str]]] = None,
                                 field: Optional[pulumi.Input[str]] = None,
                                 if_: Optional[pulumi.Input[Optional[str]]] = None,
                                 ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                 ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                 on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 pattern_definitions: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                 patterns: Optional[pulumi.Input[Sequence[str]]] = None,
                                 tag: Optional[pulumi.Input[Optional[str]]] = None,
                                 trace_match: Optional[pulumi.Input[Optional[bool]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorGrokResult]:
    """
    Extracts structured fields out of a single text field within a document. You choose which field to extract matched fields from, as well as the grok pattern you expect will match. A grok pattern is like a regular expression that supports aliased expressions that can be reused.

    This processor comes packaged with many [reusable patterns](https://github.com/elastic/elasticsearch/blob/master/libs/grok/src/main/resources/patterns).

    If you need help building patterns to match your logs, you will find the [Grok Debugger](https://www.elastic.co/guide/en/kibana/master/xpack-grokdebugger.html) tool quite useful! [The Grok Constructor](https://grokconstructor.appspot.com/) is also a useful tool.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    grok = elasticstack.ingest_processor_grok(field="message",
        patterns=[
            "%{FAVORITE_DOG:pet}",
            "%{FAVORITE_CAT:pet}",
        ],
        pattern_definitions={
            "FAVORITE_DOG": "beagle",
            "FAVORITE_CAT": "burmese",
        })
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[grok.json])
    ```


    :param str description: Description of the processor.
    :param str ecs_compatibility: Must be disabled or v1. If v1, the processor uses patterns with Elastic Common Schema (ECS) field names. **NOTE:** Supported only starting from version of Elasticsearch **7.16.x**.
    :param str field: The field to use for grok expression parsing
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param Mapping[str, str] pattern_definitions: A map of pattern-name and pattern tuples defining custom patterns to be used by the current processor. Patterns matching existing names will override the pre-existing definition.
    :param Sequence[str] patterns: An ordered list of grok expression to match and extract named captures with. Returns on the first expression in the list that matches.
    :param str tag: Identifier for the processor.
    :param bool trace_match: when true, `_ingest._grok_match_index` will be inserted into your matched document’s metadata with the index into the pattern found in `patterns` that matched.
    """
    ...
