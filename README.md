### Tax Calculator

Simple tax calculator for test assignment.

How To run:
1. Install docker in your machine
2. Run `docker-compose up --build`

*Make sure your machine is connected to internet. On first build, the Dockerfile will download go-dep and its dependencies*

Troubleshoot:
1. `docker-compose down -v`
2. `docker-compose up --build` once more

Live Demo:
`http://167.99.64.23:9091/shopee/tax/index`

#### A. Database Design

**table shopee_tax**

This table contains the tax definition for each tax category

| Field  | Type  | Desc  |
|---|---|---|
| id  | INT  | PRIMARY KEY, NOT NULL  |
| code  | VARCHAR  | NOT NULL  |
| description  | VARCHAR  | NULL  |
| expression  | VARCHAR  | NOT NULL  |
| min_price  | FLOAT  | NOT NULL DEFAULT 0.0  |

```
INSERT INTO shopee_tax (code, description, expression, min_price) VALUES 
    ('Food', 'tax for tobacco', '0.1 * ${price}', -1),
    ('Tobaco', 'tax for food items', '10 + (0.02 * ${price})', -1),
    ('Entertainment', 'tax for entertainment', '0.1 * (${price} - 100)', 99);
```
`expression` field contains tax calculation ruleset expression. 
It is possible to add developer assigned variable such as `${price}` , etc.

**table shopee_taxable_product**

This table contains calculated tax product snapshot

| Field  | Type  | Desc  |
|---|---|---|
| id  | INT  | PRIMARY KEY, NOT NULL  |
| product_name  | VARCHAR  | NOT NULL  |
| tax_category_id  | INT  | FK  |
| price  | FLOAT  | NOT NULL DEFAULT 0.0 |
| tax_price  | FLOAT  | NOT NULL DEFAULT 0.0  |
| total_price  | FLOAT  | NOT NULL DEFAULT 0.0  |

#### B. Code Structure

I uses the N-Layer design pattern with `observer` pattern support. To sustain Golang idiomaticity, I uses a lot of Go `closure` or anonymous function to replace the interface, or factory definition. This is the code structure of project module:

```
app
    +db
        +tax
            -db.go
            -query.go
    +domain
        +tax
            -helper.go
            -struct.go
            -
    +infra
        +lib1
        +lib2
    +model
        +tax
            -struct.go
    resource
        +uires
            -resource.go
        +repores
            -resource.go
        +usecaseres
            -resource.go
    usecase
        +tax
            -usecase_1.go
            -usecase_2.go
    server
        +http.go
        +server.go
    service
        +api
            -tax.go
```

Description:
1. `db`: db package has a single reponsibility as a database repository. It holds database accessor object, and has an ability to perform a query directly by using its connection string. This package has separated module folder based on its database accessor. i.e: `tax`, `user`, etc.
2. `model`: model package contains definition of business entity. Usually it is described as database table definition. `model` should not be accessible by other layer, **except** repository (i.e: database, and domain).
3. `domain`: domain contains of helper functions (internal logic) for a module business, structure builder, and connector between business logic (called usecase) and its model. Since business logic is unable to import model, it needs domain to access it.
4. `usecase`: usecase is the main business logic or knitting zone for the business logic. Use case has access to domain and repository directly, but not to model. This to ensure model is remain untouched, and refactor should be in domain only. 
5. `resource`: this package is used to holds all dependencies to each `usecase`, repository `db`, and server / presenter `ui`. By using dependency injection, we can make sure, the code is mockable and testable.
6. `service`: contains handler (similar to controller in MVC)
7. `server`: server definition and initialization, including routing handling.

#### C. Unit Test

Not added, (template only). Real implementation can be seen on `domain/tax/helper_test.go`