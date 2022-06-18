// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	//Caso não queira passar uma variável, usa-se _. A linguagem irá entender que se trata de um lixo.
	_, erros := fmt.Println("Hello, Dan", "32 anos.", "Morador de Recife")
	fmt.Println(erros)

	//tipo de variável int, string e bool
	a := 16
	b := "String"
	c := true
	fmt.Println(a, b, c)

	//:= grouph - ele serve para declarar uma variável.

}

var c = "fora do codeblocks"

func main() {

	//quando a variável é declarada fora do scopo, não se usa o grouph, mas o sinal de igual, apenas.
	//keywords (palavras chaves) são termos reservados
	//operadores - operands
	//expressão - qualquer coisa que "produz um resultado mas não gera uma ação" -> 10 + 10, 10 > 8. etc.
	//statement (declaração, expressão) uma instrução que forma uma ação e ela é formada de expressões
	a := 16
	b := "Meu tipo é"

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	
}

// o var tem que tá no nível do pacote. Ou seja, fora de um escopo. Assim, todos os escopos irão enxergar essa variável. 
//abaixo um exemplo de como funciona.
package main

import "fmt"

var z = 30

func main() {
	z := 20
	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Println(z)
	fmt.Println(x)
}

https://www.youtube.com/watch?v=YAGKHSNWdEE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=10
