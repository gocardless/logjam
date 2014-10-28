Logjam
------

Logjam is the log forwarder for the mesos age.

Logjam supports collecting logs using the following methods:

 - UDP Socket
 - File

Logjam will pipe all entries as JSON to a remote server (logstash).

Usage
-----

### Config File

    {
      "bind": "127.0.0.1",         // interface on host to bind (0.0.0.0 for all)
      "port": 1470,               // port to listen on locally
      "server": "10.1.1.10:1470",  // logstash server to forward to
      "buffer": "/tmp/buffer.log", // file to use for on-disk buffer
      "buffer_size": 1-24,         // entries to keep in memory buffer
      "truncate": 3600             // clean out disk buffer every x seconds
    }

### Execution

    $ logjam --config

