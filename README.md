# Discgo

A Discord API client library written in Go

This is not a full client, and is intended to provide support for certain Discord API interactions as I have the need. Perhaps in time full support will be added. Contributions are welcome.

Yes, there are other libraries that have been out for a while which are more actively developed, but I have my own reasons and goals for creating this project:
* Better support for metrics and tracing
* Better ruser of context
* Better control of logging and error handling
* A package structure I like
* Respect SemVer, even if semantic import versioning sucks
* Because I want to 

## Features
* Separate REST API client and Gateway Clients. Use one or both as you like.
  * although, you'll probably need to use the REST client to initialize the Gateway client
* Functional options for configuring your client

## Using the Gateway Client

> TBD
