package main

import (
	"flag"
	"log"
	"os"

  capsulehttp "github.com/bots-garden/capsule/services/http"
  capsulehttpecho "github.com/bots-garden/capsule/services/http-echo"
	capsulecli "github.com/bots-garden/capsule/services/cli"

)

type CapsuleFlags struct {
	mode     string
	param    string
	wasm     string
	httpPort string
}

func main() {

	//flags
	/*
		go run main.go \
			-wasm=./wasm_modules/capsule-function-template/hello.wasm \
			-mode=cli \
			-param="👋 hello world 🌍"

		go run main.go \
			-wasm=./wasm_modules/capsule-function-template/hello.wasm \
			-mode=http \
			-httpPort=7070
	*/
	capsuleModePtr := flag.String("mode", "http", "default mode is http else: cli")
	stringParamPtr := flag.String("param", "hello world", "string parameter for the cli mode")
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.String("httpPort", "8080", "http port")

	flag.Parse()

	flags := CapsuleFlags{
		*capsuleModePtr,
		*stringParamPtr,
		*wasmFilePathPtr,
		*httpPortPtr,
	}
	//fmt.Println(flags)

	//argsWithProg := os.Args
	//wasmModuleFilePath := os.Args[1:][0]
	wasmModuleFilePath := flags.wasm

	// 📂 Load from file and then Instantiate a WebAssembly module
	wasmFile, errLoadWasmFile := os.ReadFile(wasmModuleFilePath)

	if errLoadWasmFile != nil {
		log.Panicln("🔴 Error while loading the wasm file:", errLoadWasmFile)
	}

	switch what := flags.mode; what {
	case "http-gin":
    capsulehttp.Serve(flags.httpPort, wasmFile)

  case "http-echo":
    capsulehttpecho.Serve(flags.httpPort, wasmFile)

	case "cli":
		capsulecli.Execute(flags.param, wasmFile)

	default:
		log.Panicln("🔴 bad mode", *capsuleModePtr)
	}

}
