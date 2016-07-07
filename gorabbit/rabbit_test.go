package gorabbit

import (
    "testing"
    "fmt"
)

func TestUni2zg(t *testing.T) {
    unicodeStr := "သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ်ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေးဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။"
    str, err := Uni2zg(unicodeStr)
    if err != nil{
        t.Errorf("failed to convert zawgyi: %v", err)
    }
    fmt.Println(str)
}

func TestZg2uni(t *testing.T) {
    zawgyiStr := "သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝဍ္ဎနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘးဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။"
    str, err := Uni2zg(zawgyiStr)
    if err != nil{
        t.Errorf("failed to convert unicode: %v", err)
    }
    fmt.Println(str)
}
