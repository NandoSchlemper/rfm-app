package db

import (
    "rfmapp/utils"
)

func InstanceDB() {
    db_url := utils.Env
    println(db_url)
}
