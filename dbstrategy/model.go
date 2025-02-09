package dbstrategy

const (
	GoModelTemplate = `package {{.PackageName}}

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type {{.StructName}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} ` + "`{{if .GormTag}}gorm:\"{{.GormTag}}\" {{end}}" +
		`json:"{{.JSONTag}}"{{if .Comment}} // {{.Comment}}{{end}}` + "`" + `{{end}}
}

func (t *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}

func (t *{{.StructName}}) Find(c *gin.Context, tx *gorm.DB, search *{{.StructName}}) (*{{.StructName}}, error) {
	model := &{{.StructName}}{}
	err := tx.WithContext(c).Where(search).First(model).Error
	return model, err
}

func (t *{{.StructName}}) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(t).Error
}

func (t *{{.StructName}}) Update(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Model(t).Updates(t).Error
}

func (t *{{.StructName}}) Delete(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Delete(t).Error
}
`
	JavaModelTemplate = `package {{.PackageName}};

import lombok.Data;

import java.io.Serializable;

@Data
public class {{.StructName}} implements Serializable {
    {{range .Fields}}
    private {{.Type}} {{.Name}};{{if .Comment}} // {{.Comment}}{{end}}
    {{end}}
}
`
)
