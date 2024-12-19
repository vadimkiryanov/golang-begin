package channels

import (
	"fmt"
	"time"
)

// ChannelsExample демонстрирует работу с каналами в Go
func ChannelsExample() {
    // Создаем два канала для передачи строковых значений
    // make(chan string) создает канал для передачи строк
    channel_1 := make(chan string)
    channel_2 := make(chan string)

    // Запускаем первую горутину (параллельную функцию)
    go func() {
        // Бесконечный цикл
        for {
            // Отправляем сообщение в канал channel_1
            // Оператор <- используется для отправки данных в канал
            channel_1 <- "Канал 1. Прошло 200 мс"
            // Приостанавливаем выполнение на 200 миллисекунд
            time.Sleep(time.Millisecond * 200)
        }
    }()

    // Запускаем вторую горутину
    go func() {
        for {
            // Отправляем сообщение в канал channel_2
            channel_2 <- "Канал 2. Прошло 1 секунда"
            // Приостанавливаем выполнение на 1 секунду
            time.Sleep(time.Second * 1)
        }
    }()

    // Основной цикл программы
    for {
        // select позволяет работать с несколькими каналами одновременно
        // Он ждет, пока один из каналов не будет готов для чтения
        select {
        // Читаем данные из channel_1
        // Оператор <- справа от канала означает получение данных из канала
        case channel := <-channel_1:
            fmt.Println("channel_1", channel)
        // Читаем данные из channel_2
        case channel := <-channel_2:
            fmt.Println("channel_2", channel)
        }
    }
}