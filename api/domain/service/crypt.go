package service

type Crypt interface {
	Hash(password string) (string,error)
	Check(hash, password string) (bool,error)
}
