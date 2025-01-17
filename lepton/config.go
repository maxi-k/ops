package lepton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/nanovms/ops/types"
)

// NewConfig construct instance of Config with default values
func NewConfig() *types.Config {
	c := new(types.Config)

	conf := os.Getenv("OPS_DEFAULT_CONFIG")
	if conf != "" {
		data, err := ioutil.ReadFile(conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
			os.Exit(1)
		}
		err = json.Unmarshal(data, &c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error config: %v\n", err)
			os.Exit(1)
		}
	} else {
		usr, err := user.Current()
		if err != nil {
			return c
		}
		conf = usr.HomeDir + "/.opsrc"

		if _, err = os.Stat(conf); err == nil {
			data, err := ioutil.ReadFile(conf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
				os.Exit(1)
			}
			err = json.Unmarshal(data, &c)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error config: %v\n", err)
				os.Exit(1)
			}
		}
	}

	c.RunConfig.Accel = true
	c.RunConfig.Memory = "2G"
	c.VolumesDir = LocalVolumeDir

	return c
}
