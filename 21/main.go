package main

import "fmt"

// Сервис— это какой-то полезный класс, обычно сторонний.
//Клиент не может использовать этот класс напрямую, так как сервис имеет непонятный ему интерфейс.

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

//Клиентский код(порт Lightning):
//Клиент — это класс, который содержит существующую бизнес-логику программы.

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

//Клиентский интерфейс описывает протокол, через который клиент может работать с другими классами.

type computer interface {
	insertIntoLightningPort()
}

//: Неизвестный сервис(USB порт)

type windows struct{}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

//Адаптер(ноутбук на Windows)— это класс, который может одновременно работать и с клиентом, и с сервисом.
//Он реализует клиентский интерфейс и содержит ссылку на объект сервиса
//Адаптер получает вызовы от клиента через методы клиентского интерфейса, а затем переводит их в вызовы методов обёрнутого объекта в правильном формате.

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &client{}
	mac := &mac{}
	client.insertLightningConnectorIntoComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
