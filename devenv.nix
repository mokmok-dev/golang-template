{ pkgs, lib, ... }:

{
  env = {
    LOG_LEVEL = "debug";
    DATABASE_HOST = "localhost";
    DATABASE_PORT = "5432";
    DATABASE_USER = "pranc1ngpegasus";
    DATABASE_PASSWORD = "password";
    DATABASE_NAME = "golang-template";
    EMAIL_HOST = "localhost";
    EMAIL_PORT = "8025";
    EMAIL_USER = "";
    EMAIL_PASSWORD = "";
    EMAIL_SENDER = "temma.fukaya@mokmok.dev";
  };

  packages = with pkgs; [
    git
  ];

  scripts = { };

  enterShell = ''
    git --version
  '';

  languages = {
    nix.enable = true;
    go.enable = true;
  };

  pre-commit = { };

  processes = { };

  services = {
    mailhog = {
      enable = true;
    };

    postgres = {
      enable = true;
      package = pkgs.postgresql.withPackages (p: [ p.timescaledb ]);
      listen_addresses = "127.0.0.1";
      port = 5432;
      initialDatabases = [
        {
          name = "golang-template";
        }
      ];
      settings = {
        unix_socket_directories = lib.mkForce "";
      };
    };
  };
}
