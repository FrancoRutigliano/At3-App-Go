package sqlxRepository

import authRepository "at3-back/internal/auth/pkg/domain/repository"

type SqlxRepository struct {
	impl authRepository.Repository
}

type ImplSqlx struct {
}

func (r *SqlxRepository) New() {
	r.impl = ImplSqlx{}
}
