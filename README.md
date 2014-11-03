Logjam
------

Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward
these messages on to a log collecton server (such as logstash).

The motivation for logjam was a move to containerising our applications, and a need to get logs from these
applications out of the containers. We configure logjam to listen on the `docker0` (172.16.42.1) interface which is
accessible to applications running within docker.

Logjam supports collecting logs using the following methods:

 - UDP Socket
 - File

Logjam will pipe all entries as a JSON object terminated with "\n" to a remote server.

Usage
-----

### Config File

    {
      "bind": "127.0.0.1",         // interface on host to bind (0.0.0.0 for all)
      "port": 1470,                // port to listen on locally
      "server": "10.1.1.10:1470",  // logstash server to forward to
      "buffer": "/tmp/buffer.log", // file to use for on-disk buffer
      "buffer_size": 1024,         // entries to keep in memory buffer
      "truncate": 3600             // clean out disk buffer every x seconds
    }

### Execution

    $ logjam --config config.json

