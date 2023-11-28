import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import faker, { Faker } from "k6/x/faker";
import person from "./person.test.js";
import company from "./company.test.js";
import lorem from "./lorem.test.js";
import hacker from "./hacker.test.js";
import hipster from "./hipster.test.js";

export default function () {
  describe("default", (t) => {
    expect(faker, "default").not.empty;
    expect(faker, "default").to.have.property("person");
    expect(faker, "default").to.have.property("company");
    expect(faker, "default").to.have.property("lorem");
    expect(faker, "default").to.have.property("hacker");
    expect(faker, "default").to.have.property("hipster");
    expect(faker, "default").not.to.have.property("faker");
  });

  describe("Faker", (t) => {
    const f = new Faker();
    expect(f, "instance").not.empty;
    expect(f, "instance").to.have.property("person");
    expect(f, "instance").to.have.property("company");
    expect(f, "instance").to.have.property("lorem");
    expect(f, "instance").to.have.property("hacker");
    expect(f, "instance").to.have.property("hipster");
  });

  person();
  company();
  lorem();
  hacker();
  hipster();
}

export const options = { thresholds: { checks: ["rate==1"] } };
