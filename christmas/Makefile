.PHONY: build clean serve desktop

# Build WASM for web
build:
	@echo "Building WASM..."
	@GOOS=js GOARCH=wasm go build -o web/christmas.wasm .
	@if [ -f "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ]; then \
		cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/; \
	elif [ -f "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ]; then \
		cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" web/; \
	fi
	@echo "Done! Files in web/"

# Build desktop version
desktop:
	@go build -o christmas .
	@echo "Run with: ./christmas"

# Clean build artifacts
clean:
	rm -f web/christmas.wasm web/wasm_exec.js christmas

# Serve locally for testing
serve: build
	@echo "Starting server at http://localhost:8080"
	@cd web && python3 -m http.server 8080
