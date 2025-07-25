# GHT - GitHub Trending TUI

![ght](/asset/ght.png)

`GHT` is a terminal-based user interface (TUI) for browsing GitHub trending repositories. It allows you to view trending repositories by different time periods (daily, weekly, monthly) and filter by programming language.

## Features

- View GitHub trending repositories directly in your terminal
- Browse trending repositories by different time periods:
  - Today
  - This week
  - This month
- Filter repositories by programming language

## Installation && Update

### Linux && macOS

```sh
curl -fsSL https://raw.githubusercontent.com/lakerszhy/ght/main/install.sh | sh
```

### Windows

```powershell
irm https://raw.githubusercontent.com/lakerszhy/ght/main/install.ps1 | iex
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
- `↑/k`: Move up
- `↓/j`: Move down
- `left/h/pgup/pgdn/b/u`: Previous page
- `right/l/pgdn/pgup/f/d`: Next page
- `Enter` or `o`: Open the selected repository in your browser
- `Ctrl+C`: Quit the application

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [goquery](https://github.com/PuerkitoBio/goquery)