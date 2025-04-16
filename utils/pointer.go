package utils

// Get retorna um ponteiro para o valor passado
func Get[T any](value T) *T {
	return &value
}

// Coalesce retorna o primeiro valor não-nulo entre os ponteiros passados
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	
	return nil
}
