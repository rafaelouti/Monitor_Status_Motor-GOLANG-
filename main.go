package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Status do motor
type MotorStatus struct {
	temperature float64
	speed       int
	vibration   float64
}

// Função para simular a leitura dos sensores do motor
func readMotorStatus() MotorStatus {
	return MotorStatus{
		temperature: 50 + rand.Float64()*50, // Simula temperatura entre 50-100ºC
		speed:       1000 + rand.Intn(500),  // Simula RPM entre 1000-1500
		vibration:   rand.Float64() * 10,    // Simula vibração entre 0-10
	}
}

// Função para monitorar o status do motor e exibir avisos
func monitorMotor() {
	for {
		status := readMotorStatus()
		fmt.Printf("Status do Motor: Temperatura=%.2f°C, Velocidade=%d RPM, Vibração=%.2f\n",
			status.temperature, status.speed, status.vibration)

		// Condições de alerta
		if status.temperature > 80 {
			fmt.Println("Aviso: Temperatura Alta!")
		}
		if status.vibration > 5 {
			fmt.Println("Aviso: Nível de Vibração Alto!")
		}

		// Pausa entre leituras
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("Iniciando monitoramento do motor...")
	monitorMotor()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
