/*
 * SMF Configuration Factory
 */

package factory

import (
	"fmt"
	"io/ioutil"

	"github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"

	"github.com/free5gc/smf/internal/logger"
)

var (
	SmfConfig       *Config
	UERoutingConfig *RoutingConfig
	UEAddressConfig *UeAddresses
)

// TODO: Support configuration update from REST api
func InitConfigFactory(f string, cfg *Config) error {
	if f == "" {
		// Use default config path
		f = SmfDefaultConfigPath
	}

	if content, err := ioutil.ReadFile(f); err != nil {
		return fmt.Errorf("[Factory] %+v", err)
	} else {
		logger.CfgLog.Infof("Read config from [%s]", f)
		if yamlErr := yaml.Unmarshal(content, cfg); yamlErr != nil {
			return fmt.Errorf("[Factory] %+v", yamlErr)
		}
	}

	return nil
}

func InitUeAddressesConfigFactory(f string, cfg *UeAddresses) error {
	if f == "" {
		// Use default config path
		f = SmfDefaultUEAddressesPath
	}
	if content, err := ioutil.ReadFile(f); err != nil {
		return err
	} else {
		logger.CfgLog.Infof("Read config from [%s]", f)
		if yamlErr := yaml.Unmarshal(content, cfg); yamlErr != nil {
			return fmt.Errorf("[Factory] %+v", yamlErr)
		}
	}

	return nil
}

func InitRoutingConfigFactory(f string, cfg *RoutingConfig) error {
	if f == "" {
		// Use default config path
		f = SmfDefaultUERoutingPath
	}
	if content, err := ioutil.ReadFile(f); err != nil {
		return err
	} else {
		logger.CfgLog.Infof("Read config from [%s]", f)
		if yamlErr := yaml.Unmarshal(content, cfg); yamlErr != nil {
			return fmt.Errorf("[Factory] %+v", yamlErr)
		}
	}

	return nil
}

func ReadConfig(cfgPath string) (*Config, error) {
	cfg := &Config{}
	if err := InitConfigFactory(cfgPath, cfg); err != nil {
		return nil, fmt.Errorf("ReadConfig [%s] Error: %+v", cfgPath, err)
	}
	if _, err := cfg.Validate(); err != nil {
		validErrs := err.(govalidator.Errors).Errors()
		for _, validErr := range validErrs {
			logger.CfgLog.Errorf("%+v", validErr)
		}
		logger.CfgLog.Errorf("[-- PLEASE REFER TO SAMPLE CONFIG FILE COMMENTS --]")
		return nil, fmt.Errorf("Config validate Error")
	}

	return cfg, nil
}

func ReadUERoutingConfig(cfgPath string) (*RoutingConfig, error) {
	ueRoutingCfg := &RoutingConfig{}
	if err := InitRoutingConfigFactory(cfgPath, ueRoutingCfg); err != nil {
		return nil, fmt.Errorf("ReadConfig [%s] Error: %+v", cfgPath, err)
	}
	if _, err := ueRoutingCfg.Validate(); err != nil {
		validErrs := err.(govalidator.Errors).Errors()
		for _, validErr := range validErrs {
			logger.CfgLog.Errorf("%+v", validErr)
		}
		logger.CfgLog.Errorf("[-- PLEASE REFER TO SAMPLE CONFIG FILE COMMENTS --]")
		return nil, fmt.Errorf("Config validate Error")
	}

	return ueRoutingCfg, nil
}

func ReadUEAddressesConfig(cfgPath string) (*UeAddresses, error) {
	ueAddresses := &UeAddresses{}
	if err := InitUeAddressesConfigFactory(cfgPath, ueAddresses); err != nil {
		return nil, fmt.Errorf("ReadConfig [%s] Error: %+v", cfgPath, err)
	}
	//	if _, err := ueRoutingCfg.Validate(); err != nil {
	//		validErrs := err.(govalidator.Errors).Errors()
	//		for _, validErr := range validErrs {
	//			logger.CfgLog.Errorf("%+v", validErr)
	//		}
	//		logger.CfgLog.Errorf("[-- PLEASE REFER TO SAMPLE CONFIG FILE COMMENTS --]")
	//		return nil, fmt.Errorf("Config validate Error")
	//	}

	return ueAddresses, nil
}
