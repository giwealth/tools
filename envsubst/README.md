# ENVSUBST

OPERATION             | DESC
----------------------|---------------------------------------------------------
`${param}`            | parameter substitution
`${param,}`           | parameter substitution with lowercase first char
`${param,,}`          | parameter substitution with lowercase
`${param^}`           | parameter substitution with uppercase first char
`${param^^}`          | parameter substitution with uppercase
`${param:pos}`        | parameter substitution with substring
`${param:pos:len}`    | parameter substitution with substring and length
`${param=default}`    | parameter substitution with default
`${param##prefix}`    | parameter substitution with prefix removal
`${param%%suffix}`    | parameter substitution with suffix removal
`${param/old/new}`    | parameter substitution with find and replace
