package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var tempoTrabalho time.Duration
	var tempoDescanso time.Duration
	var tempoMaior time.Duration

	var rootCmd = &cobra.Command{
		// define o comando como pomodoro 

		Use:   "pomodoro",

		// descrição curta

		Short: "Pomodoro timer",

		// é executada quando o comando for chamado

		Run: func(cmd *cobra.Command, args []string) {
			tempoDeTrabalho := tempoTrabalho
			tempoDeDescanso := tempoDescanso
			tempoDescansoMaior := tempoMaior

			for {
				// tempo trabalhando/focado

				fmt.Println("Comece a trabalhar/focar")
				time.Sleep(tempoDeTrabalho)

				// descanço maior quando o tempo de trabalho for = ou maior que 4

				if tempoDeTrabalho >= 4*tempoTrabalho {
					fmt.Println("Começo do descanso maior")
					time.Sleep(tempoDescansoMaior)
					tempoDeTrabalho = tempoTrabalho * time.Minute
				} else {
					// tempo de descanso normal
					
					fmt.Println("Começo do descanso")
					time.Sleep(tempoDeDescanso)
					tempoDeTrabalho += tempoTrabalho
				}
			}
		},
	}

	// define um tempo padrão, mas caso o usuário queira pode escolher o tempo usando o --tempo-trabalho 1m ou --t 1m na forma curta

	rootCmd.Flags().DurationVarP(&tempoTrabalho, "tempo-trabalho", "t", 25*time.Minute, "Tempo em minutos para estar trabalhando/focado")
	rootCmd.Flags().DurationVarP(&tempoDescanso, "tempo-descanso", "d", 5*time.Minute, "Tempo em minutos para descanso")
	rootCmd.Flags().DurationVarP(&tempoMaior, "tempo-maior", "m", 15*time.Minute, "Tempo em minutos para um descanso maior")

	// trata o erro e mostra na linha de comando o erro
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
