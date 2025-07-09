## miru instance to jsonschema

Convert a config instance file to a JSON Schema file

### Synopsis

Convert a config instance file to a JSON Schema file.

```
miru instance to jsonschema PATH [OPTIONS]
```

### Examples

```

# Indentation level 4
miru instance to jsonschema ./path/to/config/instance.json -i 4

# Exclude type inference
miru instance to jsonschema ./path/to/config/instance.json --types=false

# Insert the config type slug
miru instance to jsonschema ./path/to/config/instance.json --config-type-slug=my-config-type

# Save to a file
miru instance to jsonschema ./path/to/config/instance.yaml -o ./path/to/jsonschema.yaml
```

### Options

```
      --config-type-slug string   Include the specified miru config type slug in the output JSON Schema file
  -d, --defaults                  Include default values in the output JSON Schema file (default true)
  -h, --help                      help for jsonschema
  -i, --indent uint               Indentation level for the output JSON Schema file (default 2)
  -o, --output string             Output file path (default: stdout)
  -t, --types                     Include type information in the output JSON Schema file (default true)
```

### Options inherited from parent commands

```
  -q, --quiet   suppress output except errors
```

### SEE ALSO

* [miru instance to](miru_instance_to.md)	 - Convert a config instance file to a different format

