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

from flyteadmin.models.admin_email_notification import AdminEmailNotification  # noqa: F401,E501
from flyteadmin.models.admin_pager_duty_notification import AdminPagerDutyNotification  # noqa: F401,E501
from flyteadmin.models.admin_slack_notification import AdminSlackNotification  # noqa: F401,E501
from flyteadmin.models.admin_sns_message import AdminSNSMessage  # noqa: F401,E501
from flyteadmin.models.core_workflow_execution_phase import CoreWorkflowExecutionPhase  # noqa: F401,E501


class AdminNotification(object):
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
        'phases': 'list[CoreWorkflowExecutionPhase]',
        'email': 'AdminEmailNotification',
        'pager_duty': 'AdminPagerDutyNotification',
        'slack': 'AdminSlackNotification',
        'snsmessage': 'AdminSNSMessage'
    }

    attribute_map = {
        'phases': 'phases',
        'email': 'email',
        'pager_duty': 'pager_duty',
        'slack': 'slack',
        'snsmessage': 'snsmessage'
    }

    def __init__(self, phases=None, email=None, pager_duty=None, slack=None, snsmessage=None):  # noqa: E501
        """AdminNotification - a model defined in Swagger"""  # noqa: E501

        self._phases = None
        self._email = None
        self._pager_duty = None
        self._slack = None
        self._snsmessage = None
        self.discriminator = None

        if phases is not None:
            self.phases = phases
        if email is not None:
            self.email = email
        if pager_duty is not None:
            self.pager_duty = pager_duty
        if slack is not None:
            self.slack = slack
        if snsmessage is not None:
            self.snsmessage = snsmessage

    @property
    def phases(self):
        """Gets the phases of this AdminNotification.  # noqa: E501

        A list of phases to which users can associate the notifications to.  # noqa: E501

        :return: The phases of this AdminNotification.  # noqa: E501
        :rtype: list[CoreWorkflowExecutionPhase]
        """
        return self._phases

    @phases.setter
    def phases(self, phases):
        """Sets the phases of this AdminNotification.

        A list of phases to which users can associate the notifications to.  # noqa: E501

        :param phases: The phases of this AdminNotification.  # noqa: E501
        :type: list[CoreWorkflowExecutionPhase]
        """

        self._phases = phases

    @property
    def email(self):
        """Gets the email of this AdminNotification.  # noqa: E501


        :return: The email of this AdminNotification.  # noqa: E501
        :rtype: AdminEmailNotification
        """
        return self._email

    @email.setter
    def email(self, email):
        """Sets the email of this AdminNotification.


        :param email: The email of this AdminNotification.  # noqa: E501
        :type: AdminEmailNotification
        """

        self._email = email

    @property
    def pager_duty(self):
        """Gets the pager_duty of this AdminNotification.  # noqa: E501


        :return: The pager_duty of this AdminNotification.  # noqa: E501
        :rtype: AdminPagerDutyNotification
        """
        return self._pager_duty

    @pager_duty.setter
    def pager_duty(self, pager_duty):
        """Sets the pager_duty of this AdminNotification.


        :param pager_duty: The pager_duty of this AdminNotification.  # noqa: E501
        :type: AdminPagerDutyNotification
        """

        self._pager_duty = pager_duty

    @property
    def slack(self):
        """Gets the slack of this AdminNotification.  # noqa: E501


        :return: The slack of this AdminNotification.  # noqa: E501
        :rtype: AdminSlackNotification
        """
        return self._slack

    @slack.setter
    def slack(self, slack):
        """Sets the slack of this AdminNotification.


        :param slack: The slack of this AdminNotification.  # noqa: E501
        :type: AdminSlackNotification
        """

        self._slack = slack

    @property
    def snsmessage(self):
        """Gets the snsmessage of this AdminNotification.  # noqa: E501


        :return: The snsmessage of this AdminNotification.  # noqa: E501
        :rtype: AdminSNSMessage
        """
        return self._snsmessage

    @snsmessage.setter
    def snsmessage(self, snsmessage):
        """Sets the snsmessage of this AdminNotification.


        :param snsmessage: The snsmessage of this AdminNotification.  # noqa: E501
        :type: AdminSNSMessage
        """

        self._snsmessage = snsmessage

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
        if issubclass(AdminNotification, dict):
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
        if not isinstance(other, AdminNotification):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
