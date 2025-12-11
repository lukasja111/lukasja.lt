#!/bin/bash
set -e
cd christmas
make build
cd ..
mkdir -p static/christmas
cp christmas/web/christmas.wasm christmas/web/wasm_exec.js static/christmas/
zola build
