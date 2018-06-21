# munin-mikrotik-go
Munin modules to monitor MikroTik router activity / writen in Go, using gopkg.in/routeros.v2

Add following to your munin plugins.conf

```sh
[mikrotik_*]
env.connect_host ***.***.***.***:8729 #(or 8728 for non/TLS)
env.connect_login your_login
env.connect_password your_password
env.connect_tls true #(or false for non/TLS)
env.if_bridge br0-LAN
env.if_lan eth2-master
env.if_wan1 sfp1-uplink-1G
env.if_wan2 eth1-back-10M
```

# TODO
Monitor Wireless/client connection & latency
Full munin-node emulation to run on any box in 1 minute
