# MO-HOME-MEDIA

## Описание

MO-HOME-MEDIA - это набор инструментов, предназначенных для удобного скачивания видео контента через торренты и
просмотра
его на смарт ТВ с помощью ForkPlayer или подобных плееров.

Данный комплект решает проблему с онлайн-кинотеатрами, которые часто тормозят или недоступны с хорошим качеством видео.
Он позволяет вам скачать и просмотреть фильмы и сериалы в отличном качестве в любое время.

Если у вас есть компьютер, который работает постоянно или может быть включен, и вы любите смотреть фильмы по вечерам, то
MO-HOME-MEDIA - идеальное решение для вас.

## Установка

### Требования

* **Docker:** Необходимо установить Docker. Инструкции по установке представлены ниже.
* **WSL2 (для Windows):** Если вы используете Windows, необходимо установить Windows Subsystem for Linux 2 (WSL2).
* **ForkPlayer (необязательно):** Для просмотра скачанных файлов на смарт ТВ рекомендуется
  установить [ForkPlayer](http://forkplayer.tv/) или подобный плеер? позволяющий просматривать списки и воспроизводить
  видео файлы по прямым ссылкам в сети.

### Windows (Пропускаем если у Вас Ubuntu)

1. Откройте PowerShell с правами администратора (щелкните правой кнопкой мыши по "Пуск" и выберите "Терминал (
   Администратор)").
2. Запустите команду: `wsl --install`
3. Установите версию Ubuntu: `wsl --set-version Ubuntu-22.04 2`
4. После завершения установки, откройте терминал установленной Ubuntu и выполните инструкции из раздела  "Установка
   Docker на Ubuntu".

### Установка Docker (Пропускаем если docker уже установлен)

- Установка curl для скачивания скриптов установки
  `apt update && apt install -y curl`

- Установка docker в docker
  `curl -fsSL https://github.com/oleg-mordvintsev/mo-home-media/raw/refs/heads/main/docker-install.sh | sed 's/sudo //g' | bash`

### Установка и запуск приложения

- Установка curl для скачивания скриптов установки
  `apt update && apt install -y curl`

- Установка и запуск приложения
  `curl -fsSL https://github.com/oleg-mordvintsev/mo-home-media/raw/refs/heads/main/init.sh | sed 's/sudo //g' | bash`

- Установка происходит по умолчанию в домашнюю директорию `~/mo-home-media`
  `cd ~/mo-home-media && ./start.sh`
  `cd ~/mo-home-media && ./stop.sh`

- Приложение будет доступна с разных портов:
    - Вам необходимо определить IP-адрес вашего компьютера
        - Для Windows оптимально
            - "Пуск" -> "Выполнить" -> `cmd` -> `ipconfig`
            - В списке адаптеров ищем тот, что с указанным шлюзом
            - В блоке с указанным шлюзом находим `IPv4-адрес` это и будет ip (обычно начинается со `192.168.`)
        - Для Ubuntu
            - Открываем терминал
            - Сразу результат `ip -4 -br addr show | grep -vE '^(docker|br-|veth|virbr|lo)' | awk '{print $3}' | cut -d'/' -f1`
            - `ifconfig` - ищем то, что RUNNING и имеет заполненный broadcast (шлюз)
        - Если ip начинается с `172`, то это не тот, что нужен
    - http://ip:777 или http://localhost:777 или http://192.168.0.100:777 - пример адреса, который необходимо ввести в
      поисковой строке на главной странице Fork Player
    - http://ip:888 или http://localhost:888 или http://192.168.0.100:888 - адрес, который можете открыть, к примеру, со
      смартфона, и торрент клиент сразу будет

### TEST Docker in Docker

- Запуск контейнера с Ubuntu 22.04 с пробросом сокета, чтобы docker в docker смог запустить приложение
  `docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock ubuntu:22.04 bash`

- Установка curl для скачивания скриптов установки
  `apt update && apt install -y curl`

- Установка docker в docker
  `curl -fsSL https://github.com/oleg-mordvintsev/mo-home-media/raw/refs/heads/main/docker-install.sh | sed 's/sudo //g' | bash`
- Установка и запуск приложения
  `curl -fsSL https://github.com/oleg-mordvintsev/mo-home-media/raw/refs/heads/main/init.sh | sed 's/sudo //g' | bash`

- Установка происходит в директорию `/mo-home-media`
    - `cd /mo-home-media && ./start.sh`
    - `cd /mo-home-media && ./stop.sh`

# TODO:

- Подключить логирование
- Покрыть тестами
