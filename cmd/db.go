package cmd

import (
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/cryption"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"time"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Run server",
	Long: `With this command you can
	up : create database table
	down :  drop database table
	seed :  insert data table
	`,
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up table",
	Long:  "Up table",
	Run: func(cmd *cobra.Command, args []string) {
		up()
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down table",
	Long:  "remove public schema, create public schema, restore the default grants",
	Run: func(cmd *cobra.Command, args []string) {
		down()
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed data table",
	Long:  "Seed data table",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Down, up, seed table",
	Long:  "Down, up, seed table",
	Run: func(cmd *cobra.Command, args []string) {
		down()
		up()
		seed()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(upCmd)
	dbCmd.AddCommand(downCmd)
	dbCmd.AddCommand(resetCmd)
	dbCmd.AddCommand(seedCmd)
}

func up() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	// table
	err = conn.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Usercompany{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Transaction{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Player{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Gor{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Game{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Gameplayer{})
	if err != nil {
		panic(err)
	}

	// view
	vUser := conn.Model(&model.User{}).
		Select("users.*, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join users u1 on u1.id = users.create_by").
		Joins("left join users u2 on u2.id = users.update_by").
		Joins("left join users u3 on u3.id = users.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_USER, gorm.ViewOption{
		Replace: true,
		Query:   vUser,
	})
	if err != nil {
		panic(err)
	}

	vCompany := conn.Model(&model.Company{}).
		Select("companies.*, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name, (select count(*) from gors g where g.company_id = companies.id and g.delete_dt is null) AS total_gor, (select count(*) from players p where p.company_id = companies.id and p.delete_dt is null) AS total_player").
		Joins("left join users u1 on u1.id = companies.create_by").
		Joins("left join users u2 on u2.id = companies.update_by").
		Joins("left join users u3 on u3.id = companies.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_COMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vCompany,
	})
	if err != nil {
		panic(err)
	}

	vUsercompany := conn.Model(&model.Usercompany{}).
		Select("usercompanies.*, companies.name as company_name, users.fullname as user_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join companies companies on companies.id = usercompanies.company_id").
		Joins("left join users users on users.id = usercompanies.user_id").
		Joins("left join users u1 on u1.id = usercompanies.create_by").
		Joins("left join users u2 on u2.id = usercompanies.update_by").
		Joins("left join users u3 on u3.id = usercompanies.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_USERCOMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vUsercompany,
	})
	if err != nil {
		panic(err)
	}

	vTransaction := conn.Model(&model.Transaction{}).
		Select("transactions.*, companies.name as company_name, u1.fullname as create_name").
		Joins("left join companies companies on companies.id = transactions.company_id").
		Joins("left join users u1 on u1.id = transactions.create_by")

	err = conn.Migrator().CreateView(model.VIEW_TRANSACTION, gorm.ViewOption{
		Replace: true,
		Query:   vTransaction,
	})
	if err != nil {
		panic(err)
	}

	vGor := conn.Model(&model.Gor{}).
		Select("gors.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join companies companies on companies.id = gors.company_id").
		Joins("left join users u1 on u1.id = gors.create_by").
		Joins("left join users u2 on u2.id = gors.update_by").
		Joins("left join users u3 on u3.id = gors.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_GOR, gorm.ViewOption{
		Replace: true,
		Query:   vGor,
	})
	if err != nil {
		panic(err)
	}

	vGame := conn.Model(&model.Game{}).
		Select("games.*, gors.name as gor_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join gors gors on gors.id = games.gor_id").
		Joins("left join users u1 on u1.id = games.create_by").
		Joins("left join users u2 on u2.id = games.update_by").
		Joins("left join users u3 on u3.id = games.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_GAME, gorm.ViewOption{
		Replace: true,
		Query:   vGame,
	})
	if err != nil {
		panic(err)
	}

	vPlayer := conn.Model(&model.Player{}).
		Select("players.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join companies companies on companies.id = players.company_id").
		Joins("left join users u1 on u1.id = players.create_by").
		Joins("left join users u2 on u2.id = players.update_by").
		Joins("left join users u3 on u3.id = players.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_PLAYER, gorm.ViewOption{
		Replace: true,
		Query:   vPlayer,
	})
	if err != nil {
		panic(err)
	}

	vGameplayer := conn.Model(&model.Gameplayer{}).
		Select("gameplayers.*, players.name as player_name, games.name as game_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join players players on players.id = gameplayers.player_id").
		Joins("left join games games on games.id = gameplayers.game_id").
		Joins("left join users u1 on u1.id = gameplayers.create_by").
		Joins("left join users u2 on u2.id = gameplayers.update_by").
		Joins("left join users u3 on u3.id = gameplayers.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_GAMEPLAYER, gorm.ViewOption{
		Replace: true,
		Query:   vGameplayer,
	})
	if err != nil {
		panic(err)
	}

}

