k6/x/faker
==========

xk6-faker random fake data generator

Faker
-----

This is Faker's main class containing all modules that can be used to generate data.

Please have a look at the individual modules and methods for more information and examples.

<details><summary><em>Example</em></summary>

```ts
import { Faker } from "k6/x/faker"

const faker = new Faker(11)

export default function() {
  console.log(faker.person.firstName()) // 'Josiah'
}
```

</details>

### Faker()

```ts
constructor(seed?: int64);
```

-	`seed` random seed value for deterministic generator

Creates a new instance of Faker.

Optionally, the value of the random seed can be set as a constructor parameter. This is intended to allow for consistent values in a tests, so you might want to use hardcoded values as the seed.

Please note that generated values are dependent on both the seed and the number of calls that have been made.

Setting seed to 0 (or omitting it) will use seed derived from system entropy.

<details><summary><em>Example</em></summary>

```ts
const consistentFaker = new Faker(11)
const semiRandomFaker = new Faker()
```

</details>

### Faker.person

```ts
readonly person: Person;
```

API to generate people's personal information such as names and job titles.

### Faker.company

```ts
readonly company: Company;
```

API to generate company related entries.

### Faker.hacker

```ts
readonly hacker: Hacker;
```

API to generate hacker/IT words and phrases.

### Faker.hipster

```ts
readonly hipster: Hipster;
```

API to generate hipster words, phrases and paragraphs.

### Faker.lorem

```ts
readonly lorem: Lorem;
```

API to generate random words, sentences, paragraphs, questions and quotes.

Person
------

API to generate people's personal information such as names and job titles.

### Person.firstName()

```ts
firstName(): string;
```

Generates a random first name.

*Returns a random first name*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.firstName() // 'Josiah'
```

</details>

### Person.lastName()

```ts
lastName(): string;
```

Generates a random last name.

*Returns a random last name*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.lastName() // 'Abshire'
```

</details>

### Person.prefix()

```ts
prefix(): string;
```

Generates a random person prefix.

*Returns a random person prefix.*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.prefix() // 'Mr.'
```

</details>

### Person.suffix()

```ts
suffix(): string;
```

Generates a random person suffix.

*Returns a random person suffix*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.suffix() // 'Sr.'
```

</details>

### Person.sexType()

```ts
sexType(): string;
```

Generates a random sex type.

*Returns a random sex type `male`|`female`*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.sexType() // 'male'
```

</details>

### Person.jobTitle()

```ts
jobTitle(): string;
```

Generates a random job title.

*Returns a random job title*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.jobTitle() // 'Representative'
```

</details>

### Person.jobLevel()

```ts
jobLevel(): string;
```

Generates a random job level.

*Returns a random job level*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.jobLevel() // 'Identity'
```

</details>

### Person.jobDescriptor()

```ts
jobDescriptor(): string;
```

Generates a random job descriptor.

*Returns a random job descriptor*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.person.jobDescriptor() // 'Internal'
```

</details>

Company
-------

API to generate company related entries.

### Company.name()

```ts
name(): string;
```

Generates a random company name string.

*Returns a random company name string*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.company.name() // 'Xatori'
```

</details>

### Company.suffix()

```ts
suffix(): string;
```

Generates a random company suffix string.

*Returns a random company suffix string*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.company.suffix() // 'LLC'
```

</details>

### Company.buzzWord()

```ts
buzzWord(): string;
```

Generates a random company buzz word string.

*Returns a random company buzz word string*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.company.buzzWord() // 'Reverse-engineered'
```

</details>

### Company.bs()

```ts
bs(): string;
```

Generates a random company bs string.

*Returns a random company bs string*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.company.bs() // '24-7'
```

</details>

Hacker
------

API to generate hacker/IT words and phrases.

### Hacker.abbreviation()

```ts
abbreviation(): string;
```

Generates a random hacker/IT abbreviation.

*Returns a random hacker/IT abbreviation*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.abbreviation() // 'Xatori'
```

</details>

### Hacker.adjective()

```ts
adjective(): string;
```

Generates a random hacker/IT adjective.

*Returns a random hacker/IT adjective*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.adjective() // 'Xatori'
```

</details>

### Hacker.ingverb()

```ts
ingverb(): string;
```

Generates a random hacker/IT verb for continuous actions (en: ing suffix; e.g. hacking).

*Returns a random hacker/IT verb for continuous actions*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.ingverb() // 'Xatori'
```

</details>

### Hacker.noun()

```ts
noun(): string;
```

Generates a random hacker/IT noun.

*Returns a random hacker/IT noun*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.noun() // 'Xatori'
```

</details>

### Hacker.phrase()

```ts
phrase(): string;
```

Generates a random hacker/IT phrase.

*Returns a random hacker/IT phrase*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.phrase() // 'Xatori'
```

</details>

### Hacker.verb()

```ts
verb(): string;
```

Generates a random hacker/IT verb.

*Returns a random hacker/IT verb*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hacker.verb() // 'Xatori'
```

</details>

Hipster
-------

API to generate hipster words, phrases and paragraphs.

### Hipster.word()

```ts
word(): string;
```

Generates a single hipster word.

*Returns a single hipster word*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hipster.word() // 'Xatori'
```

</details>

### Hipster.sentence()

```ts
sentence(wordCount: int): string;
```

-	`wordCount` the number of words

Generates a random hipster sentence.

*Returns a random hipster sentence*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hipster.sentence(4) // 'Xatori'
```

</details>

### Hipster.paragraph()

```ts
paragraph(paragraphCount: int, sentenceCount: int, wordCount: int, separator: string): string;
```

-	`paragraphCount` the number of paragraphs to generate

-	`sentenceCount` the number of sentences to generate

-	`wordCount` the number of words, that should be in the sentence

-	`separator` the paragraph separator

Generates a random hipster paragraphs.

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.hipster.paragraph(1, 2, 4, "\n") // 'Xatori'
```

</details>

Lorem
-----

API to generate random words, sentences, paragraphs, questions and quotes.

### Lorem.paragraph()

```ts
paragraph(paragraphCount: int, sentenceCount: int, wordCount: int, separator: string): string;
```

-	`paragraphCount` the number of paragraphs to generate

-	`sentenceCount` the number of sentences to generate

-	`wordCount` the number of words, that should be in the sentence

-	`separator` the paragraph separator

Generates the given number of paragraphs.

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.lorem.paragraph(1, 2, 4, "\n") // 'It regularly hourly stairs. Stack poorly twist troop.'
```

</details>

### Lorem.sentence()

```ts
sentence(wordCount: int): string;
```

-	`wordCount` the number of words

Generates a space separated list of words beginning with a capital letter and ending with a period.

*Returns a random sentence*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.lorem.sentence(4) // 'It regularly hourly stairs.'
```

</details>

### Lorem.word()

```ts
word(): string;
```

Generates a random word.

*Returns a random word*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.lorem.word() // 'it'
```

</details>

### Lorem.question()

```ts
question(): string;
```

Generates a random question.

*Returns a random question*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.lorem.question() // 'Forage pinterest direct trade pug skateboard food truck flannel cold-pressed?'
```

</details>

### Lorem.quote()

```ts
quote(): string;
```

Generates a random quote from a random person.

*Returns a random quote from a random person*

<details><summary><em>Example</em></summary>

```ts
const faker = new Faker(11)

faker.lorem.quote() // '"Forage pinterest direct trade pug skateboard food truck flannel cold-pressed." - Lukas Ledner'
```

</details>
