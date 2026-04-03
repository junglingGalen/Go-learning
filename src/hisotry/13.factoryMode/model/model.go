package model

type student struct {
	Name string
	score float64
}

//因为student是小写首字母，因此只能在model使用
//我们通过工厂模式来解决

func NewStudent(name string, score float64) *student {
	return &student{
		Name : name,
		score : score,
	}
}

func (s *student) Score() float64{
	return s.score
}