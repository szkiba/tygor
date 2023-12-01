k6/x/toml
=========

xk6-toml enables k6 tests to comfortably encode and decode TOML values.

Usage
-----

Import an entire module's contents: `JavaScript import * as TOML from "k6/x/toml";`

Import a single export from a module: `JavaScript import { parse } from "k6/x/toml";`

API
===

### parse()

```ts
export declare function parse(text: string): object;
```

The parse() method parses a TOML string, constructing the JavaScript object described by the string.

### stringify()

```ts
export declare function stringify(value: object): string;
```

The stringify() method converts a JavaScript object to a TOML string.
