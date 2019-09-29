package api

var Outage map[int]bool

func init() {
	Outage = make(map[int]bool)

	Outage[212122] = true
	Outage[206522] = true
	Outage[210000] = true
	Outage[212000] = true
	Outage[212001] = true
	Outage[212002] = true
	Outage[212003] = true
}
