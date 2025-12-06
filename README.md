DotmanGO

DotmanGO is a lightweight, fast, and minimalistic dotfiles manager written in Go.
It helps you organize, version, and synchronize your configuration files without relying on Git or complex setup.

DotmanGO focuses on simplicity, portability, and predictability.

✨ Features

📂 Manage dotfiles inside a single directory: ~/.dotfiles

🔄 Preserve relative paths when adding files

🔗 Optional symbolic link creation (configurable)

🧽 Non-destructive operations — safe by default

⚡ Extremely fast (written in Go)

🧰 Simple CLI interface powered by Cobra

🛠 Does not require Git (but you can use Git in .dotfiles if you want)

📦 Installation
go install github.com/FranciscoHuenchunir/dotmango@latest


Or download a binary from Releases (si en el futuro los publicas).

🚀 Quick Start
1. Initialize DotmanGO
dotman init


This creates:

~/.dotfiles/
~/.config/dotman/config.yaml

2. Add a file or directory
dotman add ~/.config/nvim/init.lua


This will:

Move the file into ~/.dotfiles/.config/nvim/init.lua

Optionally create a symlink at the original location (if enabled in config)

3. List tracked dotfiles
dotman list

4. Remove a dotfile from management
dotman remove ~/.config/nvim/init.lua

⚙️ Configuration

DotmanGO stores its configuration at:

~/.config/dotman/config.yaml


Example:

symlink_on_add: false


Set symlink_on_add: true to automatically create symlinks when adding files.

📘 Commands Overview
Command	Description
dotman init	Initializes the .dotfiles directory
dotman add <path>	Adds a file or directory to dotman
dotman list	Shows all tracked dotfiles
dotman remove <path>	Stops tracking a dotfile
dotman sync	Synchronize symlinks (planned)
dotman status	Shows changes (planned)
📚 Philosophy

DotmanGO follows three core principles:

Simple logic → Move files; optionally symlink them.

Predictable behavior → No magical Git tricks.

User freedom → You decide when to use Git, symlinks, or plain file management.

🛠 Build from source
git clone https://github.com/FranciscoHuenchunir/dotmango
cd dotmango
go build -o dotman

🤝 Contributing

PRs are welcome. Issues are welcome.
DotmanGO is still evolving, so any idea is appreciated.
