# Этап 1: Сборка
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Копируем модули
COPY go.mod go.sum* ./
RUN if [ -f go.mod ]; then go mod download; fi

# Копируем исходники (только то, что нужно для сборки)
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Этап 2: Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Копируем бинарник
COPY --from=builder /build/main .

# Копируем шаблоны в директорию относительно WORKDIR
COPY --from=builder /build/internal/delivery/http/templates ./templates

# Создаём непривилегированного пользователя
RUN adduser -D -u 1000 appuser
USER appuser

# Порты
ENV PORT=1777
EXPOSE $PORT

# Запуск
CMD ["./main"]