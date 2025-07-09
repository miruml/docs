## miru completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(miru completion zsh)

To load completions for every new session, execute once:

#### Linux:

	miru completion zsh > "${fpath[1]}/_miru"

#### macOS:

	miru completion zsh > $(brew --prefix)/share/zsh/site-functions/_miru

You will need to start a new shell for this setup to take effect.


```
miru completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -q, --quiet   suppress output except errors
```

### SEE ALSO

* [miru completion](miru_completion.md)	 - Generate the autocompletion script for the specified shell

