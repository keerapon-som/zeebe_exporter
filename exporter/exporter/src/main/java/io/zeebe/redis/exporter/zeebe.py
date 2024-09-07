import json
import logging
import grpc
from zeebe_grpc import gateway_pb2, gateway_pb2_grpc

with grpc.insecure_channel("localhost:26500") as channel:
    stub = gateway_pb2_grpc.GatewayStub(channel)

    # print the topology of the zeebe cluster
    topology = stub.Topology(gateway_pb2.TopologyRequest())
    print(topology)

    # # deploy a process definition
    # with open("bpmn/echo.bpmn", "rb") as process_definition_file:
    #     process_definition = process_definition_file.read()
    #     process = gateway_pb2.ProcessRequestObject(
    #         name="echo.bpmn",
    #         definition=process_definition
    #     )
    # stub.DeployProcess(
    #     gateway_pb2.DeployProcessRequest(
    #         processes=[process]
    #     )
    # )

    # # start a process instance
    # variables = {
    #     "message": "testperformanceNa"
    # }
    # for i in range(1000):
    #     res = stub.CreateProcessInstance(
    #         gateway_pb2.CreateProcessInstanceRequest(
    #             bpmnProcessId="testperformanceNa",
    #             version=-1,
    #             variables=json.dumps(variables)
    #         )
    #     )
    #     print("This is response", str(res), " ", str(i))

    # # start a worker
    # activate_jobs_response = stub.ActivateJobs(
    #     gateway_pb2.ActivateJobsRequest(
    #         type="jobna",
    #         worker="Python worker",
    #         timeout=60000,
    #         maxJobsToActivate=10000
    #     )
    # )
    # xx = 0
    # for response in activate_jobs_response:
    #     for job in response.jobs:
    #         try:
    #             xx = xx + 1
    #             print(job.variables)
    #             print("This is job key xx", str(xx))
    #             stub.CompleteJob(gateway_pb2.CompleteJobRequest(jobKey=job.key, variables=json.dumps({})))
    #             logging.info("Job Completed")
    #         except Exception as e:
    #             stub.FailJob(gateway_pb2.FailJobRequest(jobKey=job.key))
    #             logging.info(f"Job Failed {e}")