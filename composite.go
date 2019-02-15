package composite

import "fmt"

// ============================================================== //

// type MenuComponent interface {
// 	Name() string
// 	Description() string
// 	Price() float32
// 	Print()

// 	Add(MenuComponent)
// 	Remove(int)
// 	Child(int) MenuComponent
// }

// type MenuItem struct {
// 	name        string
// 	description string
// 	price       float32
// }

// type Menu struct {
// 	name        string
// 	description string
// 	children    []MenuComponent
// }

// ================================================================ //
//																	//
//					Component-Operation								//
//																	//
// ================================================================ //

type MenuComponent interface {
	Price() float32
	Print()
}

// ================================================================ //
//																	//
//							抽離屬性								 //
//						Component-Attributes						//
//																	//
// ================================================================ //
type MenuDesc struct {
	name        string
	description string
}

func (m *MenuDesc) Name() string {
	return m.name
}

func (m *MenuDesc) Description() string {
	return m.description
}

// ================================================================ //
//																	//
//							Left									//
//																	//
// ================================================================ //

type MenuItem struct {
	MenuDesc
	price float32
}

func NewMenuItem(name, description string, price float32) MenuComponent {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
		price: price,
	}
}

func (m *MenuItem) Price() float32 {
	return m.price
}

func (m *MenuItem) Print() {
	fmt.Printf("  %s, ￥%.2f\n", m.name, m.price)
	fmt.Printf("    -- %s\n", m.description)
}

// ================================================================ //
//																	//
//							Composite								//
//																	//
// ================================================================ //

type Group interface {
	Add(c MenuComponent)
	Remove(idx int)
	Child(idx int) MenuComponent
}

type MenuGroup struct {
	children []MenuComponent
}

type Menu struct {
	MenuDesc
	MenuGroup
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}

// for component //
func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ￥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
	fmt.Println()
}

// for group //
func (m *Menu) Add(c MenuComponent) {
	m.children = append(m.children, c)
}

func (m *Menu) Remove(idx int) {
	// check index
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}

func (m *Menu) Child(idx int) MenuComponent {
	// check index
	return m.children[idx]
}

// ================================================================ //
//																	//
//							Extra									//
//																	//
// ================================================================ //

type Product interface {
	Price() float32
}

// ============================================================== //
// func (m *MenuItem) Name() string {
// 	return m.name
// }

// func (m *MenuItem) Description() string {
// 	return m.description
// }

// func (m *MenuItem) Price() float32 {
// 	return m.price
// }

// func (m *MenuItem) Print() {
// 	fmt.Printf("  %s, $%.2f\n", m.name, m.price)
// 	fmt.Printf("    -- %s\n", m.description)
// }

// func (m *MenuItem) Add(MenuComponent) {
// 	panic("not implement")
// }

// func (m *MenuItem) Remove(int) {
// 	panic("not implement")
// }

// func (m *MenuItem) Child(int) MenuComponent {
// 	panic("not implement")
// }

// func NewMenu(name, description string) MenuComponent {
// 	return &Menu{
// 		name:        name,
// 		description: description,
// 	}
// }

// func (m *Menu) Name() string {
// 	return m.name
// }

// func (m *Menu) Description() string {
// 	return m.description
// }

// func (m *Menu) Price() (price float32) {
// 	for _, v := range m.children {
// 		price += v.Price()
// 	}
// 	return
// }

// func (m *Menu) Print() {
// 	fmt.Printf("%s, %s, $%.2f\n", m.name, m.description, m.Price())
// 	fmt.Println("------------------------")
// 	for _, v := range m.children {
// 		v.Print()
// 	}
// 	fmt.Println()
// }

// func (m *Menu) Add(c MenuComponent) {
// 	m.children = append(m.children, c)
// }

// func (m *Menu) Remove(idx int) {
// 	m.children = append(m.children[:idx], m.children[idx+1:]...)
// }

// func (m *Menu) Child(idx int) MenuComponent {
// 	return m.children[idx]
// }

// func NewMenuItem(name, description string, price float32) MenuComponent {
// 	return &MenuItem{
// 		name:        name,
// 		description: description,
// 		price:       price,
// 	}
// }
