задача из Яндекс LMS - Яндекс Лицея |SQL|

Необходимо определить фильмы-драйверы внутри каждой категории — те, которые формируют наибольший спрос.

В результате должны быть поля:

category_name — название категории;
title — название фильма;
rentals_count — количество аренд фильма;
rank_in_category — порядковый номер фильма по популярности внутри своей категории (1 — самый популярный).
Выводите в результате только первые три фильма в каждой категории.

Формат вывода
Пример формата вывода требуемого запроса (тестирующая система преобразует вывод запроса в формат TABLE):

+---------------+------------------------+---------------+------------------+
| category_name |         title          | rentals_count | rank_in_category |
+---------------+------------------------+---------------+------------------+
| Action        | RUGRATS SHAKESPEARE    | 30            | 1                |
| Action        | SUSPECTS QUILLS        | 30            | 2                |
| Action        | HANDICAP BOONDOCK      | 28            | 3                |
| Animation     | JUGGLER HARDLY         | 32            | 1                |
| Animation     | DOGMA FAMILY           | 30            | 2                |
| Animation     | STORM HAPPINESS        | 29            | 3                |
| Children      | ROBBERS JOON           | 31            | 1                |
| Children      | IDOLS SNATCHERS        | 30            | 2                |
| Children      | SWEETHEARTS SUSPECTS   | 29            | 3                |
...