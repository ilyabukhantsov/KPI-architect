# Крок 1: Використовуємо образ із JDK 21
FROM eclipse-temurin:21-jdk-alpine

# Робоча директорія всередині контейнера
WORKDIR /app

# Копіюємо зібраний jar-файл із папки target у контейнер
# Переконайся, що назва збігається з тією, що в pom.xml (або використовуй *)
COPY target/*.jar app.jar

# Відкриваємо порт 8080
EXPOSE 8080

# Команда для запуску додатка
ENTRYPOINT ["java", "-jar", "app.jar"]