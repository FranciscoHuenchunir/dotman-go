# DotmanGO

A lightweight, fast, and minimalist dotfiles manager written in Go. DotmanGO helps you organize and synchronize your configuration files without relying on Git or complex setups, focusing on simplicity, portability, and predictability.

## Why DotmanGO?

I tried GNU Stow — confusing directory structure.  
I tried using a bare Git repository, but it didn’t fully convince me.

**I just wanted:**
- Put my dotfiles in one place
- Have them work in their original locations
- Not think about it anymore

**DotmanGO does exactly that. Nothing more, nothing less.**

## ✨ Features

- **📂 Centralized management**: Keep all your dotfiles organized in `~/.dotfiles`
- **🔄 Path preservation**: Maintains relative directory structure when adding files
- **🔗 Flexible symlinking**: Optional symbolic link creation (user-configurable)
- **🧽 Safe operations**: Non-destructive by default — your files are protected
- **⚡ Lightning fast**: Written in Go for maximum performance
- **🧰 Intuitive CLI**: Simple interface powered by Cobra
- **🛠 Git-independent**: No Git required (though you can use it if you want)

## 📦 Installation

### Using Go

```bash
go install github.com/FranciscoHuenchunir/dotman-go@latest
```

### Build from Source

```bash
git clone https://github.com/FranciscoHuenchunir/dotman-go
cd dotman-go
go build -o dotman
```

## 🚀 Quick Start

### 1. Initialize DotmanGO

```bash
dotman init
```

This creates the following structure:

```
~/.dotfiles/
```

### 2. Add a File or Directory

```bash
cd ~/.config/ 
dotman add nvim
dotman add starship.toml
```

What happens:
- File moves to `~/.dotfiles/.config/nvim/init.lua`
- Original path structure is preserved
- Symlink created at original location (if enabled in config)

## 📘 Commands

| Command | Description |
|---------|-------------|
| `dotman init` | Initialize `.dotfiles` directory, create `.dotignore` and config file |
| `dotman add <path>` | Add a file or directory to DotmanGO management |
| `dotman sync` | Synchronize symlinks *(planned)* |

## 📚 Philosophy

DotmanGO is built on three core principles:

1. **Simple logic** → Move files and optionally create symlinks — nothing more, nothing less
2. **Predictable behavior** → No hidden Git operations or magic tricks
3. **User freedom** → You choose when to use Git, symlinks, or plain file management

## 🎯 Use Cases

- **Developers**: Sync configs across machines without Git complexity
- **System admins**: Quick backup and restore of critical config files
- **Minimalists**: Simple dotfile management without unnecessary features
- **Students**: Understand exactly what happens to your files

## 🤝 Contributing

Contributions are welcome! Whether you want to:
- Report bugs
- Suggest features
- Submit pull requests
- Improve documentation

Feel free to open an issue or PR. DotmanGO is actively evolving, and community input helps shape its future.

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

Built with Go and [Cobra](https://github.com/spf13/cobra).

---

**Made with ❤️ by [Francisco Huenchunir](https://github.com/FranciscoHuenchunir)**
