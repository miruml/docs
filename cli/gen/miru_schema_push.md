## miru schema push

Upload one or more config schemas to miru

### Synopsis

Upload one or more config schemas to miru.

```
miru schema push PATH... [OPTIONS]
```

### Examples

```

# Single Schema
miru schema push ./path/to/config/schema.yaml

# Multiple Schemas
miru schema push \
	./path/to/config/schema1.yaml \
	./path/to/config/schema2.yaml
```

### Options

```
  -h, --help   help for push
```

### Options inherited from parent commands

```
  -q, --quiet   suppress output except errors
```

### SEE ALSO

* [miru schema](miru_schema.md)	 - Manage config schemas

