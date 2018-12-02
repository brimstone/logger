package logger_test

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/brimstone/logger"
)

func parseSig(sig string) ([]string, error) {
	sig = strings.TrimLeft(sig, "func(")
	sig = strings.TrimRight(sig, ")")
	args := strings.Split(sig, ", ")
	args = args[1:]
	return args, nil
}

func CompareTypes(a interface{}, b interface{}) bool {
	success := true
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	bMethods := make(map[string][]string)
	for i := 1; i < bType.NumMethod(); i++ {
		bMethod := bType.Method(i)
		bName := bMethod.Name
		bSig, err := parseSig(bMethod.Type.String())
		if err != nil {
			panic(err)
		}
		bMethods[bName] = bSig
	}

	for i := 1; i < aType.NumMethod(); i++ {
		aMethod := aType.Method(i)
		aName := aMethod.Name
		aSig, err := parseSig(aMethod.Type.String())
		if err != nil {
			panic(err)
		}
		bSig, ok := bMethods[aName]
		// If the method in A doesn't exist in B, it's a problem
		if !ok {
			fmt.Printf("Method %s%s is missing\n", aName, aSig)
			success = false
			continue
		}
		// If signatures don't have the same number of arguments
		if len(aSig) != len(bSig) {
			fmt.Printf("Method %s(%s) has mismatched interfaces, got: %s\n", aName, strings.Join(aSig, ", "), strings.Join(bSig, ", "))
			success = false
			continue
		}
		// If signatures don't match
		for aArgI, aArgV := range aSig {
			if aArgV != bSig[aArgI] {
				fmt.Printf("Method %s(%s) has mismatched interfaces, got: %s\n", aName, strings.Join(aSig, ", "), strings.Join(bSig, ", "))
				success = false
			}
		}

	}
	return success
}
func TestInterface(t *testing.T) {
	stdLogger := log.New(os.Stdout, "logger: ", log.Lshortfile)

	brimstoneLogger := logger.New()
	if !CompareTypes(stdLogger, brimstoneLogger) {
		t.Fatal("Types don't match")
	}
}
