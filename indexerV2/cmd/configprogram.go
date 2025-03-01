package cmd

type ConfigProgram struct {
	isTracer         bool
	isProf           bool
	workersMails     int
	workersUploaders int
	batchSize        int
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
