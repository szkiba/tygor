/**
 * xk6-faker random fake data generator
 */
export as namespace faker;

type int64 = number;

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
   * Module to generate people's personal information such as names and job titles.
   */
  readonly person: Person;
}

/**
 * Module to generate people's personal information such as names and job titles.
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

/** Default Faker instance. */
declare const faker: Faker;

/** Default Faker instance */
export default faker;
