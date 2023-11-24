import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import faker, { Faker } from "k6/x/faker";
import person from "./person.test.js";

export default function () {
  describe("default", (t) => {
    expect(faker, "default").not.empty;
    expect(faker, "default").to.have.property("person");
    expect(faker, "default").not.to.have.property("faker");
  });

  describe("Faker", (t) => {
    const f = new Faker(11);
    expect(f, "instance").not.empty;
    expect(f, "instance").to.have.property("person");
  });

  person();
}

export const options = { thresholds: { checks: ["rate==1"] } };
