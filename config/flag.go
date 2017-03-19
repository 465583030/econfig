package config

import (
	"time"

	"github.com/spf13/pflag"
)

// GlobalFlag all flags of command
type GlobalFlag struct {
	// full path of config file
	CfgFile string

	// redis, mysql, etcd, consul etc..
	Backend string

	// process which app
	App string

	Verbose bool

	LogFile string
}

type ServeFlag struct {
	Interval  time.Duration
	Advertise string
	RPCAddr   string
	RPCAuth   string
	Bind      string
	Node      string
	Join      string
	HttpPort  int
}

func NewServeFlag(fs *pflag.FlagSet) *ServeFlag {
	gf := &ServeFlag{}
	gf.Interval = 60 * time.Second
	fs.DurationVar(&gf.Interval, "interval", gf.Interval, "Seconds for polling interval")
	fs.StringVar(&gf.Node, "node", "", "Name of the node,must be unique in the cluster. By default this is the hostname of the machine.")
	fs.StringVar(&gf.Bind, "bind", "0.0.0.0:7946", `The address that econfig will bind to for communication with other econfig nodes."`)
	fs.StringVar(&gf.Advertise, "advertise", "", `The advertise flag is used to change the address that we advertise to other nodes in the cluster.By default, the bind address is advertised.`)
	fs.StringVar(&gf.RPCAddr, "rpc-addr", "127.0.0.1:7373", `The address and port used for RPC communications`)
	fs.StringVar(&gf.RPCAuth, "rpc-auth", "", `RPC auth token`)
	fs.StringVar(&gf.Join, "join", "", "Address of another agent to join upon starting up")
	fs.IntVar(&gf.HttpPort, "http-port", 8080, "The port where the webserver.")
	return gf
}

func NewGlobalFlag(fs *pflag.FlagSet) *GlobalFlag {
	gf := &GlobalFlag{}
	fs.StringVar(&gf.LogFile, "logfile", "", "output log to the logfile. By default output to stdout")
	fs.StringVarP(&gf.CfgFile, "config", "c", "", "config file")
	fs.StringVar(&gf.Backend, "backend", "redis", "Backend storage to use, redis, mysql, postgres, etcd, consul.")
	fs.StringVar(&gf.App, "app", "", "command work only this app")
	fs.BoolVarP(&gf.Verbose, "verbose", "v", false, "if ouput detail log")
	return gf
}
