# miner-exporter

## Introdution
There are too many mining software and different apis, and it makes hard to metrify if different mining software used.
`miner-exporter` helps to collect mining statistics from each miner's api and provide as generalized form.

## Exported model
```
{
	"miner": {
		"name": "TeamRedMiner",
    "version": "0.8.6.2",
		"algorithm": "Ethash",
		"address": "0xdEf68436047BD07a65F5E9ffbd5944c052088293",
		"pool": "stratum+tcp://asia2.ethermine.org:5555",
		"uptime": 1169377
	},
	"devices": [{
		"gpu_id": 0,
		"name": "Vega 20 [Radeon VII]",
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
		"name": "Vega 20 [Radeon VII]",
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
$ miner-exporter --miner-type={gminer/trex/teamredminer} --host={your-miner-host} --protocol={http/https/tcp}
```

## Contribution
Please feel free to make pull request and issue. 🙂
