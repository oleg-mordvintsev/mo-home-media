#!/bin/bash

# MO-HOME-MEDIA Установщик

mkdir mo-home-media
cd mo-home-media

if [ -d ".git" ]; then
    echo "✓ Репозиторий уже клонирован. Обновление..."
    git pull
else
    echo "Клонирование репозитория..."
    git clone https://github.com/oleg-mordvintsev/mo-home-media.git .
fi

echo "Создание необходимых директорий и установка прав..."
mkdir data
chmod -R 755 data

echo "Установка прав на скрипты..."
chmod +x start.sh
chmod +x stop.sh
chmod +x docker-install.sh

# Проверка установленного Docker
if ! command -v docker &> /dev/null; then
    echo "ОШИБКА: Docker не установлен. Пожалуйста, сначала установите Docker."
    echo "Инструкции по установке: https://docs.docker.com/engine/install/"
    echo "Так же можно попробовать установить с помощью скрипта docker-install.sh"
    exit 1
fi

# Проверка Docker Compose
if ! docker compose version &> /dev/null && ! command -v docker-compose &> /dev/null; then
    echo "ОШИБКА: Docker Compose не установлен. Пожалуйста, сначала установите Docker Compose."
    exit 1
fi

echo "✓ Docker и Docker Compose установлены."

# Создание файла окружения для учетных данных если не существует
if [ ! -f ".env" ]; then
    echo "Создание .env файла..."
    cat > .env << EOF
PROJECT_NAME=mo-home-media
PORT_XML_PLAYER=777
PORT_TORRENT_WEB=888
DIR_TEMPLATES=templates
DIR_DATA=data
EOF
    echo "✓ Создан .env файл"
    echo "Пожалуйста, проверьте и отредактируйте .env файл для вашей конфигурации, если требуется!"
else
    echo "✓ .env файл уже существует"
fi

./start.sh

echo ""
echo "=== Установка завершена, приложение запущено ==="
echo ""
echo "Следующие шаги:"
echo "1. Для запуска и остановки используйте:"
echo "   ./start.sh"
echo "   ./stop.sh"
echo ""
echo "2. Откройте веб-интерфейс:"
echo "   http://ваш-сервер:777 - XML данные для телевизора"
echo "   http://ваш-сервер:888 - Торрент клиент для ПК или смартфона"
echo ""
echo "3. Для просмотра логов:"
echo "   docker-compose logs -f"
