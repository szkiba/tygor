import { extract } from "././extractor"
import { Declaration } from "./declaration"
import { Kind } from "./kind"

import jmespath from "jmespath"

export function map(decs: Declaration[], kind?: Kind): Record<string, Declaration> {
  return decs
    .filter((dec) => dec.name && (!kind || dec.kind == kind))
    .map((dec) => ({ [dec.name!]: dec }))
    .reduce((prev, current) => {
      return { ...prev, ...current }
    }, {} as Record<string, Declaration>)
}

export function select(name: string, source: string, query?: string): Declaration | Declaration[] | undefined {
  return search(parse(name, source), query ?? "[0]")
}

export function parse(name: string, source: string): Declaration[] {
  return JSON.parse(extract(name, source))
}

export function search(declarations: Declaration[], query: string): Declaration | Declaration[] | undefined {
  let res = jmespath.search(declarations, query)

  if (!res) {
    return undefined
  }

  if (Array.isArray(res)) {
    return res as Declaration[]
  }

  return res as Declaration
}
