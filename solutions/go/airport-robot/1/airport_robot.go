package airportrobot
import "fmt"


type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type germanGreeter struct {}

func (l germanGreeter) Greet(name string) string {
    return fmt.Sprintf("Hallo %s!", name);
}
func (l germanGreeter) LanguageName() string {
    return "German";
}
func SayHello(name string, greeter Greeter) string {
    var languageName string = greeter.LanguageName();
    var greeting string = greeter.Greet(name);
    return fmt.Sprintf("I can speak %s: %s", languageName, greeting);
}

type Italian struct {}
func (l Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name);
}
func (l Italian) LanguageName() string {
    return "Italian";
}

type Portuguese struct {}
func (l Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name);
}
func (l Portuguese) LanguageName() string {
    return "Portuguese";
}