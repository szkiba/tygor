/**
 * xk6-faker random fake data generator
 */
export as namespace faker;

type int64 = number;
type int = number;

/**
 * This is Faker's main class containing all modules that can be used to generate data.
 *
 * Please have a look at the individual modules and methods for more information and examples.
 *
 * @example
 * ```ts
 * import { Faker } from "k6/x/faker"
 *
 * const faker = new Faker(11)
 *
 * export default function() {
 *   console.log(faker.person.firstName()) // 'Josiah'
 * }
 * ```
 */
export declare class Faker {
  /**
   * Creates a new instance of Faker.
   *
   * Optionally, the value of the random seed can be set as a constructor parameter.
   * This is intended to allow for consistent values in a tests,
   * so you might want to use hardcoded values as the seed.
   *
   * Please note that generated values are dependent on both the seed and the number
   * of calls that have been made.
   *
   * Setting seed to 0 (or omitting it) will use seed derived from system entropy.
   *
   * @param seed random seed value for deterministic generator
   *
   * @example
   * ```ts
   * const consistentFaker = new Faker(11)
   * const semiRandomFaker = new Faker()
   * ```
   */
  constructor(seed?: int64);

  /**
   * API to generate people's personal information such as names and job titles.
   */
  readonly person: Person;

  /**
   * API to generate company related entries.
   */
  readonly company: Company;

  /**
   * API to generate hacker/IT words and phrases.
   */
  readonly hacker: Hacker;

  /**
   * API to generate hipster words, phrases and paragraphs.
   */
  readonly hipster: Hipster;

  /**
   * API to generate random words, sentences, paragraphs, questions and quotes.
   */
  readonly lorem: Lorem;
}

/**
 * API to generate people's personal information such as names and job titles.
 */
export declare interface Person {
  /**
   * Generates a random first name.
   *
   * @returns a random first name
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.firstName() // 'Josiah'
   * ```
   */
  firstName(): string;

  /**
   * Generates a random last name.
   *
   * @returns a random last name
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.lastName() // 'Abshire'
   * ```
   */
  lastName(): string;

  /**
   * Generates a random person prefix.
   *
   * @returns a random person prefix.
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.prefix() // 'Mr.'
   * ```
   */
  prefix(): string;

  /**
   * Generates a random person suffix.
   *
   * @returns  a random person suffix
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.suffix() // 'Sr.'
   * ```
   */
  suffix(): string;

  /**
   * Generates a random sex type.
   *
   * @returns a random sex type `male`|`female`
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.sexType() // 'male'
   * ```
   */
  sexType(): string;

  /**
   * Generates a random job title.
   *
   * @returns a random job title
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.jobTitle() // 'Representative'
   * ```
   */
  jobTitle(): string;

  /**
   * Generates a random job level.
   *
   * @returns a random job level
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.jobLevel() // 'Identity'
   * ```
   */
  jobLevel(): string;

  /**
   * Generates a random job descriptor.
   *
   * @returns a random job descriptor
   *
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.person.jobDescriptor() // 'Internal'
   * ```
   */
  jobDescriptor(): string;
}

/**
 * API to generate company related entries.
 */
export declare interface Company {
  /**
   * Generates a random company name string.
   *
   * @returns a random company name string
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.company.name() // 'Xatori'
   * ```
   */
  name(): string;

  /**
   * Generates a random company suffix string.
   *
   * @returns a random company suffix string
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.company.suffix() // 'LLC'
   * ```
   */
  suffix(): string;

  /**
   * Generates a random company buzz word string.
   *
   * @returns a random company buzz word string
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.company.buzzWord() // 'Reverse-engineered'
   * ```
   */
  buzzWord(): string;

  /**
   * Generates a random company bs string.
   *
   * @returns a random company bs string
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.company.bs() // '24-7'
   * ```
   */
  bs(): string;
}

/**
 * API to generate hacker/IT words and phrases.
 */
