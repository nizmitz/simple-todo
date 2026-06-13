{
  description = "go-todo dev environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.aarch64-darwin; # or aarch64-darwin for Apple Silicon
    in
    {
      devShells.aarch64-darwin.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
          pkgs.gopls
          pkgs.gotools
          pkgs.nixd
          pkgs.golangci-lint
          pkgs.pre-commit
        ];
      };
    };
}
