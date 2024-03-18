# Conduit Processor Examples

This repository contains two standalone processors that can be used in a [Conduit](https://github.com/ConduitIO/conduit) 
pipeline, one uses the simple way to create a processor, the other uses the full processor approach.

To build the processors and get the WASM files, run `make` under the package containing the processor. The resulting 
WASM file will be created under the same package.

Example:
````
cd simple
make
````

For more details on how to build your own processor, how to run it, or how it works, check [Standalone Processors](https://conduit.io/docs/processors/standalone).
