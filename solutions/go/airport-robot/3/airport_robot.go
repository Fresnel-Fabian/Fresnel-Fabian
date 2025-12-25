package airportrobot
import "fmt"


type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type German struct {}

func (l German) Greet(name string) string {
    return fmt.Sprintf("Hallo %s!", name);
}

func (l German) LanguageName() string {
    return "German";
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

func SayHello(name string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name));
}