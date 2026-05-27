# 🌧️ Gorain 🌧️

Gorain is a small terminal animation written in Go. It draws falling rain drops directly in the terminal using [`tcell`](https://github.com/gdamore/tcell), with randomized drop positions, speed, length, symbols, and color intensity.

## Features

- Smooth terminal rendering with `tcell`
- Randomized rain drops for a natural-looking animation
- Automatic redraw on terminal resize
- Graceful exit with `Esc` or `Ctrl+C`
- Simple file logging to `log.log`

## Requirements

- Go 1.26 or newer
- A terminal with support for ANSI-style rendering

## Getting Started

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

## Controls

| Key      | Action               |
| -------- | -------------------- |
| `Esc`    | Exit the application |
| `Ctrl+C` | Exit the application |

## Project Structure

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

## Logging

Gorain writes runtime logs to `log.log` in the current working directory. The file is created automatically if it does not exist.

## Future plans

- [ ] Change drawing speed based on monitor frequency
- [ ] Add options for configuration

## License

This project is distributed under the MIT License. See [LICENSE](LICENSE) for details.
