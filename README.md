# ｔｙｇｏｒ

**API-First approach k6 extension development**

The functionality of [k6](https://k6.io) can be extended using [JavaScript Extensions](https://k6.io/docs/extensions/get-started/create/javascript-extensions/), which can be created in the go programming language. **tygor** allows you to develop these extensions using an [API-First approach](https://swagger.io/resources/articles/adopting-an-api-first-approach/). A [TypeScript declaration file](https://www.typescriptlang.org/docs/handbook/declaration-files/introduction.html) can be used as [IDL](https://en.wikipedia.org/wiki/Interface_description_language) to define the JavaScript API of the extension.

From the TypeScript declaration file, tygor generates the go interfaces needed to implement the API, as well as the binding code between the go implementation and the JavaScript runtime. In addition, tygor is also able to generate a skeleton implementation to help create a go implementation.

**Features**

- uses a TypeScript declaration file to define the JavaScript API
- generates go interfaces matching JavaScript API
- generates bindings between JavaScript and go
- generates api documentation in Markdown or HTML format
- inserts the API documentation into an outer document (eg: README.md)
- all generated output can be updated when the API definition changes
- enables to focus only on implementing the extensions's business logic
- a single binary without dependencies

Currently, tygor is still in a relatively early stage of development, but it is already usable. The binding code generation will change in the future (e.g. due to optimization), but this will probably not affect the go interfaces to be implemented. That is, it will be sufficient to regenerate the binding code with the new version of tygor. See the [roadmap](#roadmap) section for more information.


## Crash course
<details>
<summary>Click to expand...</summary>

In short, for this TypeScript API:

```ts
export as namespace hitchhiker;

type int = number;

export declare class Guide {
  question: string;
  readonly answer: int;

  constructor(question: string);
  check(value: int): boolean;
}

declare const defaultGuide: Guide;

export default defaultGuide;
```

..these generated interfaces have to be implemented in go language:

```go
type goGuide interface {
  checkMethod(valueArg int) (bool, error)
  questionGetter() (string, error)
  questionSetter(v string) error
  answerGetter() (int, error)
}

type goModule interface { 
  newGuide(questionArg string) (goGuide, error)
  defaultGuideGetter() (goGuide, error)
}

type goModuleConstructor func(vu modules.VU) goModule
```

...and that's it! The rest is handled by the go code generated by tygor.

After that, the extension will be usable in [k6](https://k6.io/) like this:

```js
import guide, { Guide } from "k6/x/hitchhiker"

export default function() {
  console.log(guide.answer)                           // 42
  console.log(guide.check(13))                        // false
  console.log(guide.check(42))                        // true

  const other = new Guide("What is life all about?")
  console.log(other.answer)                           // 42
  console.log(other.question)                         // What is life all about?

  other.question = "Why are we here?"

  console.log(other.question)                         // Why are we here?
}
```
</details>

See the complete example in the [examples/hitchhiker](examples/hitchhiker) directory. There are more examples in the [examples](examples) directory.

Check out the [intro slides](https://ivan.szkiba.hu/tygor/intro/) for a quick introduction.

## Install

Precompiled binaries can be downloaded and installed from the [Releases](https://github.com/szkiba/tygor/releases) page.

If you have a go development environment (probably you do), the installation can also be done with the following command:

```bash
go install github.com/szkiba/tygor@latest
```

## Usage

Check [CLI Reference](#cli-reference) section for detailed command line usage.

The following example generates the `hitchhiker_bindings.go` file containing the go interfaces and bindings and the `hitchhiker_skeleton.go` file containing the skeleton implementation from the `hitchhiker.d.ts` declaration file in the current directory. By default, generated files are placed in the same directory as the declaration file.

```bash
$ # generate both bindings and skeletons
$ tygor --skeleton hitchhiker.d.ts       
$ ls
hitchhiker_bindings.go
hitchhiker_skeleton.go
hitchhiker.d.ts
```

The skeleton file can be used as a sample for the implementation. Since it contains a special go build tag (`//go:build skeleton`), its presence will not interfere with the real implementation. To start the implementation, simply copy the skeleton file under a different name (or rename it) and delete the comments at the beginning of the file. If the declaration file changes, the bindings and skeleton can be regenerated at any time, and the skeleton can be used to help implement the changes.

In the above example, the implementation can be started simply by copying the `hitchhiker_skeleton.go` file in the `hitchhiker.go` file.

```bash
$ cp hitchhiker_skeleton.go hitchhiker.go
```

Don't forget to delete the following two comment lines from the beginning of the `hitchhiker.go` file

```go
// Code generated by tygor; DO NOT EDIT.
//go:build skeleton
```

See also the [tygor](#tygor) command.

The command below inserts the generated API documentation into the `README.md` file at the location marked with marker comments:

```bash
# generate and inject API documentation
$ tygor doc --inject README.md hitchhiker.d.ts
```

See also the [tygor doc](#tygor-doc) command.

## Declarations

This section describes the TypeScript declarations that can be used in the API definition.

### `interface`

One of the typical uses of the interface declaration is to describe the class implemented by the extension, which cannot be instantiated from the JavaScript code. The other typical use is the description of the object that contains the optional function/method parameters.

The interface declaration can contain [property](#property) and [method](#method) declarations and its name typically begins with a capital letter.

A go interface is created from the interface declaration, with a name consisting of the `go` prefix and the declared interface name.

<details><summary>TypeScript</summary>

```ts
export declare interface Interface1 {
  // property declarations
  // method declarations
}
```

</details>

<details><summary>go</summary>

```go
type goInterface1 interface {
  // methods
  // property getters
  // property setters
}
```

</details>

#### `property`

Getter and setter methods are created from the property declaration in the containing go interface. In the case of a `readonly` property, only a getter method is created. The getter method name consists of the property name and the `Getter` suffix, while the setter method name consists of the property name and the `Setter` suffix. The getter method returns the value of the property and, in case of an error, an error value. The setter method returns an error value in case of an error. The property types are mapped as described in the [type](#type) section.

The property name typically starts with a lowercase letter.

<details><summary>TypeScript</summary>

```ts
export declare interface Interface1 {
  prop1 : number;
  readonly prop2 : string;
}
```

</details>

<details><summary>go</summary>

```go
type goInterface1 interface {
  prop1Getter() (float64, error)
  prop1Setter(v float64) error
  prop2Getter() (string, error)
}
```

</details>

#### `method`

A method is created from the method declaration in the containing go interface. The name of the go method is the declared name plus the `Method` suffix. The parameters of the go method correspond to the parameters of the declared method. The parameter types are mapped as described in the [type](#type) section.

The method name typically starts with a lowercase letter.

<details><summary>TypeScript</summary>

```ts
export declare interface Interface1 {
  method1(arg1:number, arg2:boolean) : number;
}
```

</details>

<details><summary>go</summary>

```go
type Interface1 interface {
  method1Method(arg1Arg float64, arg2Arg bool) (float64, error)
}
```

</details>

### `class`

A typical use of a class declaration is to describe classes implemented by an extension that can be instantiated from JavaScript code.

The class declaration can contain [constructor](#constructor), [property](#property) and [method](#method) declarations and its name typically begins with a capital letter.

<details><summary>TypeScript</summary>

```ts
export declare class Class1 {
  // property declarations
  // method declarations
  // constructor declaration
}
```

</details>

<details><summary>go</summary>

```go
type goClass1 interface {
  // methods
  // property getters
  // property setters
}

type goModule interface {
  newClass1() (goClass1, error)
}
```

</details>

#### `constructor`

Factory methods are created from the constructor declarations in the module's go interface (`goModule`). The parameters of the factory method correspond to the parameters of the constructor, and its return value is the go interface belonging to the class declaration or an error in case of an error. The name of the factory method consists of the `new` prefix and the declared name of the class.

<details><summary>TypeScript</summary>

```ts
export declare interface Class1 {
  // properties and methods
  constructor(arg1:number, arg2:string);
}
```

</details>

<details><summary>go</summary>

```go
type goClass1 interface {
  // property setters, getters and methods
}

type goModule interface {
  newClass1(arg1Arg float64, arg2Arg string) (goClass1, error)
}
```

</details>

### `namespace`

The name of the k6 extension can be specified using the `export as namespace` declaration. Using local or nested namespace declarations is not supported. The generated `register()` function uses the namespace name to register the extension under the `k6/x/` path.

An interface declaration named `Module` is implicitly created from the variables and functions of the namespace. Variable declarations become property declarations (constant variables become readonly properties), and function declarations become method declarations in the implicit `Module` interface.

The go interface (`goModule`) belonging to the `Module` interface, also contains the factory methods of the go interfaces belonging to the class declarations. These methods are used in the generated code to instantiate go interfaces. The `goModule` interface is instantiated using a `goModuleConstructor` type function. This function must be implemented by the extension developer. The [modules.VU](https://pkg.go.dev/go.k6.io/k6/js/modules#VU) interface can be used as a parameter.

<details><summary>TypeScript</summary>

```ts
export as namespace module1;              // export declare interface Module {
export declare var variable1: boolean;    //   variable1: boolean;
export declare const variable2: string;   //   readonly variable2: string;
export declare function func1(): number;  //   func1(): number;
                                          // }
```

</details>

<details><summary>go</summary>

```go
type goModule interface {
  variable1Getter() (bool, error)
  variable1Setter(v bool) error
  variable2Getter() (string, error)
  func1Method() (float64, error)

  // factory methods for classes
}

type goModuleConstructor func(vu modules.VU) goModule

func register(ctor goModuleConstructor) {
  // ...
	modules.Register("k6/x/module1", m)
}

```
</details>

<details><summary>skeleton</summary>

```go
type goModuleImpl struct {}

var _ goModule = (*goModuleImpl)(nil)

func newModule(_ modules.VU) goModule {
	return new(goModuleImpl)
}

func init() {
	register(newModule)
}
```

</details>

### `type`

The type alias declaration can be used to define a mapping different from the default type mapping. Currently, it can be defined for the `number` type in the form of a type alias to which go type it should be mapped.

<details>
  <summary>TypeScript</summary>

```ts
type int = number;

export declare interface Interface1 {
  prop1 : int;
}
```

</details>

<details>
  <summary>go</summary>

```go
type Interface1 interface {
  prop1Getter() (int, error)
  prop1Setter(v int) error
}
```

</details>


Default type mappings:

js/ts | go
-----------|--------
number     | float64
string     | string
boolean    | bool
ArrayBuffer| []byte
Date       | time.Time
any        | interface{}
object     | interface{}

Supported type aliases:

```ts
type int    = number;
type int8   = number;
type int16  = number;
type int32  = number;
type int64  = number;
type uint   = number;
type uint8  = number;
type uint16 = number;
type uint32 = number;
type uint64 = number;

type float32 = number;
type float64 = number;

type rune = number;
type byte = number;
```

## Roadmap

Currently, tygor is still in a relatively early stage of development. Many features have not yet been implemented, and there are still many opportunities to optimize the generated binding code. The following (non-exhaustive) list contains planned future developments:

- array type support (`Array<T>`)
- record type support (at least `Record<string, T>`)
- property adapter optimization (for properties of `interface` or `class` type)
- improving the go code generator
- improving the documentation generator

## How It Works

`tygor` runs the TypeScript compiler using a built-in JavaScript interpreter ([goja](https://github.com/dop251/goja)). [Using the TypeScript Compiler API](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API), the extractor (implemented in TypeScript) generates a JSON string from the declaration file, which contains the declarations and their [TSDoc](https://tsdoc.org/) documentation comments. This JSON string is parsed by the go code and this is how the API model is created.

The generator subcommands generate output in different formats from the API model.

**doc**

The [`doc`](#tygor-doc) subcommand generates Markdown/HTML documentation from the API model. The generation is done using go template. The [slim-sprig](https://github.com/go-task/slim-sprig) template function library used in the template. The generated Markdown text will be formatted using [blackfriday](https://github.com/russross/blackfriday) (with [markdownfmt](https://github.com/shurcooL/markdownfmt) as renderer). The HTML output is generated from the markdown output using [blackfriday](https://github.com/russross/blackfriday).

You can specify your own Markdown template using the `--template` flag. The [default Markdown template](internal/doc/doc.gtpl) is a good starting point for creating your own Markdown template. 

Both Markdown and HTML output can be inserted into an outer document, in a place marked by so-called marker comments:

```html
<!-- begin:api -->
generated API documentation goes here
<!-- end:api -->
```

The [default HTML outer document](internal/doc/outer.html) is a good starting point for creating your own HTML outer document.

Documentation for extensions usually includes common sections. For example, how to build k6 with the extension, or download pre-built k6 binaries, etc.

For different extensions, these boilerplate documentation sections differ almost only in the extension name and the repository URL. Consequently, these sections can be easily generated.

The [`doc`](#tygor-doc) subcommand can generate these boilerplate sections if the necessary parameters (eg repository name) are specified or detected. Thus, the extension developer does not have to write these sections, and if the tooling changes (e.g. the xk6 tool changes or improves), they are simply re-generable.

By default, GitHub repository and generateable boilerplate sections are automatically detected. This is done by examining the git configuration, the GitHub workflows configuration, and the examples directory.

**parse**

The [`parse`](#tygor-parse) subcommand simply displays (or writes to a file) the API model in JSON format. With its use, the API model can be processed by external programs without the complexity of TypeScript parsing.

**gen**

The [`gen`](#tygor-gen) subcommand generates go source code from the API model using the [Jennifer](https://github.com/dave/jennifer) go source code generator.

The go interfaces to be implemented and the JavaScript binding code are placed in the file with the `_bindings.go` suffix. And the file with the suffix `_skeleton.go` contains the skeleton implementations (this is optional). The call to register the extension is placed in the `init()` function of the skeleton file.

The generated binding code performs bidirectional mapping between JavaScript and go objects.

# CLI Reference
<!-- begin:cli -->
## tygor

CLI tool that enables the development of k6 extensions with an API-First approach.

### Synopsis

The functionality of k6 can be extended using JavaScript Extensions, which can be created in the go programming language. Tygor allows you to develop these extensions using an API-First approach. A TypeScript declaration file can be used as IDL to define the JavaScript API of the extension.

From the TypeScript declaration file, tygor generates the go interfaces needed to implement the API, as well as the binding code between the go implementation and the JavaScript runtime. In addition, tygor is also able to generate a skeleton implementation to help create a go implementation.

The skeleton file can be used as a sample for the implementation. Since it contains a special go build tag (//go:build skeleton), its presence will not interfere with the real implementation. To start the implementation, simply copy the skeleton file under a different name (or rename it) and delete the comments at the beginning of the file. If the declaration file changes, the bindings and skeleton can be regenerated at any time, and the skeleton can be used to help implement the changes.

The only mandatory argument is the name of the declaration file (which file name must end with a .d.ts suffix). In addition, different flags can be used to modify the generation output. 

The tygor command generates go source code by default, but it can also generate other outputs. Other outputs can be generated using subcommands. Using it without the subcommand is equivalent to using the gen subcommand.

Use the -h flag to get detailed help on subcommands and flags.


```
tygor file [flags]
```

### Examples

```
$ tygor --skeleton hitchhiker.d.ts
```

### Options

```
  -h, --help             help for tygor
  -o, --output string    output directory (default: same as input)
  -p, --package string   go package name (default: module name)
  -s, --skeleton         enable skeleton generation (default: disabled)
```

### SEE ALSO

* [tygor doc](#tygor-doc)	 - Generate documentation from k6 extension's API definition.
* [tygor gen](#tygor-gen)	 - Generate golang source code from k6 extension's API definition.
* [tygor parse](#tygor-parse)	 - Convert k6 extension's API definition to JSON data model.

---
## tygor doc

Generate documentation from k6 extension's API definition.

### Synopsis

From the TypeScript declaration file, tygor doc subcommand generates API documentation.

API documentation is generated to standard output in Markdown format by default. If the --html flag is used, the output format will be HTML.

The output can also be saved to a file using the --output flag. In this case, the default format is determined from the file extension: in the case of .htm and .html extensions, it will be in HTML format, otherwise it will be in Markdown format. Using the --html flag, the HTML format can also be forced for other file extensions.

API documentation can also be inserted (and updated) into an existing Markdown or HTML document using the --inject flag. The insertion takes place in the place marked by so-called marker comments:

    <!-- begin:api -->
    generated API documentation goes here
    <!-- end:api -->

The generated API documentation starts at heading level 1 by default. The starting heading level can be specified by using the --heading flag, which can be useful, for example, when inserting into an outer document.

The documentation may include the usual extension documentation sections, such as build instructions, download instructions, a link to the examples folder, etc. The required GitHub repository can be specified using the --github-repo flag. Otherwise, the tygor doc subcommand tries to guess the GitHub repository from the git configuration (if it exists). This automation can be disabled with the --no-auto flag.
By default, GitHub repository and generateable boilerplate sections are automatically detected. This is done by examining the git configuration, the GitHub workflows configuration, and the examples directory.

The only mandatory argument to the doc subcommand is the name of the declaration file (which file name must end with a .d.ts suffix).


```
tygor doc file [flags]
```

### Examples

```
$ tygor doc -o README.md hitchhiker.d.ts
```

### Options

```
      --github-repo string   GitHub repository (owner/name)
      --heading uint         initial heading level (default 1)
  -h, --help                 help for doc
      --html                 enable HTML output (default: based on file ext)
  -i, --inject string        inject into outer file
      --link-examples        enable examples folder link
      --link-packages        enable GitHub container packages link
      --link-releases        enable GitHub releases link
      --no-auto              disable automatic GitHub repo and link flags detection
  -o, --output string        output file (default: standard output)
  -t, --template string      go template file for markdown generation
```

### SEE ALSO

* [tygor](#tygor)	 - CLI tool that enables the development of k6 extensions with an API-First approach.

---
## tygor gen

Generate golang source code from k6 extension's API definition.

### Synopsis

From the TypeScript declaration file, tygor gen subcommand generates the go interfaces needed to implement the API, as well as the binding code between the go implementation and the JavaScript runtime. In addition, tygor gen subcommand is also able to generate a skeleton implementation to help create a go implementation.

The skeleton file can be used as a sample for the implementation. Since it contains a special go build tag (//go:build skeleton), its presence will not interfere with the real implementation. To start the implementation, simply copy the skeleton file under a different name (or rename it) and delete the comments at the beginning of the file. If the declaration file changes, the bindings and skeleton can be regenerated at any time, and the skeleton can be used to help implement the changes.

The only mandatory argument is the name of the declaration file (which file name must end with a .d.ts suffix).


```
tygor gen file [flags]
```

### Examples

```
$ tygor gen --skeleton hitchhiker.d.ts
```

### Options

```
  -h, --help             help for gen
  -o, --output string    output directory (default: same as input)
  -p, --package string   go package name (default: module name)
  -s, --skeleton         enable skeleton generation (default: disabled)
```

### SEE ALSO

* [tygor](#tygor)	 - CLI tool that enables the development of k6 extensions with an API-First approach.

---
## tygor help

Help about any command

### Synopsis

Help provides help for any command in the application.
Simply type tygor help [path to command] for full details.

```
tygor help [command] [flags]
```

### Options

```
  -h, --help   help for help
```

### SEE ALSO

* [tygor](#tygor)	 - CLI tool that enables the development of k6 extensions with an API-First approach.

---
## tygor parse

Convert k6 extension's API definition to JSON data model.

### Synopsis

From the TypeScript declaration file,  tygor parse subcommand generates the API model in JSON format. The API model can be processed by external programs without the complexity of TypeScript parsing.

The only mandatory argument of the tygor parse subcommand is the name of the declaration file (which file name must end with a .d.ts suffix).		


```
tygor parse file [flags]
```

### Examples

```
$ tygor parse hitchhiker.d.ts | jq
```

### Options

```
  -h, --help            help for parse
  -o, --output string   output file (default: standard output)
```

### SEE ALSO

* [tygor](#tygor)	 - CLI tool that enables the development of k6 extensions with an API-First approach.

<!-- end:cli -->
