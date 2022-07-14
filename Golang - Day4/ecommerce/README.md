# Setup Guide


# 1. Download

Download repo to your local machine:
 > git clone https://github.com/thepanshuyadav/bootcamp.git


# 2. Database

1. **Create schema**
	>CREATE SCHEMA `ecommerce` ;
2. **Create tables**

	2.1 Customer
	>CREATE TABLE `ecommerce`.`customer` (
  `customer_id` VARBINARY(45) NOT NULL,
  `customer_name` VARCHAR(45) NOT NULL,
  `customer_address` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`customer_id`));

	2.2 Retailer
	>CREATE TABLE `ecommerce`.`retailer` (
  `retailer_id` VARBINARY(45) NOT NULL,
  `retailer_name` VARCHAR(45) NOT NULL,
  `retailer_rating` INT NOT NULL,
  PRIMARY KEY (`retailer_id`));

	2.3 Product
	>CREATE TABLE `ecommerce`.`product` (
  `product_id` VARBINARY(45) NOT NULL,
  `product_name` VARCHAR(45) NOT NULL,
  `product_price` INT NOT NULL,
  `product_quantity` INT NOT NULL,
  `retailer_id` VARBINARY(45) NULL,
  PRIMARY KEY (`product_id`),
  INDEX `retailer_id_idx` (`retailer_id` ASC) VISIBLE,
  CONSTRAINT `retailer_id`
    FOREIGN KEY (`retailer_id`)
    REFERENCES `ecommerce`.`retailer` (`retailer_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

	2.4 Order
	>CREATE TABLE `ecommerce`.`order` (
  `order_id` VARBINARY(45) NOT NULL,
  `customer_id` VARBINARY(45) NOT NULL,
  PRIMARY KEY (`order_id`),
  INDEX `customer_id_idx` (`customer_id` ASC) VISIBLE,
  CONSTRAINT `customer_id`
    FOREIGN KEY (`customer_id`)
    REFERENCES `ecommerce`.`customer` (`customer_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

	2.5 Order detail

	>CREATE TABLE `ecommerce`.`order_detail` (
  `order_id` VARBINARY(45) NOT NULL,
  `product_id` VARBINARY(45) NOT NULL,
  `quantity_ordered` INT NOT NULL,
  `order_status` VARCHAR(45) NOT NULL,
  INDEX `order_id_idx` (`order_id` ASC) VISIBLE,
  INDEX `product_id_idx` (`product_id` ASC) VISIBLE,
  CONSTRAINT `order_id`
    FOREIGN KEY (`order_id`)
    REFERENCES `ecommerce`.`order` (`order_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `product_id`
    FOREIGN KEY (`product_id`)
    REFERENCES `ecommerce`.`product` (`product_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);


# 3. Go Packages
GORM
> github.com/jinzhu/gorm

MySQL Driver
> github.com/go-sql-driver/mysql

Gin Web Framework
> github.com/gin-gonic/gin

UUID
> github.com/google/uuid

ozzo-validation
> github.com/go-ozzo/ozzo-validation

Testify
> github.com/stretchr/testify/

mockery
> github.com/vektra/mockery

## Run Unit Tests

	go test ./... -v -coverpkg=./...

# 4. Go Routes

### Authentication
Use Basic Auth as authentication header while making API calls
Ex credentials `"customer1": "pass"`

### Customer Routes
1. **GET** : Find customer using uuid

	URL : */customer/:id*
	> Ex  /customer/43042c47-c9f2-431d-b8e2-effe36e10578
	
	Response Body
	```
	{

		"customer_id":  "43042c47-c9f2-431d-b8e2-effe36e10578",

		"customer_name":  "Name",

		"customer_address":  "Home"

	}
	```

2. **POST** : Add customer 

	URL : */customer/add*
Request Body
	```
	{

		"customer_name":"Name",

		"customer_address":"Home"

	}
	```


	Response Body
	```
	{

		"customer_id":  "43042c47-c9f2-431d-b8e2-effe36e10578",

		"customer_name":  "Name",

		"customer_address":  "Home"

	}
	```

### Retailer Routes
Similar to customer routes
1. **GET** : Find retailer using uuid
	URL : */retailer/:id*
	
2. **POST** : Add retailer
	URL : */retailer/add*

### Product Routes
1. **GET** : Find product using uuid
	URL : */product/:id*
	
2. **GET** : Get all products
	URL : */product/all*
	
3. **POST** : Add new product
	URL : */product/add*

4. **PATCH** : Update product
URL : */product/:id*
Ex 
	> /product/bc446c6c-2d34-4680-84ec-9912376a47c8

	Request Body
	```
	{

		"product_price"  :  70,

		"product_quantity"  :  4

	}
	```

	Response Body
	```
	{

		"product_id":  "bc446c6c-2d34-4680-84ec-9912376a47c8",

		"product_name":  "Product 3",

		"product_price":  70,

		"product_quantity":  4,

		"retailer_id":  "2004e491-72d1-404e-ba20-893f318caf2d"

	}
	```


### Order Routes
1. **GET** : Find order using uuid
	URL : */order/:id*
	Ex
	> /order/7afec1b7-1527-456c-b808-54e5fc098e57

	
2. **POST** : Add new order (containing multiple orders)
	URL : */order/add*
Ex Request body:
```
{
	"customer_id"  :  "af1f362b-862a-4b68-8297-f113e018d77d",
	"order_detail"  : [
		{
			"product_id"  :  "bb61baf1-6467-4804-aba6-d675ab08d2ad",
			"product_quantity"  :  4
		},
		{
			"product_id"  :  "bc446c6c-2d34-4680-84ec-9912376a47c8",
			"product_quantity"  :  8
		}
	]
}
```