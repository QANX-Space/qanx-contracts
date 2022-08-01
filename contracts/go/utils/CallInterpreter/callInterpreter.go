package callintrpr

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Takes in a smart contract and calls it functions based on the arguments given
func Interpret(contract interface{}, args []string) {
	methodName := args[0]
	methodArguments := args[1:]

	method := reflect.ValueOf(contract).MethodByName(methodName)

	if !method.IsValid() {
		os.Stdout.WriteString(fmt.Sprintf("CallInterpreter: Method \"%v\" does not exist\n", methodName))
		os.Exit(1)
	}

	methodType := method.Type()

	inputs := getMethodInputs(methodType, methodArguments)

	if methodType.NumOut() > 0 {
		fmt.Printf("OUT=%v\n", method.Call(inputs))
	} else {
		method.Call(inputs)
	}

	os.Exit(0)
}

// Turns string arguments in to typed arguments
func getMethodInputs(methodType reflect.Type, methodArguments []string) []reflect.Value {
	inputs := make([]reflect.Value, methodType.NumIn())

	if len(inputs) != len(methodArguments) {
		os.Stdout.WriteString(fmt.Sprintf("CallInterpreter: Expected %v arguments received %v\n", len(inputs), len(methodArguments)))
		os.Exit(1)
	}

	for i := 0; i < methodType.NumIn(); i++ {
		t := methodType.In(i)
		arg := methodArguments[i]

		switch t.Kind() {
		default:
		case reflect.String:
			inputs[i] = reflect.ValueOf(arg)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(arg, 10, 64)

			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("CallInterpreter: Expected \"%v\" to be of type %v\n", arg, t))
				os.Exit(1)
			}

			inputs[i] = reflect.ValueOf(n).Convert(t)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n, err := strconv.ParseUint(arg, 10, 64)

			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("CallInterpreter: Expected \"%v\" to be of type %v\n", arg, t))
				os.Exit(1)
			}

			inputs[i] = reflect.ValueOf(n).Convert(t)
		case reflect.Bool:
			b, err := strconv.ParseBool(arg)

			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("CallInterpreter: Expected \"%v\" to be of type %v\n", arg, t))
				os.Exit(1)
			}

			inputs[i] = reflect.ValueOf(b).Convert(t)
		}
	}

	return inputs
}
