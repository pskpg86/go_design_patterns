//https://blog.ralch.com/articles/design-patterns/golang-builder/

package car

// Message is the Product object in Builder Design Pattern
type Message struct {
	// Message Body
	Body []byte
	// Message Format
	Format string
}
