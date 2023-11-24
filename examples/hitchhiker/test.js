import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import guide, { Guide } from "k6/x/hitchhiker";

export default function () {
  describe("default", () => {
    expect(guide).to.have.property("answer", 42);
    expect(guide).to.have.property("question", "What's up?");
    expect(guide.check(42)).to.be.true;
    expect(guide.check(43)).to.be.false;
    expect(() => (guide.answer = 2)).to.throw(TypeError);
    guide.question = "Why are we here?";
    expect(guide).to.have.property("question", "Why are we here?");
  });

  describe("Guide", () => {
    const guide = new Guide("What is life all about?");
    expect(guide).to.have.property("answer", 42);
    expect(guide).to.have.property("question", "What is life all about?");
    expect(guide.check(42)).to.be.true;
    expect(guide.check(43)).to.be.false;
    expect(() => (guide.answer = 2)).to.throw(TypeError);
    guide.question = "Why are we here?";
    expect(guide).to.have.property("question", "Why are we here?");
  });
}

export const options = { thresholds: { checks: ["rate==1"] } };
