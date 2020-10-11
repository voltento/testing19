package repository

type Repository interface {
	Ages() ([]int, []int, error)
}
