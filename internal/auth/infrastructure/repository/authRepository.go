package infraSqlxRepository

import authRepository "at3-back/internal/auth/pkg/domain/repository"

type SqlxRepository struct {
	Impl authRepository.Repository
}

type ImplSqlx struct {
}

func (r *SqlxRepository) New() {
	r.Impl = &ImplSqlx{}
}
