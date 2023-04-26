package mock

type Message struct {
	Status string `json:"status"`
	Data   Update `json:"data"`
}

type Update struct {
	Main     int    `json:"main"`
	Version  int    `json:"version"`
	Release  int    `json:"release"`
	Channel  string `json:"channel"`
	Codename string `json:"codename"`
	Date     string `json:"date"`
}

func CheckUpdate() *Message {
	update := Update{
		Main:     0,
		Version:  0,
		Release:  1,
		Channel:  "alpha",
		Codename: "shine",
		Date:     "13/03/2023",
	}
	message := Message{
		Status: "success",
		Data:   update,
	}
	return &message
}
