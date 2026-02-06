---
page_title: "Troubleshooting"
subcategory: ""
description: |-
  troubleshooting tips
---

# Enable debug mode

Export these environment variables before running terraform cli. The log file will be available in /tmp/tf.log

```
export TF_LOG_PATH=/tmp/tf.log
export TF_LOG_PROVIDER=DEBUG
export TF_LOG_SDK=OFF
```

# capture http traffic

capture the smc api traffic to the smc can help troubleshooting

## capture at tcp level

you can use `tcpflow` to capture tcp traffic if you're using plain http.

on ubuntu, install it via:

```sh
❯ sudo apt install tcpflow
```

then

```sh
❯ sudo tcpflow -i any port 8082  -C |tee /tmp/tcp_flow.txt
```

## capture via an http proxy

e.g. charles, burp, mitmproxy, http-breakout-proxy, ...

under linux, define the following env. var 

```
export HTTP_PROXY="http://smcproxy:8080"
```

but there is a catch: it won't work if trying to contact the smc on
the localhost. In this case, the workaround is to define an alias name
in /etc/hosts and use that name in the provider initialization.
