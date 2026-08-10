задача из Яндекс LMS - Яндекс Лицея |SQL|

Даже если данные хранятся аккуратно и без артефактов, для анализа email пользователей всё равно приходится извлекать логины и домены.

Для таблицы customers выведите:

email;
email_lower — email в нижнем регистре;
domain — часть строки после @;
login — часть строки до @;
masked — email, в котором в логине все точки . заменены на _;
login_len — длину логина.
Отсортируйте результат по domain по возрастанию, затем по login_len по убыванию.

В ответе учитывайте только непустые строки с электронными почтами.

Формат вывода
Корректный вывод для тестовой сборки (тестирующая система преобразует вывод запроса в формат CSV):

email,email_lower,domain,login,masked,login_len
sinklitikiya.gorbacheva@bk.ru,sinklitikiya.gorbacheva@bk.ru,bk.ru,sinklitikiya.gorbacheva,sinklitikiya_gorbacheva@bk.ru,23
aleksandra.danilova81@bk.ru,aleksandra.danilova81@bk.ru,bk.ru,aleksandra.danilova81,aleksandra_danilova81@bk.ru,21
mayya-kalashnikova46@bk.ru,mayya-kalashnikova46@bk.ru,bk.ru,mayya-kalashnikova46,mayya-kalashnikova46@bk.ru,20
evpraksiya.polyakova@bk.ru,evpraksiya.polyakova@bk.ru,bk.ru,evpraksiya.polyakova,evpraksiya_polyakova@bk.ru,20
irina.scherbakova21@bk.ru,irina.scherbakova21@bk.ru,bk.ru,irina.scherbakova21,irina_scherbakova21@bk.ru,19
pelageya_matveeva41@bk.ru,pelageya_matveeva41@bk.ru,bk.ru,pelageya_matveeva41,pelageya_matveeva41@bk.ru,19
aristarkh.likhachev@bk.ru,aristarkh.likhachev@bk.ru,bk.ru,aristarkh.likhachev,aristarkh_likhachev@bk.ru,19
klavdiya.molchanova@bk.ru,klavdiya.molchanova@bk.ru,bk.ru,klavdiya.molchanova,klavdiya_molchanova@bk.ru,19
...
Поскольку весь вывод запроса очень большой, приведена только его часть.