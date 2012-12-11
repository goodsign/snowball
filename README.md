Description
====

Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality. For more detailed info see http://snowball.tartarus.org/

Installing
====

```
go get github.com/goodsign/snowball
go test github.com/goodsign/snowball (Must PASS)
```

Done! Use it in your go files. (import 'github.com/goodsign/snowball')

Usage
====

```go
  stemmer, err := NewWordStemmer(algorithm, encoding)
  
  if nil != err {
    /*...handle error...*/
  }
  defer stemmer.Close() 

  wordStem, err := stemmer.Stem(word)
  if nil != err {
    /*...handle error...*/
  }

  /* Use wordStem */

```
Usage notes
-----------

According to Snowball documentation:

```
Creating a stemmer is a relatively expensive operation - the expected
usage pattern is that a new stemmer is created when needed, used
to stem many words, and deleted after some time.
```

Algorithms & encodings
----

File **modules.txt** contains all the main algorithms for each language, in UTF-8, and also with
the most commonly used encoding.

```
Language        Encodings               Algorithms

danish          UTF_8,ISO_8859_1        danish,da,dan
dutch           UTF_8,ISO_8859_1        dutch,nl,dut,nld
english         UTF_8,ISO_8859_1        english,en,eng
finnish         UTF_8,ISO_8859_1        finnish,fi,fin
french          UTF_8,ISO_8859_1        french,fr,fre,fra
german          UTF_8,ISO_8859_1        german,de,ger,deu
hungarian       UTF_8,ISO_8859_1        hungarian,hu,hun
italian         UTF_8,ISO_8859_1        italian,it,ita
norwegian       UTF_8,ISO_8859_1        norwegian,no,nor
portuguese      UTF_8,ISO_8859_1        portuguese,pt,por
romanian        UTF_8,ISO_8859_2        romanian,ro,rum,ron
russian         UTF_8,KOI8_R            russian,ru,rus
spanish         UTF_8,ISO_8859_1        spanish,es,esl,spa
swedish         UTF_8,ISO_8859_1        swedish,sv,swe
turkish         UTF_8                   turkish,tr,tur
```

Thread-safety
====

The original Snowball documentation says:

```
Stemmers are re-entrant, but not threadsafe.  In other words, if
you wish to access the same stemmer object from multiple threads,
you must ensure that all access is protected by a mutex or similar
device.
```

Thus this Go wrapper uses **sync.Mutex** for each stem operation, so it is thread safe.

Snowball Licence
==========

The Snowball library is released under the [BSD Licence](http://opensource.org/licenses/bsd-license.php)

Licence
==========

The goodsign/snowball binding is released under the [BSD Licence](http://opensource.org/licenses/bsd-license.php)

[LICENCE file](https://github.com/goodsign/libtextcat/blob/master/LICENCE)