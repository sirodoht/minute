with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    go
  ];
  shellHook = ''
    export GOPATH="$(pwd)/go";
  '';
}
