package handler

import (
  "fmt"
  "net/http"
  interfaces "project/pkg/helper/interface"
  services "project/pkg/usecase/interface"
  response "project/pkg/utils/Response"
  models "project/pkg/utils/models"
  "strconv"

  "github.com/gin-gonic/gin"
)

type AdminHandler struct {
  adminUseCase services.AdminUseCase
  helper       interfaces.Helper
}

func NewAdminHandler(usecase services.AdminUseCase, h interfaces.Helper) *AdminHandler {
  return &AdminHandler{
    adminUseCase: usecase,
    helper:       h,
  }
}

// ADMIN_LOGIN

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
  var adminDetails models.AdminLogin
  if err := c.BindJSON(&adminDetails); err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "details not in the correct format", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  admin, err := ad.adminUseCase.LoginHandler(adminDetails)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }
  c.Set("Access", admin.AccessToken)
  // c.Set("Refresh", admin.RefreshToken)

  successRes := response.ClientResponse(http.StatusOK, "Admin authenticated succesfully", admin, nil)
  c.JSON(http.StatusOK, successRes)
}

// DISPLAY USERS

func (ad *AdminHandler) GetUsers(c *gin.Context) {

  users, err := ad.adminUseCase.GetUsers()
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve details", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }
  successRes := response.ClientResponse(http.StatusOK, "successfully retrived the users", users, nil)
  c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) BlockUser(c *gin.Context) {
  id := c.Query("id")
  userID, err := strconv.Atoi(id)
  if err != nil {
    response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
  }
  err = ad.adminUseCase.BlockUser(userID)

  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return

  }

  successRes := response.ClientResponse(http.StatusOK, "successfully blocked the user ", nil, nil)
  c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
  id := c.Query("id")
  userID, err := strconv.Atoi(id)
  if err != nil {
    response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
  }
  err = ad.adminUseCase.UnBlockUser(userID)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }
  successRess := response.ClientResponse(http.StatusOK, "successfully unblocked the user", nil, nil)
  c.JSON(http.StatusOK, successRess)

}

func (ad *AdminHandler) OrderReturnApprove(c *gin.Context) {

  idstr := c.Query("order_id")

  orderID, err := strconv.Atoi(idstr)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  err = ad.adminUseCase.OrderReturnApprove(orderID)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error approving order status", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "successfully approved order return", nil, nil)
  c.JSON(http.StatusOK, successRes)

}

func (ad *AdminHandler) GetAllOrderDetails(c *gin.Context) {

  orders, err := ad.adminUseCase.GetAllOrderDetails()
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error fetching orders", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "successfully retrieved the orders", orders, nil)
  c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) EditOrderStatus(c *gin.Context) {

  idStr := c.Query("order_id")
  status := c.Query("status")

  orderID, err := strconv.Atoi(idStr)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error strconv Atoi", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  err = ad.adminUseCase.EditOrderStatus(orderID, status)
  if err != nil {
    errREs := response.ClientResponse(http.StatusBadRequest, "error editing order status", nil, err.Error())
    c.JSON(http.StatusBadRequest, errREs)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "successfully edited order status", nil, nil)
  c.JSON(http.StatusOK, successRes)

}

func (ad *AdminHandler) MakePaymentStatusAsPaid(c *gin.Context) {

  idStr := c.Query("order_id")

  orderID, err := strconv.Atoi(idStr)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error strconv Atoi", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  err = ad.adminUseCase.MakePaymentStatusAsPaid(orderID)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error updating payment-status", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "successfully updated payment-status as PAID", nil, nil)
  c.JSON(http.StatusOK, successRes)

}

func (ad *AdminHandler) NewPaymentHandler(c *gin.Context) {

  pMethod := c.Query("payment_method")

  err := ad.adminUseCase.NewPaymentMethod(pMethod)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error adding new payment method", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }
  successRes := response.ClientResponse(http.StatusBadRequest, "succesfully added new payment", nil, nil)
  c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) GetPaymentMethods(c *gin.Context) {

  data, err := ad.adminUseCase.GetAllPaymentMethods()
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error fetching the payment methods", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "succesfully fetched all payment methods ", data, nil)
  c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) DeletePaymentMethod(c *gin.Context) {

  idStr := c.Query("payment_id")

  paymentID, err := strconv.Atoi(idStr)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  err = ad.adminUseCase.DeletePaymentMethod(paymentID)
  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error deleting paymentmethod", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  successRes := response.ClientResponse(http.StatusOK, "succesfully deleted this payment method ", nil, nil)
  c.JSON(http.StatusOK, successRes)

}

