/**
 * xk6-toml enables k6 tests to comfortably encode and decode TOML values.
 *
 * ## Usage
 *
 * Import an entire module's contents:
 * ```JavaScript
 * import * as TOML from "k6/x/toml";
 * ```
 *
 * Import a single export from a module:
 * ```JavaScript
 * import { parse } from "k6/x/toml";
 * ```
 */
export as namespace toml;

/**
 * The parse() method parses a TOML string, constructing the JavaScript object described by the string.
 *
 * @param text The string to parse as TOML
 * @returns The Object corresponding to the given TOML text.
 */
export declare function parse(text: string): object;

/**
 * The stringify() method converts a JavaScript object to a TOML string.
 *
 * @param value The value to convert to a TOML string
 * @returns A TOML string representing the given object
 */
export declare function stringify(value: object): string;
