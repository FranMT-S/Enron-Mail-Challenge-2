package cmd

import (
	"flag"
	"fmt"
)

func defineFlag(config *ConfigProgram) {
	flag.IntVar(&config.batchSize,
		"batch",
		defaultBatchSize,
		fmt.Sprintf("Batch size for inserting emails into the database,min:%v, max:%v, default:%v", minBatchSize, maxBatchSize, defaultBatchSize),
	)

	flag.IntVar(&config.workersMails,
		"wm",
		defaultMailWorkers,
		fmt.Sprintf("'Workers mails' number of process index emails, default is num cpu, min:%v, max:%v, default:%v", minMailWorkers, maxMailWorkers, defaultMailWorkers),
	)

	flag.IntVar(&config.workersUploaders,
		"wu",
		defaultUploadWorkers,
		fmt.Sprintf("'workers upload' number of process to upload the mails, min:%v, max:%v, default:%v", minUploadWorkers, maxUploadWorkers, defaultUploadWorkers),
	)

	flag.BoolVar(&config.isTracer,
		"trace",
		false,
		"Activate memory and cpu profiling. It is recommended not to use it together with trace profiling for best result. default: false.",
	)

	flag.BoolVar(&config.isProf,
		"prof",
		false,
		"Activate trace profiling. It is recommended not to use it together with CPU and memory profiling  for best result.default: false.",
	)

	flag.Parse()
}
