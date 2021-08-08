package controller

import (
	"encoding/json"
	"fmt"
	"focusapi/model"
	"focusapi/service"
	"focusapi/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// controller struct
type InstanceController struct {
	apiVersion string
	Service    *service.InstanceService
}

// get controller
func (uc *InstanceController) getCtl() *InstanceController {
	var svc *service.InstanceService
	return &InstanceController{"v1", svc}
}

// get all apis
func (uc *InstanceController) GetAllInstances(c *gin.Context) {
	var currentPageInt, pageSizeInt = util.CURRENT_PAGE, util.PAGE_SIZE
	var totalRows, totalPages int64
	pageSizeInt = viper.GetInt("PAGE_SIZE")
	currentPage, cpExist := c.GetQuery("currentpage")
	if cpExist {
		currentPageInt, _ = strconv.Atoi(currentPage)
	}

	pageSize, psExist := c.GetQuery("pagesize")
	if psExist {
		pageSizeInt, _ = strconv.Atoi(pageSize)
	}

	// data option setting
	dataOrder, dataOrderExist := c.GetQuery("dataOrder")
	if !dataOrderExist {
		dataOrder = "id desc"
	}

	dataSelect, dataSelectExist := c.GetQuery("dataSelect")
	if !dataSelectExist {
		dataSelect = ""
	}

	dataWhereMap := map[string]interface{}{}
	dataWhere, dataWhereExist := c.GetQuery("dataWhere")
	if dataWhereExist {
		err := json.Unmarshal([]byte(dataWhere), &dataWhereMap)
		if err != nil {
			util.SendError(c, err.Error())
			return
		}
	}

	dataLimitInt := 0
	dataLimit, dataLimitExist := c.GetQuery("dataLimit")
	if dataLimitExist {
		dataLimitInt, _ = strconv.Atoi(dataLimit)
	}

	daoOpt := model.DAOOption{
		Select: dataSelect,
		Order:  dataOrder,
		Where:  dataWhereMap, //map[string]interface{}{},
		Limit:  dataLimitInt,
	}

	apis, err := uc.getCtl().Service.FindAllInstanceByPagesWithKeys(
		map[string]interface{}{},
		map[string]interface{}{},
		currentPageInt,
		pageSizeInt,
		&totalRows,
		daoOpt)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	c.HTML(http.StatusOK, "apis.tmpl", gin.H{
		"title": "用户管理",
		"apis":  apis,
		"pages": util.Pagination{
			PageSize:        pageSizeInt,
			CurrentPage:     currentPageInt,
			TotalRows:       totalRows,
			TotalPages:      totalPages,
			PreCurrentPage:  currentPageInt - 1,
			NextCurrentPage: currentPageInt + 1,
		},
	})
}

// search apis by pages with keys
func (uc *InstanceController) SearchInstancesByKeys(c *gin.Context) {
	var currentPageInt, pageSizeInt = util.CURRENT_PAGE, util.PAGE_SIZE
	var totalRows, totalPages int64
	pageSizeInt = viper.GetInt("PAGE_SIZE")
	currentPage, cpExist := c.GetQuery("currentpage")
	if cpExist {
		currentPageInt, _ = strconv.Atoi(currentPage)
	}

	pageSize, psExist := c.GetQuery("pagesize")
	if psExist {
		pageSizeInt, _ = strconv.Atoi(pageSize)
	}

	//search option
	keys := make(map[string]interface{})
	keyOpts := make(map[string]interface{})
	//get search key
	searchKey, searchKeyExist := c.GetQuery("searchKey")
	if searchKeyExist {
		keys["name"] = searchKey
		keys["email"] = searchKey
		keys["age"] = searchKey
		keys["role"] = searchKey
	} else {
		//search key
		name, nameExist := c.GetQuery("name")
		if nameExist {
			keys["name"] = name
		}

		email, emailExist := c.GetQuery("email")
		if emailExist {
			keys["email"] = email
		}

		age, ageExist := c.GetQuery("age")
		if ageExist {
			keys["age"] = age
		}

		role, roleExist := c.GetQuery("role")
		if roleExist {
			keys["role"] = role
		}
	}

	//search value options
	searchKeyOpt, searchKeyOptExist := c.GetQuery("searchKeyOpt")
	if searchKeyOptExist {
		keyOpts["searchKey"] = searchKeyOpt
	}

	nameOpt, nameOptExist := c.GetQuery("nameOpt")
	if nameOptExist {
		keyOpts["name"] = nameOpt
	}

	emailOpt, emailOptExist := c.GetQuery("emailOpt")
	if emailOptExist {
		keyOpts["email"] = emailOpt
	}

	ageOpt, ageOptExist := c.GetQuery("ageOpt")
	if ageOptExist {
		keyOpts["age"] = ageOpt
	}

	roleOpt, roleOptExist := c.GetQuery("roleOpt")
	if roleOptExist {
		keyOpts["role"] = roleOpt
	}

	// data option setting
	dataOrder, dataOrderExist := c.GetQuery("dataOrder")
	if !dataOrderExist {
		dataOrder = "id desc"
	}

	dataSelect, dataSelectExist := c.GetQuery("dataSelect")
	if !dataSelectExist {
		dataSelect = ""
	}

	dataWhereMap := map[string]interface{}{}
	dataWhere, dataWhereExist := c.GetQuery("dataWhere")
	if dataWhereExist {
		err := json.Unmarshal([]byte(dataWhere), &dataWhereMap)
		if err != nil {
			util.SendError(c, err.Error())
			return
		}
	}

	dataLimitInt := 0
	dataLimit, dataLimitExist := c.GetQuery("dataLimit")
	if dataLimitExist {
		dataLimitInt, _ = strconv.Atoi(dataLimit)
	}

	daoOpt := model.DAOOption{
		Select: dataSelect,
		Order:  dataOrder,
		Where:  dataWhereMap, //map[string]interface{}{},
		Limit:  dataLimitInt,
	}

	apis, err := uc.getCtl().Service.SearchInstanceByPagesWithKeys(keys,
		keyOpts,
		currentPageInt,
		pageSizeInt,
		&totalRows,
		daoOpt)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	c.HTML(http.StatusOK, "apis.tmpl", gin.H{
		"title": "用户管理",
		"apis":  apis,
		"pages": util.Pagination{
			PageSize:        pageSizeInt,
			CurrentPage:     currentPageInt,
			TotalRows:       totalRows,
			TotalPages:      totalPages,
			PreCurrentPage:  currentPageInt - 1,
			NextCurrentPage: currentPageInt + 1,
		},
	})
}

// get api by id
func (uc *InstanceController) GetInstanceByID(c *gin.Context) {
	id, exists := c.Params.Get("id")
	if !exists {
		util.SendError(c, "id is null")
		return
	}

	idUint64, errConv := strconv.ParseUint(id, 10, 64)
	if errConv != nil {
		util.SendError(c, "id conv failed")
		return
	}

	api, err := uc.getCtl().Service.FindInstanceById(idUint64)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    api,
	})
}

