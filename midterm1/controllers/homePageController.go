package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {

	welcomeMessage := "Welcome to home page!"
	examleMessage := "Use links like this: http://127.0.0.1:8080/cashiers/  or http://localhost:8080/cashiers/"
	loginPage := "/cashiers/:cashierId/login"
	logoutPage := "/cashiers/:cashierId/logout"
	getCashiersList := "/cashiers"
	getCashierById := "/cashiers/:cashierId"
	getCategories := "/categories"
	getCategoriesById := "/categories/:categoryId"
	getProducts := "/products"
	getProductsById := "/products/:productId"
	getPayments := "/payments"
	getPaymentsById := "/payments/:paymentId"
	getOrders := "/orders"
	getOrdersById := "/orders/:orderId"
	getComments := "/comments"
	getRatings := "/ratings"

	rememberMessage := "Do not forget to change method GET, POST, DELETE, PUT, when you create requests on Postman!"

	return c.Status(200).JSON(fiber.Map{
		"success":                  true,
		"A. WelcomeMessage":        welcomeMessage,
		"B. Example":               examleMessage,
		"C. Do not forget!":        rememberMessage,
		"D. POST method To login":  loginPage,
		"E. POST method To logout": logoutPage,
		"F. POST,GET method To create cashier and get cashier list":                                    getCashiersList,
		"G. GET,PUT,DELETE method To get cashier by id, update cashier by id and delete cashier by id": getCashierById,
		"H. POST,GET method To create category and get category":                                       getCategories,
		"I. GET method To get category by id":                                                          getCategoriesById,
		"J. POST,GET method To create product and get products":                                        getProducts,
		"K. GET method To get product by id":                                                           getProductsById,
		"L. POST,GET method To create payments and get payments":                                       getPayments,
		"M' GET method To get payments by id":                                                          getPaymentsById,
		"N. POST, GET method To create order and get orders":                                           getOrders,
		"O. GET method To get order by id":                                                             getOrdersById,
		"P. POST,GET method To create comments and get comments":                                       getComments,
		"Q. POST,GET method To create ratings and get ratings":                                         getRatings,
	})

}
