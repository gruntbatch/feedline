feedline
========

What is the simplest hack that will result in a fully functional, highly opinionated RSS reader, complete with a web interface? It's _feedline_!

Features
--------

  - Use any existing `.opml` file
  - System-controlled dark mode (macOS only)

Setup
-----

_feedline_ looks for a `.feedline` folder in your home directory. It expects to find a `subscriptions.opml` file, and a `read` folder for marking entries as read.

The `.feedline` folder is designed to be synced using services such as Dropbox or iCloud, or even a self-hosted solution, such as Syncthing.

After building the `feedline` executable, launch it and run it in the background. If the templates and the executable are seperated, you'll need to supply the `-webdir` flag. You can tell the program what address to use with the `-addr` flag.