// add api tmpl
func (uc *InstanceController) AddInstance(c *gin.Context) {
	c.HTML(http.StatusOK, "api_add.tmpl", gin.H{
		"title": "添加 API ",
	})
}

// create api
func (uc *InstanceController) CreateInstance(c *gin.Context) {
	name := c.PostForm("name")
	status := 1
	api := &model.Instance{
		Name:      name,
		Status:    status,
		BaseModel: model.BaseModel{},
	}

	errCreate := uc.getCtl().Service.CreateInstance(api)
	if errCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errCreate,
		})
	}

	//redirect
	c.Redirect(http.StatusMovedPermanently, "/apis")
}

// update api
func (uc *InstanceController) UpdateInstance(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}

	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		util.SendError(c, "get uid error !")
		return
	}

	api, err := uc.getCtl().Service.FindInstanceById(uidUint64)
	if err != nil {
		panic(" get api error !")
	}

	//update data
	updateDataEnabled, updateDataExist := c.GetPostForm("updatedata")
	if updateDataExist && updateDataEnabled == "true" {
		c.HTML(http.StatusOK, "api_update.tmpl", gin.H{
			"title": "更新 API",
			"api":   api,
		})
		return
	}

	api.ID = uidUint64
	name, nameExist := c.GetPostForm("name")
	if nameExist {
		api.Name = name
	}

	status, statusExist := c.GetPostForm("status")
	if statusExist {
		api.Status, _ = strconv.Atoi(status)
	}

	//update api
	rowsAffected, updateErr := uc.getCtl().Service.UpdateInstance(uidUint64, api)
	if updateErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": updateErr,
		})
	}

	log.Println("rows affected: ", rowsAffected)
	//redirect
	c.Redirect(http.StatusMovedPermanently, "/apis")
}

func (uc *InstanceController) DeleteInstance(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}
	fmt.Println("uid", uid)
	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	rowsAffected, delErr := uc.getCtl().Service.DeleteInstance(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete api error",
		})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": rowsAffected,
	// })

	log.Println("rows affected: ", rowsAffected)
	//redirect
	c.Redirect(http.StatusMovedPermanently, "/apis")
}
