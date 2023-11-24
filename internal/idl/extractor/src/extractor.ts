import ts from "typescript"
import {
  Declaration,
  VariableDeclaration,
  UnknownDeclaration,
  ClassDeclaration,
  FunctionDeclaration,
  InterfaceDeclaration,
  MethodDeclaration,
  ConstructorDeclaration,
  PropertyDeclaration,
  HasParameters,
  NamespaceDeclaration
} from "./declaration.ts"
import { Kind } from "./kind.ts"
import { Modifier } from "./modifier.ts"

export function extract(filename: string, source: string): string {
  return new Extractor(filename, preprocess(source)).extract()
}

const scriptTarget = ts.ScriptTarget.ES5

class Extractor {
  sourceFile: ts.SourceFile
  program: ts.Program
  checker: ts.TypeChecker
  output: Declaration[]
  defaultExport?: ts.ExportAssignment

  constructor(filename: string, source: string) {
    this.sourceFile = ts.createSourceFile(filename, source, scriptTarget, true, ts.ScriptKind.TS)
    this.program = createProgram(this.sourceFile)
    this.checker = this.program.getTypeChecker()
    this.output = []
    this.defaultExport = undefined
  }

  extract(): string {
    ts.forEachChild(this.sourceFile, (child) => this.visit(child))

    if (this.defaultExport) this.defaultExportHandler(this.defaultExport)

    return JSON.stringify(this.output, undefined, 2)
  }

  defaultExportHandler(node: ts.ExportAssignment): void {
    const name = node.expression.getText(this.sourceFile)

    this.output
      .filter((e) => e.name == name)
      .forEach((decl) => {
        if (!decl.modifiers) decl.modifiers = []
        decl.modifiers.push(Modifier.Default)

        if (!decl.doc) {
          this.getDocs(node, decl)
        }
      })
  }

  visit(node: ts.Node): void {
    if (ts.isNamespaceExportDeclaration(node)) this.namespaceHandler(node)
    else if (ts.isClassDeclaration(node)) this.classHandler(node)
    else if (ts.isInterfaceDeclaration(node)) this.interfaceHandler(node)
    else if (ts.isVariableStatement(node)) this.variableListHandler(node)
    else if (ts.isFunctionDeclaration(node)) this.functionHandler(node)
    else if (ts.isNamespaceExportDeclaration(node)) this.namespaceHandler(node)
    else if (ts.isExportAssignment(node)) this.exportHandler(node)
    else if (ts.isModuleDeclaration(node)) ts.forEachChild(node, (child) => this.visit(child))
    else if (ts.isTypeAliasDeclaration(node)) undefined // noop
    else if (node.kind != ts.SyntaxKind.EndOfFileToken) this.unknownHandler(node)
  }

  push(dec: Declaration) {
    this.output.push(dec)
  }

  serializeParameters(node: ts.SignatureDeclaration): Declaration[] {
    let sig = this.checker.getSignatureFromDeclaration(node)

    if (!sig) {
      return []
    }

    const decs = [] as Declaration[]

    const params = node.parameters

    params.forEach((param) => {
      let dec = this.serializeSymbol(param, Kind.Parameter)
      if (dec.kind && dec.kind != Kind.Unknown) {
        decs.push(dec)
      }
    })

    return decs
  }

  serializeConstructor(member: ts.ConstructorDeclaration): ConstructorDeclaration {
    const sig = this.checker.getSignatureFromDeclaration(member)

    const dec = {
      kind: Kind.Constructor,
      doc: this.serializeDoc(sig!),
      source: member.getText(this.sourceFile),
      tags: this.serializeTags(sig!)
    } as ConstructorDeclaration

    let params = this.serializeParameters(member)
    if (params.length != 0) {
      dec.parameters = params
    }

    return dec
  }

  serializeDoc(from: ts.Symbol | ts.Signature): string | undefined {
    const parts = from.getDocumentationComment(this.checker)

    if (parts.length == 0) {
      return undefined
    }

    return ts.displayPartsToString(parts)
  }

