## Практическая работа №13 Вуйко Ярослава, ЭФМО-01-25

Профилирование Go-приложения (pprof). Измерение времени работы функций. 16.12.2025

## Цели работы
1.	Научиться подключать и использовать профилировщик pprof для анализа CPU, памяти, блокировок и горутин.
2.	Освоить базовые техники измерения времени выполнения функций (ручные таймеры, бенчмарки).
3.	Научиться читать отчёты go tool pprof, строить графы вызовов и находить “узкие места”.
4.	Сформировать практические навыки оптимизации кода на основании метрик.



## Структура проекта

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── work/
│       ├── note.go
│       ├── show_test.go
│       ├── show.go
│       └── timer.go
├── go.mod
├── go.sum
└── README.md
```



### Аннотация метода CreateNote
<img width="425" height="622" alt="2025-12-16_14-22-57" src="https://github.com/user-attachments/assets/42b89445-62bd-48ed-a660-93f1acf1c922" />
<img width="560" height="575" alt="2025-12-16_14-23-26" src="https://github.com/user-attachments/assets/46d21fc3-d65e-47a3-bf8c-df1aa89240cb" />

### /debug/pprof/
<img width="1736" height="775" alt="2025-12-16_14-23-42" src="https://github.com/user-attachments/assets/ac395242-282b-4752-978d-72bf54a25649" />


### Файлы goroutine, heap, profile
<img width="315" height="80" alt="2025-12-16_14-23-56" src="https://github.com/user-attachments/assets/390b4146-b017-4a15-a5cb-8b2199347502" />



### Top
<img width="742" height="400" alt="top" src="https://github.com/user-attachments/assets/91ed7625-98f6-4900-8f3f-b3b0d307a989" />


### Flame graph
<img width="1831" height="826" alt="flame graph" src="https://github.com/user-attachments/assets/a2b9807e-657a-4a9b-904f-a8dbd00ffdbe" />


### Heap
<img width="536" height="846" alt="heap1" src="https://github.com/user-attachments/assets/73654b37-b336-4af9-a0b0-76d58ce63e87" />
<img width="722" height="841" alt="heap2" src="https://github.com/user-attachments/assets/b5cc13d2-85a5-4fc5-9c76-c1cf3384a2b2" />
<img width="533" height="565" alt="heap3" src="https://github.com/user-attachments/assets/ff536fad-8190-4f0f-81b3-271b8e72646c" />


### alloc_space
<img width="1643" height="834" alt="inuse_space" src="https://github.com/user-attachments/assets/a70b8c3d-78fd-4e80-8c47-ae287bfcb02c" />
<img width="1787" height="838" alt="alloc_space1" src="https://github.com/user-attachments/assets/d94422e6-a094-44b9-b81c-3ba5c58362d3" />
<img width="1701" height="857" alt="alloc_space2" src="https://github.com/user-attachments/assets/24171fca-bc5e-4c3f-acf8-682cccc4d69f" />
<img width="1605" height="784" alt="alloc_space3" src="https://github.com/user-attachments/assets/c1aaa6e4-1b8f-4fc3-b328-6a544c9f6381" />
<img width="355" height="848" alt="alloc_space4" src="https://github.com/user-attachments/assets/cc28f505-a33c-4139-8d4c-933bdc16ec0c" />

### issue_space
<img width="1643" height="834" alt="inuse_space" src="https://github.com/user-attachments/assets/2881dbec-a590-4d41-a555-db1f28e97c91" />

### Анализ до оптимизации
Функция work.Fib потребляет 98.37% процессорного времени. Памяти код почти не тратит. Те небольшие выделения памяти, которые видны в отчётах это не программа, а внутренние процессы Go и операционной системы.


### Длительность запросов в логах
<img width="618" height="332" alt="2025-12-16_16-28-15" src="https://github.com/user-attachments/assets/bd3126d5-1713-4c63-a52e-2061e8fe32be" />

### Бенчмарки
<img width="959" height="232" alt="2025-12-16_16-58-13" src="https://github.com/user-attachments/assets/f9df1d2f-01fb-480b-8d64-9233810c4a0a" />

### Top после оптимизации
<img width="592" height="439" alt="2025-12-16_16-46-45" src="https://github.com/user-attachments/assets/8cacaf9f-7f5d-4ca4-85d2-176e9dc4b182" />

### Flame graph после оптимизации
<img width="203" height="861" alt="2025-12-16_16-47-07" src="https://github.com/user-attachments/assets/12c803aa-d101-4e13-b2d9-819b901435aa" />
<img width="202" height="838" alt="2025-12-16_16-47-23" src="https://github.com/user-attachments/assets/7fafe50e-12ef-4143-bafd-da44d36ae47c" />
<img width="213" height="842" alt="2025-12-16_16-47-33" src="https://github.com/user-attachments/assets/9f43fbfe-9b85-4e76-ac86-a9ed06d4e9d7" />
<img width="480" height="717" alt="2025-12-16_16-47-42" src="https://github.com/user-attachments/assets/603ffeb9-3661-47d4-9fef-0d3b8135f312" />

### Heap после оптимизации
<img width="1847" height="330" alt="2025-12-16_16-47-57" src="https://github.com/user-attachments/assets/ad781095-82c7-4d33-bfa4-20130d2c3ba3" />

### issue_space после оптимизации
<img width="1834" height="870" alt="inuse_space1" src="https://github.com/user-attachments/assets/9bcca7a2-93c2-4a33-b544-b012e1dfcd58" />
<img width="1819" height="785" alt="2025-12-16_16-51-11" src="https://github.com/user-attachments/assets/d8a8ca28-a71d-44c7-b5ed-956396ee3215" />
<img width="1810" height="855" alt="2025-12-16_16-51-22" src="https://github.com/user-attachments/assets/801b9a97-d017-4e3b-bade-ca9a8b4391af" />
<img width="1544" height="857" alt="2025-12-16_16-51-36" src="https://github.com/user-attachments/assets/889e6dff-d1c4-4ddf-a22c-f54659a709c5" />
<img width="1560" height="712" alt="inuse_space5" src="https://github.com/user-attachments/assets/d5c37e00-724a-47c9-949e-c46a1f4d979f" />

### alloc_space после оптимизации
<img width="1827" height="865" alt="alloc_space01" src="https://github.com/user-attachments/assets/7823f602-3fce-4c40-b9c5-275d2a1831af" />
<img width="1831" height="832" alt="2025-12-16_16-53-47" src="https://github.com/user-attachments/assets/29cba886-3df3-4161-97c6-06f3518ff37f" />
<img width="1841" height="847" alt="alloc_space03" src="https://github.com/user-attachments/assets/d9b2b032-dded-4524-8f94-4687e7a05e21" />

### Требования
- Go 1.21 или выше



### Анализ после оптимизации
Относительно памяти после оптимизации ничего не изменилось.
Функция work.FibFast потребляет менее 0.1%. Относительно памяти практически ничего не изменилось.



| Метрика | До оптимизации (рекурсивный Fib) | После оптимизации (итеративный FibFast) | Улучшение |
|---------|-----------------------------------|-----------------------------------------|-----------|
| **Время выполнения** | 9038731 ns/op | ~16.71 ns/op | **В 564920 раз быстрее** |
| **Память (B/op)** | 0 B/op | 0 B/op | Без изменений |
| **Аллокации (allocs/op)** | 0 allocs/op | 0 allocs/op | Без изменений |
| **CPU нагрузка** | 98.37% | < 0.01% | Снижение на ~98% |
| **Потребление памяти** | ~1.5 MB | ~1.5 MB | Без изменений |

###  Выводы
Узким местом была экспоненциальная рекурсивная функция Fib, потреблявшая 98% CPU-времени. Оптимизация через переход на итеративный алгоритм FibFast устранила это узкое место, ускорив вычисления и увеличив пропускную способность сервера без увеличения потребления памяти.

## Ответы на контрольные вопросы:
1. Чем профилирование отличается от логирования и простых таймеров?
Профилирование даёт полную картину работы программы, а логирование и таймеры показывают только отдельные, заранее заданные точки измерений.

2. Что показывает CPU-профиль и как его интерпретировать (top, list, graph)?
CPU-профиль показывает, какие функции потребляют процессорное время: top — список самых "тяжёлых" функций, list — исходный код с затратами по строкам, graph — визуализацию вызовов.

3. В чём разница allocs и inuse в Heap-профиле?
alloc_space показывает общий объём выделенной памяти за время работы, а inuse_space — текущую используемую память в момент снятия профиля.

4. Как включить и для чего анализировать block и mutex профили?
Включить через runtime.SetBlockProfileRate(1) и runtime.SetMutexProfileFraction(1), анализировать для обнаружения мест, где горутины простаивают из-за блокировок.

5. Какие метрики дают бенчмарки testing.B, что означает ns/op, B/op, allocs/op?
Бенчмарки дают стабильные метрики: ns/op — наносекунды на операцию, B/op — байты памяти на операцию, allocs/op — количество аллокаций на операцию.

6. Почему важно сравнивать оптимизации на одинаковой нагрузке и фиксировать «до/после»?
Сравнение на одинаковой нагрузке обеспечивает объективность, а фиксация «до/после» подтверждает, что оптимизация действительно дала эффект.

7. Какой порядок действий при поиске узкого места в производительности?
Сначала измеряем базовые метрики, интерпретируем отчёт (hot path, тяжёлые аллокации), вносим минимальную оптимизацию, снова измеряем и фиксируем результат.




