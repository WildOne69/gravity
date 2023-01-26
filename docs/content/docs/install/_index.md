---
title: "Installation"
weight: 5
---

### Installation

#### docker-compose

```yaml
---
version: '3.4'

services:
  gravity:
    # Important for this to be static and unique
    hostname: gravity1
    image: ghcr.io/beryju/gravity:latest
    restart: unless-stopped
    network_mode: host
    volumes:
      - data:/data
    # The default log level of info logs DHCP and DNS queries, so ensure
    # the logs aren't filling up the disk
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"

volumes:
  data:
    driver: local
```

### Configuration

The following environment variables can be set:

- `DEBUG`: Enable debug mode
- `LOG_LEVEL`: Log level, defaults to `info`.
- `DATA_PATH`: Path to store etcd data, defaults to `./data`
- `ETCD_PREFIX`: Global etcd prefix, defaults to `/gravity`
- `ETCD_ENDPOINT`: etcd Client endpoint, defaults to `localhost:2379` when using embedded etcd
- `ETCD_JOIN_CLUSTER`: Used when joining a node to a cluster, value is given by join API endpoint
- `BOOTSTRAP_ROLES`: Configure while roles this instance should bootstrap, defaults to `dns;dhcp;api;etcd;discovery;backup;monitoring;debug;tsdb`.
- `INSTANCE_IDENTIFIER`: Unique identifier of an instance, should ideally not change. Defaults to hostname, but needs to be set in containers.
- `INSTANCE_IP`: This instance's reachable IP, when running in docker this should be the hosts IP
- `INSTANCE_LISTEN`: By default the instance will listen on `INSTANCE_IP`, but can be set to override that (set to 0.0.0.0 in docker)
- `LISTEN_ONLY`: Enable listen-only mode which will not reply to any DHCP packets and not run discovery
- `ADMIN_PASSWORD`: Optionally set a default password for the admin user, if not set a random one will be generated
- `ADMIN_TOKEN`: Optionally set a token to be created on first start, if not set no token will be created

## First time use

A default admin user is created on the first startup. You can find the credentials printed to stdout.