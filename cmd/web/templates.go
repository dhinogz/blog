package main

import (
	"blog.dhinojosa.io/internal/models"
)

type templateData struct {
	Article  *models.Article
	Articles []*models.Article
}
