package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Metadata struct {
	Name string
}

var (
	Probes         []Probe
	Interval, Port int
	Metas          map[string]map[string]any
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	viper.OnConfigChange(
		func(e fsnotify.Event) {
			if err := update(); err != nil {
				fmt.Println("update probes fail", err)
			}
		})
	viper.WatchConfig()

	if err != nil {
		return err
	}
	Interval = viper.GetInt("interval")
	Port = viper.GetInt("port")
	return update()
}

func update() error {
	// update probe
	var tmp []Probe
	probes := viper.Get("probe").([]any)
	for _, probe := range probes {
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
			Parse: parser,
			Fetch: fetcher,
		})
	}
	Probes = tmp

	//update metadata
	var tmpMeta = make(map[string]map[string]any)
	metas := viper.Get("metadata").(map[string]any)
	for name, meta := range metas {
		meta, ok := meta.(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid metadata:%v", meta)
		}
		tmpMeta[name] = meta
	}
	Metas = tmpMeta

	return nil
}
