# miner-exporter

## Introdution
There are too many mining software and different apis, and it makes hard to metrify if different mining software used in mixed.
`miner-exporter` helps to collect mining statistics from each miner's api and provide as generalized form.

## Supported output format
- Prometheus (use `--output-format=prometheus` to enable)
- JSON

## Supported Miners
- GMiner (https://github.com/develsoftware/GMinerRelease)
- T-Rex (https://github.com/trexminer/T-Rex)
- Team Red Miner (https://github.com/todxx/teamredminer)
- NBMiner (https://github.com/NebuTech/NBMiner)

## Exported model (JSON)
```
{
  "miner": {
    "name": "TeamRedMiner",
    "version": "0.8.6.2",
    "algorithm": "ethash",
    "address": "0xdEf68436047BD07a65F5E9ffbd5944c052088293",
    "pool": "asia2.ethermine.org:5555",
    "uptime": 1169377,
    "worker": "J41",
  },
  "devices": [{
    "gpu_id": 0,
    "name": "VII",
    "hashrate": 101780000,
    "fan_speed": 97,
    "core_temp": 52,
    "memory_temp": 63,
    "power_consumption": 213,
    "share_accepted": 32153,
    "share_rejected": 0,
    "share_stale": 0
  }, {
    "gpu_id": 1,
    "name": "Vega20[VII]",
    "hashrate": 102910000,
    "fan_speed": 97,
    "core_temp": 55,
    "memory_temp": 67,
    "power_consumption": 229,
    "share_accepted": 32539,
    "share_rejected": 0,
    "share_stale": 0
  }]
}
```

## Build
```
# make sure that you have installed go 1.18
$ env GOOS={your-os} GOARCH={your-arch} go build
```

## Run
```
$ miner-exporter --miner-type={gminer/trex/teamredminer/nbminer} --miner-host={your-miner-host} --miner-protocol={http/https/tcp} --miner-port={your-miner-port} --listen-port={port-to-listen} --output-format={json/prometheus}
```

## Contribution
Please feel free to make pull request and issue. 🙂
