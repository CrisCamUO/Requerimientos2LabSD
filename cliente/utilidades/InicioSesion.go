package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// Usuario representa un usuario simple en memoria
type Usuario struct {
	Id       int
	Nickname string
	Password string // en claro para la práctica, en producción hash
}

// Mapa de usuarios (simulado). Cambia o agrega según necesites.
var usuarios = []Usuario{
	{Id: 1, Nickname: "tati", Password: "1234"},
	{Id: 2, Nickname: "juan", Password: "abcd"},
	{Id: 3, Nickname: "cris", Password: "1234"},
}

// IniciarSesion muestra prompts y devuelve el nickname y el id del usuario autenticado.
// Si falla devuelve ("", 0).
func IniciarSesion() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("👤 Nickname: ")
	nickRaw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error leyendo nickname:", err)
		return "", 0
	}
	nick := strings.TrimSpace(nickRaw)

	// Pedimos contraseña sin eco
	fmt.Print("🔑 Contraseña: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("") // nueva línea después de la contraseña
	if err != nil {
		// fallback a lectura normal si hay error con ReadPassword
		fmt.Print("\n(Advertencia) No se pudo ocultar la contraseña, se leerá visible.\nContraseña: ")
		passRaw, _ := reader.ReadString('\n')
		bytePassword = []byte(strings.TrimSpace(passRaw))
	}

	password := strings.TrimSpace(string(bytePassword))

	// Validar contra el "almacén" en memoria
	for _, u := range usuarios {
		if u.Nickname == nick && u.Password == password {
			fmt.Printf("✅ Bienvenido %s (id=%d)\n", u.Nickname, u.Id)
			return u.Nickname, u.Id
		}
	}

	fmt.Println("❌ Credenciales inválidas.")
	return "", 0
}
