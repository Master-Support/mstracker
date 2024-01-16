package mensageria

import (
	"example.com/mstracker/enum/erros_enum"
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
)

func enviarEmail() error {
	host := "smtp.gmail.com"
	port := 587
	user := "mstracker.teste@gmail.com"
	password := "x z n y y b u a d m z j j x j s"

	msg := gomail.NewMessage()
	msg.SetHeader("From", "mstracker.teste@gmail.com")
	msg.SetHeader("To", "gestao.equipamentos@vibetecnologia.com")
	msg.SetHeader("Subject", "TESTE")
	msg.SetBody("text/html", "<h1>TESTANDO</h1>")

	msg.Attach("foo.txt", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write([]byte("Content of foo.txt"))
		return err
	}))

	dialer := gomail.NewDialer(host, port, user, password)

	client, err := dialer.Dial()
	if err != nil {
		return fmt.Errorf(erros_enum.ErrClienteSMTP)
	}

	if err := gomail.Send(client, msg); err != nil {
		return fmt.Errorf(erros_enum.ErrEnviarEmail)
	}

	fmt.Println("E-mail enviado com sucesso!")
	return nil
}
