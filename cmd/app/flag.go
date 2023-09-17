package app

import "flag"

type FlagArgument struct {
	ConfigPath *string
	Env        *string
	Upgrade    *bool
}

func getFlagArgument() *FlagArgument {
	return &FlagArgument{
		ConfigPath: flag.String("config", "", "application config path"),
		Env:        flag.String("env", "", "application env"),
		Upgrade:    flag.Bool("upgrade", false, "upgrade database"),
	}
}
