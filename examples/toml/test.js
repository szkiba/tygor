import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import { parse, stringify } from "k6/x/toml";

const TOML = { parse, stringify };

const object = `
foo = "bar"
sizes = ["small", "medium", "large"]
`;

export default function () {
  describe("parse", (t) => {
    const obj = TOML.parse(object);

    expect(obj, "parsed").to.have.property("foo", "bar");
    expect(obj.sizes[0], "sizes[0]").to.be.equal("small");
  });

  describe("stringify", (t) => {
    const obj = { sizes: ["small", "medium", "large"] };
    const str = TOML.stringify(obj);

    expect(str, "string").to.be.equal(`sizes = ["small", "medium", "large"]\n`);
  });
}

export const options = { thresholds: { checks: ["rate==1"] } };
