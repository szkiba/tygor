export as namespace hitchhiker;

type int = number;

export declare class Guide {
  question: string;
  readonly answer: int;

  constructor(question: string);
  check(value: int): boolean;
}

declare const defaultGuide: Guide;
export default defaultGuide;

export function check(value: int): boolean;
