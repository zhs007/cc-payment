package err

import "errors"

var (
	// ErrConfigLogOutputType - invalid config.yaml:log.outputtype
	ErrConfigLogOutputType = errors.New("invalid config.yaml:log.outputtype")
	// ErrConfigLogLevel - invalid config.yaml:log.loglevel
	ErrConfigLogLevel = errors.New("invalid config.yaml:log.loglevel")
	// ErrNotLoadConfig - not load config.yaml
	ErrNotLoadConfig = errors.New("not load config.yaml")
)
