Registry ACL
=
When this plugin is installed, SOMETHING WILL HAPPEN TODO.
In order to use this plugin you need to be running at least Docker 1.10 which
has support for authorization plugins.

Building
-
```sh
$ export GOPATH=~ # optional if you already have this
$ mkdir -p ~/src/github.com/projectatomic # optional, from now on I'm assuming GOPATH=~
$ cd ~/src/github.com/projectatomic && git clone https://github.com/projectatomic/registryacl
$ cd registryacl
$ make
```
Installing
-
Either:
```sh
$ sudo make install
$ systemctl enable registryacl-plugin
```
Running
-
Specify `--authz-plugin=registryacl` in the `docker daemon` command line
flags (either in the systemd unit file or `/etc/sysconfig/docker` under `$OPTIONS`
or when manually starting the daemon)
The plugin must be started before `docker` (done automatically via systemd unit file).
If you're not using the systemd unit file:
```sh
$ registryacl &
```
Just restart `docker` and you're good to go!
License
-
MIT
