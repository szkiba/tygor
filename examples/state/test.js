import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import state from "k6/x/state";

export default function () {
  describe("activeVUs", (t) => {
    expect(state, "state").to.have.property("activeVUs");
    expect(state.activeVUs, "state.activeVUs").above(0);
  });

  describe("iteration", (t) => {
    expect(state, "state").to.have.property("iteration");
    expect(state.iteration, "state.iteration").above(-1);
  });

  describe("vuID", (t) => {
    expect(state, "state").to.have.property("vuID");
    expect(state.vuID, "state.vuID").above(0);
  });

  describe("vuIDFromRuntime", (t) => {
    expect(state, "state").to.have.property("vuIDFromRuntime");
    expect(state.vuIDFromRuntime, "state.vuIDFromRuntime").above(0);
  });
}

export const options = { thresholds: { checks: ["rate==1"] } };
