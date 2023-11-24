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

Module to generate people's personal information such as names and job titles.

Person
------

Module to generate people's personal information such as names and job titles.

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
