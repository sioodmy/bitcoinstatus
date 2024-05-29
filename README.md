# Bitcoin status for discord

Changes your discord status to current Bitcoin price fetched from binance API, updates randomly every
100-1440 seconds.

## Usage 

Add NixOS module to your configuration and enable service.

```nix
services.bitcoinstatus = {
  enable = true;
  tokenFile = secrets.discordtoken.path;
};
```

## Warning

This is against Discord ToS and you will (or not) get banned.

Use at your own risk!

