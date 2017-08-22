// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/paths"
)

const (
	DefaultInputType = "log"
)

type Config struct {
	Period time.Duration `config:"period"`

	// *************************add mydelf ***************************
	Prospectors      []*common.Config `config:"prospectors"`
	Modules          []*common.Config `config:"modules"`
	ConfigDir        string           `config:"config_dir"`
	ProspectorReload *common.Config   `config:"config.prospectors"`
	SpoolSize        uint64           `config:"spool_size" validate:"min=1"`
	PublishAsync     bool             `config:"publish_async"`
	IdleTimeout      time.Duration    `config:"idle_timeout" validate:"nonzero,min=0s"`
	RegistryFile     string           `config:"registry_file"`
	ShutdownTimeout  time.Duration    `config:"shutdown_timeout"`
}

var (
		DefaultConfig = Config{
		Period: 1 * time.Second,
		RegistryFile:    "registry",
		SpoolSize:       2048,
		ShutdownTimeout: 0,
	}
)
// getConfigFiles returns list of config files.
// In case path is a file, it will be directly returned.
// In case it is a directory, it will fetch all .yml files inside this directory
func getConfigFiles(path string) (configFiles []string, err error) {

	// Check if path is valid file or dir
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Create empty slice for config file list
	configFiles = make([]string, 0)

	if stat.IsDir() {
		files, err := filepath.Glob(path + "/*.yml")

		if err != nil {
			return nil, err
		}

		configFiles = append(configFiles, files...)

	} else {
		// Only 1 config file
		configFiles = append(configFiles, path)
	}

	return configFiles, nil
}

func mergeConfigFiles(configFiles []string, config *Config)error{

	for _, file := range configFiles{
		logp.Info("Additional configs loaded from :%s", file)

		tmpConfig := struct{
			Filebeat Config
		}{}

		cfgfile.Read(&tmpConfig, file)
		config.Prospectors = append(config.Prospectors, tmpConfig.Filebeat.Prospectors...)
	}

	return nil
}

/**********  add myself  ***********/
func (config *Config)FetchConfigs()(error){
	configDir := config.ConfigDir

	// if option not set do nothing
	if configDir == "" {
		fmt.Println("option not set")
		return nil
	}

	configDir = paths.Resolve(paths.Config, configDir)
	println(configDir)

	// Check if optional configDir is set to fetch additional config files
	// before logp id done, It's not need me.
	logp.Info("Additional config files are fetched from: %s", configDir)


	// 对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，
	// 退出程序并返回状态 1
	// 。但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用，
	configFiles, err := getConfigFiles(configDir)
	if err != nil {
		log.Fatal("this is a error", configFiles, err)
		return err
	}

	err = mergeConfigFiles(configFiles, config)
	fmt.Println(configDir, configFiles, err)
	return nil
}
