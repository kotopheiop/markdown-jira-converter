# Markdown to Jira/Confluence Converter

Этот проект предназначен для конвертации файлов между синтаксисами Markdown и Jira/Confluence.

## Описание

Скрипт позволяет преобразовывать файлы из формата Markdown в Jira/Confluence синтаксис и обратно. Это полезно для автоматической генерации документации, которая может использоваться в системах управления проектами, таких как Jira или Confluence.

Поддерживаются следующие операции:

- Конвертация Markdown в Jira/Confluence.
- Конвертация Jira/Confluence в Markdown.

## Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/your-username/md-jira-converter.git
cd md-jira-converter
```

2. Убедитесь, что у вас установлен Go версии 1.16 или выше.

3. Соберите проект:
```bash
   go build -o md-jira-converter
```

## Использование
### Конвертация Markdown в Jira/Confluence:
```bash
./md-jira-converter -input <путь_к_входному_файлу> -output <путь_к_выходному_файлу> -from md -to jira
```
### Конвертация Jira/Confluence в Markdown:
```bash
./md-jira-converter -input <путь_к_входному_файлу> -output <путь_к_выходному_файлу> -from jira -to md
```

## Параметры:
- `-input` — путь к входному файлу (Markdown или Jira).
- `-output` — путь к выходному файлу, в который будет записан результат.
- `-from` — исходный формат: md или jira.
- `-to` — целевой формат: md или jira.

## Пример:
Конвертация файла README.md из Markdown в формат Jira:
```bash
./md-jira-converter -input README.md -output README.jira -from md -to jira
```
Конвертация файла README.jira из Jira в формат Markdown:
```bash
./md-jira-converter -input README.jira -output README.md -from jira -to md
```

## Как это работает
Скрипт использует регулярные выражения для преобразования различных элементов синтаксиса Markdown и Jira/Confluence:

- Преобразует заголовки.
- Конвертирует жирный текст, курсив и зачеркивания.
- Обрабатывает ссылки и списки.
- Поддерживает многострочные блоки кода, включая их конвертацию с указанием языка.
