package cmd

import "runtime"

var (
	defaultMailWorkers   = runtime.NumCPU()
	minMailWorkers       = 1
	maxMailWorkers       = 30
	defaultBatchSize     = 1000
	minBatchSize         = 100
	maxBatchSize         = 2000
	defaultUploadWorkers = 4
	minUploadWorkers     = 1
	maxUploadWorkers     = 8
)
