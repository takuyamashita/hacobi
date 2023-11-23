package domain

type RandomStringRepositoryIntf interface {
	Generate(length int) (URLSafeString string, err error)
}

type MailIntf interface {
	Send(to string, subject string, body string) error
}
