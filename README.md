# tsanuma

Chess.com client for Golang

## Player

## games.AvailableArchives(player)

`AvailableArchives` will provide you a _url_ based struct to all the games listed by month availables by the given player.

```golang
games, err := games.AvailableArchives(player)
```
