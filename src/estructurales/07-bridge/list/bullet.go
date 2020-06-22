package list

import "fmt"

// Bullet imprime el listado con viñetas.-
type Bullet struct {
	bullet rune
}

func NewBullet(b rune) *Bullet {
	return &Bullet{
		bullet: b,
	}
}

func (b *Bullet) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", string(b.bullet), v)
	}
} 
