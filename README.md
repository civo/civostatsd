# Civo Statsd

[![Build Status](https://img.shields.io/travis/absolutedevops/civostatsd.svg?style=flat-square&label=build)](https://travis-ci.org/absolutedevops/civostatsd)

## Introduction

This utility is a small Go application that runs in the background on Civo instances, collecting CPU, disk and memory
usage and reporting them back to the Civo API so that they can be reported on in the control panel.  It's open-source
because we don't want users to think we're doing something underhanded and if you like, you can simply completely
remove it from your instance with no detrimental side-effects (except for a warning in your control panel to say that
we can't access the stats).

**STATUS:** This project is currently under active development and maintainance.

## Compiling

To compile this project on a Mac for the Linux AMD64 target that we use, you can run the following command:

```
GOOS="linux" GOARCH="amd64" go build  -ldflags="-s -w" -o civostatsd main.go
```

This produces a stripped binary which is much smaller than default (6.5MB down to 4.7MB). Then you can use the `upx`
utility (`brew install upx`) to compress it further, down to ~1.7MB.

## Removing Civo Statsd

To remove this utility from your instance, you can simply SSH in and run this command:

```
civostatsd --uninstall
```

## Internals

Every minute the daemon looks up the current CPU usage (as an overall percentage), the memory usage (excluding caches)
and the disk usage and sends them all over HTTPS to the Civo API using the unique token stored in `/etc/civostatsd.json`.

It retrieves the stats using [shirou](https://github.com/shirou/gopsutil) library.