// remove public schema
func down() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Exec("DROP SCHEMA public CASCADE").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("CREATE SCHEMA public").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO postgres").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO public").Error
	if err != nil {
		panic(err)
	}
}

func seed() {
	now := time.Now()
	password, err := cryption.EncryptAES64("123456")
	if err != nil {
		panic(err)
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	userID := utils.GetUniqueID()
	btcCompanyID := utils.GetUniqueID()
	blpCompanyID := utils.GetUniqueID()

	users := []model.User{
		{
			ID:          userID,
			Role:        constant.RoleAdmin,
			Email:       "jihanlugas2@gmail.com",
			Username:    "jihanlugas",
			NoHp:        "6287770333043",
			Fullname:    "Jihan Lugas",
			Passwd:      password,
			PassVersion: 1,
			IsActive:    true,
			PhotoID:     "",
			LastLoginDt: nil,
			CreateBy:    userID,
			CreateDt:    now,
			UpdateBy:    userID,
			UpdateDt:    now,
		},
	}
	tx.Create(&users)

	companies := []model.Company{
		{ID: blpCompanyID, Name: "BLP Pekanbaru", Description: "BLP Pekanbaru Company", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{ID: btcCompanyID, Name: "BTC Pekanbaru", Description: "BTC Pekanbaru Company", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "BTC Bandung", Description: "BTC Bandung Company", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Djarum", Description: "Persatuan Badminton Djarum", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Gudang Garam", Description: "Persatuan Badminton Gudang Garam", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Sampurna", Description: "Persatuan Badminton Sampurna", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Blur", Description: "Persatuan Badminton Blur", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Panam", Description: "Persatuan Badminton Panam", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Dragon Ball", Description: "Persatuan Badminton Dragon Ball", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Anak Mama", Description: "Persatuan Badminton Anak Mama", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Bangsa Panam", Description: "Persatuan Badminton Bangsa Panam", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Dunia Sukses", Description: "Persatuan Badminton Dunia Sukses", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Cloudeka", Description: "Persatuan Badminton Cloudeka", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB New World", Description: "Persatuan Badminton New World", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{Name: "PB Konoha", Description: "Persatuan Badminton Konoha", Balance: 0, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
	}
	tx.Create(&companies)

	//usercompanies := []model.Usercompany{
	//	{
	//		UserID:           userID,
	//		CompanyID:        companyID,
	//		IsDefaultCompany: true,
	//		IsCreator:        true,
	//		CreateBy:         userID,
	//		CreateDt:         now,
	//		UpdateBy:         userID,
	//		UpdateDt:         now,
	//		DeleteBy:         "",
	//		DeleteDt:         nil,
	//	},
	//}
	//tx.Create(&usercompanies)

	players := []model.Player{
		{CompanyID: btcCompanyID, Name: "Monkey D. Luffy", Email: "luffy@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: btcCompanyID, Name: "Roronoa Zoro", Email: "zoro@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: btcCompanyID, Name: "Sakazuki", Email: "sakazuki@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: btcCompanyID, Name: "Isho", Email: "isho@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: blpCompanyID, Name: "Itadori Yuji", Email: "yuji@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: blpCompanyID, Name: "Fushiguro Megumi", Email: "megumi@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: blpCompanyID, Name: "Ryomen Sukuna", Email: "sukuna@gmail.com", NoHp: "", Address: "", IsActive: true, PhotoID: "", CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
	}
	tx.Create(&players)

	gors := []model.Gor{
		{CompanyID: btcCompanyID, Name: "Gor Wahyu", Description: "Gor Wahyu Gobah", Address: "Jl. Sumatra", NormalGamePrice: 8000, RubberGamePrice: 11000, BallPrice: 3000, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: btcCompanyID, Name: "Gor PRS", Description: "Gor Panam Raya Square", Address: "Jl. HR. Subrantas", NormalGamePrice: 7000, RubberGamePrice: 10000, BallPrice: 3000, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
		{CompanyID: blpCompanyID, Name: "Gor Wahyu", Description: "Gor Wahyu Gobah", Address: "Jl. Sumatra", NormalGamePrice: 7000, RubberGamePrice: 10000, BallPrice: 3000, CreateBy: userID, CreateDt: now, UpdateBy: userID, UpdateDt: now},
	}
	tx.Create(&gors)

	err = tx.Commit().Error
	if err != nil {
		panic(err)
	}

}
