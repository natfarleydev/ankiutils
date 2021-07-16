# ankiutils
A set of personal utility functions for use with AnkiConnect

It is designed for use in shell scripting, keyboard shortcuts, etc., e.g. `ankiutils addCard` will bring up an an Anki card for access to Anki.

`ankutils help` brings up:
```
A set of tools for Anki for my own personal use.
	
Initially, this is just to insert quick GURPS advantages. It will likely grow.

Usage:
  ankiutils [command]

Available Commands:
  addCard              Brings up the Add Card dialog from a currently open Anki session
  addGURPSAdvantage    Sets up an Add Card dialog to add a GURPS Advantage
  addGURPSDisadvantage Sets up an Add Card dialog to add a GURPS Disadvantage
  completion           Generate completion script
  help                 Help about any command

Flags:
      --config string   config file (default is $HOME/.ankiutils.yaml)
  -h, --help            help for ankiutils
  -t, --toggle          Help message for toggle

Use "ankiutils [command] --help" for more information about a command.
```

# Contributions

Feel free to fork and use adapt the commands. PRs that extend the general functionality are likely to be accepted, but this project remains for personal use so it may change unexpectedly. 
