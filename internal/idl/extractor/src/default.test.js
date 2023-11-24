/// <reference types="@types/jest" />

import { select } from "./test-helpers"
import { Modifier } from "./modifier"

const script_d_ts = `//js

export as namespace lorem;

export declare class Ipsum {
	fact: string;
}

declare const ipsomium : Ipsum;

/**
 * Duis velit sit laborum do consectetur ex culpa adipisicing.
 */
export default ipsomium;
//!js`

describe("default", () => {
  const decls = select("script.d.ts", script_d_ts, "[?kind=='VARIABLE']")

  test("declaration", () => {
    expect(decls).toHaveLength(1)

    let decl = decls[0]

    expect(decl.name).toEqual("ipsomium")
    expect(decl.modifiers).toContain(Modifier.Default)
    expect(decl.doc).toEqual("Duis velit sit laborum do consectetur ex culpa adipisicing.")
  })
})
