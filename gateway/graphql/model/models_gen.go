// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BaseResponse interface {
	IsBaseResponse()
	GetCode() string
	GetStatus() ResponseStatus
	GetError() *string
}

type AuthMutation struct {
	Login         *AuthResponse `json:"login"`
	Logout        *AuthResponse `json:"logout"`
	OpenAuthorize *AuthResponse `json:"openAuthorize"`
	SignUp        *AuthResponse `json:"signUp"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type CreateProductRequest struct {
	Name      string    `json:"name"`
	Something []*string `json:"something,omitempty"`
}

type ListProductResponse struct {
	Products   []*Product          `json:"products,omitempty"`
	Pagination *PaginationResponse `json:"pagination"`
}

type ListUserResponse struct {
	Products   []*User             `json:"products,omitempty"`
	Pagination *PaginationResponse `json:"pagination"`
}

type Mutation struct {
}

type NormalLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NormalSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OpenAuthorizeRequest struct {
	Token    string                `json:"token"`
	Provider OpenAuthorizeProvider `json:"provider"`
}

type PaginationResponse struct {
	Total  string `json:"total"`
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ProductMutation struct {
	Create bool `json:"create"`
}

type ProductQuery struct {
	Get  *Product             `json:"get,omitempty"`
	List *ListProductResponse `json:"list"`
}

type Query struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UserQuery struct {
	Get  *User             `json:"get,omitempty"`
	List *ListUserResponse `json:"list"`
}

type LoginResponse string

const (
	LoginResponseNormal LoginResponse = "NORMAL"
	LoginResponseOauth  LoginResponse = "OAUTH"
)

var AllLoginResponse = []LoginResponse{
	LoginResponseNormal,
	LoginResponseOauth,
}

func (e LoginResponse) IsValid() bool {
	switch e {
	case LoginResponseNormal, LoginResponseOauth:
		return true
	}
	return false
}

func (e LoginResponse) String() string {
	return string(e)
}

func (e *LoginResponse) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LoginResponse(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LoginResponse", str)
	}
	return nil
}

func (e LoginResponse) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OpenAuthorizeProvider string

const (
	OpenAuthorizeProviderFacebook OpenAuthorizeProvider = "FACEBOOK"
	OpenAuthorizeProviderGoogle   OpenAuthorizeProvider = "GOOGLE"
)

var AllOpenAuthorizeProvider = []OpenAuthorizeProvider{
	OpenAuthorizeProviderFacebook,
	OpenAuthorizeProviderGoogle,
}

func (e OpenAuthorizeProvider) IsValid() bool {
	switch e {
	case OpenAuthorizeProviderFacebook, OpenAuthorizeProviderGoogle:
		return true
	}
	return false
}

func (e OpenAuthorizeProvider) String() string {
	return string(e)
}

func (e *OpenAuthorizeProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OpenAuthorizeProvider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OpenAuthorizeProvider", str)
	}
	return nil
}

func (e OpenAuthorizeProvider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResponseStatus string

const (
	ResponseStatusSuccess ResponseStatus = "SUCCESS"
)

var AllResponseStatus = []ResponseStatus{
	ResponseStatusSuccess,
}

func (e ResponseStatus) IsValid() bool {
	switch e {
	case ResponseStatusSuccess:
		return true
	}
	return false
}

func (e ResponseStatus) String() string {
	return string(e)
}

func (e *ResponseStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResponseStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResponseStatus", str)
	}
	return nil
}

func (e ResponseStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
