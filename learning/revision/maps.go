package revision

import "fmt"

func Map() {
	a := map[string]string{}
	a["name"] = "Chetan"
	fmt.Println(a)

	b := make(map[string]string)
	b["name"] = "Bindu"
	fmt.Println(b["name"])

	// Iterate over maps
	c := make(map[string]string)
	c["name"] = "Chetan"
	c["age"] = "31"
	for keys, value := range c {
		fmt.Println(keys, value)
	}

	// if key exits
	d := make(map[string]string)
	d["name"] = "Chetan"
	d["age"] = "31"
	temp := make(map[string]string)

	for dkey, dvalue := range d {
		if tempValue, boolCheck := temp[dvalue]; boolCheck {
			fmt.Println(tempValue)
		} else {
			temp[dkey] = dvalue
		}
	}
	fmt.Println(temp)

	// delete from map
	delete(d, "name")
}
