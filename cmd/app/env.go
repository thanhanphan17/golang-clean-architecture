package app

import (
	"os"
	"strconv"
)

// setEnv sets environment variables based on the provided flag arguments.
func setEnv(flagArg *FlagArgument) {
	setEnvVar("CONFIG_PATH", *flagArg.ConfigPath)
	setEnvVar("CONFIG_ENV", *flagArg.Env)
	setEnvVar("CONFIG_ALLOW_MIGRATION", strconv.FormatBool(*flagArg.Upgrade))
}

// setEnvVar sets the environment variable with the given name and value.
// It panics if there is an error setting the environment variable.
func setEnvVar(name, value string) {
	if err := os.Setenv(name, value); err != nil {
		panic(err)
	}
}
