package main

/*
import (
	"fmt"
	"time"
)

// Project representa un proyecto creativo
type Project struct {
	Name        string
	Description string
	Team        []TeamMember
	Milestones  []Milestone
}

// TeamMember representa un miembro del equipo
type TeamMember struct {
	Name          string
	Role          string
	Specialization string
}

// Milestone representa un hito en el proyecto con fecha límite
type Milestone struct {
	Description string
	Deadline    time.Time
}

// Función para crear un nuevo proyecto
func createNewProject(name, description string) *Project {
	return &Project{
		Name:        name,
		Description: description,
	}
}

// Función para asignar un miembro del equipo a un proyecto
func assignTeamMember(project *Project, member TeamMember) {
	project.Team = append(project.Team, member)
}

// Función para añadir un hito y fecha límite al proyecto
func trackProgress(project *Project, milestone Milestone) {
	project.Milestones = append(project.Milestones, milestone)
}

// Función para generar un informe de estado del proyecto
func generateProjectReport(project Project) {
	fmt.Printf("Informe de estado del proyecto: %s\n", project.Name)
	fmt.Printf("Descripción: %s\n", project.Description)
	fmt.Println("Miembros del equipo:")
	for _, member := range project.Team {
		fmt.Printf("- %s (%s)\n", member.Name, member.Role)
	}
	fmt.Println("Hitos y fechas límite:")
	for _, milestone := range project.Milestones {
		fmt.Printf("- %s (Fecha límite: %s)\n", milestone.Description, milestone.Deadline.Format("02-01-2006"))
	}
}

func main() {
	// Crear un nuevo proyecto
	project := createNewProject("Proyecto de Diseño Gráfico", "Desarrollo de una nueva identidad visual")

	// Asignar miembros del equipo al proyecto
	assignTeamMember(project, TeamMember{Name: "Ana García", Role: "Diseñadora", Specialization: "Diseño Gráfico"})
	assignTeamMember(project, TeamMember{Name: "Juan Pérez", Role: "Desarrollador Frontend", Specialization: "Diseño Web"})

	// Seguimiento del progreso del proyecto
	trackProgress(project, Milestone{Description: "Entrega de Propuestas de Diseño", Deadline: time.Date(2024, time.April, 30, 0, 0, 0, 0, time.UTC)})
	trackProgress(project, Milestone{Description: "Desarrollo de Prototipo Interactivo", Deadline: time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC)})

	// Generar informe de estado del proyecto
	generateProjectReport(*project)
}
*/


import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Character representa un personaje en el juego
type Character struct {
	Name        string
	Health      int
	MaxHealth   int
	Strength    int
	Defense     int
	Experience  int
	Level       int
	CurrentRoom *Room
}

// Room representa una ubicación en el mundo del juego
type Room struct {
	Name        string
	Description string
	Enemies     []*Enemy
}

// Enemy representa un enemigo en el juego
type Enemy struct {
	Name     string
	Health   int
	Strength int
	Defense  int
}

// Función para simular un combate por turnos
func battle(player *Character, enemy *Enemy) bool {
	fmt.Printf("¡Comienza el combate! %s vs %s\n", player.Name, enemy.Name)

	for player.Health > 0 && enemy.Health > 0 {
		// Turno del jugador
		playerAttack := player.Strength - enemy.Defense
		if playerAttack < 0 {
			playerAttack = 0
		}
		enemy.Health -= playerAttack
		fmt.Printf("%s golpea a %s y hace %d de daño\n", player.Name, enemy.Name, playerAttack)

		// Verificar si el enemigo fue derrotado
		if enemy.Health <= 0 {
			fmt.Printf("%s derrotó a %s!\n", player.Name, enemy.Name)
			player.Experience += 10
			return true
		}

		// Turno del enemigo
		enemyAttack := enemy.Strength - player.Defense
		if enemyAttack < 0 {
			enemyAttack = 0
		}
		player.Health -= enemyAttack
		fmt.Printf("%s golpea a %s y hace %d de daño\n", enemy.Name, player.Name, enemyAttack)

		// Verificar si el jugador fue derrotado
		if player.Health <= 0 {
			fmt.Printf("%s fue derrotado por %s :(\n", player.Name, enemy.Name)
			return false
		}

		// Esperar un momento antes del siguiente turno
		time.Sleep(1 * time.Second)
	}

	return false
}

// Función para que el jugador explore una habitación
func exploreRoom(player *Character) {
	currentRoom := player.CurrentRoom
	fmt.Println(currentRoom.Description)

	if len(currentRoom.Enemies) > 0 {
		for _, enemy := range currentRoom.Enemies {
			if !battle(player, enemy) {
				break
			}
		}
	}

	player.CurrentRoom = nil
}

// Función para manejar la interacción del jugador con el juego
func playGame(player *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("---------------------------------")
		fmt.Printf("¡Bienvenido/a, %s! Estás en %s\n", player.Name, player.CurrentRoom.Name)
		fmt.Println("¿Qué deseas hacer? (explorar / salir)")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "explorar":
			if player.CurrentRoom != nil {
				exploreRoom(player)
			} else {
				fmt.Println("No hay una habitación para explorar.")
			}
		case "salir":
			fmt.Println("Gracias por jugar. ¡Hasta luego!")
			return
		default:
			fmt.Println("Comando no reconocido. Intenta de nuevo.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Inicializar al jugador
	player := &Character{
		Name:       "Héroe",
		MaxHealth:  100,
		Health:     100,
		Strength:   15,
		Defense:    5,
		Experience: 0,
		Level:      1,
	}

	// Crear habitaciones y enemigos
	room1 := &Room{
		Name:        "Bosque Oscuro",
		Description: "Te encuentras en un oscuro bosque. ¡Cuidado con los enemigos!",
		Enemies: []*Enemy{
			{Name: "Goblin", Health: 30, Strength: 8, Defense: 2},
			{Name: "Lobo", Health: 40, Strength: 10, Defense: 4},
		},
	}
	room2 := &Room{
		Name:        "Cueva Misteriosa",
		Description: "Una cueva oscura y misteriosa se abre ante ti. Puede haber peligros dentro...",
		Enemies: []*Enemy{
			{Name: "Esqueleto", Health: 50, Strength: 12, Defense: 6},
		},
	}

	// Asignar la ubicación inicial al jugador
	player.CurrentRoom = room1

	// Iniciar el juego
	fmt.Println("¡Bienvenido/a al juego de rol!")
	playGame(player)
}
