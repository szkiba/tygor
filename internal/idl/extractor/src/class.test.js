/// <reference types="@types/jest" />

import { Kind } from "./kind"
import { Modifier } from "./modifier"
import { select, map, search } from "./test-helpers"

const script_d_ts = `//js

/**
 * Exercitation non duis qui ad.
 * 
 * Lorem dolore nostrud deserunt proident non.
 * 
 * @category consectetur
 */
export declare class Lorem {
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

  /**
   * Dolor Lorem eu aliquip fugiat non labore nostrud.
   * 
   * @param non Mollit minim nisi ea culpa incididunt adipisicing.
   * @param occaecat Excepteur velit est Lorem voluptate consequat sit cillum quis.
   * @param anim Excepteur do do nisi aliqua velit laborum.
   */
  constructor(non: string, occaecat: boolean, anim: number);

  /**
   * Nisi tempor amet culpa aliquip dolor quis ea id.
   * 
   * @param mollit Consequat culpa nostrud eiusmod ut ipsum elit tempor.
   * @param magna Do id deserunt sunt ad ut ipsum eu et qui veniam.
   * @param officia Enim irure id culpa amet est velit.
   */
  constructor(mollit: number, magna: string, officia: boolean);
}

//!js`

describe("class", () => {
  const decl = select("script.d.ts", script_d_ts, "[?kind=='CLASS']|[0]")

  test("declaration", () => {
    expect(decl).toHaveProperty("name", "Lorem")
    expect(decl).toHaveProperty("kind", Kind.Class)
    expect(decl).toHaveProperty("doc", "Exercitation non duis qui ad.\n\nLorem dolore nostrud deserunt proident non.")
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

  test("constructors", () => {
    expect(decl.constructors).toHaveLength(2)

    let ctor = search(decl.constructors, "[?doc=='Dolor Lorem eu aliquip fugiat non labore nostrud.']|[0]")
    expect(ctor).toBeTruthy

    let param = search(ctor.parameters, "[?name=='non']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "string")
    expect(param).toHaveProperty("doc", "Mollit minim nisi ea culpa incididunt adipisicing.")
    expect(param.tags).toHaveProperty("param", ["non Mollit minim nisi ea culpa incididunt adipisicing."])

    param = search(ctor.parameters, "[?name=='occaecat']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "boolean")
    expect(param).toHaveProperty("doc", "Excepteur velit est Lorem voluptate consequat sit cillum quis.")
    expect(param.tags).toHaveProperty("param", ["occaecat Excepteur velit est Lorem voluptate consequat sit cillum quis."])

    param = search(ctor.parameters, "[?name=='anim']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "number")
    expect(param).toHaveProperty("doc", "Excepteur do do nisi aliqua velit laborum.")
    expect(param.tags).toHaveProperty("param", ["anim Excepteur do do nisi aliqua velit laborum."])

    ctor = search(decl.constructors, "[?doc=='Nisi tempor amet culpa aliquip dolor quis ea id.']|[0]")
    expect(ctor).toBeTruthy

    param = search(ctor.parameters, "[?name=='mollit']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "number")
    expect(param).toHaveProperty("doc", "Consequat culpa nostrud eiusmod ut ipsum elit tempor.")
    expect(param.tags).toHaveProperty("param", ["mollit Consequat culpa nostrud eiusmod ut ipsum elit tempor."])

    param = search(ctor.parameters, "[?name=='magna']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "string")
    expect(param).toHaveProperty("doc", "Do id deserunt sunt ad ut ipsum eu et qui veniam.")
    expect(param.tags).toHaveProperty("param", ["magna Do id deserunt sunt ad ut ipsum eu et qui veniam."])

    param = search(ctor.parameters, "[?name=='officia']|[0]")
    expect(param).toBeTruthy
    expect(param).toHaveProperty("kind", Kind.Parameter)
    expect(param).toHaveProperty("type", "boolean")
    expect(param).toHaveProperty("doc", "Enim irure id culpa amet est velit.")
    expect(param.tags).toHaveProperty("param", ["officia Enim irure id culpa amet est velit."])
  })
})
