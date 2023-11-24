/// <reference types="@types/jest" />

import { preprocess } from "./extractor"

describe("preprocess", () => {
  test("match", () => {
    const src = `//js
    type int = number;
    type int8 = number;
    type int16 = number;
    type int32 = number;
    type int64 = number;
    type uint = number;
    type uint8 = number;
    type uint16 = number;
    type uint32 = number;
    type uint64 = number;
    type byte = number;
    type rune = number;
    type float32 = number;
    type float64 = number;
//!js
`
    const expected = `//js
type int=number & {};
type int8=number & {};
type int16=number & {};
type int32=number & {};
type int64=number & {};
type uint=number & {};
type uint8=number & {};
type uint16=number & {};
type uint32=number & {};
type uint64=number & {};
type byte=number & {};
type rune=number & {};
type float32=number & {};
type float64=number & {};
//!js
`
    expect(preprocess(src)).toEqual(expected)
  })

  test("no match", () => {
    const src = `//js
    type int11 = number;
    type uint3 = number;
    type byte2 = number;
    type rune4 = number;
    type float5 = number;
//!js
`
    const expected = `//js
    type int11 = number;
    type uint3 = number;
    type byte2 = number;
    type rune4 = number;
    type float5 = number;
//!js
`
    expect(preprocess(src)).toEqual(expected)
  })
})
