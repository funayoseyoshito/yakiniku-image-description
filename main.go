package main

import (
	"fmt"

	"github.com/funayoseyoshito/yakiniku-image-description/db"
	"github.com/funayoseyoshito/yakiniku-image-description/lib"
)

func main() {
	dbSet := db.NewDatabaseSet(
		lib.Config.Database.User,
		lib.Config.Database.Password,
		lib.Config.Database.Host,
		lib.Config.Database.Port,
		lib.Config.Database.Name)

	var image db.Images
	defer dbSet.Connection().Close()

	var cnt int = 0
	for i := 0; ; {
		rows := dbSet.SelectProcessingRows(i)

		if !rows.Next() {
			break
		}

		for {
			dbSet.Connection().ScanRows(rows, &image)
			if image.Description != "" {

				fmt.Println(image.Description)
				dbSet.Connection().Create(
					&db.ImageDescriptions{
						OriginID:    image.OriginID,
						Description: image.Description})
				cnt++
			}

			if !rows.Next() {
				break
			}
		}
		i += db.SelectLimit
	}
	fmt.Println("処理件数", cnt)
}
