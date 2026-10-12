
# Telegram Bridge

Это плагин GitHub Actions, который отправляет уведомления об изменениях в GitHub в чат Telegram через бота. Он помогает ускорить процесс разработки.

## Что умеет этот плагин?

Если использовать предоставленный файл [main.yaml](https://github.com/GokulDas027/TelegramBridge/blob/master/main.yaml) (рекомендуется), он будет отправлять уведомления о следующих событиях:

- Отправка изменений в репозиторий (Push).
- Pull Request — запросы на слияние (открытие, закрытие, редактирование, готовность к проверке, запрос проверки, повторное открытие).
- Комментарии к Pull Request (создание).
- Issues — задачи (открытие, редактирование, закрепление, закрытие, повторное открытие, назначение исполнителя, добавление метки).
- Комментарии к Issue (создание).

> Ограничения установлены не потому, что возможности плагина ограничены, а для предотвращения чрезмерного количества сообщений. Их легко изменить в файле `main.yaml` своего репозитория. Подробнее см. в [документации по триггерам GitHub Actions](https://developer.github.com/webhooks/).

## Как использовать плагин?

Плагин доступен в [GitHub Marketplace](https://github.com/marketplace/actions/telegrambridge).

Также его можно использовать с помощью файла [main.yaml](https://github.com/GokulDas027/TelegramBridge/blob/master/main.yaml):

- Поместите его в каталог `<repository>/.github/workflows/main.yml` своего репозитория.

Подключите Telegram-бота:

- **Token** — токен бота, полученный от [BotFather](https://core.telegram.org/bots).
- **Id** — идентификатор канала или вашего пользователя, в зависимости от того, куда нужно отправлять сообщения.

### [Пошаговая инструкция](https://github.com/GokulDas027/TelegramBridge/blob/master/instruction.md)

## Как выглядят уведомления?

![Пример сообщения](Screenshot.png)

## Что дальше?

Это решение можно настроить под свои потребности. Изменить его можно в файле [main.go](https://github.com/GokulDas027/TelegramBridge/blob/master/main.go). Соответствующие изменения также необходимо внести в файл `main.yaml` GitHub Action.

Благодарность [Yanzay](https://github.com/yanzay) и [Athul](https://github.com/athul).
