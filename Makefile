all: docs/${DEMO}/main.wasm docs/${DEMO}/wasm_exec.js
docs/${DEMO}/main.wasm: docs/${DEMO} cmd/${DEMO}/main.go
	GOOS=js GOARCH=wasm go build -o docs/${DEMO}/main.wasm cmd/${DEMO}/main.go
docs/${DEMO}/wasm_exec.js: docs/${DEMO} $(shell go env GOROOT)/misc/wasm/wasm_exec.js
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js docs/${DEMO}/
docs/${DEMO}:
	mkdir -p docs/${DEMO}

clean:
	rm -rf docs/${DEMO}/main.wasm docs/${DEMO}/wasm_exec.js