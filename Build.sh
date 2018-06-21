#!/bin/sh

# Build for different OSes [ linux/freebsd/windows/darwin ]
OS="freebsd"   

# Build for different architectures [ arm/arm64/386/amd64 ]
ARCH="amd64"

env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_caps_clients mikrotik_caps_clients.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_caps_ifaces mikrotik_caps_ifaces.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_cpu_load mikrotik_cpu_load.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_dhcp_leases mikrotik_dhcp_leases.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_memory_load mikrotik_memory_load.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_firewall_connections mikrotik_firewall_connections.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_bridge_load mikrotik_bridge_load.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_lan_load mikrotik_lan_load.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_wan1_load mikrotik_wan1_load.go
env GOOS=${OS} GOARCH=${ARCH} go build -o bin/mikrotik_wan2_load mikrotik_wan2_load.go
