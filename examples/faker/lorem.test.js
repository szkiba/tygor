import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { Faker } from "k6/x/faker";

export default function () {
  describe("lorem", (t) => {
    let faker = new Faker(11);
    expect(faker.lorem.paragraph(1, 2, 4, "\n"), "paragraph()").to.equal(
      "It regularly hourly stairs. Stack poorly twist troop."
    );
    faker = new Faker(11);
    expect(faker.lorem.sentence(4), "sentence()").to.equal("It regularly hourly stairs.");
    faker = new Faker(11);
    expect(faker.lorem.word(), "word()").to.equal("it");
    faker = new Faker(11);
    expect(faker.lorem.question(), "question()").to.equal(
      "Forage pinterest direct trade pug skateboard food truck flannel cold-pressed?"
    );
    faker = new Faker(11);
    expect(faker.lorem.quote(), "quote()").to.equal(
      '"Forage pinterest direct trade pug skateboard food truck flannel cold-pressed." - Lukas Ledner'
    );
  });
}
