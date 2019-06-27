# Syntax

```yaml
expression: slice
slice:
    ("has": string)?
    ("between": [between])?
    ("plugin": string | [string] )?
    (   ("value": any) |
        ("slice": slice | [slice] ) |
        ("array": array) |
        ("object": object) |
        ("convert": string) )?
between:
    (backward: true | false)?
    (prefix: string | [string] ))?
    (suffix:  string | [string] )?
    (trim: true | false)?
array:
    ("separator":  string | [string] )?
    ("omit": true | false )?
    ("item": slice | [slice | null] )?
object:
    ( $name: slice | [slice | null] )*
```
