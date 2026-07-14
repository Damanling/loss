Проект loss:
трейкер личных финансов, после записи трат или пополнений он анализирует и помогает распределить их по категориям для того чтобы лучше отслеживать свои финансы.


Как запустить:
go run cmd/main.go

Команды:
hello
echo

Ошибки: 
usage: 
loss hello
loss echo args
exit status 1
 Если вылазиют ошибки, то ваш запрос не соответсует данным командам.
 
 Структура папок:
── Loss
    ├──.github
        └──ci.yml
    ├──cmd/loss
        └──main.go
    ├──internal
        └──.getkeep
    ├──.editorconfig
    ├──.gitignore
    ├──.golangli.yml
    ├──go.mod
    ├──README.md
 