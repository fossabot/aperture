---
title: Blueprints
description: Policies and Dashboards pre-packaged as reusable Blueprints
keywords:
  - jsonnet
  - grafana
sidebar_position: 3
---

# Policy Blueprints

## Introduction

Aperture comes with a pre-packaged list of [Aperture Policies][policies] and Grafana Dashboards that
can be used both as a guide for creating new Policies, and as ready-to-use
Blueprints for generating Aperture Policies customized to a [Service][service].

All Dashboards and Policies are written using the [Jsonnet][jsonnet-lang]
language, and can be used both as jsonnet mixins or as standalone Blueprints.

[jsonnet-lang]: https://jsonnet.org

## Initial Setup

All Blueprints are available from a separate [repository][aperture-blueprints].
See the repository [README.md][blueprints-readme] for the list of required tools
and instructions on installing jsonnet dependencies with the help of a [jsonnet
bundler][jb].

The Blueprint Generator (used to generate Policy files from Blueprints) also
depends on Python 3.8+ and [jsonnet][go-jsonnet].

[aperture-blueprints]: https://github.com/fluxninja/aperture-blueprints
[blueprints-readme]: https://github.com/fluxninja/aperture-blueprints/blob/main/README.md
[jb]: https://github.com/jsonnet-bundler/jsonnet-bundler
[go-jsonnet]: https://github.com/google/go-jsonnet

## Generating Aperture Policies and Grafana Dashboards

The simplest way to use the Blueprints repository is to render Blueprints into
Policy and dashboard files.

To generate files, `scripts/aperture-generate.py` can be used:

```sh
$ ./scripts/aperture-generate.py --help
usage: aperture-generate.py [-h] [--verbose] [--output OUTPUT] [--config CONFIG] BLUEPRINT

Aperture Policies & Dashboards generator utility.

This utility can be used to generate Aperture Policies and Grafana Dashboards "in-place". Check [aperture-blueprint's README.md](https://github.com/fluxninja/aperture-blueprints/blob/main/README.md) for more
details.

positional arguments:
  BLUEPRINT        Aperture blueprint path

options:
  -h, --help       show this help message and exit
  --verbose        Whether to log verbose messages to stderr
  --output OUTPUT  Output directory for json files
  --config CONFIG  jsonnet file with blueprint configuration
```

This script takes as options an output directory path where files will be
saved and a path to a `config.libsonnet` file with local blueprint
configuration. It also takes the BLUEPRINT argument, which is a path to the
blueprint under the `blueprints/` directory.

Under the `blueprints/` directory, the currently available Blueprints can be
found. Each blueprint consists of at least two files: `config.libsonnet` and
`main.libsonnet`. `main.libsonnet` bundles actual Policy and dashboard code
(available under `lib/1.0`) into Blueprints, and `config.libsonnet` comes with
the default configuration for the given Policy. This can be overridden by the
`--config` option passed to the `aperture-generate.py` script.

Custom configurations will be merged with Blueprints' `config.libsonnet`
resulting in the final configuration, according to jsonnet language rules: keys
can be overwritten by reusing them in the custom configuration and nested
objects can be merged by using `+:` operator. Check the `examples/` directory
for more information.

The full command using the demoapp-latency-grand example looks like this:

```sh
jb install
./scripts/aperture-generate.py --output _gen --config examples/demoapp-latency-gradient.jsonnet Blueprints/latency-gradient
```

## Using aperture-blueprints as a jsonnet mixins library

An alternate way of using the aperture-blueprints repository is to import it
from another jsonnet project and render Policies or Dashboards directly in
jsonnet.

For example, to create a ConfigMap with Aperture Policies that can be loaded by
the controller, you need to install aperture-blueprints with the jsonnet
bundler:

```sh
jb install github.com/fluxninja/aperture-blueprints@main
```

Additionally, for this example to work install the k8s-libsonnet dependency:

```sh
jb install github.com/jsonnet-libs/k8s-libsonnet/1.24@main
```

Finally, you can create a ConfigMap resource with Policy like this:

```jsonnet
local k = import "github.com/jsonnet-libs/k8s-libsonnet/1.24/main.libsonnet";

local latencyGradientPolicy = import "github.com/fluxninja/aperture-blueprints/lib/1.0/policies/latency-gradient.libsonnet";

local policy = latencyGradientPolicy({
  policyName: "service1-demo-app",
  serviceSelector+: {
    service: "service1-demo-app.demoapp.svc.cluster.local"
  },
}).policy;

[
    k.core.v1.configMap.new("policies")
 + k.core.v1.configMap.metadata.withLabels({ "fluxninja.com/validate": "true"})
 + k.core.v1.configMap.withData({
   "service1-demo-app.yaml": std.manifestYamlDoc(policy, quote_keys=false)
 })
]
```

And then, render it with [jsonnet][jsonnet]:

```sh
jsonnet --yaml-stream -J vendor [example file].jsonnet
```

This can be also integrated with other Kubernetes deployment tools like
[tanka][tk].

[jsonnet]: https://github.com/google/go-jsonnet
[tk]: https://grafana.com/oss/tanka/
[policies]: /concepts/policy/policy.md
[service]: /concepts/service.md
