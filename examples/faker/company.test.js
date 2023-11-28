import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { Faker } from "k6/x/faker";

export default function () {
  describe("company", (t) => {
    let faker = new Faker(11);
    expect(faker.company.name(), "name()").to.equal("Xatori");
    faker = new Faker(11);
    expect(faker.company.suffix(), "suffix()").to.equal("LLC");
    faker = new Faker(11);
    expect(faker.company.buzzWord(), "buzzWord()").to.equal("Reverse-engineered");
    faker = new Faker(11);
    expect(faker.company.bs(), "bs()").to.equal("24-7");
  });
}
