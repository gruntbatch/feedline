feedline
========

What is the simplest hack that will result in a fully functional RSS reader, complete with a web interface and sync capabilities? It's _feedline_!

Setup
-----

_feedline_ looks for a `.feedline` folder in your home directory. It expects to find a `subscriptions.opml` file, and a `read` folder for marking entries as read.

The `.feedline` folder is designed to be synced using services such as Dropbox or iCloud, or even a self-hosted solution, such as Syncthing.

After building the `feedline` executable, launch it and run it in the background. Navigating to `localhost:8080` will get you your RSS feed.
