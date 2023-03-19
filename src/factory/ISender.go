package factory

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
