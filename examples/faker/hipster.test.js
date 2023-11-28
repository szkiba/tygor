import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { Faker } from "k6/x/faker";

export default function () {
  describe("hipster", (t) => {
    let faker = new Faker(11);
    expect(faker.hipster.word(), "word()").to.equal("offal");
    faker = new Faker(11);
    expect(faker.hipster.sentence(4), "sentence()").to.equal("Offal forage pinterest direct trade.");
    faker = new Faker(11);
    expect(faker.hipster.paragraph(1, 2, 4, "\n"), "paragraph()").to.equal(
      "Offal forage pinterest direct trade. Pug skateboard food truck flannel."
    );
  });
}
