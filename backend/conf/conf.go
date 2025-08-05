package conf

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Uptime = UptimeCfg{
		// Interval for uptime is statically set to
		Interval: time.Minute,
	}
	Influxdb struct {
		Addr  string
		Token string
		Org   string
	}
	Probes           []Probe
	Port             int
	Interval         time.Duration
	ProbeTimeout     time.Duration
	UpdateCallBack   func()
	MetadataRedirect string
	Analysis         bool
	Trace            Tracer
)

func Init() error {
	viper.SetDefault("port", 8787)
	viper.SetDefault("interval", 10)
	viper.SetDefault("probeTimeout", 8)
	viper.SetDefault("uptime.recordDuration", "48h")
	viper.SetDefault("analysis", true)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("netm")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config file: %v", err)
	}
	viper.OnConfigChange(
		func(_ fsnotify.Event) {
			if err := update(); err != nil {
				fmt.Println("update probes fail", err)
			}
		})
	viper.WatchConfig()

	Port = viper.GetInt("port")
	Influxdb.Addr = viper.GetString("influxdb.addr")
	Influxdb.Token = viper.GetString("influxdb.token")
	Influxdb.Org = viper.GetString("influxdb.org")

	return update()
}

func update() error {
	// update probe
	var tmp []Probe
	if probes, ok := viper.Get("probe").(map[string]any); ok {
		for name, probe := range probes {
			probe, ok := probe.(map[string]any)
			if !ok {
				return fmt.Errorf("parse config error: %v", probe)
			}
			parser, ok := probe["parse"].(map[string]any)
			if !ok {
				return fmt.Errorf("parse config error: invalid field parse")
			}
			fetcher, ok := probe["fetch"].(map[string]any)
			if !ok {
				return fmt.Errorf("parse config error: invalid field fetch")
			}
			drawer, ok := probe["draw"].(map[string]any)
			if !ok {
				return fmt.Errorf("parse config error: invalid field draw")
			}
			tmp = append(tmp, Probe{
				Name:  name,
				Parse: parser,
				Fetch: fetcher,
				Draw:  drawer,
			})
		}
		Probes = tmp
	} else {
		Probes = nil
	}

	Interval = time.Duration(viper.GetInt("interval")) * time.Second
	ProbeTimeout = time.Duration(viper.GetInt("probeTimeout")) * time.Second
	MetadataRedirect = viper.GetString("metadataRedirect")
	Analysis = viper.GetBool("analysis")

	dur, err := time.ParseDuration(viper.GetString("uptime.store-duration"))
	if err != nil {
		return fmt.Errorf("parse uptime.store-duration fail:%v", err)
	}
	Uptime.StoreDuration = dur

	uptimeInterval := time.Duration(viper.GetInt("uptime.interval")) * time.Second
	Uptime.Interval = uptimeInterval

	Trace.Endpoint = viper.GetString("trace.endpoint")

	if UpdateCallBack != nil {
		UpdateCallBack()
	}
	log.Println("update config success")
	return nil
}
