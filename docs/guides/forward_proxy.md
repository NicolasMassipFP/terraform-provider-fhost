---
page_title: "use a forward proxy"
subcategory: ""
description: |-
  use a forward proxy
---

# use a forward proxy

In case the host running terraform cli needs to go through a forward proxy to contact the SMC host:
under linux, define the following env. var 

```
export HTTP_PROXY="http://smcproxy:8080"
```
