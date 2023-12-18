# 1. Какой самый эффективный способ конкатенации строк?

Старый способ с использованием пакета bytes:

```
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
    }

    fmt.Println(buffer.String())
}
```

Новый способ начина с Go 1.10, используя пакет strings:

```
package main

import (
    "strings"
    "fmt"
)

func main() {
    var sb strings.Builder

    for i := 0; i < 1000; i++ {
        sb.WriteString("a")
    }

    fmt.Println(sb.String())
}
```

Также есть метод с copy, который может быть быстрее, если длина строки известна  
Метод с оператором + тоже может быть эффективным благодаря оптимизации компилятора,  
если известно количество конкатенируемых строк при компиляции.  


# 2. Что такое интерфейсы, как они применяются в Go?

Интерфейс - это абстрактный тип, который описывает поведение других типов
Если конкретно, то он определяет и описывает набор методов, которые нужно реализовать, чтобы удовлетворить контракт (интерфейс). Применяется для создания абстракций.


# 3. Чем отличаются RWMutex от Mutex?

RWMutex нужен, когда у нас есть объект, который нельзя параллельно писать,
но можно параллельно читать. Например, стандартный тип map.
Перед записью в защищаемый мьютексом объект делается .Lock(),
а вызовы .Lock() и .RLock() в других горутинах будут ждать,
пока мы не отпустим мьютекс через .Unlock().
Перед чтением защищаемого объекта делается .RLock()
и только вызовы .Lock() в других горутинах блокируются, вызовы .RLock() спокойно проходят.
Когда отпускаем мьютекс через .RUnlock(), ждущие вызовы .Lock() по-очереди могут забирать мьютекс на себя.
Таких образом обеспечивается параллельное чтение объекта несколькими горутинами,
что улучшает производительность

Можно было бы обойтись только RWMutex, ведь он умеет все то, что и обычный Mutex
Но обычный Mutex проще в реализации и работает быстрее. RW используется только
когда нужно параллельно читать


# 4. Чем отличаются буферизированные и не буферизированные каналы?

Буферизированные каналы имеют определенную емкость (количество элементов, которые они могут содержать).
Они позволяют отправителю отправлять данные в канал, даже если нет никого, кто их читает.
Они не блокируют отправителя до тех пор, пока буфер не заполнен.

Небуферизированные каналы же наоборот, блокируют выполнение горутины,
пока другая горутина не считает из него значение
И только после чтения блокировка снимется и можно будет отправить новое значение

Если покопаться еще глубже, то небуферизированные каналы не используют буфер
при передаче значения другой горутине, а отправляют данные прямо ей на стек


# 5. Какой размер у структуры struct{}{}?

Структура struct{} в Go не содержит полей, поэтому ее размер в памяти равен нулю.


# 6. Есть ли в Go перегрузка методов или операторов?

Нет


# 7. В какой последовательности будут выведены элементы map[int]int?

Пример:  
m[0]=1  
m[1]=124  
m[2]=281  

В map порядок не гарантируется
Поэтому неизвестно



# 8. В чем разница make и new?

make используется для создания слайсов, мап, каналов
и возвращает готовые к использованию экземпляры этих типов.

new используется для выделения памяти (в куче) и инициализации указателей на типы данных
и возвращает указатели на нулевые значения.


# 9. Сколько существует способов задать переменную типа slice или map?

1)
```
s := []int{1, 2, 3, 4}
m := map[string]int{"one": 1, "two": 2}
```

2)
```
s := make([]int, 10)
m := make(map[string]int)
```

# 10. Что выведет данная программа и почему?

```
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```

Выведется 1 два раза  
p - это переменная, хранящая адрес переменной a  
Первый вывод делает дереференс и выводит значение по этому адресу, поэтому выводится первая 1  
Функция update принимает указатель на инт (адрес интовой переменной)  
Передаем в качестве этого адреса наше p  
В функции update выделяется память под локальную переменную p, хранящую переданный адрес (указатель передался по значению)  
Далее мы просто делаем присвоение этой локальной переменной, таким образом стирая
хранящийся там адрес  
С самим адресом мы никаких манипуляций не производили  
update закончился и мы снова делаем дереференс p - выводится вторая 1  



# 11. Что выведет данная программа и почему?

```
func main() {
  wg := sync.WaitGroup{}

  for i := 0; i < 5; i++ {
     wg.Add(1)

     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }

  wg.Wait()
  fmt.Println("exit")
}
```

Выведет числа 0, 1, 2, 3, 4 в произвольном порядке  
Программа "зависнет" и произойдет deadlock  
Потому что мы передаем wg sync.WaitGroup по значению, а не по ссылке  
И поэтому wg скопируется, создастся локальная копия wg в каждой горутине  
И уже у этих копий будет вызываться Done, что никак не повлияет на wg, ожидающий в main горутине  
Таким образом, дойдя до wg.Wait() программа никогда не дождется, когда у настоящего wg счетчик обнулится  



# 12. Что выведет данная программа и почему?

```
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```

Выведет 0, потому что внутри блока if была объявлена и проинициализирована
новая переменная n. Блок if создает свою область видимости, в которой могут определяться переменные
и они существовуют только в пределах if


# 13. Что выведет данная программа и почему?

```
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```

Выведет
```
[100 2 3 4 5]
```

Строчка v[0] = 100  
изменит первый элемент оригинального слайса, т.к. слайсы передаются по ссылке

Строчка v = append(v, b)  
присвоит локальной переменной v новое значение - слайс [100, 2, 3, 4, 5]




# 14. Что выведет данная программа и почему?

```
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```

Выведет
```
[b b a]
[a a]
```

Передаем слайс в функцию  
Сразу же переопределяем значение локальной переменной новым слайсом [a a a]  
Затем в этом новом слайсе меняем первые 2 элемента на b и получаем [b b a]  
Тут же происходит первый вывод [b b a]  
Затем выходим из функции и происходит второй вывод старого слайса - [a a]  