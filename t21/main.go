package main

import (
	"encoding/json"

	"github.com/kodeyeen/test-wildberries-l1/t21/adapters"
	// это может быть сторонний пакет, который мы не можем редактировать
	// либо пакет, который будет больно менять, т.к. от него много что зависит
	"github.com/kodeyeen/test-wildberries-l1/t21/deliveryclub"
)

// у нас есть общий интерфейс сервисов по доставке еды
type DeliveryServiceAPI interface {
	// мы ожидаем в реализациях увидеть метод FetchProducts для получения списка продуктов
	FetchProducts() json.RawMessage
}

func main() {
	// есть некий клиентский код, который работает с API сервисов по доставке
	// необязательно Delivery Club или Yandex Eda.
	// В будущем могут добавиться и другие сервисы
	// поэтому с помощью интерфейса код унифицирован для любых сервисов
	// но проблема в том, что эти сервисы сторонние, их разрабатывали разные люди
	// они не договаривались между собой ни о чем, поэтому их пакеты хоть и
	// предлают похожий функционал, однако могут быть отличия, например,
	// в названиях и сигнатурах методов
	// Паттерн Адаптер позволяет буквально адаптировать сторонний код под наш интерфейс

	dcAPI := deliveryclub.NewAPI("api.delivery-club.ru")
	someLogicWithDeliveryService(dcAPI) // у dcAPI есть метод FetchProducts с нужной сигнатурой

	// yeAPI := yandexeda.NewAPI("api.yandexeda.ru/foods")
	// someLogicWithDeliveryService(yeAPI) // ошибка: yeAPI не реализует интерфейс

	// написали адаптер, который оборачивает имеющийся функционал
	// и подгоняет его под наш код
	// иногда этот паттерн называют Wrapper
	yeAPI := adapters.NewYandexEdaAPIAdapter("api.yandexeda.ru")
	someLogicWithDeliveryService(yeAPI)
}

// какая-то функция в клиентском коде
func someLogicWithDeliveryService(service DeliveryServiceAPI) {
	service.FetchProducts()
}
