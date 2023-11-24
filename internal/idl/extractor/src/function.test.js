/// <reference types="@types/jest" />

import { Kind } from "./kind"
import { map, select } from "./test-helpers"

const script_d_ts = `//js

/**
 * Ea laborum dolore aliqua incididunt ex commodo.
 * 
 * @returns Officia laborum tempor qui velit ipsum excepteur minim irure.
 */
export declare function lorem(): string;

/**
 * Deserunt eu exercitation incididunt mollit esse ad nisi nostrud.
 * 
 * @param enim Fugiat quis ipsum cupidatat amet velit nisi.
 * @param quis Culpa labore excepteur tempor magna sunt.
 * @param consequat Ea anim aliqua elit cupidatat enim eiusmod esse ea enim.
 */
export declare function irure(enim:boolean, quis:string, consequat: number): void;

//!js`

describe("function", () => {
  const decls = select("script.d.ts", script_d_ts, "[?kind=='FUNCTION']")
  const rec = map(decls)

  test("declaration", () => {
    expect(decls).toHaveLength(2)

    expect(rec).toHaveProperty("lorem")
    expect(rec).toHaveProperty("irure")

    let func = rec["lorem"]
    expect(func).toHaveProperty("kind", Kind.Function)
    expect(func).toHaveProperty("type", "string")
    expect(func).toHaveProperty("doc", "Ea laborum dolore aliqua incididunt ex commodo.")
    expect(func).not.toHaveProperty("parameters")
    expect(func.tags).toHaveProperty("returns", ["Officia laborum tempor qui velit ipsum excepteur minim irure."])
    expect(func).not.toHaveProperty("modifiers")

    func = rec["irure"]
    expect(func).toHaveProperty("kind", Kind.Function)
    expect(func).toHaveProperty("type", "void")
    expect(func).toHaveProperty("doc", "Deserunt eu exercitation incididunt mollit esse ad nisi nostrud.")
    expect(func.tags).toHaveProperty("param", [
      "enim Fugiat quis ipsum cupidatat amet velit nisi.",
      "quis Culpa labore excepteur tempor magna sunt.",
      "consequat Ea anim aliqua elit cupidatat enim eiusmod esse ea enim."
    ])
    expect(func).not.toHaveProperty("modifiers")
  })
})
