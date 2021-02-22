package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type SoporteConocimientoIdioma struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	Institucion        int                 `orm:"column(institucion)"`
	Documento          int                 `orm:"column(documento)"`
	Descripcion        string              `orm:"column(descripcion);null"`
	ConocimientoIdioma *ConocimientoIdioma `orm:"column(conocimiento_idioma);rel(fk)"`
	FechaCreacion      string              `orm:"column(fecha_creacion);null"`
	FechaModificacion  string              `orm:"column(fecha_modificacion);null"`
}

func (t *SoporteConocimientoIdioma) TableName() string {
	return "soporte_conocimiento_idioma"
}

func init() {
	orm.RegisterModel(new(SoporteConocimientoIdioma))
}

// AddSoporteConocimientoIdioma insert a new SoporteConocimientoIdioma into database and returns
// last inserted Id on success.
func AddSoporteConocimientoIdioma(m *SoporteConocimientoIdioma) (id int64, err error) {
	m.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSoporteConocimientoIdiomaById retrieves SoporteConocimientoIdioma by Id. Returns error if
// Id doesn't exist
func GetSoporteConocimientoIdiomaById(id int) (v *SoporteConocimientoIdioma, err error) {
	o := orm.NewOrm()
	v = &SoporteConocimientoIdioma{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSoporteConocimientoIdioma retrieves all SoporteConocimientoIdioma matches certain condition. Returns empty list if
// no records exist
func GetAllSoporteConocimientoIdioma(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SoporteConocimientoIdioma))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SoporteConocimientoIdioma
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSoporteConocimientoIdioma updates SoporteConocimientoIdioma by Id and returns error if
// the record to be updated doesn't exist
func UpdateSoporteConocimientoIdiomaById(m *SoporteConocimientoIdioma) (err error) {
	o := orm.NewOrm()
	v := SoporteConocimientoIdioma{Id: m.Id}
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Institucion", "Documento", "Descripcion", "ConocimientoIdioma", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSoporteConocimientoIdioma deletes SoporteConocimientoIdioma by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSoporteConocimientoIdioma(id int) (err error) {
	o := orm.NewOrm()
	v := SoporteConocimientoIdioma{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SoporteConocimientoIdioma{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
