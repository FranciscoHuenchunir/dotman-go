package model

type Config struct {
	SymlinkOnAdd bool   `toml:"symlink_on_add"`
	DotfilesDir  string `toml:"dotfiles_dir"`
	IgnoreFile   string `toml:"ignore_file"`
}
