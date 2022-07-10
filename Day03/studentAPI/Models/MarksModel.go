package Models

type Marks struct {
	Id	uint   	`json:"id" gorm:"primary_key"`
	Subject    	string `json:"subject"`
	Marks		int	`json:"marks"`
}

func (m *Marks) TableName() string {
	return "marks"
}
