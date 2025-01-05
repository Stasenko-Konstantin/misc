#!/bin/sh

cargo build --release
cp target/release/count ~/bin/count