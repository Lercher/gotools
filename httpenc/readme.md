# httpenc

This is an explorative program to examine an
http like serialization format
for [NSQ](https://nsq.io) messages.
It would be a standards
conform procedure to store metadata with
a queued message.

*Metadata* in this program's sense is something
like

* correlation ID
* response topic
* sender machine name
* identity information (e.g. login name)
* tenant ID
* checksum or size
* MIME type
* *you name it*

that is normally not part of the payload.

## Design

Using an `http.Header` we can collect
metadata and write the header canonicalized
to an writer.

The other way round `net/textproto` provides
the parser `.ReadMIMEHeader()` for an `http.Header`.

In [main.go](main.go) we explore
how we need to combine the header with the
body and how to properly assemble the pre-existing
parts.