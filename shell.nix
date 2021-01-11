{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/a3ab47ec9067b5f9fccda506fc8641484c3d8e73.tar.gz") {} }:

with pkgs;

mkShell {
  # Sets the build inputs, i.e. what will be available in our
  # local environment.
  buildInputs = [
    go
  ];

  shellHook = ''
    export GO111MODULE=on
    export GOPRIVATE="github.com/teamexos"
  '';
}
