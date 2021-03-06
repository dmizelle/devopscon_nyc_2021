let
  sources = import ./nix/sources.nix;
  pkgs = import sources.nixpkgs {
    overlays = [
      (_: _: { inherit sources; })
    ];
  };
in
  with pkgs;
  pkgs.mkShell {
    buildInputs = [
      go
    ];
  }
