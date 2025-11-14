package solidprinciples

//Dependency Inversion Principle (DIP)
//This one states that high level modules should not depend on state level modules,but both should depend on abstactions

//In the below example the Notification service is directly dependent on Email service
type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	// Send email
}

type NotificationService struct {
	emailService *EmailService
}

func (n *NotificationService) Notify(to string, message string) {
	n.emailService.Send(to, message)
}

//So for this we be dependent on the MessageSender interface instead

type MessageSender interface {
	Send(to string, message string)
}

type EmailServicee struct{}

func (e *EmailServicee) Send(to string, message string) {
	// Send email
}

type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
	// Send SMS
}

type NotificationServicee struct {
	messageSender MessageSender
}

func (n *NotificationServicee) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}
