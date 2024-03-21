with (import <nixpkgs> {});
rec {
  mopc-frontend = mkYarnPackage {
    name = "mopc-frontend";
    src = ./.;
    packageJSON = ./package.json;
    yarnLock = ./yarn.lock;
    yarnNix = ./yarn.nix;
  };
}