// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"actionflow/model"
	"actionflow/util"
)

// GetAccount godoc
// @Summary Get account
// @Description get account by ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /accounts/{id} [get]
func (c *Controller) GetAccount(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	account, err := model.AccountOne(id)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}
