package myPackage

import "fmt"

type car struct {
	barnd string
	year  int
}

// Struct público
type CarPublic struct {
	Barnd string
	Year  int
}

// Funcion pública
func Mensaje() {
	fmt.Println("Soy una función Pública!")
}

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC PC) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}

// Para modificar el valor de la dirección de memoria
// a donde apunta la variable
func (myPC *PC) AumentarRam() {
	myPC.Ram = myPC.Ram * 2
}

// --- Para imprimir un output personalizado de un Struct --- //
// Sprintf no imprime en consola directamente, sino que
// genera un strig con el resultado que quiero
func (myPC PC) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es un %s", myPC.Ram, myPC.Disk, myPC.Brand)
}
