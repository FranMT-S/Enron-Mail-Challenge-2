package cmd

type ConfigProgram struct {
	isTracer         bool // activate tracer profiling
	isProf           bool // activate mem and cpu profiling
	workersMails     int  // workers to index mails
	workersUploaders int  // workers to upload a batch of mails
	batchSize        int  // batch size
}

func validateFields(config *ConfigProgram) {
	if config.batchSize < minBatchSize {
		config.batchSize = minBatchSize
	}

	if config.batchSize > maxBatchSize {
		config.batchSize = maxBatchSize
	}

	if config.workersMails < minMailWorkers {
		config.workersMails = minMailWorkers
	}

	if config.workersMails > maxMailWorkers {
		config.workersMails = maxMailWorkers
	}

	if config.workersUploaders < minUploadWorkers {
		config.workersUploaders = minUploadWorkers
	}

	if config.workersUploaders > maxUploadWorkers {
		config.workersUploaders = maxUploadWorkers
	}
}
