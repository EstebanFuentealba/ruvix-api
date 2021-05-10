package auth

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/cagodoy/ruvix-api/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

// NewService ...
func NewService(store Store, us users.Service, mt *MailingTemplates) *AuthService {
	return &AuthService{
		Store:            store,
		UserService:      us,
		MailingTemplates: mt,
	}
}

// AuthService ...
type AuthService struct {
	Store            Store
	UserService      users.Service
	MailingTemplates *MailingTemplates
}

// GetByToken ...
func (as *AuthService) GetByToken(token string) (*Auth, error) {
	// validate token param
	if token == "" {
		return nil, errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.Store.Get(&Query{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Login ...
func (as *AuthService) Login(email, password string) (*Response, error) {
	// validate email param
	if email == "" {
		return nil, errors.New("invalid email")
	}

	// validate password param
	if password == "" {
		return nil, errors.New("invalid password")
	}

	// verify if password is valid
	err := as.UserService.VerifyPassword(email, password)
	if err != nil {
		return nil, err
	}

	// get user by email
	user, err := as.UserService.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// create token difinition
	t := Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(), // one year of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("env variable JWT_SECRET must be defined")
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create auth definition
	a := &Auth{
		Token:     tokenStr,
		Kind:      KindUser,
		Blacklist: false,
		UserID:    user.ID,
	}

	// save auth store
	err = as.Store.Create(a)
	if err != nil {
		return nil, err
	}

	// prepare metatoken
	mt := &MetaToken{
		Token: tokenStr,
	}

	// prepare response
	res := &Response{
		Data: user,
		Meta: mt,
	}

	return res, nil
}

// Signup ...
func (as *AuthService) Signup(u *users.User) (*Response, error) {
	// validate user existence
	if u == nil {
		return nil, errors.New("invalid user")
	}

	// validate user params
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return nil, errors.New("invalid user params")
	}

	// create new user
	err := as.UserService.Create(u)
	if err != nil {
		return nil, err
	}

	// create token definition
	t := Token{
		UserID: u.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(4320 * time.Hour).Unix(), // six months of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("env variable JWT_SECRET must be defined")
	}

	// generate login jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create signup auth definition
	a := &Auth{
		Token:     tokenStr,
		Kind:      KindUser,
		Blacklist: false,
		UserID:    u.ID,
	}

	// save auth store
	err = as.Store.Create(a)
	if err != nil {
		return nil, err
	}

	// create verify token definition
	vt := Token{
		UserID: u.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(2160 * time.Hour).Unix(), // 3 months of expiration
		},
	}

	// generate verify jwt token
	verificationToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), vt)
	verificationTokenStr, err := verificationToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create verification auth definition
	va := &Auth{
		Token:     verificationTokenStr,
		Kind:      KindVerifyPassword,
		Blacklist: false,
		UserID:    u.ID,
	}

	// save auth store
	err = as.Store.Create(va)
	if err != nil {
		return nil, err
	}

	// prepare metatoken
	mt := &MetaToken{
		Token:             tokenStr,
		VerificationToken: verificationTokenStr,
	}

	// prepare response
	res := &Response{
		Data: u,
		Meta: mt,
	}

	// send signup email with token and url
	if as.MailingTemplates.Signup != nil {
		err = as.MailingTemplates.Signup(u)
		if err != nil {
			return nil, err
		}
	}

	// send email verification with token and url
	if as.MailingTemplates.VerifyEmail != nil {
		err = as.MailingTemplates.VerifyEmail(u, verificationTokenStr)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// VerifyToken ...
func (as *AuthService) VerifyToken(token string, kind string) (*Token, error) {
	// validate token param
	if token == "" {
		return nil, errors.New("invalid token")
	}

	// validate kind param
	if kind == "" {
		return nil, errors.New("invalid kind")
	}

	// get Auth by token from store
	a, err := as.Store.Get(&Query{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	// validate auth kind
	if a.Kind != kind {
		return nil, errors.New("invalid kind")
	}

	// check if token is blacklisted
	if a.Blacklist {
		return nil, errors.New("token is blacklisted")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	decode := new(Token)
	err = json.Unmarshal(data, decode)
	if err != nil {
		return nil, err
	}

	return decode, nil
}

// VerifyEmail ...
func (as *AuthService) VerifyEmail(token string) error {
	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.Store.Get(&Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// validate auth kind
	if a.Kind != KindVerifyPassword {
		return errors.New("invalid kind")
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// update verified to true
	err = as.UserService.Update(&users.User{
		ID:       a.UserID,
		Verified: true,
	})
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.Store.Update(a)
	if err != nil {
		return err
	}

	return nil
}

// Logout ...
func (as *AuthService) Logout(token string) error {
	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.Store.Get(&Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// validate auth kind is user
	if a.Kind != KindUser {
		return errors.New("invalid auth kind")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	var decode Token
	err = json.Unmarshal(data, &decode)
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.Store.Update(a)
	if err != nil {
		return err
	}

	return nil
}

// ForgotPassword ...
func (as *AuthService) ForgotPassword(e string) (string, error) {
	// validate email param
	if e == "" {
		return "", errors.New("invalid email")
	}

	// check if email exist on users service
	user, err := as.UserService.GetByEmail(e)
	if err != nil {
		return "", err
	}

	// create temporal token difinition
	t := Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // 2 hours of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	// create auth definition
	a := &Auth{
		Token:     tokenStr,
		Kind:      KindForgotPassword,
		Blacklist: false,
		UserID:    user.ID,
	}

	// save auth store
	err = as.Store.Create(a)
	if err != nil {
		return "", err
	}

	// send email with token and url
	if as.MailingTemplates.ForgotPassword != nil {
		err = as.MailingTemplates.ForgotPassword(user, tokenStr)
		if err != nil {
			return "", err
		}
	}

	return tokenStr, nil
}

// RecoverPassword ...
func (as *AuthService) RecoverPassword(newPassword, token string) error {
	// validate newPassword param
	if newPassword == "" {
		return errors.New("invalid newPassword")
	}

	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.Store.Get(&Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// validate auth kind is forgot-password
	if a.Kind != KindForgotPassword {
		return errors.New("invalid auth kind")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	var decode Token
	err = json.Unmarshal(data, &decode)
	if err != nil {
		return err
	}

	user := &users.User{
		ID:       decode.UserID,
		Password: newPassword,
	}

	// update user password
	err = as.UserService.Update(user)
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.Store.Update(a)
	if err != nil {
		return err
	}

	// send password changed email
	if as.MailingTemplates.PasswordChanged != nil {
		err = as.MailingTemplates.PasswordChanged(user)
		if err != nil {
			return err
		}
	}

	return nil
}
