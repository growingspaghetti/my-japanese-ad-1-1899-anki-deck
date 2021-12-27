just one of my test scripts to use Wikipedia APIs
なんでもない自分用のankiデッキを作るスクリプト

* [sections of a page](https://ja.wikipedia.org/w/api.php?action=parse&format=json&page=1%E5%B9%B4&prop=sections&disabletoc=1)
* [content of a section](https://ja.wikipedia.org/w/api.php?action=parse&pageid=18479&format=json&prop=text&wrapoutputclass&section=2&disablelimitreport&disableeditsection)

```
(cd cmd && go build -o ../millenium) && ./millenium
```