export declare interface Hacker {
  /**
   * Generates a random hacker/IT abbreviation.
   *
   * @returns a random hacker/IT abbreviation
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.abbreviation() // 'Xatori'
   * ```
   */
  abbreviation(): string;

  /**
   * Generates a random hacker/IT adjective.
   *
   * @returns a random hacker/IT adjective
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.adjective() // 'Xatori'
   * ```
   */
  adjective(): string;

  /**
   * Generates a random hacker/IT verb for continuous actions (en: ing suffix; e.g. hacking).
   *
   * @returns a random hacker/IT verb for continuous actions
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.ingverb() // 'Xatori'
   * ```
   */
  ingverb(): string;

  /**
   * Generates a random hacker/IT noun.
   *
   * @returns a random hacker/IT noun
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.noun() // 'Xatori'
   * ```
   */
  noun(): string;

  /**
   * Generates a random hacker/IT phrase.
   *
   * @returns a random hacker/IT phrase
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.phrase() // 'Xatori'
   * ```
   */
  phrase(): string;

  /**
   * Generates a random hacker/IT verb.
   *
   * @returns a random hacker/IT verb
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hacker.verb() // 'Xatori'
   * ```
   */
  verb(): string;
}

/**
 * API to generate hipster words, phrases and paragraphs.
 */
export declare interface Hipster {
  /**
   * Generates a single hipster word.
   *
   * @returns a single hipster word
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hipster.word() // 'Xatori'
   * ```
   */
  word(): string;

  /**
   * Generates a random hipster sentence.
   *
   * @returns a random hipster sentence
   * @param wordCount the number of words
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hipster.sentence(4) // 'Xatori'
   * ```
   */
  sentence(wordCount: int): string;

  /**
   * Generates a random hipster paragraphs.
   *
   * @param paragraphCount the number of paragraphs to generate
   * @param sentenceCount the number of sentences to generate
   * @param wordCount the number of words, that should be in the sentence
   * @param separator the paragraph separator
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.hipster.paragraph(1, 2, 4, "\n") // 'Xatori'
   * ```
   */
  paragraph(paragraphCount: int, sentenceCount: int, wordCount: int, separator: string): string;
}

/**
 * API to generate random words, sentences, paragraphs, questions and quotes.
 */
export declare interface Lorem {
  /**
   * Generates the given number of paragraphs.
   *
   * @param paragraphCount the number of paragraphs to generate
   * @param sentenceCount the number of sentences to generate
   * @param wordCount the number of words, that should be in the sentence
   * @param separator the paragraph separator
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.lorem.paragraph(1, 2, 4, "\n") // 'It regularly hourly stairs. Stack poorly twist troop.'
   * ```
   */
  paragraph(paragraphCount: int, sentenceCount: int, wordCount: int, separator: string): string;

  /**
   * Generates a space separated list of words beginning with a capital letter and ending with a period.
   *
   * @param wordCount the number of words
   * @returns a random sentence
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.lorem.sentence(4) // 'It regularly hourly stairs.'
   * ```
   */
  sentence(wordCount: int): string;

  /**
   * Generates a random word.
   *
   * @returns a random word
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.lorem.word() // 'it'
   * ```
   */
  word(): string;

  /**
   * Generates a random question.
   *
   * @returns a random question
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.lorem.question() // 'Forage pinterest direct trade pug skateboard food truck flannel cold-pressed?'
   * ```
   */
  question(): string;

  /**
   * Generates a random quote from a random person.
   *
   * @returns a random quote from a random person
   * @example
   * ```ts
   * const faker = new Faker(11)
   *
   * faker.lorem.quote() // '"Forage pinterest direct trade pug skateboard food truck flannel cold-pressed." - Lukas Ledner'
   * ```
   */
  quote(): string;
}

/** Default Faker instance. */
declare const faker: Faker;

/** Default Faker instance */
export default faker;
