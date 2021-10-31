package main

import "github.com/gin-gonic/gin"

// f) выбор текстового файла с заметками из архива файлов по разделам и его отображение
// Текстовый файл будет представлен в в виде json с следующей структурой
// id       		  - id заметки
// user				  - автор заметки
// create_datetime 	  - время создания
// last_edit_datetime - время последнего редактирования
// text				  - текст заметки в формате Markdown
// tags				  - список тегов

func main() {
	r := gin.Default()
	routes.prepareRoutes(r)
	r.Run()
}
