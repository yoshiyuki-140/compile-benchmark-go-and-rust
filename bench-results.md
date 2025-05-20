| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `cd go && go clean && go build -o guessing_game` | 493.3 ± 31.7 | 463.9 | 541.0 | 1.00 |
| `cd rust && cargo clean && cargo build --release` | 4190.5 ± 306.1 | 3930.1 | 4712.4 | 8.50 ± 0.83 |
| `cd rust && cargo clean && cargo build` | 3863.9 ± 281.1 | 3672.4 | 4358.0 | 7.83 ± 0.76 |
