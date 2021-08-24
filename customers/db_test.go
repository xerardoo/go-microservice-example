package customers_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xerardoo/go-microservice-example/customers"
	"gorm.io/gorm"
	"regexp"
	"syreclabs.com/go/faker"
	"testing"
)

func TestCustomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Customer Suite")
}

var _ = Describe("Customer", func() {
	var db *sql.DB
	var mock sqlmock.Sqlmock
	var repo *Repos
	var customer Customer
	var ctx context.Context

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())
		repo = &Repos{Db: db}
		ctx = context.Background()

		customer = Customer{
			Model:       gorm.Model{ID: uint(faker.Number().NumberInt32(4))},
			FirstName:   faker.Name().FirstName(),
			LastName:    faker.Name().LastName(),
			PhoneNumber: faker.PhoneNumber().String(),
			Address:     faker.Address().String(),
			City:        faker.Address().City(),
			State:       faker.Address().State(),
			ZipCode:     faker.Address().ZipCode(),
		}
	})

	AfterEach(func() {
		defer db.Close()
		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Get All Customers", func() {

		It("Empty", func() {
			query := "SELECT * FROM customers"
			rows := sqlmock.NewRows(nil)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := repo.GetAllCustomers(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(c).To(BeEmpty())
		})

		It("Exist", func() {
			query := "SELECT * FROM `customers`"
			newId := uint(faker.Number().NumberInt32(4))
			rows := sqlmock.
				NewRows([]string{"id", "first_name", "last_name", "phone", "address", "city", "state", "zip_code"}).
				AddRow(customer.ID, customer.FirstName, customer.LastName, customer.PhoneNumber, customer.Address, customer.City, customer.State, customer.ZipCode).
				AddRow(newId, customer.FirstName, customer.LastName, customer.PhoneNumber, customer.Address, customer.City, customer.State, customer.ZipCode)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := repo.GetAllCustomers(ctx)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c[0].ID).Should(Equal(customer.ID))
			Expect(c[0].FirstName).Should(Equal(customer.FirstName))
			Expect(c[1].ID).Should(Equal(newId))
		})
	})
})
