package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"gopkg.in/routeros.v2"
)

var tlsconfig = &tls.Config{
	InsecureSkipVerify: true,
}

type Config struct {
	Host     string
	Login    string
	Password string
	TLS      string
}

func main() {

	var Option string

	if len(os.Args) > 1 {
		Option = os.Args[1]
	} else {
		PrintStats()
		os.Exit(0)
	}

	RunWithOption(Option)

}

func RunWithOption(option string) {

	switch option {
	case "autoconf":
		fmt.Println("yes")
		return
	case "config":
		PrintConf()
		return
	default:
		PrintStats()
	}
	return

}

func PrintConf() {
	fmt.Println("graph_title Active TCP connections")
	fmt.Println("graph_vlabel conn")
	fmt.Println("graph_category network")
	fmt.Println("graph_info This graph shows active TCP connections on MikroTik device")
	fmt.Println("graph_args -l 0")
	fmt.Println("graph_scale no")
	fmt.Println("conn.label Connections")
	fmt.Println("conn.type GAUGE")
	fmt.Println("conn.draw AREA")
	fmt.Println("conn.graph yes")

}

func PrintStats() {

	Conf := getEnv()

	c, err := dial(Conf)

	reply, err := c.Run("/ip/firewall/connection/print", "", "", "")
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	for _, _ = range reply.Re {
		i++
	}

	fmt.Printf("conn.value %d\n", i)

}

func dial(Conf Config) (*routeros.Client, error) {
	if Conf.TLS == "true" {
		return routeros.DialTLS(Conf.Host, Conf.Login, Conf.Password, tlsconfig)
	}
	return routeros.Dial(Conf.Host, Conf.Login, Conf.Password)
}

func getEnv() Config {

	host := os.Getenv("connect_host")
	login := os.Getenv("connect_login")
	pass := os.Getenv("connect_password")
	tls := os.Getenv("connect_tls")

	var Conf Config

	Conf = Config{
		Host:     host,
		Login:    login,
		Password: pass,
		TLS:      tls,
	}

	return Conf
}
