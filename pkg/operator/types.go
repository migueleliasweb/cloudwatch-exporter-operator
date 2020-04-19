package operator

// ConfigInstance Defines all configuration defined in a ConfigMap
type ConfigInstance struct {
	CloudWatchExporterConfig CloudWatchExporterConfig
}

// CloudWatchExporterConfig The overall configuration for CloudWatchExporter
type CloudWatchExporterConfig struct {
	Region  string
	RoleArn string
	Metrics []CloudWatchExporterMetric
}

// CloudWatchExporterMetric A single metric from CloudWatch
type CloudWatchExporterMetric struct {
	Region                  string
	AWSNamespace            string
	AWSMetricName           string
	AWSDimensions           []string
	AWSDimensionSelect      map[string]string
	AWSDimensionSelectRegex map[string]string
	AWSTagSelect            AWSTagSelectFields
	AWSStatistics           []string
	AWSExtendedStatistics   []string
	DelaySeconds            int
	RangeSeconds            int
	PeriodSeconds           int
	SetTimestamp            bool
}

// AWSTagSelectFields Fields of AWSTagSelect
type AWSTagSelectFields struct {
	TagSelections         map[string]string
	ResourceTypeSelection string
	ResourceIDDimension   string
}
