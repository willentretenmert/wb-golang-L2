Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Проблема в том, что горутина, отправляющая числа в канал, никогда не закрывает канал.
В результате, после отправки всех чисел от 0 до 9, горутина завершается, 
но канал остается открытым. Поскольку нет больше данных для чтения из канала и канал не закрыт, 
основная горутина (main) будет бесконечно ожидать данные из канала, что приведет к deadlock.

Go runtime обнаруживает эту ситуацию и завершает программу с сообщением об ошибке, 
указывая на проблему deadlock.
```