package model

type tea struct {
    name string
    rank int
    isOxidized bool
}

func (t tea) Prepare() string {
    if t.rank > 2 {
        if t.isOxidized {
            return "The " + t.name + " tea needs boiling water, wait a second please."
        } else {
            return t.name + " shall be prepared for you."
        }
    }
    return "Oh, " + t.name + " tea is one of my favourites! I'll make some for myself as well!"
}

type herb struct {
    name string
    power int
}

func (h herb) Prepare() string {
    if h.power > 9000 {
        return "You are too weak for such an intense herb."
    }
    return "This herb has so many medical properties that I forgot them all!"
}

type Prepareer interface {
    Prepare() string
}
func Box() []Prepareer {
    tea1 := &tea{name: "white", rank: 2}
    tea2 := &tea{name: "green", rank: 1}
    tea3 := &tea{name: "black", rank: 3, isOxidized: true}
    herb1 := &herb{name: "Basil", power: 12}

    box := []Prepareer{tea1, tea2, tea3, herb1}
    return box
}
