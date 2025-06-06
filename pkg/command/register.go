package command

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func RegisterCommand(cmd *SCommand) *cobra.Command {
	min_args_count := cmd.GetArgsCount(true)
	max_args_count := cmd.GetArgsCount(false)
	args_count := cobra.ExactArgs(min_args_count)

	if min_args_count != max_args_count {
		args_count = cobra.RangeArgs(min_args_count, max_args_count)
	}

	command := &cobra.Command{
		Use:   cmd.Name,
		Short: cmd.Description,
		Long:  cmd.Description,
		Args:  args_count,
		Run: func(cc *cobra.Command, args []string) {
			for index, arg := range args {
				switch v := any(cmd.Args[index].Value).(type) {
				case *string:
					*v = arg
				case *bool:
					if val, err := strconv.ParseBool(arg); err == nil {
						*v = val
					} else {
						panic(fmt.Sprintf("parse %s command arg %s error, should is %T, but get %T", cmd.Name, cmd.Args[index].Name, v, val))
					}
				case *int:
					if val, err := strconv.ParseInt(arg, 10, 64); err == nil {
						*v = int(val)
					} else {
						panic(fmt.Sprintf("parse %s command arg %s error, should is %T, but get %T", cmd.Name, cmd.Args[index].Name, v, val))
					}
				case *float64:
					if val, err := strconv.ParseFloat(arg, 64); err == nil {
						*v = val
					} else {
						panic(fmt.Sprintf("parse %s command arg %s error, should is %T, but get %T", cmd.Name, cmd.Args[index].Name, v, val))
					}
				case *float32:
					if val, err := strconv.ParseFloat(arg, 32); err == nil {
						*v = float32(val)
					} else {
						panic(fmt.Sprintf("parse %s command arg %s error, should is %T, but get %T", cmd.Name, cmd.Args[index].Name, v, val))
					}
				}
			}
			cmd.Handle(cmd)
		},
	}

	for _, flag := range cmd.Flags {
		switch v := any(flag.Value).(type) {
		case *bool:
			command.Flags().BoolVar(v, flag.Name, *v, flag.Description)
		case *[]bool:
			command.Flags().BoolSliceVar(v, flag.Name, *v, flag.Description)
		case *int:
			command.Flags().IntVar(v, flag.Name, *v, flag.Description)
		case *[]int:
			command.Flags().IntSliceVar(v, flag.Name, *v, flag.Description)
		case *float32:
			command.Flags().Float32Var(v, flag.Name, *v, flag.Description)
		case *[]float32:
			command.Flags().Float32SliceVar(v, flag.Name, *v, flag.Description)
		case *float64:
			command.Flags().Float64Var(v, flag.Name, *v, flag.Description)
		case *[]float64:
			command.Flags().Float64SliceVar(v, flag.Name, *v, flag.Description)
		case *time.Duration:
			command.Flags().DurationVar(v, flag.Name, *v, flag.Description)
		case *[]time.Duration:
			command.Flags().DurationSliceVar(v, flag.Name, *v, flag.Description)
		case *string:
			command.Flags().StringVar(v, flag.Name, *v, flag.Description)
		case *[]string:
			command.Flags().StringSliceVar(v, flag.Name, *v, flag.Description)
		default:
			panic(fmt.Sprintf("not support command %s flag %s type: %T", cmd.Name, flag.Name, v))
		}
	}

	return command
}
