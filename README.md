# botCore - Core part of GoBot by Stanzinofree

This part is the entry point, it receive the request, understand context and query gateway to obtain answer from the right module.

## Strucutre

It runs as webservice that expone API

## Route exposed

/ -> Entry Point show version, name, release and if there are updates (to implement when gateway and release server is up, to know is mockup)
/healt -> Show status of the service
/api/vX -> root of API versioned
/api/vX/doc -> documentation for API vX
