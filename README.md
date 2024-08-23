# HOME-MEDIA

## Описание

Комплект собран для удобного скачивания видео контента через торренты и просмотра на смарт ТВ в ForkPlayer и подобных
плеерах сидя в удобном кресле.

Онлайн кинотеатры часто тормозят, или вообще не доступны с хорошим качеством видео контента, а данный комплект позволяет
удобно скачать и просмотреть.

Если Вы любите под вечер посмотреть что-то интересное в отличном качестве и у Вас есть ПК, который, к примеру, постоянно
включен или может быть включен, то данный вариант Вам просто необходим.

## Установка

- Убедитесь, что у Вас уже установлен [Docker](#install-docker-ubuntu)
- Если у Вас [Windows](#install-windows), то необходимо
  установить [WSL2](https://learn.microsoft.com/ru-ru/windows/wsl/install)
  и [Docker для windows](https://docs.docker.com/desktop/install/windows-install/)

- Создание директории
  `mkdir home-media && cd home-media`

- Получение с репозитория
  `git clone git@github.com:oleg-mordvintsev/home-media.git .`

- Директория видео и других файлов
  `mkdir data`

- Устанавливаем права на запуск
  `chmod +x start && chmod +x stop`

- Запуск
  `./start`

- Остановка
  `./stop`

## Install Docker Ubuntu

- Инструкция взята с официального сайта https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository
- `sudo apt-get update`
- `sudo apt-get install ca-certificates curl`
- `sudo install -m 0755 -d /etc/apt/keyrings`
- `sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc`
- `sudo chmod a+r /etc/apt/keyrings/docker.asc`
- `echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`
- `sudo apt-get update`
- `sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin`

## Install Windows

- В PowerShell (Правой кнопкой по "Пуск", Терминал)
- В терминале
  `wsl --install`
  `wsl --set-version Ubuntu-24.04 2`
- Далее действия уже в терминале установленной Ubuntu в разделе [Установка](#установка)

## Определение ip

- Для Ubuntu `ifconfig`, для Windows `ipconfig`. Если Вы во внутренней сети, то Ваш ip будет начинаться с `192.168.0` 