package domain

type RandomStringRepositoryIntf interface {
	Generate(length int) (URLSafeString string, err error)
}
