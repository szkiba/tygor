import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { Faker } from "k6/x/faker";

export default function () {
  describe("person", (t) => {
    let faker = new Faker(11);
    expect(faker.person.firstName(), "firstName()").to.equal("Josiah");
    faker = new Faker(11);
    expect(faker.person.lastName(), "lastName()").to.equal("Abshire");
    faker = new Faker(11);
    expect(faker.person.prefix(), "prefix()").to.equal("Mr.");
    faker = new Faker(11);
    expect(faker.person.suffix(), "suffix()").to.equal("Sr.");
    faker = new Faker(11);
    expect(faker.person.jobTitle(), "jobTitle()").to.equal("Representative");
    faker = new Faker(11);
    expect(faker.person.jobDescriptor(), "jobDescriptor()").to.equal("Internal");
    faker = new Faker(11);
    expect(faker.person.jobLevel(), "jobLevel()").to.equal("Identity");
    faker = new Faker(11);
    expect(faker.person.sexType(), "sexType()").to.equal("male");
  });
}
