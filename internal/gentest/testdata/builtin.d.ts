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

export declare interface Builtin {
  i: int;
  i8: int8;
  i16: int16;
  i32: int32;
  i64: int64;

  u: uint;
  u8: uint8;
  u16: uint16;
  u32: uint32;
  u64: uint64;

  b: byte;
  r: rune;

  f32: float32;
  f64: float64;

  n: number;
  s: string;

  a: any;
  o: object;

  d: Date;
  ab: ArrayBuffer;
}