  serializeTags(from: ts.Symbol | ts.Signature): Record<string, string[]> | undefined {
    const tags = from.getJsDocTags(this.checker)

    if (tags.length == 0) {
      return undefined
    }

    let rec = {} as Record<string, string[]>

    tags.forEach((taginfo) => {
      if (!rec[taginfo.name]) {
        rec[taginfo.name] = []
      }

      rec[taginfo.name].push(ts.displayPartsToString(taginfo.text))
    })

    return rec
  }

  serializeSymbol(node: ts.NamedDeclaration | undefined, kind: Kind): Declaration {
    if (!node || !node.name) return {} as UnknownDeclaration

    const symbol = this.checker.getSymbolAtLocation(node.name)
    if (!symbol) return {} as UnknownDeclaration

    const dec = {
      name: symbol.getName(),
      kind: kind,
      doc: this.serializeDoc(symbol),
      tags: this.serializeTags(symbol)
    } as Declaration

    if (!ts.isClassDeclaration(node) && !ts.isInterfaceDeclaration(node)) {
      dec.source = node.getText(this.sourceFile)
    }

    let type: ts.Type

    if (ts.canHaveModifiers(node)) {
      const mods = ts.getCombinedModifierFlags(node)
      let modifiers = []

      if (mods & ts.ModifierFlags.Readonly) {
        modifiers.push(Modifier.Readonly)
      }

      if (mods & ts.ModifierFlags.Const) {
        modifiers.push(Modifier.Const)
      }

      if (mods & ts.ModifierFlags.Async) {
        modifiers.push(Modifier.Async)
      }

      if (modifiers.length != 0) {
        dec.modifiers = modifiers
      }
    }

    if (ts.isMethodSignature(node) || ts.isFunctionDeclaration(node) || ts.isMethodDeclaration(node)) {
      type = this.checker.getReturnTypeOfSignature(this.checker.getSignatureFromDeclaration(node)!)
      let pars = this.serializeParameters(node)
      if (pars.length != 0) {
        ;(dec as HasParameters).parameters = pars
      }
    } else {
      type = this.checker.getTypeAtLocation(node.name)
    }

    dec.type = this.checker.typeToString(type)

    return dec
  }

  getDocs(node: ts.Node, decl: Declaration): void {
    const dnode = node as any as { jsDoc: ts.JSDoc[] }

    let concat = (buff: string, comment?: string | ts.JSDocComment | ts.NodeArray<ts.JSDocComment>): string => {
      if (!comment) return buff
      if (buff.length > 0) buff += "\n"
      if (typeof comment == "string") {
        buff += comment
      } else if (Array.isArray(comment)) {
        comment.forEach((c) => (buff += concat(buff, c)))
      } else {
        buff += (comment as ts.JSDocComment).text
      }
      return buff
    }

    let tags: Record<string, string[]> = {}

    let docs: string = ""

    if (dnode.jsDoc) {
      dnode.jsDoc.forEach((doc) => {
        if (docs.length > 0) docs += "\n"
        docs += concat("", doc.comment)

        if (!doc.tags) {
          return
        }

        doc.tags!.forEach((tag) => {
          let arr = tags[tag.tagName.text]
          if (!arr) {
            arr = [] as string[]
            tags[tag.tagName.text] = arr
          }

          arr.push(concat("", tag.comment))
        })
      })
    }

    if (docs.length != 0) {
      decl.doc = docs
    }

    if (Object.keys(tags).length > 0) {
      decl.tags = tags
    }
  }

  namespaceHandler(node: ts.NamespaceExportDeclaration): void {
    const decl = { name: node.name.getText(this.sourceFile), kind: Kind.Namespace } as NamespaceDeclaration

    this.getDocs(node, decl)

    this.push(decl)
  }

  exportHandler(node: ts.ExportAssignment): void {
    this.defaultExport = node
  }

  variableListHandler(node: ts.VariableStatement): void {
    node.declarationList.forEachChild((child) => {
      if (ts.isVariableDeclaration(child)) {
        this.variableHandler(child)
      }
    })
  }

