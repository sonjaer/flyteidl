# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from flyteidl.admin import event_pb2 as flyteidl_dot_admin_dot_event__pb2
from flyteidl.admin import execution_pb2 as flyteidl_dot_admin_dot_execution__pb2
from flyteidl.admin import launch_plan_pb2 as flyteidl_dot_admin_dot_launch__plan__pb2
from flyteidl.admin import project_pb2 as flyteidl_dot_admin_dot_project__pb2
from flyteidl.admin import task_pb2 as flyteidl_dot_admin_dot_task__pb2
from flyteidl.admin import workflow_pb2 as flyteidl_dot_admin_dot_workflow__pb2


class AdminServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateTask = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateTask',
        request_serializer=flyteidl_dot_admin_dot_task__pb2.TaskCreateRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_task__pb2.TaskCreateResponse.FromString,
        )
    self.GetTask = channel.unary_unary(
        '/flyteidl.service.AdminService/GetTask',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_task__pb2.Task.FromString,
        )
    self.ListTaskIds = channel.unary_unary(
        '/flyteidl.service.AdminService/ListTaskIds',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.IdentifierListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_common__pb2.IdentifierList.FromString,
        )
    self.ListTasks = channel.unary_unary(
        '/flyteidl.service.AdminService/ListTasks',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_task__pb2.TaskList.FromString,
        )
    self.CreateWorkflow = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateWorkflow',
        request_serializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowCreateRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowCreateResponse.FromString,
        )
    self.GetWorkflow = channel.unary_unary(
        '/flyteidl.service.AdminService/GetWorkflow',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_workflow__pb2.Workflow.FromString,
        )
    self.ListWorkflowIds = channel.unary_unary(
        '/flyteidl.service.AdminService/ListWorkflowIds',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.IdentifierListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_common__pb2.IdentifierList.FromString,
        )
    self.ListWorkflows = channel.unary_unary(
        '/flyteidl.service.AdminService/ListWorkflows',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowList.FromString,
        )
    self.CreateLaunchPlan = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateLaunchPlan',
        request_serializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanCreateRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanCreateResponse.FromString,
        )
    self.GetLaunchPlan = channel.unary_unary(
        '/flyteidl.service.AdminService/GetLaunchPlan',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlan.FromString,
        )
    self.UpdateLaunchPlan = channel.unary_unary(
        '/flyteidl.service.AdminService/UpdateLaunchPlan',
        request_serializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanUpdateRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanUpdateResponse.FromString,
        )
    self.CreateExecution = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateExecution',
        request_serializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionCreateRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionCreateResponse.FromString,
        )
    self.GetExecution = channel.unary_unary(
        '/flyteidl.service.AdminService/GetExecution',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_execution__pb2.Execution.FromString,
        )
    self.ListExecutions = channel.unary_unary(
        '/flyteidl.service.AdminService/ListExecutions',
        request_serializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionList.FromString,
        )
    self.RegisterProject = channel.unary_unary(
        '/flyteidl.service.AdminService/RegisterProject',
        request_serializer=flyteidl_dot_admin_dot_project__pb2.Project.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_project__pb2.ProjectRegisterResponse.FromString,
        )
    self.ListProjects = channel.unary_unary(
        '/flyteidl.service.AdminService/ListProjects',
        request_serializer=flyteidl_dot_admin_dot_project__pb2.ProjectListRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_project__pb2.Projects.FromString,
        )
    self.CreateWorkflowEvent = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateWorkflowEvent',
        request_serializer=flyteidl_dot_admin_dot_event__pb2.WorkflowExecutionEventRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_event__pb2.WorkflowExecutionEventResponse.FromString,
        )
    self.CreateNodeEvent = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateNodeEvent',
        request_serializer=flyteidl_dot_admin_dot_event__pb2.NodeExecutionEventRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_event__pb2.NodeExecutionEventResponse.FromString,
        )
    self.CreateTaskEvent = channel.unary_unary(
        '/flyteidl.service.AdminService/CreateTaskEvent',
        request_serializer=flyteidl_dot_admin_dot_event__pb2.TaskExecutionEventRequest.SerializeToString,
        response_deserializer=flyteidl_dot_admin_dot_event__pb2.TaskExecutionEventResponse.FromString,
        )


class AdminServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateTask(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetTask(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListTaskIds(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListTasks(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateWorkflow(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetWorkflow(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListWorkflowIds(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListWorkflows(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateLaunchPlan(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetLaunchPlan(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateLaunchPlan(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateExecution(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetExecution(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListExecutions(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RegisterProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListProjects(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateWorkflowEvent(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNodeEvent(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateTaskEvent(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AdminServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateTask': grpc.unary_unary_rpc_method_handler(
          servicer.CreateTask,
          request_deserializer=flyteidl_dot_admin_dot_task__pb2.TaskCreateRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_task__pb2.TaskCreateResponse.SerializeToString,
      ),
      'GetTask': grpc.unary_unary_rpc_method_handler(
          servicer.GetTask,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_task__pb2.Task.SerializeToString,
      ),
      'ListTaskIds': grpc.unary_unary_rpc_method_handler(
          servicer.ListTaskIds,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.IdentifierListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_common__pb2.IdentifierList.SerializeToString,
      ),
      'ListTasks': grpc.unary_unary_rpc_method_handler(
          servicer.ListTasks,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_task__pb2.TaskList.SerializeToString,
      ),
      'CreateWorkflow': grpc.unary_unary_rpc_method_handler(
          servicer.CreateWorkflow,
          request_deserializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowCreateRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowCreateResponse.SerializeToString,
      ),
      'GetWorkflow': grpc.unary_unary_rpc_method_handler(
          servicer.GetWorkflow,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_workflow__pb2.Workflow.SerializeToString,
      ),
      'ListWorkflowIds': grpc.unary_unary_rpc_method_handler(
          servicer.ListWorkflowIds,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.IdentifierListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_common__pb2.IdentifierList.SerializeToString,
      ),
      'ListWorkflows': grpc.unary_unary_rpc_method_handler(
          servicer.ListWorkflows,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_workflow__pb2.WorkflowList.SerializeToString,
      ),
      'CreateLaunchPlan': grpc.unary_unary_rpc_method_handler(
          servicer.CreateLaunchPlan,
          request_deserializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanCreateRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanCreateResponse.SerializeToString,
      ),
      'GetLaunchPlan': grpc.unary_unary_rpc_method_handler(
          servicer.GetLaunchPlan,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlan.SerializeToString,
      ),
      'UpdateLaunchPlan': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateLaunchPlan,
          request_deserializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanUpdateRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_launch__plan__pb2.LaunchPlanUpdateResponse.SerializeToString,
      ),
      'CreateExecution': grpc.unary_unary_rpc_method_handler(
          servicer.CreateExecution,
          request_deserializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionCreateRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionCreateResponse.SerializeToString,
      ),
      'GetExecution': grpc.unary_unary_rpc_method_handler(
          servicer.GetExecution,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ObjectGetRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_execution__pb2.Execution.SerializeToString,
      ),
      'ListExecutions': grpc.unary_unary_rpc_method_handler(
          servicer.ListExecutions,
          request_deserializer=flyteidl_dot_admin_dot_common__pb2.ResourceListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_execution__pb2.ExecutionList.SerializeToString,
      ),
      'RegisterProject': grpc.unary_unary_rpc_method_handler(
          servicer.RegisterProject,
          request_deserializer=flyteidl_dot_admin_dot_project__pb2.Project.FromString,
          response_serializer=flyteidl_dot_admin_dot_project__pb2.ProjectRegisterResponse.SerializeToString,
      ),
      'ListProjects': grpc.unary_unary_rpc_method_handler(
          servicer.ListProjects,
          request_deserializer=flyteidl_dot_admin_dot_project__pb2.ProjectListRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_project__pb2.Projects.SerializeToString,
      ),
      'CreateWorkflowEvent': grpc.unary_unary_rpc_method_handler(
          servicer.CreateWorkflowEvent,
          request_deserializer=flyteidl_dot_admin_dot_event__pb2.WorkflowExecutionEventRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_event__pb2.WorkflowExecutionEventResponse.SerializeToString,
      ),
      'CreateNodeEvent': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNodeEvent,
          request_deserializer=flyteidl_dot_admin_dot_event__pb2.NodeExecutionEventRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_event__pb2.NodeExecutionEventResponse.SerializeToString,
      ),
      'CreateTaskEvent': grpc.unary_unary_rpc_method_handler(
          servicer.CreateTaskEvent,
          request_deserializer=flyteidl_dot_admin_dot_event__pb2.TaskExecutionEventRequest.FromString,
          response_serializer=flyteidl_dot_admin_dot_event__pb2.TaskExecutionEventResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'flyteidl.service.AdminService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
