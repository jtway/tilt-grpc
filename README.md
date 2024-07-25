# tilt-proxy

Unofficial Golang based service to read data from Tilt Hydrometers and expose
data via various endpoints.

## Motivation

Needing to keep a phone or old iPad near my fermentor during fermentation is
not an option for me. Plus while it integrates with [BrewFather] and can log
data to a google log, the interval is 15 minutes, and doesn't expose promtheus
metrics.

Additionally, other exporters, and integrations, are either not in go, or did
not fully cover what I wanted. While the [go-tilt-exporter], which I created,
based on [go-tilt], was closer, it was all in one. The goal here is to go
further, and to create a service to provide the base for multiple exported
formats.

## Design

### Bluetooth and Tilt Handling

The Bluetooth, and tilt handling, is still loosely based off of [go-tilt].
However, it has been reworked to continually scan, until cancelled, and
communicate captured tilt events using channels.

### Exporters

The exporter interface provides the simple base for providing multiple means
of exposing the Tilt data. For me, the goal is to create another service
for the [BrewFather] integration. This will be similar to what is in
[go-tilt-exporter]. This will allow me to have the Bluetooth handling on a
raspberry PI and not worry about dealing with it from a container using either
docker or my home k3s (k8s) cluster.
