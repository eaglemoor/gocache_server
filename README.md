# gocache

[![Build Status](https://travis-ci.org/eaglemoor/gocache_server.svg?branch=master)](https://travis-ci.org/eaglemoor/gocache_server)
[![codecov](https://codecov.io/gh/eaglemoor/gocache_server/branch/master/graph/badge.svg)](https://codecov.io/gh/eaglemoor/gocache_server)

Go cache server (test task)

## TASK

Простая имплементация похожего на Redis кеша в памяти

Необходимый функционал:

• Key-value хранилище строк, списков, словарей
• TTL на каждый ключ
• Операторы
• GET
• SET
• REMOVE
• KEYS

• Дополнительные операции (получить значение по индексу из списка, получить значение по ключу из словаря)
• Golang API Client к кешу
• API tcp(telnet)/REST API (реализовать одно на выбор)

Предоставить несколько тестов, документацию по API, документацию по развертыванию, несколько кейсов и примеров использования кеша вызовов http/tcp api.

Не обязательные требования:

• Сохранение на диск
• Масштабирование (на серверной стороне, на клиентское выбирайте сами)
• Авторизация
• Нагрузочные тесты

