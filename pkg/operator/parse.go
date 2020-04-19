package operator

// BuildConfigInstance BuildConfigInstance is the output from a single ConfigMap.
// Mind that a single ConfigMap can contain multiple configurations.
func BuildConfigInstance(configMapData map[string]string) (ConfigInstance, error) {

	errs := []error{}

	configInstance := ConfigInstance{
		CloudWatchExporterConfig: CloudWatchExporterConfig{},
	}

	backwardCompatibleKey := "config.yml"

	if configValue, keyExists := configMapData[backwardCompatibleKey]; keyExists {

	} else if false {

	}

	return configInstance, nil
}
