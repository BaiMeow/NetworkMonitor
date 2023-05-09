package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var Probes []Probe
var Interval int

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	probes := viper.Get("probe").([]any)
	Interval = viper.GetInt("interval")
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
		Probes = append(Probes, Probe{
			Parse: parser,
			Fetch: fetcher,
		})
	}
	return nil
}
