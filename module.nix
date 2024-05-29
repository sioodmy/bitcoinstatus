inputs: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (pkgs.stdenv.hostPlatform) system;
  cfg = config.services.bitcoinstatus;

  package = inputs.self.packages.${system}.default;
  inherit (lib) mkOption mkEnableOption types mkIf;
in {
  options.services.bitcoinstatus = {
    enable = mkEnableOption "Bitcoin status for discord";
    package = mkOption {
      type = types.package;
      default = package;
      example = package;
      description = "bitcoinstatus package";
    };
    tokenFile = mkOption {
      type = types.str;
      description = "Discord user token";
    };
  };
  config = mkIf cfg.enable {
    systemd.services.bitcoinstatus = {
      description = "Bitcoin status for discord";
      wantedBy = ["multi-user.target"];
      wants = ["network.target"];
      after = [
        "network-online.target"
        "NetworkManager.service"
        "systemd-resolved.service"
      ];
      serviceConfig = {
        ExecStart = ''${cfg.package}/bin/bitcoinstatus'';
        Restart = "always";
        Environment = ''TOKEN_PATH=${cfg.tokenFile}'';
      };
    };
  };
}
