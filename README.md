# 🌧️ Gorain 🌧️

<img width="1280" height="720" alt="gorain_opt" src="https://github.com/user-attachments/assets/53873eb6-2612-431a-a306-5b0ec6495565" />


<p>
  <a href="#english">English</a> |
  <a href="#russian">Русский</a>
</p>


<a id="english"></a>

## English

Gorain is a small terminal animation written in Go. It draws falling rain drops directly in the terminal using [`tcell`](https://github.com/gdamore/tcell), with randomized drop positions, speed, length, symbols, and color intensity.

### Features

- Smooth terminal rendering with `tcell`
- Randomized rain drops for a natural-looking animation
- Automatic redraw on terminal resize
- Graceful exit with `Esc` or `Ctrl+C`
- Simple file logging to `log.log`

### Requirements

- Go 1.26 or newer
- A terminal with support for ANSI-style rendering

### Getting Started

Clone the repository and run the app:

```bash
git clone https://github.com/talismanch1k/gorain.git
cd gorain
go run .
```

To build a binary:

```bash
go build -o gorain .
./gorain
```

### Controls

| Key      | Action               |
| -------- | -------------------- |
| `Esc`    | Exit the application |
| `Ctrl+C` | Exit the application |

### Project Structure

```text
.
├── main.go                  # Application entry point
├── internal/
│   ├── app/                 # Terminal screen rendering and rain drop logic
│   ├── closer/              # Small helpers for closing resources
│   └── logger/              # slog setup and file logging
├── go.mod
├── go.sum
└── LICENSE
```

### Logging

Gorain writes runtime logs to `log.log` in the current working directory. The file is created automatically if it does not exist.

### Future plans

- [ ] Change drawing speed based on monitor frequency
- [ ] Add options for configuration

### License

This project is distributed under the MIT License. See [LICENSE](LICENSE) for details.

<a id="russian"></a>

## Русский

Gorain - небольшая терминальная анимация дождя, написанная на Go. Приложение рисует капли прямо в терминале с помощью [`tcell`](https://github.com/gdamore/tcell): у каждой капли случайные позиция, скорость, длина, символ и интенсивность цвета.

### Возможности

- плавная отрисовка в терминале через `tcell`;
- случайная генерация капель для более живой анимации;
- корректная перерисовка при изменении размера окна;
- выход по `Esc` или `Ctrl+C`;
- простое логирование в файл `log.log`.

### Требования

- Go 1.26 или новее;
- терминал с поддержкой ANSI-отрисовки.

### Запуск

Склонируйте репозиторий и запустите приложение:

```bash
git clone https://github.com/talismanch1k/gorain.git
cd gorain
go run .
```

Чтобы собрать исполняемый файл:

```bash
go build -o gorain .
./gorain
```

### Управление

| Клавиша  | Действие            |
| -------- | ------------------- |
| `Esc`    | выйти из приложения |
| `Ctrl+C` | выйти из приложения |

### Структура проекта

```text
.
├── main.go                  # точка входа
├── internal/
│   ├── app/                 # отрисовка экрана и логика капель
│   ├── closer/              # вспомогательные функции закрытия ресурсов
│   └── logger/              # настройка slog и файлового логирования
├── go.mod
├── go.sum
└── LICENSE
```

### Логирование

Во время работы Gorain пишет служебные сообщения в `log.log` в текущей директории. Если файла нет, он будет создан автоматически.

### Планы

- [ ] менять скорость отрисовки с учетом частоты обновления монитора;
- [ ] добавить параметры конфигурации.

### Лицензия

Проект распространяется под лицензией MIT. Подробнее см. в [LICENSE](LICENSE).
