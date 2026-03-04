{
  description = "Go environment";
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
      supportedSystems = [ "x86_64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          rpnc = pkgs.buildGoModule {
            pname = "rpnc";
            inherit version;
            src = pkgs.fetchFromGitHub {
              owner = "tlinden";
              repo = "rpnc";
              rev = "v2.1.0";
              sha256 = "sha256-CojP1a19b2zKfUMp+wN7FFs+SzSoc8sYqKvXTg4RnOA=";
            };
            vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
          };
        }
      );

      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              go-tools
              golangci-lint
            ];
          };
        }
      );

      defaultPackage = forAllSystems (system: self.packages.${system}.rpnc);
    };
}
