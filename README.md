# Experiments with spf13's comprehensive command line app building library

## Initialising a cobra command line application
`cobra-cli init`

## Adding a cmd
Use cobra-cli add.

## cobra framework execution tree
1. main.go: cmd.Execute
1. cmd/root.go: rootCmd.Execute
1. cobra/command.go: ExecuteC
1. ExecuteC runs prehooks, posthook, parent cmd, child cmd checks then cmd.execute
1. execute parses options, preRun, postRun and Run.

## related configuration management project viper
`https://github.com/spf13/viper`
