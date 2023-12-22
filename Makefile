all: dist/${DEMO}/main.wasm dist/${DEMO}/wasm_exec.js
dist/${DEMO}/main.wasm: dist/${DEMO} cmd/${DEMO}/main.go
	GOOS=js GOARCH=wasm go build -o dist/${DEMO}/main.wasm cmd/${DEMO}/main.go
dist/${DEMO}/wasm_exec.js: dist/${DEMO} $(shell go env GOROOT)/misc/wasm/wasm_exec.js
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js dist/${DEMO}/
dist/${DEMO}:
	mkdir -p dist/${DEMO}

clean:
	rm -rf dist/${DEMO}/main.wasm dist/${DEMO}/wasm_exec.js