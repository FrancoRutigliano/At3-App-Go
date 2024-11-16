package authUseCases

import (
	infraSqlxRepository "at3-back/internal/auth/infrastructure/repository"
	authUseCaseImpl "at3-back/internal/auth/pkg/useCases/useCaseImpl"
	"at3-back/internal/shared/infrastructure/data"
	"at3-back/internal/shared/infrastructure/service"
	"context"
	"log"
	"os"
)

type AuthImpl struct {
	Impl authUseCaseImpl.IauthUseCase
}

func (a *AuthImpl) New() {
	var ctx = context.Background()
	var repository infraSqlxRepository.SqlxRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Fatal("Failed to instance DB conn: Error:", err)
	}

	redis, err := data.NewRedisConnnection()
	if err != nil {
		log.Fatal("Failed to instance Redis conn: Error:", err)
	}

	var emailService service.EmailService

	err = emailService.New(os.Getenv("SMTP_HOST"), "465", "no-reply@atomico3.io", os.Getenv("MAIL_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	var jwtService service.JwtService

	jS := jwtService.New()

	a.Impl = &authUseCaseImpl.Auth{
		Repository:   repository,
		Db:           db,
		Redis:        redis,
		Ctx:          ctx,
		EmailService: emailService,
		JwtService:   *jS,
	}
}
