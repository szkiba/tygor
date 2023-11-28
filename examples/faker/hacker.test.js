import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { Faker } from "k6/x/faker";

export default function () {
  describe("hacker", (t) => {
    let faker = new Faker(11);
    expect(faker.hacker.abbreviation(), "abbreviation()").to.equal("GB");
    faker = new Faker(11);
    expect(faker.hacker.adjective(), "adjective()").to.equal("auxiliary");
    faker = new Faker(11);
    expect(faker.hacker.ingverb(), "ingverb()").to.equal("quantifying");
    faker = new Faker(11);
    expect(faker.hacker.noun(), "noun()").to.equal("application");
    faker = new Faker(11);
    expect(faker.hacker.phrase(), "phrase()").to.equal(
      "Try to transpile the EXE sensor, maybe it will deconstruct the wireless interface!"
    );
    faker = new Faker(11);
    expect(faker.hacker.verb(), "verb()").to.equal("read");
  });
}
