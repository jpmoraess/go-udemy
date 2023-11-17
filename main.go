package main

import "fmt"

func main() {
	/**
	################ SECAO 1: FUNDAMENTOS ################
	**/

	/*
		var peso int
		peso = 10
		fmt.Println(peso)
		script := "teste"
		fmt.Println(script)
	*/

	/*
		texto := "teste"
		numero := 10
		metro := 1.5
		isMale := true
		fmt.Println(texto, numero, metro, isMale)
	*/

	/*
		var texto string
		var numero int
		var metro float32
		var isMale bool
		fmt.Println(texto, numero, metro, isMale)
	*/

	/*
		num1 := 10.0
		num2 := 20.0
		result := num1 + num2
		result := num1 - num2
		result := num1 * num2
		result := num1 / num2
		fmt.Println(result)
		fmt.Println(reflect.TypeOf(result))
	*/

	/*
		const taxa = 10
		fmt.Println(taxa)
		fmt.Println(reflect.TypeOf(taxa))
	*/

	/*
		var numero int8
		numero = 127
		fmt.Println(numero)
	*/

	/*
		var numero int8 = 127
		var numeroInt int
		numeroInt = int(numero)
		fmt.Println(numeroInt)
		var numero2 int = 127
		var numeroInt8 int8
		numeroInt8 = int8(numero2)
		fmt.Println(numeroInt8)
	*/

	/*
		b, _ := strconv.ParseBool("true")
		fmt.Println(b)
		i, _ := strconv.ParseInt("-42", 10, 64)
		fmt.Println(i)
		texto := "42.55"
		in, _ := strconv.ParseFloat(texto, 64)
		fmt.Println(in)
	*/

	/**
	################ SECAO 2: FLUXO DE CONTROLE ################
	**/

	/*
		salario := 1000.00
		var salarioComBonus float64
		salarioComBonus = salario
		if salario <= 1000.00 {
			salarioComBonus = salarioComBonus + 100
			salarioComBonus += 100
		}
		fmt.Println("Salario: ", salarioComBonus)
	*/

	/*
		var isCar bool = false
		var value = 1000.00
		if isCar {
			value += 55.50
		} else {
			value += 20.50
		}
		fmt.Println("Final value: ", value)
	*/

	/*
		salario := 1000.00
		if salario < 1000.00 {
			salario = salario - (salario * 0.08)
		} else if salario <= 1200.00 {
			salario = salario - (salario * 0.10)
		} else {
			salario = salario - (salario * 0.10)
		}
		fmt.Println("Salario final: ", salario)
	*/

	/*
		salario := 900.00
		tipo := "PJ"
		if salario < 1000.00 && tipo == "CLT" {
			salario = salario - (salario * 0.08)
		} else if salario < 1000.00 && tipo == "PJ" {
			salario = salario - (salario * 0.05)
		} else if salario <= 1200.00 {
			salario = salario - (salario * 0.10)
		} else {
			salario = salario - (salario * 0.15)
		}
		fmt.Println("Salario final: ", salario)
	*/

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		texto := "palavra"
		tamanho := len(texto)
		fmt.Println("Quantidade de letras:", tamanho)
		for i := 0; i < tamanho; i++ {
			if string(texto[i]) == "r" {
				continue
			}
			fmt.Println(string(texto[i]))
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		texto := "palavra"
		tamanho := len(texto)
		fmt.Println("Quantidade de letras:", tamanho)
		i := 0
		for i < tamanho {
			if string(texto[i]) == "r" {
				continue
			}
			fmt.Println(string(texto[i]))
			i++
		}
	*/

	/*
		for i := 1; i <= 10; i++ {
			fmt.Printf("Tabuada do [%d]\n", i)
			for j := 1; j <= 10; j++ {
				fmt.Printf("[%d]x[%d]=[%d]\n", i, j, i*j)
			}
		}
	*/

	/**
	################ SECAO 3: SLICES E ARRAYS ################
	**/

	/*
		lista := []int{4, 9, 6, 7}
		lista = append(lista, 8)

		listaDeString := []string{"Go", "C#", "Java"}
		listaDeString = append(listaDeString, "Kotlin")
		fmt.Println(listaDeString)
		fmt.Println("Lista:", lista)
		fmt.Println("Lista pos1:", lista[0])
		fmt.Println("Lista pos2:", lista[1])
		fmt.Println("Tamanho da lista:", len(lista))

		lista2 := make([]float64, 1)
		lista2[0] = 5.0
		lista2 = append(lista2, 2.0)
		lista2 = append(lista2, 3.0)
		fmt.Println("Lista2:", lista2)
		fmt.Printf("%T\n", lista2)

		var soma float64
		for i := 0; i < len(lista2); i++ {
			soma += lista2[i]
		}
		media := soma / float64(len(lista2))
		fmt.Println("Media:", media)
	*/

	/*
		var lista = []int{2, 10, 9, 4, 5, 8, 1, 3}
		var listaMenorQueCinco = make([]int, 0)
		for i := 0; i < len(lista); i++ {
			if lista[i] < 5 {
				listaMenorQueCinco = append(listaMenorQueCinco, lista[i])
			}
		}
		fmt.Println(listaMenorQueCinco)

		lista2 := lista[:3]
		lista3 := lista[4:]
		ultimoItem := lista[len(lista)-1:]
		fmt.Println(lista2)
		fmt.Println(lista3)
		fmt.Println(ultimoItem)
	*/

	/*
		//mapa := map[string]int{"SP": 10000000, "CG": 900000}
		mapa := make(map[string]int)
		mapa["SP"] = 10000000
		mapa["CG"] = 900000

		valor, existe := mapa["RJ"]
		if existe {
			fmt.Println(valor)
		}
		fmt.Println(mapa)
		fmt.Println(mapa["SP"])

		if valor, existe := mapa["RJ"]; existe {
			fmt.Println(valor)
		} else {
			fmt.Println("Chave não encontrada...")
		}
	*/

	/*
		mapa := make(map[string]int)
		mapa["SP"] = 100000
		mapa["MG"] = 10000
		mapa["RJ"] = 6000000
		delete(mapa, "MG")

		for chave, valor := range mapa {
			fmt.Println("chave:", chave, "valor:", valor)
		}
	*/

	/**
	################ SECAO 4: EXERCICIOS ################
	**/

	/**
	* 1) Criar um array com 2 posi;óes de inteiros e armazenar em uma variavel a soma total da lista.
	* A variavel deve ser imprimida no console
	**/

	/**
	array := [2]int{10, 30}
	var soma int
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}
	fmt.Println("Soma: ", soma)
	**/

	/**
	* 2) Dado um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1"que vao de 1 a 10, efetuar a soma em duas variaveis, a primeira numeros de 1 a 5
	* e a segunda de 6 a 10.
	* Imprimir os dois resultados
	**/

	/*
		numeros := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
		var total1Ate5 int
		var total6ate10 int
		for i := 0; i < len(numeros); i++ {
			if numeros[i] <= 5 {
				total1Ate5 += numeros[i]
			} else {
				total6ate10 += numeros[i]
			}
		}
		fmt.Println("Total de 1 ate 5:", total1Ate5)
		fmt.Println("Total de 6 ate 10:", total6ate10)
	*/

	/**
	################ SECAO 5: FUNCOES ################
	**/

	/*
		ImprimeMensagem("João Dev Java/Kotlin está aprendendo GO")
		ImprimeMensagem("João está aprendendo GO com facilidade")
	*/

	/*
		resultado := Soma(1, 2)
		fmt.Println("Resultado:", resultado)
	*/

	soma, subtracao, multiplicacao, divisao := Operacao(10, 5)
	fmt.Println(soma, subtracao, multiplicacao, divisao)
}

/*
func Operacao(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subtracao := num1 - num2
	multiplicacao := num1 * num2
	divisao := num1 / num2
	return soma, subtracao, multiplicacao, divisao
}
*/

func Operacao(num1 int, num2 int) (soma int, subtracao int, multiplicacao int, divisao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	multiplicacao = num1 * num2
	divisao = num1 / num2
	return
}

func Soma(num1 int, num2 int) int {
	return num1 + num2
}

func ImprimeMensagem(mensagem string) {
	mensagem += "!!!"
	fmt.Println(mensagem)
}