func (ah *AdminHandler) FilteredSalesReport(c *gin.Context) {

  timePeriod := c.Query("period")
  salesReport, err := ah.adminUseCase.FilteredSalesReport(timePeriod)
  if err != nil {
    errorRes := response.ClientResponse(http.StatusInternalServerError, "error fetching filtered salesReport ", nil, err.Error())
    c.JSON(http.StatusInternalServerError, errorRes)
    return
  }
  message := "Successfully retrieved current " + timePeriod + " Data"

  success := response.ClientResponse(http.StatusOK, message, salesReport, nil)
  c.JSON(http.StatusOK, success)
}

func (a *AdminHandler) PrintSalesByDate(c *gin.Context) {
  year := c.Query("year")
  yearInt, err := strconv.Atoi(year)

  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, " error year conversion from string", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  month := c.Query("month")
  monthInt, err := strconv.Atoi(month)

  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error month conversion from string", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  day := c.Query("day")
  dayInt, err := strconv.Atoi(day)

  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "error day conversion from string ", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }

  body, err := a.adminUseCase.SalesByDate(dayInt, monthInt, yearInt)

  if err != nil {
    errRes := response.ClientResponse(http.StatusBadRequest, "failed to fetch sales data", nil, err.Error())
    c.JSON(http.StatusBadRequest, errRes)
    return
  }
  fmt.Println("body", body)

  download := c.Query("download")
  if download == "pdf" {
    pdf, err := a.adminUseCase.PrintSalesReport(body)
    if err != nil {
      errRes := response.ClientResponse(http.StatusBadGateway, "error fetching sales data", nil, err.Error())
      c.JSON(http.StatusBadRequest, errRes)
      return
    }
    c.Header("Content-Disposition", "attachment;filename=totalsalesreport.pdf")

    pdfFilePath := "../salesReport/totalsalesreport.pdf"

    err = pdf.OutputFileAndClose(pdfFilePath)
    if err != nil {
      errRes := response.ClientResponse(http.StatusBadGateway, "error opening file", nil, err.Error())
      c.JSON(http.StatusBadRequest, errRes)
      return
    }

    c.Header("Content-Disposition", "attachment; filename=total_sales_report.pdf")
    c.Header("Content-Type", "application/pdf")

    c.File(pdfFilePath)

    c.Header("Content-Type", "application/pdf")

    err = pdf.Output(c.Writer)
    if err != nil {
      errRes := response.ClientResponse(http.StatusBadGateway, "error Printing", nil, err.Error())
      c.JSON(http.StatusBadRequest, errRes)
      return
    }
  } else {

    excel, err := a.helper.ConvertToExel(body)
    if err != nil {
      errRes := response.ClientResponse(http.StatusBadGateway, "error Printing SAles Report", nil, err.Error())
      c.JSON(http.StatusBadRequest, errRes)
      return
    }

    fileName := "sales_report.xlsx"

    c.Header("Content-Disposition", "attachment; filename="+fileName)
    c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

    if err := excel.Write(c.Writer); err != nil {
      errRes := response.ClientResponse(http.StatusBadGateway, "Error in serving the sales report", nil, err.Error())
      c.JSON(http.StatusBadRequest, errRes)
      return
    }
  }

  succesRes := response.ClientResponse(http.StatusOK, "successFully printed sales report", body, nil)
  c.JSON(http.StatusOK, succesRes)
}

func (ad *AdminHandler)ChangePassword (c *gin.Context){

  var adminPassChange  models.AdminPasswordChange


  if err:=c.BindJSON(&adminPassChange);err!=nil{
  
    errRes:=response.ClientResponse(http.StatusBadRequest,"error BindingJson Invalid Format",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)

    return
  }
  id,exist:= c.Get("id")
  fmt.Println("LOGG ADMIN JWT ID",id)

  if !exist{
    
    errRes:=response.ClientResponse(http.StatusBadRequest,"error getting admin Id",nil,fmt.Errorf("No admin id exist"))
    c.JSON(http.StatusInternalServerError,errRes)
     return 
  }
    
  err:=ad.adminUseCase.ChangePassword(adminPassChange,id.(int))
  if err!=nil{

    errRes:=response.ClientResponse(http.StatusBadRequest,"error BindingJson Invalid Format",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)
    return

  }

  successRes:=response.ClientResponse(http.StatusOK,"successfully changed your password",nil,nil)
  c.JSON(http.StatusOK,successRes)


}
