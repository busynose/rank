package migrate

import (
	"rank/internal/app/model"

	fake "github.com/brianvoe/gofakeit/v6"
	"gitlab.com/makeblock-go/mysql"
)

const (
	matchNumber = 5
	userNumber = 100
)

func InitModel(){
	mysql.GetDB().DropTable("match","user","match_user")
	mysql.GetDB().AutoMigrate(&model.User{}, &model.Match{}, &model.MatchUser{})
}

func InitData() {
	fake.Seed(0)
	var matchIds []int64
	for i:=0;i<matchNumber;i++ {
		match := model.Match{
			Name: fake.AppName(),
		}
		mysql.GetDB().Create(&match)

		matchIds = append(matchIds, match.ID)
	}
	for i:=0;i<userNumber;i++{
		user :=  model.User{
			Name: fake.Name(),
		}
		mysql.GetDB().Create(&user)
		for _, mid := range matchIds {
			mysql.GetDB().Create(&model.MatchUser{
				UserID:  user.ID,
				MatchID: mid,
				Score:   0,
			})
		}
	}
}