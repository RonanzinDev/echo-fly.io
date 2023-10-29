{nixpkgs ? import <nixpkgs> {}}:
with nixpkgs;

stdenv.mkDerivation {
    name = "echo-fly.io";
    buildInputs = [
        go
    ];
    shellHook = ''
        go mod tidy
    '';
}
