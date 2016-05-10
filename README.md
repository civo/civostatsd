# Civo Statsd

[![Build Status](https://img.shields.io/travis/absolutedevops/civostatsd.svg?style=flat-square&label=build)](https://travis-ci.org/absolutedevops/civostatsd)

This utility is a small Go application that runs in the background on Civo instances, collecting CPU, disk and memory
usage and reporting them back to the Civo API so that they can be reported on in the control panel.  It's open-source
because we don't want users to think we're doing something underhanded and if you like, you can simply completely
remove it from your instance with no detrimental side-effects (except for a warning in your control panel to say that
we can't access the stats).

## Removing Civo Statsd

To remove this utility from your instance, you can simply SSH in and run this command:

```
civostatsd --uninstall
```

## Internals

To be documented as they're decided...
