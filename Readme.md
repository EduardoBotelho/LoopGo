# Programa de Laços em Go (Loop Program in Go)

## Descrição (Description)

Este programa em Go demonstra o uso de laços `for` para simular a exibição de horas e minutos em um relógio. O programa itera sobre um intervalo de horas (0 a 12) e outro intervalo de minutos (0 a 60), exibindo os valores correspondentes. (This Go program demonstrates the use of `for` loops to simulate displaying hours and minutes on a clock. The program iterates over a range of hours (0 to 12) and another range of minutes (0 to 60), displaying the corresponding values.)

## Como Funciona (How It Works)

1. **Iteração das Horas (Hour Iteration):**
   - Um laço `for` é usado para iterar de 0 a 12.
   - Em cada iteração, o valor atual da hora é exibido. (A `for` loop is used to iterate from 0 to 12. In each iteration, the current hour value is displayed.)

2. **Iteração dos Minutos (Minute Iteration):**
   - Outro laço `for` é usado para iterar de 0 a 60.
   - Em cada iteração, o valor atual dos minutos é exibido. (Another `for` loop is used to iterate from 0 to 60. In each iteration, the current minute value is displayed.)

3. O programa imprime os resultados no console. (The program prints the results to the console.)

## Saída Esperada (Expected Output)

Ao executar o programa, a saída será: (When running the program, the output will be:)

```
Horas: 0
Horas: 1
Horas: 2
...
Horas: 12
minutos: 0
minutos: 1
minutos: 2
...
minutos: 60
```

## Como Executar (How to Run)

1. Certifique-se de ter o Go instalado na sua máquina. (Ensure you have Go installed on your machine.)
2. Salve o código em um arquivo chamado `main.go`. (Save the code in a file named `main.go`.)
3. Abra o terminal, navegue até o diretório onde o arquivo está salvo e execute: (Open the terminal, navigate to the directory where the file is saved, and run:)

   ```bash
   go run main.go
   ```

4. Você verá a saída no terminal. (You will see the output in the terminal.)

## Código (Code)

```go
package main

import "fmt"

func main() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Print("Horas: ", horas)
		fmt.Print("\n")
	}
	for min := 0; min <= 60; min++ {
		fmt.Print("minutos: ", min)
		fmt.Print("\n")
	}
}
```

## Licença (License)

Este programa é gratuito para uso e modificação. (This program is free for use and modification.)
