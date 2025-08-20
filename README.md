# tsacl
Access Control CLI for Tailscale.

Handles Tailscale ACL via the command line.
This is done via API calls to Tailscales API via a locally stored API key.
(Generation of API keys is done via the admin panel and will not be supported here)

This repo is under construction and in version 0.0.1.

For installation, see below.

## Techniques
Building this app as a GO app. It relies on Cobra CLI for generating commands,
possibly Oauth, JSON manipulation, 

# Background
Using Tailscale a lot, I've found that it is not very wonderful to interact with
the raw ACL file. While Tailscale has now implemented a good GUI for this
there lacks an CLI for it. Hence, I am doing this for myself and others who
might be interested in this solution.

This is not a commercial product and mostly for my own use.
The source code is open source and you are free to use, spread and build on it.
Let me know if you want to contribute!

I am not affiliated with Tailscale, but love their product.
I will not sell this functionality to others and Tailscale owns all rights
to their own property (which I will not capitalize off). This is just a
personal project.

## Near future
I am working with this as a side project. Updates to this will be sparse
and bug fixes might be slow. If you find anything to fix, you can report it
in "Issues". I have no idea when this will be done (if ever), but the near future
will see tests with getting the ACL from Tailscale API (probably via Oauth).

## Structure
Commands will be structured as follows:

tsacl <group> <command> <subcommand>

Flags and options will follow naturally.
Help sections will be added for each command with natural explanations of what is happening.

# Installation
As for now this application is best installed by downloading this repo
and installing it as a Go build directly. You can then access the CLI
from ./tsacl

If you want to invoke it as a "real" program, you need to add it to your paths,
to /bin, to /opt or as an alias.

## How to Install:
Go and git needs to be installed.

````
git clone https://github.com/cfossto/tsacl.git
cd tsacl
go build
````

The app is now installed and ready to be configured when you start the program.