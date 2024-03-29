package controllers

import (
	"fmt"
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/middleware"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Products struct {
	Products     models.Product
	CategoriesId string `json:"categories_Id"`
}
type ProdDiscount struct {
	Id         int      `json:"id" gorm:"primaryKey"`
	Sku        string   `json:"sku"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Image      string   `json:"image"`
	CategoryId int      `json:"categoryId"`
	Discount   Discount `json:"discount"`
}
type Discount struct {
	Qty       int    `json:"qty"`
	Types     string `json:"type"`
	Result    int    `json:"result"`
	ExpiredAt int    `json:"expiredAt"`
}

func CreateProduct(c *fiber.Ctx) error {
	var data ProdDiscount
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Product error in post request %v", err)
	}
	var p []models.Product
	db.DB.Find(&p)

	discount := models.Discount{
		Qty:       data.Discount.Qty,
		Type:      data.Discount.Types,
		Result:    data.Discount.Result,
		ExpiredAt: data.Discount.ExpiredAt,
	}
	db.DB.Create(&discount)
	product := models.Product{
		Name:       data.Name,
		Image:      data.Image,
		CategoryId: data.CategoryId,
		DiscountId: discount.Id,
		Price:      data.Price,
		Stock:      data.Stock,
	}
	db.DB.Create(&product)

	db.DB.Table("products").Where("id = ?", product.Id).Update("sku", "SK00"+strconv.Itoa(product.Id))

	fmt.Println("--------------------------------------->")
	fmt.Println("------------Product Creation Done----------->", product.Id)
	fmt.Println("--------------------------------------->")
	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    product,
	}
	return (c.JSON(Response))

}

func GetProductDetails(c *fiber.Ctx) error {
	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	//Token authenticate

	productId := c.Params("productId")
	productsRes := make([]*models.ProductResult, 0)
	var products []models.Product
	var ratings []models.Rating
	db.DB.Where("id = ? ", productId).Find(&products)

	var averageRating float64

	var category models.Category
	var discount models.Discount
	var rating models.Rating
	for i := 0; i < len(products); i++ {
		db.DB.Where("id = ?", products[i].CategoryId).Find(&category)

		db.DB.Where("id = ?", products[i].DiscountId).Find(&discount)
		db.DB.Find(&ratings, "product_rating")
		db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)

		//productsRes =
		productsRes = append(productsRes,
			&models.ProductResult{
				Id:            products[i].Id,
				Sku:           products[i].Sku,
				Name:          products[i].Name,
				Stock:         products[i].Stock,
				Price:         products[i].Price,
				Image:         products[i].Image,
				Category:      category,
				Discount:      discount,
				ProductRating: averageRating,
			},
		)
	}

	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    productsRes,
	}
	return (c.JSON(Response))
}

func UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	var product models.Product

	db.DB.Find(&product, "id = ?", productId)

	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
			"error":   map[string]interface{}{},
		})
	}

	var updateProductData models.Product
	c.BodyParser(&updateProductData)

	if updateProductData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Product name is required",
			"error":   map[string]interface{}{},
		})
	}

	product.Name = updateProductData.Name
	product.CategoryId = updateProductData.CategoryId
	product.Image = updateProductData.Image
	product.Price = updateProductData.Price
	product.Stock = updateProductData.Stock

	db.DB.Save(&product)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    product,
	})

	//Validation rules
	//if data.Name == "" {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": "Product Name is required",
	//	})
	//}
	//
	//if data.Price <= 0 {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": "Price field is required",
	//	})
	//}
	//
	//if data.CategoryId <= 0 {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": "Category Id field is required",
	//	})
	//}
	//if data.Image == "" {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": "Image field is required",
	//	})
	//}
	//if data.Stock <= 0 {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": " Stock field is required",
	//	})
	//}

	//product := models.Product{
	//	Name:       data.Name,
	//	Image:      data.Image,
	//	CategoryId: data.CategoryId,
	//	Price:      data.Price,
	//	Stock:      data.Stock,
	//}
	//var products models.Product
	//db.DB.Model(products).Where("id = ?", productId).Updates(&product)

	//return c.Status(200).JSON(fiber.Map{
	//	"success": true,
	//	"message": "Success",
	//	"data":    product,
	//})
}

func ProductList(c *fiber.Ctx) error {
	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	//Token authenticate

	limit := c.Query("limit")
	skip := c.Query("skip")
	categoryId := c.Query("categoryId")
	productName := c.Query("name")
	price := c.Query("price")
	intLimit, _ := strconv.Atoi(limit)
	intSkip, _ := strconv.Atoi(skip)
	var products []models.Product
	var ratings []models.Rating

	productsRes := make([]*models.ProductResult, 0)

	if productName == "" {
		//filter by price
		var count int64
		var averageRating float64
		db.DB.Where("price = ?", price).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)

		var category models.Category
		var discount models.Discount
		var rating models.Rating

		for i := 0; i < len(products); i++ {

			db.DB.Table("categories").Where("id = ?", products[i].CategoryId).Find(&category)

			db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
			db.DB.Find(&ratings, "product_rating")
			db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)

			//ratingCount = int64(len(ratings))
			//db.DB.Where("id = ?", products[i].RatingId).Limit(intLimit).Offset(intSkip).Find(&rating).Count(&count)
			count = int64(len(products))
			//productsRes =
			productsRes = append(productsRes,
				&models.ProductResult{
					Id:            products[i].Id,
					Sku:           products[i].Sku,
					Name:          products[i].Name,
					Stock:         products[i].Stock,
					Price:         products[i].Price,
					Image:         products[i].Image,
					Category:      category,
					Discount:      discount,
					ProductRating: averageRating,
					Rating:        rating,
				},
			)
		}

		meta := map[string]interface{}{
			"total":  count,
			"limit":  limit,
			"skip":   skip,
			"Rating": averageRating,
		}

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Success",
			"data": map[string]interface{}{
				"products": productsRes,
				"meta":     meta,
			},
		})
		//} else if categoryId != "" {
		//	//filter by categoryId
		//	var count int64
		//	var averageRating float64
		//	db.DB.Where("category_id = ?", categoryId).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		//	//db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_rating = ?", proRating)
		//
		//	var category models.Category
		//	var discount models.Discount
		//	var rating models.Rating
		//
		//	for i := 0; i < len(products); i++ {
		//
		//		db.DB.Table("categories").Where("id = ?", products[i].CategoryId).Find(&category)
		//		//db.DB.Table("discounts").Where("id = ?", products[i].DiscountId).Find(&discount)
		//		//db.DB.Table("ratings").Where("id = ?", products[i].ProductRating).Find(&ratings)
		//
		//		db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
		//		db.DB.Find(&ratings, "product_rating")
		//		db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)
		//
		//		//ratingCount = int64(len(ratings))
		//		//db.DB.Where("id = ?", products[i].RatingId).Limit(intLimit).Offset(intSkip).Find(&rating).Count(&count)
		//		count = int64(len(products))
		//		//productsRes =
		//		productsRes = append(productsRes,
		//			&models.ProductResult{
		//				Id:            products[i].Id,
		//				Sku:           products[i].Sku,
		//				Name:          products[i].Name,
		//				Stock:         products[i].Stock,
		//				Price:         products[i].Price,
		//				Image:         products[i].Image,
		//				Category:      category,
		//				Discount:      discount,
		//				ProductRating: averageRating,
		//				Rating:        rating,
		//			},
		//		)
		//	}
		//
		//	meta := map[string]interface{}{
		//		"total":  count,
		//		"limit":  limit,
		//		"skip":   skip,
		//		"Rating": averageRating,
		//	}
		//
		//	return c.Status(200).JSON(fiber.Map{
		//		"success": true,
		//		"message": "Success",
		//		"data": map[string]interface{}{
		//			"products": productsRes,
		//			"meta":     meta,
		//		},
		//	})

		//} else if categoryId != "" {
		//	// filter by categoryId
		//	var count int64
		//	var averageRating float64
		//	db.DB.Where("category_id = ?", categoryId).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		//
		//	var category models.Category
		//	var discount models.Discount
		//	var rating models.Rating
		//
		//	for i := 0; i < len(products); i++ {
		//
		//		db.DB.Table("categories").Where("id = ?", products[i].CategoryId).Find(&category)
		//		//db.DB.Table("discounts").Where("id = ?", products[i].DiscountId).Find(&discount)
		//		//db.DB.Table("ratings").Where("id = ?", products[i].ProductRating).Find(&ratings)
		//
		//		db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
		//		db.DB.Find(&ratings, "product_rating")
		//		db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)
		//
		//		//ratingCount = int64(len(ratings))
		//		//db.DB.Where("id = ?", products[i].RatingId).Limit(intLimit).Offset(intSkip).Find(&rating).Count(&count)
		//		count = int64(len(products))
		//		productsRes = append(productsRes,
		//			&models.ProductResult{
		//				Id:            products[i].Id,
		//				Sku:           products[i].Sku,
		//				Name:          products[i].Name,
		//				Stock:         products[i].Stock,
		//				Price:         products[i].Price,
		//				Image:         products[i].Image,
		//				Category:      category,
		//				Discount:      discount,
		//				ProductRating: averageRating,
		//				Rating:        rating,
		//			},
		//		)
		//	}
		//
		//	meta := map[string]interface{}{
		//		"total":  count,
		//		"limit":  limit,
		//		"skip":   skip,
		//		"Rating": averageRating,
		//	}
		//
		//	return c.Status(200).JSON(fiber.Map{
		//		"success": true,
		//		"message": "Success",
		//		"data": map[string]interface{}{
		//			"products": productsRes,
		//			"meta":     meta,
		//		},
		//	})
	} else {

		var count int64
		var averageRating float64
		if categoryId != "" {
			db.DB.Where("category_id = ?", categoryId).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		} else {
			db.DB.Where(" name= ?", productName).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		}
		var category models.Category
		var discount models.Discount
		var rating models.Rating
		for i := 0; i < len(products); i++ {
			db.DB.Where("id = ?", products[i].CategoryId).Find(&category)
			db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
			db.DB.Find(&ratings, "product_rating")
			db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)
			count = int64(len(products))
			productsRes = append(productsRes,
				&models.ProductResult{
					Id:            products[i].Id,
					Sku:           products[i].Sku,
					Name:          products[i].Name,
					Stock:         products[i].Stock,
					Price:         products[i].Price,
					Image:         products[i].Image,
					Category:      category,
					Discount:      discount,
					ProductRating: averageRating,
				},
			)
		}

		meta := map[string]interface{}{
			"total":  count,
			"Rating": averageRating,
			"limit":  limit,
			"skip":   skip,
		}

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Success",
			"data": map[string]interface{}{
				"products": productsRes,
				"meta":     meta,
			},
		})
	}
}

//Старая функиция

//func ProductList(c *fiber.Ctx) error {
//	//Token authenticate
//	headerToken := c.Get("Authorization")
//	if headerToken == "" {
//		return c.Status(401).JSON(fiber.Map{
//			"success": false,
//			"message": "Unauthorized",
//			"error":   map[string]interface{}{},
//		})
//	}
//	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
//		return c.Status(401).JSON(fiber.Map{
//			"success": false,
//			"message": "Unauthorized",
//			"error":   map[string]interface{}{},
//		})
//	}
//	//Token authenticate
//
//	limit, _ := strconv.Atoi(c.Query("limit"))
//	skip, _ := strconv.Atoi(c.Query("skip"))
//	var count int64
//	var product []models.Product
//
//	db.DB.Select("*").Limit(limit).Offset(skip).Find(&product).Count(&count)
//
//	type ProductList struct {
//		ProductId int             `json:"productId"`
//		Sku       string          `json:"sku"`
//		Name      string          `json:"name"`
//		Stock     int             `json:"stock"`
//		Price     int             `json:"price"`
//		Image     string          `json:"image"`
//		Category  models.Category `json:"category"`
//		Discount  models.Discount `json:"discount"`
//		CreatedAt time.Time       `json:"createdAt"`
//	}
//	ProductResponse := make([]*ProductList, 0)
//
//	for _, v := range product {
//		category := models.Category{}
//		db.DB.Where("id = ?", v.CategoryId).Find(&category)
//		discount := models.Discount{}
//		db.DB.Where("id = ?", v.DiscountId).Find(&discount)
//
//		ProductResponse = append(ProductResponse, &ProductList{
//			ProductId: v.Id,
//			Sku:       v.Sku,
//			Name:      v.Name,
//			Stock:     v.Stock,
//			Price:     v.Price,
//			Image:     v.Image,
//			Category:  category,
//			Discount:  discount,
//		})
//
//	}
//
//	return c.Status(404).JSON(fiber.Map{
//		"success": true,
//		"message": "Sucess",
//		"data":    ProductResponse,
//		"meta": map[string]interface{}{
//			"total": count,
//			"limit": limit,
//			"skip":  skip,
//		},
//	})
//}

//func ProductList_Backup(c *fiber.Ctx) error {
//	//Token authenticate
//	headerToken := c.Get("Authorization")
//	if headerToken == "" {
//		return c.Status(401).JSON(fiber.Map{
//			"success": false,
//			"message": "Unauthorized",
//			"error":   map[string]interface{}{},
//		})
//	}
//	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
//		return c.Status(401).JSON(fiber.Map{
//			"success": false,
//			"message": "Unauthorized",
//			"error":   map[string]interface{}{},
//		})
//	}
//	//Token authenticate
//
//	limit := c.Query("limit")
//	skip := c.Query("skip")
//	categoryId := c.Query("categoryId")
//	productName := c.Query("q")
//	intLimit, _ := strconv.Atoi(limit)
//	intSkip, _ := strconv.Atoi(skip)
//	var products []models.Product
//
//	productsRes := make([]*models.ProductResult, 0)
//
//	fmt.Println(productName)
//	if productName == "" {
//		var count int64
//		fmt.Println("inside check")
//		db.DB.Where("category_Id = ?", categoryId).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
//
//		var category models.Category
//		var discount models.Discount
//		for i := 0; i < len(products); i++ {
//			db.DB.Table("categories").Where("id = ?", products[i].CategoryId).Find(&category)
//			db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
//			fmt.Println("for loop inside count:", count)
//			//productsRes =
//			productsRes = append(productsRes,
//				&models.ProductResult{
//					Id:       products[i].Id,
//					Sku:      products[i].Sku,
//					Name:     products[i].Name,
//					Stock:    products[i].Stock,
//					Price:    products[i].Price,
//					Image:    products[i].Image,
//					Category: category,
//					Discount: discount,
//				},
//			)
//		}
//		fmt.Println("if count:", count)
//		meta := map[string]interface{}{
//			"total": count,
//			"limit": limit,
//			"skip":  skip,
//		}
//		return c.Status(200).JSON(fiber.Map{
//			"success": true,
//			"message": "Success",
//			"data":    productsRes,
//			"meta":    meta,
//		})
//	} else {
//		var count int64
//
//		//
//		db.DB.Where("category_Id = ? AND name= ?", categoryId, productName).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
//
//		var category models.Category
//		var discount models.Discount
//		for i := 0; i < len(products); i++ {
//			db.DB.Where("id = ?", products[i].CategoryId).Find(&category)
//
//			db.DB.Where("id = ?", products[i].DiscountId).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
//
//			//productsRes =
//			productsRes = append(productsRes,
//				&models.ProductResult{
//					Id:       products[i].Id,
//					Sku:      products[i].Sku,
//					Name:     products[i].Name,
//					Stock:    products[i].Stock,
//					Price:    products[i].Price,
//					Image:    products[i].Image,
//					Category: category,
//					Discount: discount,
//				},
//			)
//		}
//
//		fmt.Println("else count:", count)
//
//		meta := map[string]interface{}{
//			"total": count,
//			"limit": limit,
//			"skip":  skip,
//		}
//
//		return c.Status(200).JSON(fiber.Map{
//			"success": true,
//			"message": "Success",
//			"data":    productsRes,
//			"meta":    meta,
//		})
//	}
//}

// Done
func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	var product models.Product

	db.DB.First(&product, productId)
	if product.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Delete(&product)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
	})
}
