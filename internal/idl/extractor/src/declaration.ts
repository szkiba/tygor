import { Kind } from "./kind.ts"
import { Modifier } from "./modifier.ts"

export interface Declaration {
  name?: string
  type?: string
  doc?: string
  source?: string
  readonly kind: Kind
  tags?: Record<string, string[]>
  modifiers: Array<Modifier>
}

export interface UnknownDeclaration extends Declaration {
  readonly kind: Kind.Class
}

export interface NamespaceDeclaration extends Declaration {
  readonly kind: Kind.Namespace
}

export interface ClassDeclaration extends Declaration {
  readonly kind: Kind.Class
  constructors: Array<ConstructorDeclaration>
  methods: Array<MethodDeclaration>
  properties: Array<PropertyDeclaration>
}

export interface InterfaceDeclaration extends Declaration {
  readonly kind: Kind.Interface
  methods: Array<MethodDeclaration>
  properties: Array<PropertyDeclaration>
}

export interface FunctionDeclaration extends Declaration, HasParameters {
  readonly kind: Kind.Function
}

export interface VariableDeclaration extends Declaration {
  readonly kind: Kind.Variable
}

export interface MethodDeclaration extends Declaration, HasParameters {
  readonly kind: Kind.Method
}

export interface ConstructorDeclaration extends Declaration, HasParameters {
  readonly kind: Kind.Constructor
}

export interface PropertyDeclaration extends Declaration {
  readonly kind: Kind.Property
}

export interface HasParameters extends Declaration {
  parameters: Array<Declaration>
}
