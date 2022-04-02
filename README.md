# MIYO-go

**MIYO-go** is a go package for interacting with the MIYO cube via its REST interface.

## Getting started

In order to communicate with the MIYO cube you need two pieces of information:

*   The MIYO cube's IP address (or hostname)
*   An API key

One way to get the MIYO gateway's IP address is to use UPNP discovery.
On Debian/Ubuntu, install the `gupnp-tools` package and run:

```
gssdp-discover --interface=eth0 --timeout=3 --target='upnp:rootdevice'
```

To get an API key, press the physical button at the front of the MIYO cube.
This will enable a one-time retrieval of an API key.
You can then use the sample `apikey` command to retrieve an API key:

```
go run apikey/main.go
```

Alternatively, use cURL to retrieve an API key:

```
curl "http://${MIYO_ADDRESS:?}/api/link"
```

The API key has the form `{6c6cb2ce-b24b-11ec-a61c-482ae37173b5}`,
i.e. the curly braces are part of the API key.

## Features

At the moment, the package supports the following API calls:

*   `APIKey()`

    Requests a new API key. Requires pushing the physical button on the MIYO gateway before calling this method.

*   `DeviceAll()`

    Queries a list of devices (valves and moisture sensors) from the MIYO gateway.

*   `CircuitAll()`

    Queries a list of "circuits" (irrigation areas) from the MIYO gateway.

## Author

Florian Forster &lt;ff at octo.it&gt;