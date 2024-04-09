package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"syscall/js"
	"time"

	pbs "github.com/joeldejesus1/wasm-go-1/proto/subway"
	"github.com/tarndt/wasmws"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	// console.log("Hello from Go!")
	js.Global().Get("console").Call("log", "Hello from Go!")

	ch := make(chan struct{})

	v := js.Global().Call("getTcpOverWsUrl")
	if v.Type() != js.TypeString {
		panic(fmt.Errorf("url not a string; it is a %s", v.Type().String()))
	}
	var websocketURL string = v.String() + "/grpc"

	// a grpc header differentiates grpc traffic
	// being compiled to wasm forces grpc.Dial to connect over http/1+websockets
	dialCtx, dialCancel := context.WithTimeout(ctx, 5*time.Second)
	defer dialCancel()
	conn, err := grpc.DialContext(
		dialCtx,
		"passthrough:///"+websocketURL,
		grpc.WithContextDialer(wasmws.GRPCDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	client := pbs.NewSubwayMapClient(conn)

	// in javascript, call:
	// submitDirectionRequest("A","B")
	// get returned ????
	js.Global().Set("submitDirectionRequest", SetupAccessor(func(mainArgs []js.Value) (interface{}, error) {
		// blocking is ok here
		log.Printf("in Go: submit direction request; args: %+v", mainArgs)
		if len(mainArgs) != 2 {
			return nil, errors.New("missing destination and source")
		}
		for _, arg := range mainArgs {
			if arg.Type() != js.TypeString {
				return nil, fmt.Errorf("got arg of type %s, not string", arg.Type().String())
			}
		}
		source, err := convertStation(mainArgs[0].String())
		if err != nil {
			return nil, err
		}
		destination, err := convertStation(mainArgs[1].String())
		if err != nil {
			return nil, err
		}

		return client.AskDirection(ctx, &pbs.DirectionRequest{
			Segment: &pbs.Segment{
				Source:      source,
				Destination: destination,
			},
		})
	}))

	// block forever, do not exit
	<-ch
}

func convertStation(s string) (pbs.Station, error) {
	x, present := pbs.Station_value[s]
	if !present {
		return pbs.Station_A, fmt.Errorf("station %s does not exist", s)
	}
	return pbs.Station(x), nil
}

// Convert a Go function into a Javascript promise
func SetupAccessor(cb func(a []js.Value) (interface{}, error)) js.Func {

	return js.FuncOf(func(this js.Value, mainArgs []js.Value) interface{} {

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				ans, err := cb(mainArgs)
				if err != nil {
					log.Print("invoking reject")
					reject.Invoke(js.ValueOf(err.Error()))
				} else {
					log.Print("invoking resolve")
					resolve.Invoke(js.ValueOf(ans))
				}
			}()
			// The handler of a Promise doesn't return any value
			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
