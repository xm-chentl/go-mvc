package verify

type IVerify interface {
	Args(string) IVerify
	Execute(interface{}) bool
}
