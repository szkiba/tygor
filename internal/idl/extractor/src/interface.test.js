/// <reference types="@types/jest" />

import { Kind } from "./kind"
import { Modifier } from "./modifier"
import { select, map, search } from "./test-helpers"

const script_d_ts = `//js

/**
 * Exercitation non duis qui ad.
 * 
 * Ipsum dolore nostrud deserunt proident non.
 * @category consectetur
 */
export declare interface Ipsum {
  /** Id eiusmod tempor sunt est. */
  readonly cupidatat: string;

  /** Nisi est veniam eiusmod exercitation quis. */
  laboris: string;

  /** Excepteur tempor adipisicing occaecat incididunt. */
  velit: number;

  /** Duis duis proident fugiat sunt non duis non mollit nulla tempor et.
   * 
   * Culpa et quis incididunt tempor amet.
   * 
   * @returns Velit in velit laborum ad quis esse eiusmod ea et duis.
   */
  incididunt() : number;

  /** Amet pariatur dolor minim velit ea deserunt dolor.
   * 
   * Ex tempor culpa sit Lorem quis.
   * 
   * @param est Nostrud mollit commodo exercitation minim in et.
   * In sunt aliquip ipsum magna voluptate.
   * @param ad Do fugiat laborum elit non velit ea ea nisi tempor.
   * @returns Ullamco ad et sunt aute cillum culpa aute tempor.
   */
  excepteur(est: number, ad?: string) : string;
}

//!js`

describe("interface", () => {
  const decl = select("script.d.ts", script_d_ts, "[?kind=='INTERFACE']|[0]")

  test("declaration", () => {
    expect(decl).toHaveProperty("name", "Ipsum")
    expect(decl).toHaveProperty("kind", Kind.Interface)
    expect(decl).toHaveProperty("doc", "Exercitation non duis qui ad.\n\nIpsum dolore nostrud deserunt proident non.")
    expect(decl.tags).toHaveProperty("category", ["consectetur"])
  })

  test("properties", () => {
    const dict = map(decl.properties)

    expect(dict).toHaveProperty("cupidatat")
    expect(dict).toHaveProperty("laboris")
    expect(dict).toHaveProperty("velit")

    let prop = dict["cupidatat"]
    expect(prop).toHaveProperty("kind", Kind.Property)
    expect(prop).toHaveProperty("type", "string")
    expect(prop).toHaveProperty("modifiers", [Modifier.Readonly])
    expect(prop).toHaveProperty("doc", "Id eiusmod tempor sunt est.")

    prop = dict["laboris"]
    expect(prop).toHaveProperty("kind", Kind.Property)
    expect(prop).toHaveProperty("type", "string")
    expect(prop).not.toHaveProperty("modifiers")
    expect(prop).toHaveProperty("doc", "Nisi est veniam eiusmod exercitation quis.")

    prop = dict["velit"]
    expect(prop).toHaveProperty("kind", Kind.Property)
    expect(prop).toHaveProperty("type", "number")
    expect(prop).not.toHaveProperty("modifiers")
    expect(prop).toHaveProperty("doc", "Excepteur tempor adipisicing occaecat incididunt.")
  })

  test("methods", () => {
    const dict = map(decl.methods)

    expect(dict).toHaveProperty("incididunt")
    expect(dict).toHaveProperty("excepteur")

    let method = dict["incididunt"]
    expect(method).toHaveProperty("kind", Kind.Method)
    expect(method).toHaveProperty("type", "number")
    expect(method).not.toHaveProperty("parameters")
    expect(method).toHaveProperty(
      "doc",
      "Duis duis proident fugiat sunt non duis non mollit nulla tempor et.\n\nCulpa et quis incididunt tempor amet."
    )
    expect(method.tags).toHaveProperty("returns", ["Velit in velit laborum ad quis esse eiusmod ea et duis."])

    method = dict["excepteur"]
    expect(method).toHaveProperty("kind", Kind.Method)
    expect(method).toHaveProperty("type", "string")
    expect(method).toHaveProperty("parameters")
    expect(method).toHaveProperty(
      "doc",
      "Amet pariatur dolor minim velit ea deserunt dolor.\n\nEx tempor culpa sit Lorem quis."
    )
    expect(method.tags).toHaveProperty("returns", ["Ullamco ad et sunt aute cillum culpa aute tempor."])

    let param = search(method.parameters, "[?name=='est']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "number")
    expect(param).toHaveProperty(
      "doc",
      "Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate."
    )
    expect(param.tags).toHaveProperty("param", [
      "est Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate."
    ])

    param = search(method.parameters, "[?name=='ad']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "string")
    expect(param).toHaveProperty("doc", "Do fugiat laborum elit non velit ea ea nisi tempor.")
    expect(param.tags).toHaveProperty("param", ["ad Do fugiat laborum elit non velit ea ea nisi tempor."])
  })
})
