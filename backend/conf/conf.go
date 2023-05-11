package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Metadata struct {
	Name string
}

var (
	Probes         []Probe
	Interval, Port int
	Metas          map[string]map[string]any
	UpdateCallBack func()
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	viper.OnConfigChange(
		func(_ fsnotify.Event) {
			if err := update(); err != nil {
				fmt.Println("update probes fail", err)
			}
		})
	viper.WatchConfig()

	if err != nil {
		return err
	}
	Port = viper.GetInt("port")
	return update()
}

func update() error {
	// update probe
	var tmp []Probe
	probes := viper.Get("probe").(map[string]any)
	for name, probe := range probes {
		probe, ok := probe.(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:%v", probe)
		}
		parser, ok := probe["parse"].(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid field parse")
		}
		fetcher, ok := probe["fetch"].(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid field fetch")
		}
		tmp = append(tmp, Probe{
			Name:  name,
			Parse: parser,
			Fetch: fetcher,
		})
	}
	Probes = tmp

	//update metadata
	var tmpMeta = make(map[string]map[string]any)
	metas, ok := viper.Get("metadata").(map[string]any)
	if !ok {
		return nil
	}
	for name, meta := range metas {
		meta, ok := meta.(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid metadata:%v", meta)
		}
		tmpMeta[name] = meta
	}
	Metas = tmpMeta

	Interval = viper.GetInt("interval")

	if UpdateCallBack != nil {
		UpdateCallBack()
	}
	log.Println("update config success")
	return nil
}
