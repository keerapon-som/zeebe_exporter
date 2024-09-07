
    # start a worker
    activate_jobs_response = stub.ActivateJobs(
        gateway_pb2.ActivateJobsRequest(
            type="jobna",
            worker="Python worker",
            timeout=60000,
            maxJobsToActivate=10000
        )
    )
    xx = 0
    for response in activate_jobs_response:
        for job in response.jobs:
            try:
                xx = xx + 1
                print(job.variables)
                print("This is job key xx", str(xx))
                stub.CompleteJob(gateway_pb2.CompleteJobRequest(jobKey=job.key, variables=json.dumps({})))
                logging.info("Job Completed")
            except Exception as e:
                stub.FailJob(gateway_pb2.FailJobRequest(jobKey=job.key))
                logging.info(f"Job Failed {e}")