package command

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func RegisterCommand(cmd *SCommand) *cobra.Command {
	min_args_count := cmd.GetArgsCount(false)
	max_args_count := cmd.GetArgsCount(true)
	args_count := cobra.ExactArgs(min_args_count)

	if min_args_count != max_args_count {
		args_count = cobra.RangeArgs(min_args_count, max_args_count)
	}

	panicParseError := func(cmd string, arg string, v any, getV any) {
		panic(fmt.Sprintf("parse %s command arg %s error, should is %T, but get %T", cmd, arg, v, getV))
	}

	command := &cobra.Command{
		Use:   cmd.Name,
		Short: cmd.Description,
		Long:  cmd.Description,
		Args:  args_count,
		Run: func(cc *cobra.Command, args []string) {
			for index, arg := range args {
				item := cmd.Args[index]
				name := item.GetName()

				switch v := item.GetValue().(type) {
				case *string:
					*v = arg
				case *bool:
					if val, err := strconv.ParseBool(arg); err == nil {
						*v = val
					} else {
						panicParseError(cmd.Name, name, v, val)
					}
				case *int:
					if val, err := strconv.ParseInt(arg, 10, 64); err == nil {
						*v = int(val)
					} else {
						panicParseError(cmd.Name, name, v, val)
					}
				case *float64:
					if val, err := strconv.ParseFloat(arg, 64); err == nil {
						*v = val
					} else {
						panicParseError(cmd.Name, name, v, val)
					}
				case *float32:
					if val, err := strconv.ParseFloat(arg, 32); err == nil {
						*v = float32(val)
					} else {
						panicParseError(cmd.Name, name, v, val)
					}
				}
			}
			cmd.Handle(cmd)
		},
	}

	for _, flag := range cmd.Flags {
		name := flag.GetName()
		short := flag.GetShort()
		desc := flag.GetDescription()

		switch v := any(flag.GetValue()).(type) {
		case *bool:
			command.Flags().BoolVarP(v, name, short, *v, desc)
		case *[]bool:
			command.Flags().BoolSliceVarP(v, name, short, *v, desc)
		case *int:
			command.Flags().IntVarP(v, name, short, *v, desc)
		case *[]int:
			command.Flags().IntSliceVarP(v, name, short, *v, desc)
		case *float32:
			command.Flags().Float32VarP(v, name, short, *v, desc)
		case *[]float32:
			command.Flags().Float32SliceVarP(v, name, short, *v, desc)
		case *float64:
			command.Flags().Float64VarP(v, name, short, *v, desc)
		case *[]float64:
			command.Flags().Float64SliceVarP(v, name, short, *v, desc)
		case *time.Duration:
			command.Flags().DurationVarP(v, name, short, *v, desc)
		case *[]time.Duration:
			command.Flags().DurationSliceVarP(v, name, short, *v, desc)
		case *string:
			command.Flags().StringVarP(v, name, short, *v, desc)
		case *[]string:
			command.Flags().StringSliceVarP(v, name, short, *v, desc)
		default:
			panic(fmt.Sprintf("not support command %s flag %s type: %T", cmd.Name, name, v))
		}
	}

	for _, subc := range cmd.SubCommand {
		command.AddCommand(RegisterCommand(subc))
	}

	return command
}
