nltk-go
=======

Natural Language Toolkit (NLTK) for Golang inspired by [nltk](https://github.com/nltk/nltk)

This is a working in progress (**the very beginning of this project to be more precise**), so everything might change very fast and backwards compatibility is not guaranteed.

Our focus now is making ["Floresta Sintática"](http://www.linguateca.pt/Floresta/info_floresta_English.html) available and easy to use using Golang, after that we will work on some other languages.

We started using a SQL dump of "Floresta Sintática" that we got working on a Docker postgres (see [Bosque_CP_8.0_Postgres](Bosque_CP_8.0_Postgres) for more details), however there was some issues with the SQL insert statements, so we are working on parsing an XML (Tiger-XML) version and inserting on the database.

Next Steps
----------
1. Finish the database integration/interface
2. Read the Tiger-XML
3. Convert it to UTF-8
4. Insert it in a Postgres Database
5. Convert it to a Sqlite Database (it's easier to have the whole database in one single file)
6. Support other languages

Contributing
------------

1. Fork it
2. Write some good code
3. Make a Pull-Request
4. That's it :D

Licence
-------

[GPL V3](LICENSE)