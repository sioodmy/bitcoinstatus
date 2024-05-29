{
  description = "Bitcoin status for discord";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      imports = [inputs.treefmt-nix.flakeModule];
      perSystem = {
        config,
        pkgs,
        ...
      }: {
        packages = rec {
          bitcoinstatus = pkgs.callPackage ./default.nix {};
          default = bitcoinstatus;
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            sass
            gnumake
            go
          ];
        };
        treefmt.config = {
          projectRootFile = "flake.nix";
          programs = {
            alejandra.enable = true;
            deadnix.enable = true;
            gofumpt.enable = true;
            statix.enable = true;
          };
        };
      };
      flake = {
        nixosModules.default = import ./module.nix inputs;
      };
    };
}
