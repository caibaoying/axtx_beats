package beater
//
//import (
//	"fmt"
//
//	"github.com/elastic/beats/libbeat/beat"
//	"github.com/elastic/beats/libbeat/common"
//	"github.com/elastic/beats/libbeat/logp"
//	"github.com/elastic/beats/libbeat/publisher"
//	"github.com/caibaoying/axtx_beats/config"
//	"github.com/elastic/beats/filebeat/fileset"
//
//	//"errors"
//	"flag"
//	"github.com/elastic/beats/libbeat/outputs/elasticsearch"
//	"errors"
//)
//
//var (
//	once = flag.Bool("once", false, "Run filebeat only once until all harvesters reach EOF")
//)
//
//type Axtx_beats struct {
//	done   chan struct{}
//	config *config.Config
//	client publisher.Client
//	moduleRegistry *fileset.ModuleRegistry
//}
//
//// Create beater
//
//// listen to me , New has trouble
//func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
//
//	config := config.DefaultConfig
//	if err := cfg.Unpack(&config); err != nil {
//		return nil, fmt.Errorf("Error reading config file: %v", err)
//	}
//
//	/************************************************************************/
//	// ******************  if you don't use var, it's error  *****************
//	moduleRegistry, err := fileset.NewModuleRegistry(config.Modules, b.Version)
//	fmt.Println("moduleRegistry:", moduleRegistry)
//	if err != nil {
//		return nil, err
//	}
//
//	moduleProspectors, err := moduleRegistry.GetProspectorConfigs()
//	fmt.Println("moduleProspector:", moduleProspectors, "\nerr:", err)
//
//
//	if err != nil {
//		return nil, err
//	}
//
//	// jiang suo you de peizhi wenjian hebing dao yiqi
//	if err := config.FetchConfigs(); err != nil {
//		return nil, err
//	}
//
//	config.Prospectors = append(config.Prospectors, moduleProspectors...)
//
//	fmt.Println("config.Prospectors:", config.Prospectors, "\nconfig.ProspectorReload.Enabled():", config.ProspectorReload.Enabled())
//
//	if !config.ProspectorReload.Enabled() && len(config.Prospectors) == 0 {
//		return nil, errors.New("No prospectors defined. What files do you want me to watch?")
//	}
//
//	if *once && config.ProspectorReload.Enabled() {
//		return nil, errors.New("prospector reloading and -once cannot be used together.")
//	}
//
//	/************************************************************************/
//
//	bt := &Axtx_beats{
//		done:   make(chan struct{}),
//		config: &config,
//		moduleRegistry: moduleRegistry,
//	}
//
//	/***************** reasion for error: Home is not sot *************/
//	// next problem is answer this question
//	return bt, nil
//}
//
//// 当模块配置为进行初始设置时，将调用modulesSetup
//func (fb *Axtx_beats) modulesSetup(b *beat.Beat) error {
//	esConfig := b.Config.Output["elasticsearch"]
//	if esConfig == nil || !esConfig.Enabled() {
//		logp.Warn("Filebeat is unable to load the Ingest Node pipelines for the configured" +
//			" modules because the Elasticsearch output is not configured/enabled. If you have" +
//			" already loaded the Ingest Node pipelines or are using Logstash pipelines, you" +
//			" can ignore this warning.")
//		return nil
//	}
//	esClient, err := elasticsearch.NewConnectedClient(esConfig)
//	if err != nil {
//		return fmt.Errorf("Error creating ES client: %v", err)
//	}
//	defer esClient.Close()
//
//	err = fb.moduleRegistry.LoadPipelines(esClient)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// 当模块配置为进行初始设置时，将调用 modulesSetup
//func (bt *Axtx_beats) Run(b *beat.Beat) error {
//
//	var err error
//	config := bt.config
//
//	if !bt.moduleRegistry.Empty() {
//		err = bt.modulesSetup(b)
//		if err != nil {
//			return err
//		}
//	}
//
//	fmt.Println("config:", config)
//	//logp.Info("axtx_beats is running! Hit CTRL-C to stop it.")
//	//
//	//bt.client = b.Publisher.Connect()
//	//// 设置定时器
//	//ticker := time.NewTicker(bt.config.Period)
//	//counter := 1
//	//
//	//// 死循环，相当于while
//	//for {
//	//	select {
//	//	// 接收操作符
//	//	case <-bt.done:
//	//		return nil
//	//	case <-ticker.C:
//	//	}
//	//
//	//	event := common.MapStr{
//	//		"@timestamp": common.Time(time.Now()),
//	//		"type":       b.Name,
//	//		"counter":    counter,
//	//	}
//	//	bt.client.PublishEvent(event)
//	//	logp.Info("Event sent")
//	//	counter++
//	//}
//
//	return nil
//}
//
//func (bt *Axtx_beats) Stop() {
//	bt.client.Close()
//	close(bt.done)
//}
