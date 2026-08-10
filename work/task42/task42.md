задача из Яндекс LMS - Яндекс Лицея |SQL|

Во многих системах логин пользователя берут из email, поэтому полезно уметь извлекать его прямо в запросе.

Для таблицы customers вывести:

email — исходный email;
email_lower — email в нижнем регистре;
login — часть email до символа @ в нижнем регистре.
Отсортировать результат по login по возрастанию.

Формат вывода
Корректный вывод для тестовой сборки (тестирующая система преобразует вывод запроса в формат CSV):

email,email_lower,login
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
,,
adrian.medvedev27@bk.ru,adrian.medvedev27@bk.ru,adrian.medvedev27
agafya.antonova@bk.ru,agafya.antonova@bk.ru,agafya.antonova
agafya.naumova@mail.ru,agafya.naumova@mail.ru,agafya.naumova
aGafYA.oRLOVa@mail.ru,agafya.orlova@mail.ru,agafya.orlova
agafya.samoylova@yandex.ru,agafya.samoylova@yandex.ru,agafya.samoylova
agata.lapina@mail.ru,agata.lapina@mail.ru,agata.lapina
Agata.Matveeva@gmail.com,agata.matveeva@gmail.com,agata.matveeva
agata.scherbakova@gmail.com,agata.scherbakova@gmail.com,agata.scherbakova
agataSoloveva@yandex.ru,agatasoloveva@yandex.ru,agatasoloveva
aggey.vishnyakov@gmail.com,aggey.vishnyakov@gmail.com,aggey.vishnyakov
aleksandra.danilova81@bk.ru,aleksandra.danilova81@bk.ru,aleksandra.danilova81
alevtina.savina53@yandex.ru,alevtina.savina53@yandex.ru,alevtina.savina53
alevtina.subbotina@list.ru,alevtina.subbotina@list.ru,alevtina.s
...
Обратите внимание, что из-за большого числа записей в тестовой сборке приведена только часть ответа запроса. Строки вида ,, означают, что для конкретного покупателя не записан email в базу данных - несмотря на это запрос должен работать корректно и выводить пустые строки (да, это немного некрасиво, но об этом поговорим позже).