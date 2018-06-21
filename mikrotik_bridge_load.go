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
	Iface    string
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

	fmt.Println("graph_title LocalBridge interface throughput")
	fmt.Println("graph_vlabel bits per second")
	fmt.Println("graph_category network")
	fmt.Println("graph_info This graph shows the incoming and outgoing traffic rate of an interface")
	fmt.Println("in.label inbound")
	fmt.Println("in.type DERIVE")
	fmt.Println("in.draw AREA")
	fmt.Println("in.min 0")
	fmt.Println("in.cdef in,8,*")
	fmt.Println("out.label outbound")
	fmt.Println("out.type DERIVE")
	fmt.Println("out.draw LINE1")
	fmt.Println("out.min 0")
	fmt.Println("out.cdef out,8,*")
}

func PrintStats() {

	Conf := getEnv()

	c, err := dial(Conf)

	reply, err := c.Run("/interface/print", "?name="+Conf.Iface, "", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("in.value", reply.Re[0].Map["rx-byte"])
	fmt.Println("out.value", reply.Re[0].Map["tx-byte"])

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
	iface := os.Getenv("if_bridge")

	var Conf Config

	Conf = Config{
		Host:     host,
		Login:    login,
		Password: pass,
		TLS:      tls,
		Iface:    iface,
	}

	if Conf.Iface == "" {
		log.Fatalln("Interface if_bridge undefined: set env.if_bridge in plugins.conf")
	}

	return Conf
}
