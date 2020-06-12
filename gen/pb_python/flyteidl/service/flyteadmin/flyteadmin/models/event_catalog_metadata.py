# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.core_workflow_execution_identifier import CoreWorkflowExecutionIdentifier  # noqa: F401,E501


class EventCatalogMetadata(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'dataset_id': 'str',
        'artifact_tag': 'str',
        'source_execution_id': 'CoreWorkflowExecutionIdentifier'
    }

    attribute_map = {
        'dataset_id': 'dataset_id',
        'artifact_tag': 'artifact_tag',
        'source_execution_id': 'source_execution_id'
    }

    def __init__(self, dataset_id=None, artifact_tag=None, source_execution_id=None):  # noqa: E501
        """EventCatalogMetadata - a model defined in Swagger"""  # noqa: E501

        self._dataset_id = None
        self._artifact_tag = None
        self._source_execution_id = None
        self.discriminator = None

        if dataset_id is not None:
            self.dataset_id = dataset_id
        if artifact_tag is not None:
            self.artifact_tag = artifact_tag
        if source_execution_id is not None:
            self.source_execution_id = source_execution_id

    @property
    def dataset_id(self):
        """Gets the dataset_id of this EventCatalogMetadata.  # noqa: E501


        :return: The dataset_id of this EventCatalogMetadata.  # noqa: E501
        :rtype: str
        """
        return self._dataset_id

    @dataset_id.setter
    def dataset_id(self, dataset_id):
        """Sets the dataset_id of this EventCatalogMetadata.


        :param dataset_id: The dataset_id of this EventCatalogMetadata.  # noqa: E501
        :type: str
        """

        self._dataset_id = dataset_id

    @property
    def artifact_tag(self):
        """Gets the artifact_tag of this EventCatalogMetadata.  # noqa: E501


        :return: The artifact_tag of this EventCatalogMetadata.  # noqa: E501
        :rtype: str
        """
        return self._artifact_tag

    @artifact_tag.setter
    def artifact_tag(self, artifact_tag):
        """Sets the artifact_tag of this EventCatalogMetadata.


        :param artifact_tag: The artifact_tag of this EventCatalogMetadata.  # noqa: E501
        :type: str
        """

        self._artifact_tag = artifact_tag

    @property
    def source_execution_id(self):
        """Gets the source_execution_id of this EventCatalogMetadata.  # noqa: E501


        :return: The source_execution_id of this EventCatalogMetadata.  # noqa: E501
        :rtype: CoreWorkflowExecutionIdentifier
        """
        return self._source_execution_id

    @source_execution_id.setter
    def source_execution_id(self, source_execution_id):
        """Sets the source_execution_id of this EventCatalogMetadata.


        :param source_execution_id: The source_execution_id of this EventCatalogMetadata.  # noqa: E501
        :type: CoreWorkflowExecutionIdentifier
        """

        self._source_execution_id = source_execution_id

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(EventCatalogMetadata, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, EventCatalogMetadata):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other