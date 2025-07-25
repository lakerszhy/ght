# GHT - GitHub Trending TUI

`GHT` is a terminal-based user interface (TUI) for browsing GitHub trending repositories. It allows you to view trending repositories by different time periods (daily, weekly, monthly) and filter by programming language.

## Features

- View GitHub trending repositories directly in your terminal
- Browse trending repositories by different time periods:
  - Today
  - This week
  - This month
- Filter repositories by programming language

## Installation

### Using Go

```bash
go install github.com/lakerszhy/ght@latest
```

### From Source

```bash
git clone https://github.com/lakerszhy/ght.git
cd ght
go build
```

## Usage

```bash
# View all trending repositories
ght
```

```bash
# Filter by programming language
ght -l go
```

## Keyboard Controls

- `Tab`: Switch between time periods (Today, This week, This month)
- `↑/↓`: Navigate through repositories
- `Enter` or `o`: Open the selected repository in your browser
- `Ctrl+C`: Quit the application

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea): Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles): UI components for Bubble Tea
- [Lip Gloss](https://github.com/charmbracelet/lipgloss): Style definitions for terminal applications
- [goquery](https://github.com/PuerkitoBio/goquery): HTML parsing library

## License

MIT