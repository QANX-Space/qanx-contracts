package callintrpr

import (
	"fmt"
	"math/big"
	"os"
	"reflect"
	"strconv"

	common "qanx.space/qanx-contracts/go/utils/Common"
)

// Takes in a smart contract and calls its functions based on the arguments given
func Interpret(contract interface{}, args []string) {
	if len(args) == 0 {
		os.Stderr.WriteString("CallInterpreter: Method name was not provided\n")
		os.Exit(1)
	}

	methodName := args[0]
	methodArguments := args[1:]

	method := reflect.ValueOf(contract).MethodByName(methodName)

	if !method.IsValid() {
		os.Stderr.WriteString(fmt.Sprintf("CallInterpreter: Method \"%v\" does not exist\n", methodName))
		os.Exit(1)
	}

	methodType := method.Type()
	inputs := getMethodInputs(methodType, methodArguments)

	if methodType.NumOut() > 1 { // Multiple return values
		response := method.Call(inputs)
		output := make([]string, len(response))

		for i := 0; i < len(output); i++ {
			output[i] = fmt.Sprintf("%v", response[i])
		}

		fmt.Printf("OUT=%v\n", output)
	} else if methodType.NumOut() > 0 { // Single return value
		fmt.Printf("OUT=%v\n", method.Call(inputs)[0])
	} else { // No return value
		method.Call(inputs)
	}

	os.Exit(0)
}

// Turns string arguments in to typed arguments
func getMethodInputs(methodType reflect.Type, methodArguments []string) []reflect.Value {
	inputs := make([]reflect.Value, methodType.NumIn())

	if len(inputs) != len(methodArguments) {
		os.Stderr.WriteString(fmt.Sprintf("CallInterpreter: Expected %v arguments received %v\n", len(inputs), len(methodArguments)))
		os.Exit(1)
	}

	for i := 0; i < methodType.NumIn(); i++ {
		t := methodType.In(i)
		arg := methodArguments[i]

		switch t.Kind() {
		default:
		case reflect.String:
			fmt.Println(t.Name())

			inputs[i] = reflect.ValueOf(arg)

		case reflect.Ptr:
			if t.Elem() == reflect.TypeOf(big.Int{}) {
				n, ok := common.ParseBig256(arg)

				if ok {
					inputs[i] = reflect.ValueOf(n)
					continue
				}
			}

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
