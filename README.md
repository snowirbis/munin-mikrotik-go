# munin-mikrotik-go
Munin modules to monitor MikroTik router activity / writen in Go

Add following to your munin plugins.conf

[mikrotik_*]
env.connect_host ***.***.***.***:8729 (or 8728 for non/TLS)
env.connect_login your_login
env.connect_password your_password
env.connect_tls true (or false for non/TLS)
