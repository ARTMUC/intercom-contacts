package response

type Pagination[T any] struct {
	Rows  []T `json:"rows,omitempty"`
	Count int `json:"count,omitempty"`
}

func NewPaginationReponse[T any](t []T, count int) *Pagination[T] {
	res := Pagination[T]{Rows: make([]T, len(t)), Count: count}
	for i, e := range t {
		res.Rows[i] = T(e)
	}

	return &res
}

func NewPaginationReponseWithMap[T any, U any](input []T, count int, mappers ...func(T) U) *Pagination[U] {
	output := make([]U, 0)

	for _, v := range input {
		var result U
		for _, mapper := range mappers {
			result = mapper(v)
		}

		output = append(output, result)
	}

	return &Pagination[U]{Rows: output, Count: count}
}
