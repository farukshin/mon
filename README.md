# Mon

![GitHub Release](https://img.shields.io/github/v/release/farukshin/mon)
![GitHub build status](https://github.com/farukshin/mon/actions/workflows/mon.yml/badge.svg)
![Codecov](https://img.shields.io/codecov/c/github/farukshin/mon)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/farukshin/mon/total?color=green)
![GitHub License](https://img.shields.io/github/license/farukshin/mon)


Система мониторинга

* [Installation](#Installation)
* * [Install from Releases](#InstallationFromReleases)
* * [Install from Source](#InstallationFromSource)
* [Sensors](#Sensors)
* * [Add Sensor](#AddSensor)
* * [Edit Sensor](#EditSensor)
* * [Delete Sensor](#DeleteSensor)


<a name="Installation"></a> 

## Installation

<a name="InstallationFromReleases"></a> 

### Install from Releases

1. Get the [latest release](https://github.com/farukshin/mon/releases) version.

``` bash
$ VERSION=$(curl -s "https://api.github.com/repos/farukshin/mon/releases/latest" | jq -r '.tag_name')
```
or set a specific version:

``` bash
$ VERSION=vX.Y.Z   # Version number with a leading v
```

2. Download the release.

``` bash
$ OS=Linux       # or Darwin, Windows
$ ARCH=x86_64    # or arm64, x86_64, armv6, i386, s390x
$ curl -sL "https://github.com/farukshin/mon/releases/download/${VERSION}/mon_${OS}_${ARCH}.tar.gz" > mon.tar.gz
```

3. Unpack it in the PATH.

``` bash
$ tar -zxvf mon.tar.gz mon
```

<a name="InstallationFromSource"></a> 

## Install from source

``` bash
git clone https://github.com/farukshin/mon.git
cd mon
go build .
./mon start
```

``` bash
curl localhost:1616
```

<a name="AddSensor"></a> 

## Add Sensor

From CLI

``` bash
./mon sensors list
./mon sensors add --type=httpcode --target=google.com
# > 27e92b12-0933-4b82-b2b9-96c1b64745a2
```

From API

``` bash
MON_SRV="localhost:1616"
curl -X POST $MON_SRV/api/sensors/add \
    -d '{"kind":"httpcode", "target":"google.com", "time":60}'
# > {ok:true, uid:"27e92b12-0933-4b82-b2b9-96c1b64745a2"}
```

<a name="EditSensor"></a> 

## Edit sensor

From CLI

``` bash
./mon sensors list
./mon sensors edit --uid=27e92b12-0933-4b82-b2b9-96c1b64745a2 --target=farukshin.com
# > 27e92b12-0933-4b82-b2b9-96c1b64745a2 change target, from google.com, to farukshin.com
```
From API

``` bash
MON_SRV="localhost:1616"
curl -X POST $MON_SRV/api/sensors/edit \
    -d '{"uid":"7e92b12-0933-4b82-b2b9-96c1b64745a2", "target":"farukshin.com"}'
# > {ok:true, uid:"27e92b12-0933-4b82-b2b9-96c1b64745a2"}
```


<a name="DeleteSensor"></a> 

## Delete sensor

From CLI

``` bash
./mon sensors list
./mon sensors delete --uid=27e92b12-0933-4b82-b2b9-96c1b64745a2
# > success
```

From API

``` bash
MON_SRV="localhost:1616"
curl -X POST $MON_SRV/api/sensors/delete \
    -d '{"uid":"7e92b12-0933-4b82-b2b9-96c1b64745a2"}'
# > {ok:true}
```