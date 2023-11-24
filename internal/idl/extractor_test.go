package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_class(t *testing.T) {
	t.Parallel()

	const script = `//js

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

	actual := testExtract(t, "script.d.ts", script, "[?kind=='CLASS']")
	expected := Declarations{
		&Declaration{
			Kind: KindClass, Name: "Lorem", Type: "Lorem",
			Doc:  "Exercitation non duis qui ad.\n\nLorem dolore nostrud deserunt proident non.",
			Tags: Tags{"category": []string{"consectetur"}},
			Methods: Declarations{
				&Declaration{
					Kind: KindMethod, Name: "incididunt", Type: "number",
					Doc: "Duis duis proident fugiat sunt non duis non mollit nulla tempor et.\n\nCulpa et quis incididunt tempor amet.",
					Tags: Tags{
						"returns": []string{
							"Velit in velit laborum ad quis esse eiusmod ea et duis.",
						},
					},
					Source: "incididunt() : number;",
				},
				&Declaration{
					Kind: KindMethod, Name: "excepteur", Type: "string",
					Doc: "Amet pariatur dolor minim velit ea deserunt dolor.\n\nEx tempor culpa sit Lorem quis.",
					Tags: Tags{
						"param": []string{
							"est Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
							"ad Do fugiat laborum elit non velit ea ea nisi tempor.",
						},
						"returns": []string{"Ullamco ad et sunt aute cillum culpa aute tempor."},
					},
					Source: "excepteur(est: number, ad?: string) : string;",
					Parameters: Declarations{
						&Declaration{
							Kind: KindParameter, Name: "est", Type: "number",
							Doc: "Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
							Tags: Tags{
								"param": []string{
									"est Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
								},
							},
							Source: "est: number",
						},
						&Declaration{
							Kind: KindParameter, Name: "ad", Type: "string",
							Doc: "Do fugiat laborum elit non velit ea ea nisi tempor.",
							Tags: Tags{
								"param": []string{
									"ad Do fugiat laborum elit non velit ea ea nisi tempor.",
								},
							},
							Source: "ad?: string",
						},
					},
				},
			},
			Properties: Declarations{
				&Declaration{
					Kind: KindProperty, Name: "cupidatat", Type: "string", Modifiers: []Modifier{ModifierReadonly},
					Doc:    "Id eiusmod tempor sunt est.",
					Source: "readonly cupidatat: string;",
				},
				&Declaration{
					Kind: KindProperty, Name: "laboris", Type: "string",
					Doc:    "Nisi est veniam eiusmod exercitation quis.",
					Source: "laboris: string;",
				},
				&Declaration{
					Kind: KindProperty, Name: "velit", Type: "number",
					Doc:    "Excepteur tempor adipisicing occaecat incididunt.",
					Source: "velit: number;",
				},
			},
			Constructors: Declarations{
				&Declaration{
					Kind: KindConstructor,
					Doc:  "Dolor Lorem eu aliquip fugiat non labore nostrud.",
					Tags: Tags{
						"param": []string{
							"non Mollit minim nisi ea culpa incididunt adipisicing.",
							"occaecat Excepteur velit est Lorem voluptate consequat sit cillum quis.",
							"anim Excepteur do do nisi aliqua velit laborum.",
						},
					},
					Source: "constructor(non: string, occaecat: boolean, anim: number);",

					Parameters: Declarations{
						&Declaration{
							Kind: KindParameter, Name: "non", Type: "string",
							Doc: "Mollit minim nisi ea culpa incididunt adipisicing.",
							Tags: Tags{
								"param": []string{
									"non Mollit minim nisi ea culpa incididunt adipisicing.",
								},
							},
							Source: "non: string",
						},
						&Declaration{
							Kind: KindParameter, Name: "occaecat", Type: "boolean",
							Doc: "Excepteur velit est Lorem voluptate consequat sit cillum quis.",
							Tags: Tags{
								"param": []string{
									"occaecat Excepteur velit est Lorem voluptate consequat sit cillum quis.",
								},
							},
							Source: "occaecat: boolean",
						},
						&Declaration{
							Kind: KindParameter, Name: "anim", Type: "number",
							Doc: "Excepteur do do nisi aliqua velit laborum.",
							Tags: Tags{
								"param": []string{
									"anim Excepteur do do nisi aliqua velit laborum.",
								},
							},
							Source: "anim: number",
						},
					},
				},

				&Declaration{
					Kind: KindConstructor,
					Doc:  "Nisi tempor amet culpa aliquip dolor quis ea id.",
					Tags: Tags{
						"param": []string{
							"mollit Consequat culpa nostrud eiusmod ut ipsum elit tempor.",
							"magna Do id deserunt sunt ad ut ipsum eu et qui veniam.",
							"officia Enim irure id culpa amet est velit.",
						},
					},
					Source: "constructor(mollit: number, magna: string, officia: boolean);",

					Parameters: Declarations{
						&Declaration{
							Kind: KindParameter, Name: "mollit", Type: "number",
							Doc: "Consequat culpa nostrud eiusmod ut ipsum elit tempor.",
							Tags: Tags{
								"param": []string{
									"mollit Consequat culpa nostrud eiusmod ut ipsum elit tempor.",
								},
							},
							Source: "mollit: number",
						},
						&Declaration{
							Kind: KindParameter, Name: "magna", Type: "string",
							Doc: "Do id deserunt sunt ad ut ipsum eu et qui veniam.",
							Tags: Tags{
								"param": []string{
									"magna Do id deserunt sunt ad ut ipsum eu et qui veniam.",
								},
							},
							Source: "magna: string",
						},
						&Declaration{
							Kind: KindParameter, Name: "officia", Type: "boolean",
							Doc: "Enim irure id culpa amet est velit.",
							Tags: Tags{
								"param": []string{
									"officia Enim irure id culpa amet est velit.",
								},
							},
							Source: "officia: boolean",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_interface(t *testing.T) {
	t.Parallel()

	const script = `//js

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

	actual := testExtract(t, "script.d.ts", script, "[?kind=='INTERFACE']")
	expected := Declarations{
		&Declaration{
			Kind: KindInterface, Name: "Ipsum", Type: "Ipsum",
			Doc:  "Exercitation non duis qui ad.\n\nIpsum dolore nostrud deserunt proident non.",
			Tags: Tags{"category": []string{"consectetur"}},
			Methods: Declarations{
				&Declaration{
					Kind: KindMethod, Name: "incididunt", Type: "number",
					Doc: "Duis duis proident fugiat sunt non duis non mollit nulla tempor et.\n\nCulpa et quis incididunt tempor amet.",
					Tags: Tags{
						"returns": []string{
							"Velit in velit laborum ad quis esse eiusmod ea et duis.",
						},
					},
					Source: "incididunt() : number;",
				},
				&Declaration{
					Kind: KindMethod, Name: "excepteur", Type: "string",
					Doc: "Amet pariatur dolor minim velit ea deserunt dolor.\n\nEx tempor culpa sit Lorem quis.",
					Tags: Tags{
						"param": []string{
							"est Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
							"ad Do fugiat laborum elit non velit ea ea nisi tempor.",
						},
						"returns": []string{"Ullamco ad et sunt aute cillum culpa aute tempor."},
					},
					Source: "excepteur(est: number, ad?: string) : string;",
					Parameters: Declarations{
						&Declaration{
							Kind: KindParameter, Name: "est", Type: "number",
							Doc: "Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
							Tags: Tags{
								"param": []string{
									"est Nostrud mollit commodo exercitation minim in et.\nIn sunt aliquip ipsum magna voluptate.",
								},
							},
							Source: "est: number",
						},
						&Declaration{
							Kind: KindParameter, Name: "ad", Type: "string",
							Doc: "Do fugiat laborum elit non velit ea ea nisi tempor.",
							Tags: Tags{
								"param": []string{
									"ad Do fugiat laborum elit non velit ea ea nisi tempor.",
								},
							},
							Source: "ad?: string",
						},
					},
				},
			},
			Properties: Declarations{
				&Declaration{
					Kind: KindProperty, Name: "cupidatat", Type: "string", Modifiers: []Modifier{ModifierReadonly},
					Doc:    "Id eiusmod tempor sunt est.",
					Source: "readonly cupidatat: string;",
				},
				&Declaration{
					Kind: KindProperty, Name: "laboris", Type: "string",
					Doc:    "Nisi est veniam eiusmod exercitation quis.",
					Source: "laboris: string;",
				},
				&Declaration{
					Kind: KindProperty, Name: "velit", Type: "number",
					Doc:    "Excepteur tempor adipisicing occaecat incididunt.",
					Source: "velit: number;",
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_function(t *testing.T) {
	t.Parallel()

	const script = `//js

  /**
   * Ea laborum dolore aliqua incididunt ex commodo.
   * 
   * @returns Officia laborum tempor qui velit ipsum excepteur minim irure.
   */
  export declare function lorem(): string;
  
  /**
   * Deserunt eu exercitation incididunt mollit esse ad nisi nostrud.
   * 
   * @param enim Fugiat quis ipsum cupidatat amet velit nisi.
   * @param quis Culpa labore excepteur tempor magna sunt.
   * @param consequat Ea anim aliqua elit cupidatat enim eiusmod esse ea enim.
   */
  export declare function irure(enim:boolean, quis:string, consequat: number): void;
  
  //!js`

	actual := testExtract(t, "script.d.ts", script, "[?kind=='FUNCTION']")
	expected := Declarations{
		&Declaration{
			Kind: KindFunction, Name: "lorem", Type: "string",
			Doc: "Ea laborum dolore aliqua incididunt ex commodo.",
			Tags: Tags{
				"returns": []string{
					"Officia laborum tempor qui velit ipsum excepteur minim irure.",
				},
			},
			Source: "export declare function lorem(): string;",
		},
		&Declaration{
			Kind: KindFunction, Name: "irure", Type: "void",
			Doc: "Deserunt eu exercitation incididunt mollit esse ad nisi nostrud.",
			Tags: Tags{
				"param": []string{
					"enim Fugiat quis ipsum cupidatat amet velit nisi.",
					"quis Culpa labore excepteur tempor magna sunt.",
					"consequat Ea anim aliqua elit cupidatat enim eiusmod esse ea enim.",
				},
			},
			Source: "export declare function irure(enim:boolean, quis:string, consequat: number): void;",
			Parameters: Declarations{
				&Declaration{
					Kind: KindParameter, Name: "enim", Type: "boolean",
					Doc: "Fugiat quis ipsum cupidatat amet velit nisi.",
					Tags: Tags{
						"param": []string{"enim Fugiat quis ipsum cupidatat amet velit nisi."},
					},
					Source: "enim:boolean",
				},
				&Declaration{
					Kind: KindParameter, Name: "quis", Type: "string",
					Doc: "Culpa labore excepteur tempor magna sunt.",
					Tags: Tags{
						"param": []string{"quis Culpa labore excepteur tempor magna sunt."},
					},
					Source: "quis:string",
				},
				&Declaration{
					Kind: KindParameter, Name: "consequat", Type: "number",
					Doc: "Ea anim aliqua elit cupidatat enim eiusmod esse ea enim.",
					Tags: Tags{
						"param": []string{
							"consequat Ea anim aliqua elit cupidatat enim eiusmod esse ea enim.",
						},
					},
					Source: "consequat: number",
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_variable(t *testing.T) {
	t.Parallel()

	const script = `//js

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

	actual := testExtract(t, "script.d.ts", script, "[?kind=='VARIABLE']")
	expected := Declarations{
		&Declaration{
			Kind: KindVariable, Name: "laborum", Type: "string", Modifiers: []Modifier{ModifierConst},
			Doc:    "Ea laborum dolore aliqua incididunt ex commodo.",
			Source: "export declare const laborum: string;",
		},
		&Declaration{
			Kind: KindVariable, Name: "dolore", Type: "number",
			Doc:    "Occaecat elit consectetur nisi sint in est aliquip sint.",
			Source: "export declare var dolore: number;",
		},
		&Declaration{
			Kind: KindVariable, Name: "sunt", Type: "{}", /* TODO */
			Doc:    "Officia nulla excepteur in ea nostrud duis elit.",
			Source: "export declare var sunt: string[];",
		},
		&Declaration{
			Kind: KindVariable, Name: "quis", Type: "Array<string>",
			Doc:    "Officia laboris eu ullamco laboris.",
			Source: "export declare var quis: Array<string>;",
		},
		&Declaration{
			Kind: KindVariable, Name: "amet", Type: "Date",
			Doc:    "Ad do incididunt tempor occaecat ex velit commodo.",
			Source: "export declare var amet: Date;",
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_namespace(t *testing.T) {
	t.Parallel()

	const script = `//js
  /**
   * Laborum cupidatat elit ut amet.
   */
  
  /**
   * Nulla sunt enim culpa do irure.
   * 
   * @author Incididunt Aliquip
   */
  export as namespace lorem;
  
  //!js`

	actual := testExtract(t, "script.d.ts", script, "[?kind=='NAMESPACE']")
	expected := Declarations{
		&Declaration{
			Kind: KindNamespace, Name: "lorem",
			Doc:  "Laborum cupidatat elit ut amet.\nNulla sunt enim culpa do irure.",
			Tags: Tags{"author": {"Incididunt Aliquip"}},
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_default(t *testing.T) {
	t.Parallel()

	const script = `//js

  export as namespace lorem;

  export declare class Ipsum {
	fact: string;
  }

  declare const ipsorium : Ipsum;

  export default ipsorium;
//!js`

	actual := testExtract(t, "script.d.ts", script, "[*]")
	expected := Declarations{
		&Declaration{
			Kind: KindNamespace, Name: "lorem",
		},
		&Declaration{
			Kind: KindClass, Name: "Ipsum", Type: "Ipsum",
			Properties: Declarations{
				&Declaration{
					Kind: KindProperty, Name: "fact", Type: "string",
					Source: "fact: string;",
				},
			},
		},
		&Declaration{
			Kind: KindVariable, Name: "ipsorium", Type: "Ipsum",
			Modifiers: Modifiers{ModifierConst, ModifierDefault},
			Source:    "export declare const ipsorium : Ipsum;",
		},
	}

	assert.Equal(t, expected, actual)
}