  variableHandler(node: ts.VariableDeclaration): void {
    const common = this.serializeSymbol(node, Kind.Variable)
    if (common.kind == Kind.Unknown) {
      return
    }

    const dec = { ...common } as VariableDeclaration

    if (node.parent.flags & (ts.NodeFlags.Constant | ts.NodeFlags.Const)) {
      dec.modifiers = [Modifier.Const]
      dec.source = "export declare const " + dec.source + ";"
    } else {
      dec.source = "export declare var " + dec.source + ";"
    }

    this.push(dec)
  }

  functionHandler(node: ts.FunctionDeclaration): void {
    const common = this.serializeSymbol(node, Kind.Function)
    if (common.kind == Kind.Unknown) {
      return
    }

    const dec = { ...common } as FunctionDeclaration

    this.push(dec)
  }

  classHandler(node: ts.ClassDeclaration): void {
    const common = this.serializeSymbol(node, Kind.Class)
    if (common.kind == Kind.Unknown) {
      return
    }

    const dec = { ...common } as ClassDeclaration

    let methods = [] as Array<MethodDeclaration>
    let props = [] as Array<PropertyDeclaration>
    let ctors = [] as Array<ConstructorDeclaration>

    node.members.forEach((member) => {
      if (ts.isMethodDeclaration(member)) {
        const mem = this.serializeSymbol(member, Kind.Method)
        if (mem.kind != Kind.Unknown) {
          methods.push(mem as MethodDeclaration)
        }
      } else if (ts.isPropertyDeclaration(member)) {
        const mem = this.serializeSymbol(member, Kind.Property)
        if (mem.kind != Kind.Unknown) {
          props.push(mem as PropertyDeclaration)
        }
      } else if (ts.isConstructorDeclaration(member)) {
        const mem = this.serializeConstructor(member)
        ctors.push(mem as ConstructorDeclaration)
      }
    })

    if (methods.length != 0) {
      dec.methods = methods
    }

    if (props.length != 0) {
      dec.properties = props
    }

    if (ctors.length != 0) {
      dec.constructors = ctors
    }

    this.push(dec)
  }

  interfaceHandler(node: ts.InterfaceDeclaration): void {
    const common = this.serializeSymbol(node, Kind.Interface)
    if (common.kind == Kind.Unknown) {
      return
    }

    const dec = { ...common } as InterfaceDeclaration

    let methods = [] as Array<MethodDeclaration>
    let props = [] as Array<PropertyDeclaration>

    node.members.forEach((member) => {
      if (ts.isMethodSignature(member)) {
        const mem = this.serializeSymbol(member, Kind.Method)
        if (mem.kind != Kind.Unknown) {
          methods.push(mem as MethodDeclaration)
        }
      } else if (ts.isPropertySignature(member)) {
        const mem = this.serializeSymbol(member, Kind.Property)
        if (mem.kind != Kind.Unknown) {
          props.push(mem as PropertyDeclaration)
        }
      }
    })

    if (methods.length != 0) {
      dec.methods = methods
    }

    if (props.length != 0) {
      dec.properties = props
    }

    this.push(dec)
  }

  unknownHandler(node: ts.Node): void {
    this.push({ name: node.kind.toString(), kind: Kind.Unknown, modifiers: [] })
  }
}

function createProgram(sourceFile: ts.SourceFile): ts.Program {
  let options: ts.CompilerOptions = {
    target: scriptTarget,
    module: ts.ModuleKind.CommonJS,
    emitDecoratorMetadata: true,
    experimentalDecorators: true
  }

  const host: ts.CompilerHost = {
    getSourceFile: (name) => {
      if (name === sourceFile.fileName) {
        return sourceFile
      }
    },
    writeFile: () => {},
    getDefaultLibFileName: () => "",
    useCaseSensitiveFileNames: () => true,
    getCanonicalFileName: (filename: string): string => filename,
    getCurrentDirectory: () => "",
    getNewLine: () => "\n",
    getDirectories: () => [],
    fileExists: () => true,
    readFile: () => ""
  }

  return ts.createProgram([sourceFile.fileName], options, host)
}

export function preprocess(source: string): string {
  return source.replaceAll("\r", "").replace(brandable, "type $1=number & {};")
}

const brandable = /^\s*type\s+(u?int(8|16|32|64)?|float(32|64)|rune|byte)\s*=\s*number\s*;?$/gm
