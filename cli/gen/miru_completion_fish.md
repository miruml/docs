## miru completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	miru completion fish | source

To load completions for every new session, execute once:

	miru completion fish > ~/.config/fish/completions/miru.fish

You will need to start a new shell for this setup to take effect.


```
miru completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -q, --quiet   suppress output except errors
```

### SEE ALSO

* [miru completion](miru_completion.md)	 - Generate the autocompletion script for the specified shell

