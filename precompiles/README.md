# hypr-precompiles

## Compile

1. compile precompiles
``
cargo build --release
``
2. compile op-geth
``
env GO111MODULE=on go run build/ci.go install ./cmd/geth
``
3. compile all
``
make geth
``
4. get file
``
ls build/bin/geth 
ls target/release/libprecompiles.so 
``

## Functions

These functions need export as precompiles:

1. Anemoi [execute](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/anemoi.rs#L103) [gas](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/anemoi.rs#L121)
2. Anonymous Transfer [execute](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/anon/anonymous.rs#L35) [gas](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/anon/anonymous.rs#L29)
3. Shuffle
    1. Verifys [execute](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/shuffle/mod.rs#L149) [gas](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/shuffle/mod.rs#L157)
    2. Funcions [execute](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/shuffle/mod.rs#L187) [gas](https://github.com/HyprNetwork/rust-precompiles/blob/main/src/shuffle/mod.rs#L195)

## LICENSE

GPLv3.0
