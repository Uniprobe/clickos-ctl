# ClickOS Control

ClickOS is a port of the popular modular router Click, to the MiniOS operative system, that enables Click to run under the Xen hypervisor. ClickOS Control implements the necessary functionality to control a ClickOS instance.

This fork also adds in the functionality required to use Unimon monitoring elements, this includes reading exported data from the xenstore or from packets.

## Basic functionality

ClickOS Control supports the following operations:

* Install and remove a router from a running ClickOS instance; (TODO)
* Start, stop and pause installed routers; (TODO)
* Read and write routers' element handlers; (TODO)
* Read from unimon elements; (TODO)
* Control unimon elements. (TODO)

## Controlling mechanisms

ClickOS Control interacts with ClickOS instances via the Xenstore for all control and configuration. Handler data also is passed via the Xenstore. Reading unimon data can be done via the Xenstore or though packets.

# Build

This fork of ClickOS Control is written in GoLang. You can download binaries for x86_64 linux systems from the releases page or use the docker image (this has limitations).

To build, simply run:

```bash
CGO_ENABLED=1 go build
```
