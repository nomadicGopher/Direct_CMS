declare class Go {
  constructor();
  run(instance: WebAssembly.Instance): void;
  importObject: any; // ? Specify a more precise type if known
}

type WasmInstantiatedSource = WebAssembly.WebAssemblyInstantiatedSource;

const go: Go = new Go();

function init(wasmObj: WasmInstantiatedSource): void {
  go.run(wasmObj.instance);
}

if ('instantiateStreaming' in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject)
      .then((wasmObj: WasmInstantiatedSource) => {
          init(wasmObj);
      });
} else {
  fetch("go.wasm")
      .then((resp: Response) => resp.arrayBuffer())
      .then((bytes: ArrayBuffer) =>
          WebAssembly.instantiate(bytes, go.importObject)
              .then((wasmObj: WasmInstantiatedSource) => {
                  init(wasmObj);
              })
      );
}