---
author: Iván SZKIBA
marp: true
theme: uncover
---

> I just want to implement my k6 extension's API!

# ｔｙｇｏｒ

**API-First approach k6 extension development**

*Iván Szkiba*

https://github.com/szkiba/tygor

<!--
Thank you for watching the presentation about API-First approach k6 extension development.

Let's see what this presentation will be about.
-->

---

# Agenda

- Use Case
- Under The Hood

<!--
-->

---

### What I want:

- define my k6 extension's API in an IDL
- just write the implementation in golang
- generate API doc from the definition

### What I don't want:

- write boilerplate code
- understand JavaScript bindings
- use any Node.js based tool

---

> I want to define my k6 extension's API in an IDL!
 
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

<!--
-->

---

> I don't want to write boilerplate code!
 
```bash
$ # generate only bindings
$ tygor hitchhiker.d.ts                  
$ ls
hitchhiker_bindings.go
hitchhiker.d.ts
```

&nbsp;

```bash
$ # generate both bindings and skeletons
$ tygor --skeleton hitchhiker.d.ts       
$ ls
hitchhiker_bindings.go
hitchhiker_skeleton.go
hitchhiker.d.ts
```

---

> I don't want to understand how the binding is done!

```go
type goGuide interface {                              // export declare class Guide {
  checkMethod(valueArg int) (bool, error)             //   check(value: int): boolean
  questionGetter() (string, error)                    //   question: string
  questionSetter(v string) error                      //
  answerGetter() (int, error)                         //   readonly answer: int
}                                                     // }

type goModule interface { 
  newGuide(questionArg string) (goGuide, error)       // constructor(question: string)
  defaultGuideGetter() (goGuide, error)               // export default defaultGuide
}

type goModuleConstructor func(vu modules.VU) goModule // export as namespace ...

```

&nbsp;
&nbsp;

---

> I just want to implement my k6 extension's API!

```go
func init() { register(newModule) }

func newModule(_ modules.VU) goModule {
	return &goModuleImpl{defaultGuide: &goGuideImpl{question: "What's up?"}}
}

type goModuleImpl struct{ defaultGuide goGuide }

func (self *goModuleImpl) newGuide(questionArg string) (goGuide, error) {
	return &goGuideImpl{question: questionArg}, nil
}

func (self *goModuleImpl) checkMethod(valueArg int) (bool, error) {
	return self.defaultGuide.checkMethod(valueArg)
}

func (self *goModuleImpl) defaultGuideGetter() (goGuide, error) { return self.defaultGuide, nil }

type goGuideImpl struct{ question string }

func (self *goGuideImpl) checkMethod(valueArg int) (bool, error) {
	return valueArg == 42, nil
}

func (self *goGuideImpl) questionGetter() (string, error) { return self.question, nil }

func (self *goGuideImpl) questionSetter(questionArg string) error {
	self.question = questionArg
	return nil
}

func (self *goGuideImpl) answerGetter() (int, error) {
	return 42, nil
}
```

<!--
-->

---

> Let's see how it works!

```js
import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import guide, { Guide } from "k6/x/hitchhiker";

export default function () {
  describe("default", () => {
    expect(guide).to.have.property("answer", 42);
    expect(guide).to.have.property("question", "What's up?");
    expect(guide.check(42)).to.be.true;
    expect(guide.check(43)).to.be.false;
    expect(() => (guide.answer = 2)).to.throw(TypeError);
    guide.question = "Why are we here?";
    expect(guide).to.have.property("question", "Why are we here?");
  });

  describe("Guide", () => {
    const guide = new Guide("What is life all about?");
    expect(guide).to.have.property("answer", 42);
    expect(guide).to.have.property("question", "What is life all about?");
    expect(guide.check(42)).to.be.true;
    expect(guide.check(43)).to.be.false;
    expect(() => (guide.answer = 2)).to.throw(TypeError);
    guide.question = "Why are we here?";
    expect(guide).to.have.property("question", "Why are we here?");
  });
}

export const options = { thresholds: { checks: ["rate==1"] } };
```

---

> I want to generate API doc without using Node.js!

```bash
$ # generate markdown documentation                 
$ tygor doc -o README.md faker.d.ts
$ ls
README.md
faker.d.ts
```

&nbsp;

```bash
$ # generate HTML documentation                     
$ tygor doc -o index.html faker.d.ts
$ ls
index.html
faker.d.ts
```

---

> ...but I want it as part of a larger document!

```bash
$ # inject as markdown into existing documentation   
$ tygor doc --inject README.md faker.d.ts
$ ls
README.md
faker.d.ts
```

&nbsp;

```bash
$ # inject as HTML into existing documentation        
$ tygor doc --inject index.html faker.d.ts
$ ls
index.html
faker.d.ts
```


---

> Oh, I want to do something special...

&nbsp;
&nbsp;

```bash
$ # convert TypeScript declarations to JSON        
$ tygor parse faker.d.ts | jq .
```

&nbsp;
&nbsp;
&nbsp;
&nbsp;

---

# Under The Hood

<!--
A few technological details about the tygor follow.
-->

---

## Design Considerations

<!--
What were the design considerations of tygor?

-->

- support an API-First approach
- use JavaScript's de facto IDL, TypeScript
- generate API documentation from IDL
- single binary without dependencies
- re-generatable output in case of API change

---

## How It Works


<!--
A few bullet points about how tygor works. A more detailed description can be found in the readme file.

-->

- embedding a real TypeScript compiler
- embedded JavaScript engine (goja)
- model building written in TypeScript
- JSON interface between script part and go
- Jennifer as a go source code generator
- go interface and implementation are separated
- go template engine for generating Markdown
- HTML generation from Markdown

---

That's All Folks!

<!--
That's all I wanted to share with you about the tygor in brief. I hope this is just the beginning of the journey.
-->
