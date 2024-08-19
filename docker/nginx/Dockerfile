# Название контейнера с DockerHub
FROM nginx:1.17

# Используем локальный файл настройки nginx в контейнере
ADD ./default.conf /etc/nginx/conf.d/default.conf

# Рабочая директория в контейнере (используется локальная, смотри docker-compose.yml в корне проект)
WORKDIR /home-media-app