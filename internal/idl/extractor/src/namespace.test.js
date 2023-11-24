/// <reference types="@types/jest" />

import { Kind } from "./kind"
import { map, select } from "./test-helpers"

const script_d_ts = `//js

/**
 * Laborum cupidatat elit ut amet.
 */

/**
 * Nulla sunt enim culpa do irure.
 * 
 * @author Incididunt Aliquip
 */
export as namespace lorem;

//!js`

const no_doc_script_d_ts = `export as namespace lorem;`

describe("namespace", () => {
  const decls = select("script.d.ts", script_d_ts, "[?kind=='NAMESPACE']")

  const rec = map(decls)

  test("declaration", () => {
    expect(decls).toHaveLength(1)

    expect(rec).toHaveProperty("lorem")

    let ns = rec["lorem"]

    expect(ns).toBeTruthy
    expect(ns).toHaveProperty("kind", Kind.Namespace)
    expect(ns).toHaveProperty("doc", "Laborum cupidatat elit ut amet.\nNulla sunt enim culpa do irure.")
    expect(ns.tags).toHaveProperty("author", ["Incididunt Aliquip"])
  })

  test("no doc", () => {
    const nodoc = select("no_doc_script.d.ts", no_doc_script_d_ts, "[?kind=='NAMESPACE']")

    expect(nodoc).toHaveLength(1)
  })
})
