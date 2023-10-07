package main

import "fmt"

/*
Adapter pattern is a structural design pattern which allows
incompatible objects to collaborate
*/

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(cm Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	cm.insertIntoLightningPort()
}

type Computer interface {
	insertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct {
}

func (w *Windows) insertIntoUSB() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	w *Windows
}

func (wa *WindowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	wa.w.insertIntoUSB()
}

func main() {
	c := &Client{}
	m := &Mac{}
	c.InsertLightningConnectorIntoComputer(m)

	w := &Windows{}
	wa := &WindowsAdapter{w: w}
	c.InsertLightningConnectorIntoComputer(wa)
}
