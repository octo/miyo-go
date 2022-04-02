# MIYO-go

**MIYO-go** is a go package for interacting with the MIYO cube via its REST interface.

## Getting started

In order to communicate with the MIYO cube you need two pieces of information:

*   The MIYO cube's IP address (or hostname)
*   An API key

Use can use the sample code in the `setup/` directory to automatically discover this information. Press the physical button on the MIYO Cube, then run:

```
go run setup/main.go
```

The API key has the form `{6c6cb2ce-b24b-11ec-a61c-482ae37173b5}`,
i.e. the curly braces are part of the API key.

## Features

At the moment, the package supports the following API calls:

*   `FindCube()`

    Discovers the MIYO Cube on the local network using UPnP.
*   `APIKey()`

    Requests a new API key. Requires pushing the physical button on the MIYO gateway before calling this method.
*   `Devices()`

    Queries a list of devices (valves and moisture sensors) from the MIYO gateway.
*   `Areas()`

    Queries detailed information of all irrigation areas (aka. "circuits") from the MIYO gateway.

## Author

Florian Forster &lt;ff at octo.it&gt;