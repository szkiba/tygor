/// <reference types="@types/jest" />

import { Kind } from "./kind"
import { Modifier } from "./modifier"
import { map, select } from "./test-helpers"

const script_d_ts = `//js

/**
 * Ea laborum dolore aliqua incididunt ex commodo.
 */
export declare const laborum: string;

/**
 * Occaecat elit consectetur nisi sint in est aliquip sint.
 */
export declare var dolore: number;

/**
 * Officia nulla excepteur in ea nostrud duis elit.
 */
export declare var sunt: string[];

/**
 * Officia laboris eu ullamco laboris.
 */
export declare var quis: Array<string>;

/**
 * Ad do incididunt tempor occaecat ex velit commodo.
 */
export declare var amet: Date;

//!js`

describe("variable", () => {
  const decls = select("script.d.ts", script_d_ts, "[?kind=='VARIABLE']")
  const rec = map(decls)

  test("declaration", () => {
    expect(decls).toHaveLength(5)

    expect(rec).toHaveProperty("laborum")
    expect(rec).toHaveProperty("dolore")
    expect(rec).toHaveProperty("sunt")
    expect(rec).toHaveProperty("quis")
    expect(rec).toHaveProperty("amet")

    let dec = rec["laborum"]
    expect(dec).toBeTruthy
    expect(dec).toHaveProperty("kind", Kind.Variable)
    expect(dec).toHaveProperty("type", "string")
    expect(dec).toHaveProperty("modifiers", [Modifier.Const])
    expect(dec).toHaveProperty("doc", "Ea laborum dolore aliqua incididunt ex commodo.")

    dec = rec["dolore"]
    expect(dec).toBeTruthy
    expect(dec).toHaveProperty("kind", Kind.Variable)
    expect(dec).toHaveProperty("type", "number")
    expect(dec).not.toHaveProperty("modifiers")
    expect(dec).toHaveProperty("doc", "Occaecat elit consectetur nisi sint in est aliquip sint.")

    dec = rec["sunt"]
    expect(dec).toBeTruthy
    expect(dec).toHaveProperty("kind", Kind.Variable)
    // TODO   expect(dec).toHaveProperty("type", "Array<string>")
    expect(dec).not.toHaveProperty("modifiers")
    expect(dec).toHaveProperty("doc", "Officia nulla excepteur in ea nostrud duis elit.")

    dec = rec["quis"]
    expect(dec).toBeTruthy
    expect(dec).toHaveProperty("kind", Kind.Variable)
    expect(dec).toHaveProperty("type", "Array<string>")
    expect(dec).not.toHaveProperty("modifiers")
    expect(dec).toHaveProperty("doc", "Officia laboris eu ullamco laboris.")

    dec = rec["amet"]
    expect(dec).toBeTruthy
    expect(dec).toHaveProperty("kind", Kind.Variable)
    expect(dec).toHaveProperty("type", "Date")
    expect(dec).not.toHaveProperty("modifiers")
    expect(dec).toHaveProperty("doc", "Ad do incididunt tempor occaecat ex velit commodo.")
  })
})
