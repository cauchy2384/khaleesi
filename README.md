# Khaleesi
[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)

Библиотека, которая позволяет передразнивать сообщения в стиле мема "Кхалиси передразнивает". 

![Кхалиси передразнивает](https://memepedia.ru/wp-content/uploads/2017/08/%D0%BA%D1%85%D0%B0%D0%BB%D0%B8%D1%81%D0%B8-%D0%BC%D0%B5%D0%BC.jpg)

## Как использовать

Если нужно передразнить одно сообщение
```
khaleesi.Mock("Позвольте мне сражаться за вас, Кхалиси!")
// Позвёйте мьне сйажятьсйа зя вас, Кхаиси!
```

Если нужно передразнить несколько сообщений
```
kh, _ := khaleesi.New()

kh.Mock("Позвольте мне сражаться за вас, Кхалиси!")
// Позвёйте мьне сйажятьсйа зя вас, Кхаиси!

kh.Mock("Дерись за меня дракон!")
// Дейисьь зя миня дьяконь!
```

## Лицензия

Этот проект распространяется на условиях лицензии MIT. Пожалуйста, ознакомьтесь с файлом LICENSE в корне проекта для получения дополнительной информации.
