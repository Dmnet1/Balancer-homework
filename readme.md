балансировщик нагрузки

есть некоторое количество серверов, нагрузка на которых распределяется методом "по кругу"
есть балансир, который должен распределить запросы равномерно

суть в том, чтобы на каждый сервер попало равномерное количество нагрузки

У балансира есть несколько функций

Init - инициализирует собственно балансер - представьте, что устанавливает соединения с указанным колчиеством серверов.
GiveStat - даёт статистику, сколько запросов пришло на каждый из серверов.
GiveNode - эта функция вызывается, когда пришел запрос. мы получаем номер сервера, на который идти.

Задание ориентировано на понимание того, что такое состояние гонки за данными и как с этим работать.