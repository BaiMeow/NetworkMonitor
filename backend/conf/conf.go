package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Metadata struct {
	Name string
}

var (
	Probes         []Probe
	Port           int
	Interval       time.Duration
	ProbeTimeout   time.Duration
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

	Interval = time.Duration(viper.GetInt("interval")) * time.Second
	ProbeTimeout = time.Duration(viper.GetInt("probeTimeout")) * time.Second

	if UpdateCallBack != nil {
		UpdateCallBack()
	}
	log.Println("update config success")
	return nil
}
