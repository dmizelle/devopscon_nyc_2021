# DevOpsCon 2021 (Unix Signals)

This repo serves as an example of what happens
when processes don't define signal handlers.

# Getting Started

_There is a shell.nix file here, if you are running Nix / NixOS (not as in \*NIX!)_

There is a Makefile that will help you out:

`make server-no-handler` will start the process without a signal handler for SIGINT

`make server-with-handler` will start the process _with_ a signal handler for SIGINT

`make request` will run `curl localhost:8888`

The server will always wait 10 seconds before returning content.

To have the example "work", in the terminal / shell running the server, hit Control-C after
you've ran `make server-no-handler` or `make server-with-handler`.

In the `server-no-handler` example, you'll see cURL error out.
In the `server-with-handler` example, you'll see the server _wait and finish the request_ before exiting.
