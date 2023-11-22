package live_house_staff_email_authorization_domain

type RandomStringRepositoryIntf interface {
	Generate(length int) (URLSafeString string, err error)
}
