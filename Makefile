# Makefile

GO_DIR := go
RUST_DIR := rust

.PHONY: bench clean_go clean_rust build_go build_rust

bench:
	@echo "🚀 Running compile-time benchmarks with hyperfine..."
	@hyperfine \
	  --warmup 3 \
	  --runs 5 \
	  --show-output \
	  --export-markdown bench-results.md \
	  "cd $(GO_DIR) && go clean && go build -o guessing_game" \
	  "cd $(RUST_DIR) && cargo clean && cargo build --release"

# 個別に実行したいとき用
clean_go:
	@cd $(GO_DIR) && go clean

build_go:
	@cd $(GO_DIR) && go build -o guessing_game

clean_rust:
	@cd $(RUST_DIR) && cargo clean

build_rust:
	@cd $(RUST_DIR) && cargo build --release
