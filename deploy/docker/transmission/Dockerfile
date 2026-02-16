# Название контейнера с DockerHub
FROM lscr.io/linuxserver/transmission:latest

# Рабочая директория в контейнере (используется локальная, смотри docker-compose.yml в корне проект)
#WORKDIR /mo-home-media-app

# Создание директории, изменение владельца и прав доступа при сборке образа
RUN mkdir -p /downloads/incomplete && \
    chown abc:users /downloads/incomplete && \
    chmod 777 /downloads/incomplete
